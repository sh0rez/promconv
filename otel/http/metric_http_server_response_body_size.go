package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
	"shorez.de/promconv/otel/user_agent"
)

// Size of HTTP server response bodies.
type ServerResponseBodySize struct {
	*prometheus.HistogramVec
	extra ServerResponseBodySizeExtra
}

func NewServerResponseBodySize() ServerResponseBodySize {
	labels := []string{"http_request_method", "url_scheme", "error_type", "http_response_status_code", "http_route", "network_protocol_name", "network_protocol_version", "server_address", "server_port", "user_agent_synthetic_type"}
	return ServerResponseBodySize{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "server_response_body_size",
		Help:      "Size of HTTP server response bodies.",
	}, labels)}
}

func (m ServerResponseBodySize) With(requestMethod AttrRequestMethod, scheme url.AttrScheme, extra interface {
	AttrErrorType() error.AttrType
	AttrHttpResponseStatusCode() AttrResponseStatusCode
	AttrHttpRoute() AttrRoute
	AttrNetworkProtocolName() network.AttrProtocolName
	AttrNetworkProtocolVersion() network.AttrProtocolVersion
	AttrServerAddress() server.AttrAddress
	AttrServerPort() server.AttrPort
	AttrUserAgentSyntheticType() user_agent.AttrSyntheticType
}) prometheus.Observer {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(requestMethod),
		string(scheme),
		string(extra.AttrErrorType()),
		string(extra.AttrHttpResponseStatusCode()),
		string(extra.AttrHttpRoute()),
		string(extra.AttrNetworkProtocolName()),
		string(extra.AttrNetworkProtocolVersion()),
		string(extra.AttrServerAddress()),
		string(extra.AttrServerPort()),
		string(extra.AttrUserAgentSyntheticType()),
	)
}

func (a ServerResponseBodySize) WithErrorType(attr interface{ AttrErrorType() error.AttrType }) ServerResponseBodySize {
	a.extra.ErrorType = attr.AttrErrorType()
	return a
}
func (a ServerResponseBodySize) WithHttpResponseStatusCode(attr interface{ AttrHttpResponseStatusCode() AttrResponseStatusCode }) ServerResponseBodySize {
	a.extra.HttpResponseStatusCode = attr.AttrHttpResponseStatusCode()
	return a
}
func (a ServerResponseBodySize) WithHttpRoute(attr interface{ AttrHttpRoute() AttrRoute }) ServerResponseBodySize {
	a.extra.HttpRoute = attr.AttrHttpRoute()
	return a
}
func (a ServerResponseBodySize) WithNetworkProtocolName(attr interface {
	AttrNetworkProtocolName() network.AttrProtocolName
}) ServerResponseBodySize {
	a.extra.NetworkProtocolName = attr.AttrNetworkProtocolName()
	return a
}
func (a ServerResponseBodySize) WithNetworkProtocolVersion(attr interface {
	AttrNetworkProtocolVersion() network.AttrProtocolVersion
}) ServerResponseBodySize {
	a.extra.NetworkProtocolVersion = attr.AttrNetworkProtocolVersion()
	return a
}
func (a ServerResponseBodySize) WithServerAddress(attr interface{ AttrServerAddress() server.AttrAddress }) ServerResponseBodySize {
	a.extra.ServerAddress = attr.AttrServerAddress()
	return a
}
func (a ServerResponseBodySize) WithServerPort(attr interface{ AttrServerPort() server.AttrPort }) ServerResponseBodySize {
	a.extra.ServerPort = attr.AttrServerPort()
	return a
}
func (a ServerResponseBodySize) WithUserAgentSyntheticType(attr interface {
	AttrUserAgentSyntheticType() user_agent.AttrSyntheticType
}) ServerResponseBodySize {
	a.extra.UserAgentSyntheticType = attr.AttrUserAgentSyntheticType()
	return a
}

type ServerResponseBodySizeExtra struct {
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// [HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).
	HttpResponseStatusCode AttrResponseStatusCode `otel:"http.response.status_code"`
	// The matched route, that is, the path template in the format used by the respective server framework.
	HttpRoute AttrRoute `otel:"http.route"`
	// [OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.
	NetworkProtocolName network.AttrProtocolName `otel:"network.protocol.name"`
	// The actual version of the protocol used for network communication.
	NetworkProtocolVersion network.AttrProtocolVersion `otel:"network.protocol.version"`
	// Name of the local HTTP server that received the request.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// Port of the local HTTP server that received the request.
	ServerPort server.AttrPort `otel:"server.port"`
	// Specifies the category of synthetic traffic, such as tests or bots.
	UserAgentSyntheticType user_agent.AttrSyntheticType `otel:"user_agent.synthetic.type"`
}

func (a ServerResponseBodySizeExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a ServerResponseBodySizeExtra) AttrHttpResponseStatusCode() AttrResponseStatusCode {
	return a.HttpResponseStatusCode
}
func (a ServerResponseBodySizeExtra) AttrHttpRoute() AttrRoute { return a.HttpRoute }
func (a ServerResponseBodySizeExtra) AttrNetworkProtocolName() network.AttrProtocolName {
	return a.NetworkProtocolName
}
func (a ServerResponseBodySizeExtra) AttrNetworkProtocolVersion() network.AttrProtocolVersion {
	return a.NetworkProtocolVersion
}
func (a ServerResponseBodySizeExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ServerResponseBodySizeExtra) AttrServerPort() server.AttrPort       { return a.ServerPort }
func (a ServerResponseBodySizeExtra) AttrUserAgentSyntheticType() user_agent.AttrSyntheticType {
	return a.UserAgentSyntheticType
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerResponseBodySizeExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.response.body.size",
        "Type": "ServerResponseBodySize",
        "attributes": [
            {
                "brief": "HTTP request method.",
                "examples": [
                    "GET",
                    "POST",
                    "HEAD",
                ],
                "name": "http.request.method",
                "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "CONNECT method.",
                            "deprecated": none,
                            "id": "connect",
                            "note": none,
                            "stability": "stable",
                            "value": "CONNECT",
                        },
                        {
                            "brief": "DELETE method.",
                            "deprecated": none,
                            "id": "delete",
                            "note": none,
                            "stability": "stable",
                            "value": "DELETE",
                        },
                        {
                            "brief": "GET method.",
                            "deprecated": none,
                            "id": "get",
                            "note": none,
                            "stability": "stable",
                            "value": "GET",
                        },
                        {
                            "brief": "HEAD method.",
                            "deprecated": none,
                            "id": "head",
                            "note": none,
                            "stability": "stable",
                            "value": "HEAD",
                        },
                        {
                            "brief": "OPTIONS method.",
                            "deprecated": none,
                            "id": "options",
                            "note": none,
                            "stability": "stable",
                            "value": "OPTIONS",
                        },
                        {
                            "brief": "PATCH method.",
                            "deprecated": none,
                            "id": "patch",
                            "note": none,
                            "stability": "stable",
                            "value": "PATCH",
                        },
                        {
                            "brief": "POST method.",
                            "deprecated": none,
                            "id": "post",
                            "note": none,
                            "stability": "stable",
                            "value": "POST",
                        },
                        {
                            "brief": "PUT method.",
                            "deprecated": none,
                            "id": "put",
                            "note": none,
                            "stability": "stable",
                            "value": "PUT",
                        },
                        {
                            "brief": "TRACE method.",
                            "deprecated": none,
                            "id": "trace",
                            "note": none,
                            "stability": "stable",
                            "value": "TRACE",
                        },
                        {
                            "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
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
                "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                "examples": [
                    "http",
                    "https",
                ],
                "name": "url.scheme",
                "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "timeout",
                    "java.net.UnknownHostException",
                    "server_certificate_invalid",
                    "500",
                ],
                "name": "error.type",
                "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                "requirement_level": {
                    "conditionally_required": "If request has ended with an error.",
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
                "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                "examples": [
                    200,
                ],
                "name": "http.response.status_code",
                "requirement_level": {
                    "conditionally_required": "If and only if one was received/sent.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                "examples": [
                    "/users/:userID?",
                    "{controller}/{action}/{id?}",
                ],
                "name": "http.route",
                "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if it's available",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                "examples": [
                    "http",
                    "spdy",
                ],
                "name": "network.protocol.name",
                "note": "The value SHOULD be normalized to lowercase.",
                "requirement_level": {
                    "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The actual version of the protocol used for network communication.",
                "examples": [
                    "1.0",
                    "1.1",
                    "2",
                    "3",
                ],
                "name": "network.protocol.version",
                "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Name of the local HTTP server that received the request.\n",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Port of the local HTTP server that received the request.\n",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                "name": "user_agent.synthetic.type",
                "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "Bot source.",
                            "deprecated": none,
                            "id": "bot",
                            "note": none,
                            "stability": "development",
                            "value": "bot",
                        },
                        {
                            "brief": "Synthetic test source.",
                            "deprecated": none,
                            "id": "test",
                            "note": none,
                            "stability": "development",
                            "value": "test",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                    "examples": [
                        200,
                    ],
                    "name": "http.response.status_code",
                    "requirement_level": {
                        "conditionally_required": "If and only if one was received/sent.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                    "examples": [
                        "/users/:userID?",
                        "{controller}/{action}/{id?}",
                    ],
                    "name": "http.route",
                    "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                    "requirement_level": {
                        "conditionally_required": "If and only if it's available",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "timeout",
                        "java.net.UnknownHostException",
                        "server_certificate_invalid",
                        "500",
                    ],
                    "name": "error.type",
                    "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                    "requirement_level": {
                        "conditionally_required": "If request has ended with an error.",
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
                    "brief": "HTTP request method.",
                    "examples": [
                        "GET",
                        "POST",
                        "HEAD",
                    ],
                    "name": "http.request.method",
                    "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "CONNECT method.",
                                "deprecated": none,
                                "id": "connect",
                                "note": none,
                                "stability": "stable",
                                "value": "CONNECT",
                            },
                            {
                                "brief": "DELETE method.",
                                "deprecated": none,
                                "id": "delete",
                                "note": none,
                                "stability": "stable",
                                "value": "DELETE",
                            },
                            {
                                "brief": "GET method.",
                                "deprecated": none,
                                "id": "get",
                                "note": none,
                                "stability": "stable",
                                "value": "GET",
                            },
                            {
                                "brief": "HEAD method.",
                                "deprecated": none,
                                "id": "head",
                                "note": none,
                                "stability": "stable",
                                "value": "HEAD",
                            },
                            {
                                "brief": "OPTIONS method.",
                                "deprecated": none,
                                "id": "options",
                                "note": none,
                                "stability": "stable",
                                "value": "OPTIONS",
                            },
                            {
                                "brief": "PATCH method.",
                                "deprecated": none,
                                "id": "patch",
                                "note": none,
                                "stability": "stable",
                                "value": "PATCH",
                            },
                            {
                                "brief": "POST method.",
                                "deprecated": none,
                                "id": "post",
                                "note": none,
                                "stability": "stable",
                                "value": "POST",
                            },
                            {
                                "brief": "PUT method.",
                                "deprecated": none,
                                "id": "put",
                                "note": none,
                                "stability": "stable",
                                "value": "PUT",
                            },
                            {
                                "brief": "TRACE method.",
                                "deprecated": none,
                                "id": "trace",
                                "note": none,
                                "stability": "stable",
                                "value": "TRACE",
                            },
                            {
                                "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
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
                    "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                    "name": "user_agent.synthetic.type",
                    "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Bot source.",
                                "deprecated": none,
                                "id": "bot",
                                "note": none,
                                "stability": "development",
                                "value": "bot",
                            },
                            {
                                "brief": "Synthetic test source.",
                                "deprecated": none,
                                "id": "test",
                                "note": none,
                                "stability": "development",
                                "value": "test",
                            },
                        ],
                    },
                },
                {
                    "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                    "examples": [
                        "http",
                        "spdy",
                    ],
                    "name": "network.protocol.name",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": {
                        "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The actual version of the protocol used for network communication.",
                    "examples": [
                        "1.0",
                        "1.1",
                        "2",
                        "3",
                    ],
                    "name": "network.protocol.version",
                    "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Name of the local HTTP server that received the request.\n",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Port of the local HTTP server that received the request.\n",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                    "examples": [
                        "http",
                        "https",
                    ],
                    "name": "url.scheme",
                    "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "brief": "Size of HTTP server response bodies.",
            "events": [],
            "id": "metric.http.server.response.body.size",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "http.request.method": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "http.response.status_code": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "http.route": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "network.protocol.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
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
                    "server.address": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "url.scheme": {
                        "inherited_fields": [
                            "brief",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.url",
                    },
                    "user_agent.synthetic.type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.user_agent.os",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.server.response.body.size",
            "name": none,
            "note": "The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
            "root_namespace": "http",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "By",
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
