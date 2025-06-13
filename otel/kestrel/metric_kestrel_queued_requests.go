package kestrel

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
)

// Number of HTTP requests on multiplexed connections (HTTP/2 and HTTP/3) that are currently queued and are waiting to start.
type QueuedRequests struct {
	*prometheus.GaugeVec
	extra QueuedRequestsExtra
}

func NewQueuedRequests() QueuedRequests {
	labels := []string{network.AttrProtocolName("").Key(), network.AttrProtocolVersion("").Key(), network.AttrTransport("").Key(), network.AttrType("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key()}
	return QueuedRequests{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kestrel_queued_requests",
		Help: "Number of HTTP requests on multiplexed connections (HTTP/2 and HTTP/3) that are currently queued and are waiting to start.",
	}, labels)}
}

func (m QueuedRequests) With(extras ...interface {
	NetworkProtocolName() network.AttrProtocolName
	NetworkProtocolVersion() network.AttrProtocolVersion
	NetworkTransport() network.AttrTransport
	NetworkType() network.AttrType
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.NetworkProtocolName().Value(), extra.NetworkProtocolVersion().Value(), extra.NetworkTransport().Value(), extra.NetworkType().Value(), extra.ServerAddress().Value(), extra.ServerPort().Value())
}

// Deprecated: Use [QueuedRequests.With] instead
func (m QueuedRequests) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a QueuedRequests) WithNetworkProtocolName(attr interface {
	NetworkProtocolName() network.AttrProtocolName
}) QueuedRequests {
	a.extra.AttrNetworkProtocolName = attr.NetworkProtocolName()
	return a
}
func (a QueuedRequests) WithNetworkProtocolVersion(attr interface {
	NetworkProtocolVersion() network.AttrProtocolVersion
}) QueuedRequests {
	a.extra.AttrNetworkProtocolVersion = attr.NetworkProtocolVersion()
	return a
}
func (a QueuedRequests) WithNetworkTransport(attr interface{ NetworkTransport() network.AttrTransport }) QueuedRequests {
	a.extra.AttrNetworkTransport = attr.NetworkTransport()
	return a
}
func (a QueuedRequests) WithNetworkType(attr interface{ NetworkType() network.AttrType }) QueuedRequests {
	a.extra.AttrNetworkType = attr.NetworkType()
	return a
}
func (a QueuedRequests) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) QueuedRequests {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a QueuedRequests) WithServerPort(attr interface{ ServerPort() server.AttrPort }) QueuedRequests {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}

type QueuedRequestsExtra struct {
	// [OSI application layer] or non-OSI equivalent
	//
	// [OSI application layer]: https://wikipedia.org/wiki/Application_layer
	AttrNetworkProtocolName    network.AttrProtocolName    `otel:"network.protocol.name"`    // The actual version of the protocol used for network communication
	AttrNetworkProtocolVersion network.AttrProtocolVersion `otel:"network.protocol.version"` // [OSI transport layer] or [inter-process communication method]
	//
	// [OSI transport layer]: https://wikipedia.org/wiki/Transport_layer
	// [inter-process communication method]: https://wikipedia.org/wiki/Inter-process_communication
	AttrNetworkTransport network.AttrTransport `otel:"network.transport"` // [OSI network layer] or non-OSI equivalent
	//
	// [OSI network layer]: https://wikipedia.org/wiki/Network_layer
	AttrNetworkType   network.AttrType   `otel:"network.type"`   // Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
	AttrServerAddress server.AttrAddress `otel:"server.address"` // Server port number
	AttrServerPort    server.AttrPort    `otel:"server.port"`
}

func (a QueuedRequestsExtra) NetworkProtocolName() network.AttrProtocolName {
	return a.AttrNetworkProtocolName
}
func (a QueuedRequestsExtra) NetworkProtocolVersion() network.AttrProtocolVersion {
	return a.AttrNetworkProtocolVersion
}
func (a QueuedRequestsExtra) NetworkTransport() network.AttrTransport { return a.AttrNetworkTransport }
func (a QueuedRequestsExtra) NetworkType() network.AttrType           { return a.AttrNetworkType }
func (a QueuedRequestsExtra) ServerAddress() server.AttrAddress       { return a.AttrServerAddress }
func (a QueuedRequestsExtra) ServerPort() server.AttrPort             { return a.AttrServerPort }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "QueuedRequestsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "queued_requests",
        "Type": "QueuedRequests",
        "attributes": [
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
        ],
        "ctx": {
            "attributes": [
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
            ],
            "brief": "Number of HTTP requests on multiplexed connections (HTTP/2 and HTTP/3) that are currently queued and are waiting to start.",
            "events": [],
            "id": "metric.kestrel.queued_requests",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/kestrel/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "kestrel.queued_requests",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Server.Kestrel`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "kestrel",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "{request}",
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
