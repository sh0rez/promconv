package kestrel

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
)

// Number of connections that are currently active on the server.
type ActiveConnections struct {
	*prometheus.GaugeVec
	extra ActiveConnectionsExtra
}

func NewActiveConnections() ActiveConnections {
	labels := []string{"network_transport", "network_type", "server_address", "server_port"}
	return ActiveConnections{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "kestrel",
		Name:      "active_connections",
		Help:      "Number of connections that are currently active on the server.",
	}, labels)}
}

func (m ActiveConnections) With(extra interface {
	AttrNetworkTransport() network.AttrTransport
	AttrNetworkType() network.AttrType
	AttrServerAddress() server.AttrAddress
	AttrServerPort() server.AttrPort
}) prometheus.Gauge {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrNetworkTransport()),
		string(extra.AttrNetworkType()),
		string(extra.AttrServerAddress()),
		string(extra.AttrServerPort()),
	)
}

func (a ActiveConnections) WithNetworkTransport(attr interface{ AttrNetworkTransport() network.AttrTransport }) ActiveConnections {
	a.extra.NetworkTransport = attr.AttrNetworkTransport()
	return a
}
func (a ActiveConnections) WithNetworkType(attr interface{ AttrNetworkType() network.AttrType }) ActiveConnections {
	a.extra.NetworkType = attr.AttrNetworkType()
	return a
}
func (a ActiveConnections) WithServerAddress(attr interface{ AttrServerAddress() server.AttrAddress }) ActiveConnections {
	a.extra.ServerAddress = attr.AttrServerAddress()
	return a
}
func (a ActiveConnections) WithServerPort(attr interface{ AttrServerPort() server.AttrPort }) ActiveConnections {
	a.extra.ServerPort = attr.AttrServerPort()
	return a
}

type ActiveConnectionsExtra struct {
	// [OSI transport layer](https://wikipedia.org/wiki/Transport_layer) or [inter-process communication method](https://wikipedia.org/wiki/Inter-process_communication).
	NetworkTransport network.AttrTransport `otel:"network.transport"`
	// [OSI network layer](https://wikipedia.org/wiki/Network_layer) or non-OSI equivalent.
	NetworkType network.AttrType `otel:"network.type"`
	// Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
}

func (a ActiveConnectionsExtra) AttrNetworkTransport() network.AttrTransport {
	return a.NetworkTransport
}
func (a ActiveConnectionsExtra) AttrNetworkType() network.AttrType     { return a.NetworkType }
func (a ActiveConnectionsExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ActiveConnectionsExtra) AttrServerPort() server.AttrPort       { return a.ServerPort }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ActiveConnectionsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "active_connections",
        "Type": "ActiveConnections",
        "attributes": [
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
        ],
        "ctx": {
            "attributes": [
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
            "brief": "Number of connections that are currently active on the server.",
            "events": [],
            "id": "metric.kestrel.active_connections",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
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
            "metric_name": "kestrel.active_connections",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Server.Kestrel`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "kestrel",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "{connection}",
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
