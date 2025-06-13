package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
)

// Number of outbound HTTP connections that are currently active or idle on the client.
type ClientOpenConnections struct {
	*prometheus.GaugeVec
	extra ClientOpenConnectionsExtra
}

func NewClientOpenConnections() ClientOpenConnections {
	labels := []string{AttrConnectionState("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key(), network.AttrPeerAddress("").Key(), network.AttrProtocolVersion("").Key(), url.AttrScheme("").Key()}
	return ClientOpenConnections{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_client_open_connections",
		Help: "Number of outbound HTTP connections that are currently active or idle on the client.",
	}, labels)}
}

func (m ClientOpenConnections) With(connectionState AttrConnectionState, address server.AttrAddress, port server.AttrPort, extras ...interface {
	NetworkPeerAddress() network.AttrPeerAddress
	NetworkProtocolVersion() network.AttrProtocolVersion
	UrlScheme() url.AttrScheme
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(connectionState.Value(), address.Value(), port.Value(), extra.NetworkPeerAddress().Value(), extra.NetworkProtocolVersion().Value(), extra.UrlScheme().Value())
}

// Deprecated: Use [ClientOpenConnections.With] instead
func (m ClientOpenConnections) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ClientOpenConnections) WithNetworkPeerAddress(attr interface {
	NetworkPeerAddress() network.AttrPeerAddress
}) ClientOpenConnections {
	a.extra.AttrNetworkPeerAddress = attr.NetworkPeerAddress()
	return a
}
func (a ClientOpenConnections) WithNetworkProtocolVersion(attr interface {
	NetworkProtocolVersion() network.AttrProtocolVersion
}) ClientOpenConnections {
	a.extra.AttrNetworkProtocolVersion = attr.NetworkProtocolVersion()
	return a
}
func (a ClientOpenConnections) WithUrlScheme(attr interface{ UrlScheme() url.AttrScheme }) ClientOpenConnections {
	a.extra.AttrUrlScheme = attr.UrlScheme()
	return a
}

type ClientOpenConnectionsExtra struct {
	// Peer address of the network connection - IP address or Unix domain socket name
	AttrNetworkPeerAddress     network.AttrPeerAddress     `otel:"network.peer.address"`     // The actual version of the protocol used for network communication
	AttrNetworkProtocolVersion network.AttrProtocolVersion `otel:"network.protocol.version"` // The [URI scheme] component identifying the used protocol
	//
	// [URI scheme]: https://www.rfc-editor.org/rfc/rfc3986#section-3.1
	AttrUrlScheme url.AttrScheme `otel:"url.scheme"`
}

func (a ClientOpenConnectionsExtra) NetworkPeerAddress() network.AttrPeerAddress {
	return a.AttrNetworkPeerAddress
}
func (a ClientOpenConnectionsExtra) NetworkProtocolVersion() network.AttrProtocolVersion {
	return a.AttrNetworkProtocolVersion
}
func (a ClientOpenConnectionsExtra) UrlScheme() url.AttrScheme { return a.AttrUrlScheme }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientOpenConnectionsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.open_connections",
        "Type": "ClientOpenConnections",
        "attributes": [
            {
                "brief": "State of the HTTP connection in the HTTP connection pool.",
                "examples": [
                    "active",
                    "idle",
                ],
                "name": "http.connection.state",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "active state.",
                            "deprecated": none,
                            "id": "active",
                            "note": none,
                            "stability": "development",
                            "value": "active",
                        },
                        {
                            "brief": "idle state.",
                            "deprecated": none,
                            "id": "idle",
                            "note": none,
                            "stability": "development",
                            "value": "idle",
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
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                "examples": [
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "network.peer.address",
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
                "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                "examples": [
                    "http",
                    "https",
                ],
                "name": "url.scheme",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Peer address of the network connection - IP address or Unix domain socket name.",
                    "examples": [
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "network.peer.address",
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
                    "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                    "examples": [
                        "http",
                        "https",
                    ],
                    "name": "url.scheme",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "State of the HTTP connection in the HTTP connection pool.",
                    "examples": [
                        "active",
                        "idle",
                    ],
                    "name": "http.connection.state",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "active state.",
                                "deprecated": none,
                                "id": "active",
                                "note": none,
                                "stability": "development",
                                "value": "active",
                            },
                            {
                                "brief": "idle state.",
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                        ],
                    },
                },
                {
                    "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "int",
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
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "brief": "Number of outbound HTTP connections that are currently active or idle on the client.",
            "events": [],
            "id": "metric.http.client.open_connections",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "http.connection.state": {
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
                    "network.peer.address": {
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
                    "network.protocol.version": {
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
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "url.scheme": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.url",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.open_connections",
            "name": none,
            "root_namespace": "http",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{connection}",
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
