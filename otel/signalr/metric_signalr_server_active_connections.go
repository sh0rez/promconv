package signalr

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of connections that are currently active on the server.
type ServerActiveConnections struct {
	*prometheus.GaugeVec
	extra ServerActiveConnectionsExtra
}

func NewServerActiveConnections() ServerActiveConnections {
	labels := []string{"signalr_connection_status", "signalr_transport"}
	return ServerActiveConnections{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "signalr",
		Name:      "server_active_connections",
		Help:      "Number of connections that are currently active on the server.",
	}, labels)}
}

func (m ServerActiveConnections) With(extra interface {
	AttrSignalrConnectionStatus() AttrConnectionStatus
	AttrSignalrTransport() AttrTransport
}) prometheus.Gauge {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrSignalrConnectionStatus()),
		string(extra.AttrSignalrTransport()),
	)
}

func (a ServerActiveConnections) WithSignalrConnectionStatus(attr interface{ AttrSignalrConnectionStatus() AttrConnectionStatus }) ServerActiveConnections {
	a.extra.SignalrConnectionStatus = attr.AttrSignalrConnectionStatus()
	return a
}
func (a ServerActiveConnections) WithSignalrTransport(attr interface{ AttrSignalrTransport() AttrTransport }) ServerActiveConnections {
	a.extra.SignalrTransport = attr.AttrSignalrTransport()
	return a
}

type ServerActiveConnectionsExtra struct {
	// SignalR HTTP connection closure status.
	SignalrConnectionStatus AttrConnectionStatus `otel:"signalr.connection.status"`
	// [SignalR transport type](https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md)
	SignalrTransport AttrTransport `otel:"signalr.transport"`
}

func (a ServerActiveConnectionsExtra) AttrSignalrConnectionStatus() AttrConnectionStatus {
	return a.SignalrConnectionStatus
}
func (a ServerActiveConnectionsExtra) AttrSignalrTransport() AttrTransport { return a.SignalrTransport }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerActiveConnectionsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.active_connections",
        "Type": "ServerActiveConnections",
        "attributes": [
            {
                "brief": "SignalR HTTP connection closure status.",
                "examples": [
                    "app_shutdown",
                    "timeout",
                ],
                "name": "signalr.connection.status",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "The connection was closed normally.",
                            "deprecated": none,
                            "id": "normal_closure",
                            "note": none,
                            "stability": "stable",
                            "value": "normal_closure",
                        },
                        {
                            "brief": "The connection was closed due to a timeout.",
                            "deprecated": none,
                            "id": "timeout",
                            "note": none,
                            "stability": "stable",
                            "value": "timeout",
                        },
                        {
                            "brief": "The connection was closed because the app is shutting down.",
                            "deprecated": none,
                            "id": "app_shutdown",
                            "note": none,
                            "stability": "stable",
                            "value": "app_shutdown",
                        },
                    ],
                },
            },
            {
                "brief": "[SignalR transport type](https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md)",
                "examples": [
                    "web_sockets",
                    "long_polling",
                ],
                "name": "signalr.transport",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "ServerSentEvents protocol",
                            "deprecated": none,
                            "id": "server_sent_events",
                            "note": none,
                            "stability": "stable",
                            "value": "server_sent_events",
                        },
                        {
                            "brief": "LongPolling protocol",
                            "deprecated": none,
                            "id": "long_polling",
                            "note": none,
                            "stability": "stable",
                            "value": "long_polling",
                        },
                        {
                            "brief": "WebSockets protocol",
                            "deprecated": none,
                            "id": "web_sockets",
                            "note": none,
                            "stability": "stable",
                            "value": "web_sockets",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "SignalR HTTP connection closure status.",
                    "examples": [
                        "app_shutdown",
                        "timeout",
                    ],
                    "name": "signalr.connection.status",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The connection was closed normally.",
                                "deprecated": none,
                                "id": "normal_closure",
                                "note": none,
                                "stability": "stable",
                                "value": "normal_closure",
                            },
                            {
                                "brief": "The connection was closed due to a timeout.",
                                "deprecated": none,
                                "id": "timeout",
                                "note": none,
                                "stability": "stable",
                                "value": "timeout",
                            },
                            {
                                "brief": "The connection was closed because the app is shutting down.",
                                "deprecated": none,
                                "id": "app_shutdown",
                                "note": none,
                                "stability": "stable",
                                "value": "app_shutdown",
                            },
                        ],
                    },
                },
                {
                    "brief": "[SignalR transport type](https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md)",
                    "examples": [
                        "web_sockets",
                        "long_polling",
                    ],
                    "name": "signalr.transport",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "ServerSentEvents protocol",
                                "deprecated": none,
                                "id": "server_sent_events",
                                "note": none,
                                "stability": "stable",
                                "value": "server_sent_events",
                            },
                            {
                                "brief": "LongPolling protocol",
                                "deprecated": none,
                                "id": "long_polling",
                                "note": none,
                                "stability": "stable",
                                "value": "long_polling",
                            },
                            {
                                "brief": "WebSockets protocol",
                                "deprecated": none,
                                "id": "web_sockets",
                                "note": none,
                                "stability": "stable",
                                "value": "web_sockets",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of connections that are currently active on the server.",
            "events": [],
            "id": "metric.signalr.server.active_connections",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "signalr.connection.status": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.signalr",
                    },
                    "signalr.transport": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.signalr",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/signalr/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "signalr.server.active_connections",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Http.Connections`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "signalr",
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
