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

// The duration of TLS handshakes on the server.
type TlsHandshakeDuration struct {
	*prometheus.HistogramVec
	extra TlsHandshakeDurationExtra
}

func NewTlsHandshakeDuration() TlsHandshakeDuration {
	labels := []string{error.AttrType("").Key(), network.AttrTransport("").Key(), network.AttrType("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key(), tls.AttrProtocolVersion("").Key()}
	return TlsHandshakeDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "kestrel_tls_handshake_duration",
		Help: "The duration of TLS handshakes on the server.",
	}, labels)}
}

func (m TlsHandshakeDuration) With(extras ...interface {
	ErrorType() error.AttrType
	NetworkTransport() network.AttrTransport
	NetworkType() network.AttrType
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
	TlsProtocolVersion() tls.AttrProtocolVersion
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(extra.ErrorType().Value(), extra.NetworkTransport().Value(), extra.NetworkType().Value(), extra.ServerAddress().Value(), extra.ServerPort().Value(), extra.TlsProtocolVersion().Value())
}

// Deprecated: Use [TlsHandshakeDuration.With] instead
func (m TlsHandshakeDuration) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a TlsHandshakeDuration) WithErrorType(attr interface{ ErrorType() error.AttrType }) TlsHandshakeDuration {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a TlsHandshakeDuration) WithNetworkTransport(attr interface{ NetworkTransport() network.AttrTransport }) TlsHandshakeDuration {
	a.extra.AttrNetworkTransport = attr.NetworkTransport()
	return a
}
func (a TlsHandshakeDuration) WithNetworkType(attr interface{ NetworkType() network.AttrType }) TlsHandshakeDuration {
	a.extra.AttrNetworkType = attr.NetworkType()
	return a
}
func (a TlsHandshakeDuration) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) TlsHandshakeDuration {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a TlsHandshakeDuration) WithServerPort(attr interface{ ServerPort() server.AttrPort }) TlsHandshakeDuration {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}
func (a TlsHandshakeDuration) WithTlsProtocolVersion(attr interface {
	TlsProtocolVersion() tls.AttrProtocolVersion
}) TlsHandshakeDuration {
	a.extra.AttrTlsProtocolVersion = attr.TlsProtocolVersion()
	return a
}

type TlsHandshakeDurationExtra struct {
	// The full name of exception type
	AttrErrorType error.AttrType `otel:"error.type"` // [OSI transport layer] or [inter-process communication method]
	//
	// [OSI transport layer]: https://wikipedia.org/wiki/Transport_layer
	// [inter-process communication method]: https://wikipedia.org/wiki/Inter-process_communication
	AttrNetworkTransport network.AttrTransport `otel:"network.transport"` // [OSI network layer] or non-OSI equivalent
	//
	// [OSI network layer]: https://wikipedia.org/wiki/Network_layer
	AttrNetworkType   network.AttrType   `otel:"network.type"`   // Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
	AttrServerAddress server.AttrAddress `otel:"server.address"` // Server port number
	AttrServerPort    server.AttrPort    `otel:"server.port"`    // Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version]
	//
	// [SSL/TLS protocol version]: https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values
	AttrTlsProtocolVersion tls.AttrProtocolVersion `otel:"tls.protocol.version"`
}

func (a TlsHandshakeDurationExtra) ErrorType() error.AttrType { return a.AttrErrorType }
func (a TlsHandshakeDurationExtra) NetworkTransport() network.AttrTransport {
	return a.AttrNetworkTransport
}
func (a TlsHandshakeDurationExtra) NetworkType() network.AttrType     { return a.AttrNetworkType }
func (a TlsHandshakeDurationExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }
func (a TlsHandshakeDurationExtra) ServerPort() server.AttrPort       { return a.AttrServerPort }
func (a TlsHandshakeDurationExtra) TlsProtocolVersion() tls.AttrProtocolVersion {
	return a.AttrTlsProtocolVersion
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "TlsHandshakeDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "tls_handshake.duration",
        "Type": "TlsHandshakeDuration",
        "attributes": [
            {
                "brief": "The full name of exception type.",
                "examples": [
                    "System.OperationCanceledException",
                    "Contoso.MyException",
                ],
                "name": "error.type",
                "note": "Captures the exception type when a TLS handshake fails.",
                "requirement_level": {
                    "conditionally_required": "if and only if an error has occurred.",
                },
                "stability": "stable",
                "type": {
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
                        "System.OperationCanceledException",
                        "Contoso.MyException",
                    ],
                    "name": "error.type",
                    "note": "Captures the exception type when a TLS handshake fails.",
                    "requirement_level": {
                        "conditionally_required": "if and only if an error has occurred.",
                    },
                    "stability": "stable",
                    "type": {
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
            ],
            "brief": "The duration of TLS handshakes on the server.",
            "events": [],
            "id": "metric.kestrel.tls_handshake.duration",
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
            "metric_name": "kestrel.tls_handshake.duration",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Server.Kestrel`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "kestrel",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "s",
        },
        "for_each_attr": <macro for_each_attr>,
        "module": "shorez.de/promconv/otel",
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
            "attribute_id",
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
            "vec.go.j2",
        ],
    },
}
*/
