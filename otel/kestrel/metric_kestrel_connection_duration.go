package kestrel

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/tls"
)

// The duration of connections on the server.
type ConnectionDuration struct {
	*prometheus.HistogramVec
}

func NewConnectionDuration() ConnectionDuration {
	labels := []string{"error_type", "network_protocol_name", "network_protocol_version", "network_transport", "network_type", "server_address", "server_port", "tls_protocol_version"}
	return ConnectionDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "kestrel",
		Name:      "connection_duration",
		Help:      "The duration of connections on the server.",
	}, labels)}
}

func (m ConnectionDuration) With(extra interface {
	AttrErrorType() error.AttrType
	AttrNetworkProtocolName() network.AttrProtocolName
	AttrNetworkProtocolVersion() network.AttrProtocolVersion
	AttrNetworkTransport() network.AttrTransport
	AttrNetworkType() network.AttrType
	AttrServerAddress() server.AttrAddress
	AttrServerPort() server.AttrPort
	AttrTlsProtocolVersion() tls.AttrProtocolVersion
}) prometheus.Observer {
	if extra == nil {
		extra = ConnectionDurationExtra{}
	}
	return m.WithLabelValues(
		string(extra.AttrErrorType()),
		string(extra.AttrNetworkProtocolName()),
		string(extra.AttrNetworkProtocolVersion()),
		string(extra.AttrNetworkTransport()),
		string(extra.AttrNetworkType()),
		string(extra.AttrServerAddress()),
		string(extra.AttrServerPort()),
		string(extra.AttrTlsProtocolVersion()),
	)
}

type ConnectionDurationExtra struct {
	// The full name of exception type.
	ErrorType error.AttrType `otel:"error.type"`
	// [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.
	NetworkProtocolName network.AttrProtocolName `otel:"network.protocol.name"`
	// The actual version of the protocol used for network communication.
	NetworkProtocolVersion network.AttrProtocolVersion `otel:"network.protocol.version"`
	// [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
	NetworkTransport network.AttrTransport `otel:"network.transport"`
	// [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.
	NetworkType network.AttrType `otel:"network.type"`
	// Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
	// Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)
	TlsProtocolVersion tls.AttrProtocolVersion `otel:"tls.protocol.version"`
}

func (a ConnectionDurationExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a ConnectionDurationExtra) AttrNetworkProtocolName() network.AttrProtocolName {
	return a.NetworkProtocolName
}
func (a ConnectionDurationExtra) AttrNetworkProtocolVersion() network.AttrProtocolVersion {
	return a.NetworkProtocolVersion
}
func (a ConnectionDurationExtra) AttrNetworkTransport() network.AttrTransport {
	return a.NetworkTransport
}
func (a ConnectionDurationExtra) AttrNetworkType() network.AttrType     { return a.NetworkType }
func (a ConnectionDurationExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ConnectionDurationExtra) AttrServerPort() server.AttrPort       { return a.ServerPort }
func (a ConnectionDurationExtra) AttrTlsProtocolVersion() tls.AttrProtocolVersion {
	return a.TlsProtocolVersion
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ConnectionDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "connection.duration",
        "Type": "ConnectionDuration",
        "attributes": [
            {
                "brief": "The full name of exception type.",
                "examples": [
                    "connection_reset",
                    "invalid_handshake",
                ],
                "name": "error.type",
                "note": "Starting from .NET 9, Kestrel `kestrel.connection.duration` metric reports\nthe following errors types when a corresponding error occurs:\n\n| Value  | Description | Stability |\n|---|---|---|\n| `aborted_by_app` | The HTTP/1.1 connection was aborted when app code aborted an HTTP request with `HttpContext.Abort()`. |\n| `app_shutdown_timeout` | The connection was aborted during app shutdown. During shutdown, the server stops accepting new connections and HTTP requests, and it is given time for active requests to complete. If the app shutdown timeout is exceeded, all remaining connections are aborted. |\n| `closed_critical_stream` | A critical control stream for an HTTP/3 connection was closed. |\n| `connection_reset` | The connection was reset while there were active HTTP/2 or HTTP/3 streams on the connection. |\n| `error_after_starting_response` | An error such as an unhandled application exception or invalid request body occurred after the response was started, causing an abort of the HTTP/1.1 connection. |\n| `error_reading_headers` | An error occurred when decoding HPACK headers in an HTTP/2 `HEADERS` frame. |\n| `error_writing_headers` | An error occurred when encoding HPACK headers in an HTTP/2 `HEADERS` frame. |\n| `flow_control_queue_size_exceeded` | The connection exceeded the outgoing flow control maximum queue size and was closed with `INTERNAL_ERROR`. This can be caused by an excessive number of HTTP/2 stream resets. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `flow_control_window_exceeded` | The client sent more data than allowed by the current flow-control window. |\n| `frame_after_stream_close` | An HTTP/2 frame was received on a closed stream. |\n| `insufficient_tls_version` | The connection doesn't have TLS 1.2 or greater, as required by HTTP/2. |\n| `invalid_body_reader_state` | An error occurred when draining the request body, aborting the HTTP/1.1 connection. This could be caused by app code reading the request body and missing a call to `PipeReader.AdvanceTo` in a finally block. |\n| `invalid_data_padding` | An HTTP/2 `HEADER` or `DATA` frame has an invalid amount of padding. |\n| `invalid_frame_length` | An HTTP/2 frame was received with an invalid frame payload length. The frame could contain a payload that is not valid for the type, or a `DATA` frame payload does not match the length specified in the frame header. |\n| `invalid_handshake` | An invalid HTTP/2 handshake was received. |\n| `invalid_http_version` | The connection received an HTTP request with the wrong version. For example, a browser sends an HTTP/1.1 request to a plain-text HTTP/2 connection. |\n| `invalid_request_headers` | The HTTP request contains invalid headers. This error can occur in a number of scenarios: a header might not be allowed by the HTTP protocol, such as a pseudo-header in the `HEADERS` frame of an HTTP/2 request. A header could also have an invalid value, such as a non-integer `content-length`, or a header name or value might contain invalid characters. |\n| `invalid_request_line` | The first line of an HTTP/1.1 request was invalid, potentially due to invalid content or exceeding the allowed limit. Configured by `KestrelServerLimits.MaxRequestLineSize`. |\n| `invalid_settings` | The connection received an HTTP/2 or HTTP/3 `SETTINGS` frame with invalid settings. |\n| `invalid_stream_id` | An HTTP/2 stream with an invalid stream ID was received. |\n| `invalid_window_update_size` | The server received an HTTP/2 `WINDOW_UPDATE` frame with a zero increment, or an increment that caused a flow-control window to exceed the maximum size. |\n| `io_error` | An `IOException` occurred while reading or writing HTTP/2 or HTTP/3 connection data. |\n| `keep_alive_timeout` | There was no activity on the connection, and the keep-alive timeout configured by `KestrelServerLimits.KeepAliveTimeout` was exceeded. |\n| `max_concurrent_connections_exceeded` | The connection exceeded the maximum concurrent connection limit. Configured by `KestrelServerLimits.MaxConcurrentConnections`. |\n| `max_frame_length_exceeded` | The connection received an HTTP/2 frame that exceeded the size limit specified by `Http2Limits.MaxFrameSize`. |\n| `max_request_body_size_exceeded` | The HTTP request body exceeded the maximum request body size limit. Configured by `KestrelServerLimits.MaxRequestBodySize`. |\n| `max_request_header_count_exceeded` | The HTTP request headers exceeded the maximum count limit. Configured by `KestrelServerLimits.MaxRequestHeaderCount`. |\n| `max_request_headers_total_size_exceeded` | The HTTP request headers exceeded the maximum total size limit. Configured by `KestrelServerLimits.MaxRequestHeadersTotalSize`. |\n| `min_request_body_data_rate` | Reading the request body timed out due to data arriving too slowly. Configured by `KestrelServerLimits.MinRequestBodyDataRate`. |\n| `min_response_data_rate` | Writing the response timed out because the client did not read it at the specified minimum data rate. Configured by `KestrelServerLimits.MinResponseDataRate`. |\n| `missing_stream_end` | The connection received an HTTP/2 `HEADERS` frame for trailers without a stream end flag. |\n| `output_queue_size_exceeded` | The connection exceeded the output queue size and was closed with `INTERNAL_ERROR`. This can be caused by an excessive number of HTTP/2 stream resets. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `request_headers_timeout` | Request headers timed out while waiting for headers to be received after the request started. Configured by `KestrelServerLimits.RequestHeadersTimeout`. |\n| `response_content_length_mismatch` | The HTTP response body sent data that didn't match the response's `content-length` header. |\n| `server_timeout` | The connection timed out with the `IConnectionTimeoutFeature`. |\n| `stream_creation_error` | The HTTP/3 connection received a stream that it wouldn't accept. For example, the client created duplicate control streams. |\n| `stream_reset_limit_exceeded` | The connection received an excessive number of HTTP/2 stream resets and was closed with `ENHANCE_YOUR_CALM`. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `stream_self_dependency` | The connection received an HTTP/2 frame that caused a frame to depend on itself. |\n| `tls_handshake_failed` | An error occurred during the TLS handshake for a connection. Only reported for HTTP/1.1 and HTTP/2 connections. The TLS handshake for HTTP/3 is internal to QUIC transport. ![Development](https://img.shields.io/badge/-development-blue) |\n| `tls_not_supported` | A TLS handshake was received by an endpoint that isn't configured to support TLS. |\n| `unexpected_end_of_request_content` | The HTTP/1.1 request body ended before the data specified by the `content-length` header or chunked transfer encoding mechanism was received. |\n| `unexpected_frame` | An unexpected HTTP/2 or HTTP/3 frame type was received. The frame type is either unknown, unsupported, or invalid for the current stream state. |\n| `unknown_stream` | An HTTP/2 frame was received on an unknown stream. |\n| `write_canceled` | The cancellation of a response body write aborted the HTTP/1.1 connection. |\n\nIn other cases, `error.type` contains the fully qualified type name of the exception.\n",
                "requirement_level": {
                    "conditionally_required": "if and only if an error has occurred.",
                },
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                "examples": [
                    "http",
                    "web_sockets",
                ],
                "name": "network.protocol.name",
                "note": "The value SHOULD be normalized to lowercase.",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The actual version of the protocol used for network communication.",
                "examples": [
                    "1.1",
                    "2",
                ],
                "name": "network.protocol.version",
                "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "[OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).\n",
                "examples": [
                    "tcp",
                    "unix",
                ],
                "name": "network.transport",
                "note": "The value SHOULD be normalized to lowercase.\n\nConsider always setting the transport when setting a port number, since\na port number is ambiguous without knowing the transport. For example\ndifferent processes could be listening on TCP port 12345 and UDP port 12345.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "TCP",
                            "deprecated": none,
                            "id": "tcp",
                            "note": none,
                            "stability": "stable",
                            "value": "tcp",
                        },
                        {
                            "brief": "UDP",
                            "deprecated": none,
                            "id": "udp",
                            "note": none,
                            "stability": "stable",
                            "value": "udp",
                        },
                        {
                            "brief": "Named or anonymous pipe.",
                            "deprecated": none,
                            "id": "pipe",
                            "note": none,
                            "stability": "stable",
                            "value": "pipe",
                        },
                        {
                            "brief": "Unix domain socket",
                            "deprecated": none,
                            "id": "unix",
                            "note": none,
                            "stability": "stable",
                            "value": "unix",
                        },
                        {
                            "brief": "QUIC",
                            "deprecated": none,
                            "id": "quic",
                            "note": none,
                            "stability": "stable",
                            "value": "quic",
                        },
                    ],
                },
            },
            {
                "brief": "[OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.",
                "examples": [
                    "ipv4",
                    "ipv6",
                ],
                "name": "network.type",
                "note": "The value SHOULD be normalized to lowercase.",
                "requirement_level": {
                    "recommended": "if the transport is `tcp` or `udp`",
                },
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "IPv4",
                            "deprecated": none,
                            "id": "ipv4",
                            "note": none,
                            "stability": "stable",
                            "value": "ipv4",
                        },
                        {
                            "brief": "IPv6",
                            "deprecated": none,
                            "id": "ipv6",
                            "note": none,
                            "stability": "stable",
                            "value": "ipv6",
                        },
                    ],
                },
            },
            {
                "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Server port number.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)\n",
                "examples": [
                    "1.2",
                    "3",
                ],
                "name": "tls.protocol.version",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The actual version of the protocol used for network communication.",
                    "examples": [
                        "1.1",
                        "2",
                    ],
                    "name": "network.protocol.version",
                    "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "[OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.",
                    "examples": [
                        "ipv4",
                        "ipv6",
                    ],
                    "name": "network.type",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": {
                        "recommended": "if the transport is `tcp` or `udp`",
                    },
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "IPv4",
                                "deprecated": none,
                                "id": "ipv4",
                                "note": none,
                                "stability": "stable",
                                "value": "ipv4",
                            },
                            {
                                "brief": "IPv6",
                                "deprecated": none,
                                "id": "ipv6",
                                "note": none,
                                "stability": "stable",
                                "value": "ipv6",
                            },
                        ],
                    },
                },
                {
                    "brief": "[OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).\n",
                    "examples": [
                        "tcp",
                        "unix",
                    ],
                    "name": "network.transport",
                    "note": "The value SHOULD be normalized to lowercase.\n\nConsider always setting the transport when setting a port number, since\na port number is ambiguous without knowing the transport. For example\ndifferent processes could be listening on TCP port 12345 and UDP port 12345.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "TCP",
                                "deprecated": none,
                                "id": "tcp",
                                "note": none,
                                "stability": "stable",
                                "value": "tcp",
                            },
                            {
                                "brief": "UDP",
                                "deprecated": none,
                                "id": "udp",
                                "note": none,
                                "stability": "stable",
                                "value": "udp",
                            },
                            {
                                "brief": "Named or anonymous pipe.",
                                "deprecated": none,
                                "id": "pipe",
                                "note": none,
                                "stability": "stable",
                                "value": "pipe",
                            },
                            {
                                "brief": "Unix domain socket",
                                "deprecated": none,
                                "id": "unix",
                                "note": none,
                                "stability": "stable",
                                "value": "unix",
                            },
                            {
                                "brief": "QUIC",
                                "deprecated": none,
                                "id": "quic",
                                "note": none,
                                "stability": "stable",
                                "value": "quic",
                            },
                        ],
                    },
                },
                {
                    "brief": "The full name of exception type.",
                    "examples": [
                        "connection_reset",
                        "invalid_handshake",
                    ],
                    "name": "error.type",
                    "note": "Starting from .NET 9, Kestrel `kestrel.connection.duration` metric reports\nthe following errors types when a corresponding error occurs:\n\n| Value  | Description | Stability |\n|---|---|---|\n| `aborted_by_app` | The HTTP/1.1 connection was aborted when app code aborted an HTTP request with `HttpContext.Abort()`. |\n| `app_shutdown_timeout` | The connection was aborted during app shutdown. During shutdown, the server stops accepting new connections and HTTP requests, and it is given time for active requests to complete. If the app shutdown timeout is exceeded, all remaining connections are aborted. |\n| `closed_critical_stream` | A critical control stream for an HTTP/3 connection was closed. |\n| `connection_reset` | The connection was reset while there were active HTTP/2 or HTTP/3 streams on the connection. |\n| `error_after_starting_response` | An error such as an unhandled application exception or invalid request body occurred after the response was started, causing an abort of the HTTP/1.1 connection. |\n| `error_reading_headers` | An error occurred when decoding HPACK headers in an HTTP/2 `HEADERS` frame. |\n| `error_writing_headers` | An error occurred when encoding HPACK headers in an HTTP/2 `HEADERS` frame. |\n| `flow_control_queue_size_exceeded` | The connection exceeded the outgoing flow control maximum queue size and was closed with `INTERNAL_ERROR`. This can be caused by an excessive number of HTTP/2 stream resets. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `flow_control_window_exceeded` | The client sent more data than allowed by the current flow-control window. |\n| `frame_after_stream_close` | An HTTP/2 frame was received on a closed stream. |\n| `insufficient_tls_version` | The connection doesn't have TLS 1.2 or greater, as required by HTTP/2. |\n| `invalid_body_reader_state` | An error occurred when draining the request body, aborting the HTTP/1.1 connection. This could be caused by app code reading the request body and missing a call to `PipeReader.AdvanceTo` in a finally block. |\n| `invalid_data_padding` | An HTTP/2 `HEADER` or `DATA` frame has an invalid amount of padding. |\n| `invalid_frame_length` | An HTTP/2 frame was received with an invalid frame payload length. The frame could contain a payload that is not valid for the type, or a `DATA` frame payload does not match the length specified in the frame header. |\n| `invalid_handshake` | An invalid HTTP/2 handshake was received. |\n| `invalid_http_version` | The connection received an HTTP request with the wrong version. For example, a browser sends an HTTP/1.1 request to a plain-text HTTP/2 connection. |\n| `invalid_request_headers` | The HTTP request contains invalid headers. This error can occur in a number of scenarios: a header might not be allowed by the HTTP protocol, such as a pseudo-header in the `HEADERS` frame of an HTTP/2 request. A header could also have an invalid value, such as a non-integer `content-length`, or a header name or value might contain invalid characters. |\n| `invalid_request_line` | The first line of an HTTP/1.1 request was invalid, potentially due to invalid content or exceeding the allowed limit. Configured by `KestrelServerLimits.MaxRequestLineSize`. |\n| `invalid_settings` | The connection received an HTTP/2 or HTTP/3 `SETTINGS` frame with invalid settings. |\n| `invalid_stream_id` | An HTTP/2 stream with an invalid stream ID was received. |\n| `invalid_window_update_size` | The server received an HTTP/2 `WINDOW_UPDATE` frame with a zero increment, or an increment that caused a flow-control window to exceed the maximum size. |\n| `io_error` | An `IOException` occurred while reading or writing HTTP/2 or HTTP/3 connection data. |\n| `keep_alive_timeout` | There was no activity on the connection, and the keep-alive timeout configured by `KestrelServerLimits.KeepAliveTimeout` was exceeded. |\n| `max_concurrent_connections_exceeded` | The connection exceeded the maximum concurrent connection limit. Configured by `KestrelServerLimits.MaxConcurrentConnections`. |\n| `max_frame_length_exceeded` | The connection received an HTTP/2 frame that exceeded the size limit specified by `Http2Limits.MaxFrameSize`. |\n| `max_request_body_size_exceeded` | The HTTP request body exceeded the maximum request body size limit. Configured by `KestrelServerLimits.MaxRequestBodySize`. |\n| `max_request_header_count_exceeded` | The HTTP request headers exceeded the maximum count limit. Configured by `KestrelServerLimits.MaxRequestHeaderCount`. |\n| `max_request_headers_total_size_exceeded` | The HTTP request headers exceeded the maximum total size limit. Configured by `KestrelServerLimits.MaxRequestHeadersTotalSize`. |\n| `min_request_body_data_rate` | Reading the request body timed out due to data arriving too slowly. Configured by `KestrelServerLimits.MinRequestBodyDataRate`. |\n| `min_response_data_rate` | Writing the response timed out because the client did not read it at the specified minimum data rate. Configured by `KestrelServerLimits.MinResponseDataRate`. |\n| `missing_stream_end` | The connection received an HTTP/2 `HEADERS` frame for trailers without a stream end flag. |\n| `output_queue_size_exceeded` | The connection exceeded the output queue size and was closed with `INTERNAL_ERROR`. This can be caused by an excessive number of HTTP/2 stream resets. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `request_headers_timeout` | Request headers timed out while waiting for headers to be received after the request started. Configured by `KestrelServerLimits.RequestHeadersTimeout`. |\n| `response_content_length_mismatch` | The HTTP response body sent data that didn't match the response's `content-length` header. |\n| `server_timeout` | The connection timed out with the `IConnectionTimeoutFeature`. |\n| `stream_creation_error` | The HTTP/3 connection received a stream that it wouldn't accept. For example, the client created duplicate control streams. |\n| `stream_reset_limit_exceeded` | The connection received an excessive number of HTTP/2 stream resets and was closed with `ENHANCE_YOUR_CALM`. For more information, see [Microsoft Security Advisory CVE-2023-44487](https://github.com/dotnet/runtime/issues/93303). |\n| `stream_self_dependency` | The connection received an HTTP/2 frame that caused a frame to depend on itself. |\n| `tls_handshake_failed` | An error occurred during the TLS handshake for a connection. Only reported for HTTP/1.1 and HTTP/2 connections. The TLS handshake for HTTP/3 is internal to QUIC transport. ![Development](https://img.shields.io/badge/-development-blue) |\n| `tls_not_supported` | A TLS handshake was received by an endpoint that isn't configured to support TLS. |\n| `unexpected_end_of_request_content` | The HTTP/1.1 request body ended before the data specified by the `content-length` header or chunked transfer encoding mechanism was received. |\n| `unexpected_frame` | An unexpected HTTP/2 or HTTP/3 frame type was received. The frame type is either unknown, unsupported, or invalid for the current stream state. |\n| `unknown_stream` | An HTTP/2 frame was received on an unknown stream. |\n| `write_canceled` | The cancellation of a response body write aborted the HTTP/1.1 connection. |\n\nIn other cases, `error.type` contains the fully qualified type name of the exception.\n",
                    "requirement_level": {
                        "conditionally_required": "if and only if an error has occurred.",
                    },
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
                },
                {
                    "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                    "examples": [
                        "http",
                        "web_sockets",
                    ],
                    "name": "network.protocol.name",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)\n",
                    "examples": [
                        "1.2",
                        "3",
                    ],
                    "name": "tls.protocol.version",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "The duration of connections on the server.",
            "events": [],
            "id": "metric.kestrel.connection.duration",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "network.protocol.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.protocol.version": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.transport": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.network",
                    },
                    "server.address": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.server",
                    },
                    "tls.protocol.version": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.tls",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/kestrel/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "kestrel.connection.duration",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Server.Kestrel`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "kestrel",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "s",
        },
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "params": {},
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "metric.go.j2",
        ],
    },
}
*/
