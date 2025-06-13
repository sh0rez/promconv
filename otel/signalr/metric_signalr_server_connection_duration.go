package signalr

import (
	"github.com/prometheus/client_golang/prometheus"
)

// The duration of connections on the server.
type ServerConnectionDuration struct {
	*prometheus.HistogramVec
	extra ServerConnectionDurationExtra
}

func NewServerConnectionDuration() ServerConnectionDuration {
	labels := []string{AttrConnectionStatus("").Key(), AttrTransport("").Key()}
	return ServerConnectionDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "signalr_server_connection_duration",
		Help: "The duration of connections on the server.",
	}, labels)}
}

func (m ServerConnectionDuration) With(extras ...interface {
	SignalrConnectionStatus() AttrConnectionStatus
	SignalrTransport() AttrTransport
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(extra.SignalrConnectionStatus().Value(), extra.SignalrTransport().Value())
}

// Deprecated: Use [ServerConnectionDuration.With] instead
func (m ServerConnectionDuration) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ServerConnectionDuration) WithConnectionStatus(attr interface{ SignalrConnectionStatus() AttrConnectionStatus }) ServerConnectionDuration {
	a.extra.AttrConnectionStatus = attr.SignalrConnectionStatus()
	return a
}
func (a ServerConnectionDuration) WithTransport(attr interface{ SignalrTransport() AttrTransport }) ServerConnectionDuration {
	a.extra.AttrTransport = attr.SignalrTransport()
	return a
}

type ServerConnectionDurationExtra struct {
	// SignalR HTTP connection closure status
	AttrConnectionStatus AttrConnectionStatus `otel:"signalr.connection.status"` // [SignalR transport type]
	//
	// [SignalR transport type]: https://github.com/dotnet/aspnetcore/blob/main/src/SignalR/docs/specs/TransportProtocols.md
	AttrTransport AttrTransport `otel:"signalr.transport"`
}

func (a ServerConnectionDurationExtra) SignalrConnectionStatus() AttrConnectionStatus {
	return a.AttrConnectionStatus
}
func (a ServerConnectionDurationExtra) SignalrTransport() AttrTransport { return a.AttrTransport }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerConnectionDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.connection.duration",
        "Type": "ServerConnectionDuration",
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
            "brief": "The duration of connections on the server.",
            "events": [],
            "id": "metric.signalr.server.connection.duration",
            "instrument": "histogram",
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
            "metric_name": "signalr.server.connection.duration",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Http.Connections`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "signalr",
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
