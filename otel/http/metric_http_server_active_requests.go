package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
)

// Number of active HTTP server requests.
type ServerActiveRequests struct {
	*prometheus.GaugeVec
	extra ServerActiveRequestsExtra
}

func NewServerActiveRequests() ServerActiveRequests {
	labels := []string{AttrRequestMethod("").Key(), url.AttrScheme("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key()}
	return ServerActiveRequests{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_server_active_requests",
		Help: "Number of active HTTP server requests.",
	}, labels)}
}

func (m ServerActiveRequests) With(requestMethod AttrRequestMethod, scheme url.AttrScheme, extras ...interface {
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(requestMethod.Value(), scheme.Value(), extra.ServerAddress().Value(), extra.ServerPort().Value())
}

// Deprecated: Use [ServerActiveRequests.With] instead
func (m ServerActiveRequests) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ServerActiveRequests) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) ServerActiveRequests {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a ServerActiveRequests) WithServerPort(attr interface{ ServerPort() server.AttrPort }) ServerActiveRequests {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}

type ServerActiveRequestsExtra struct {
	// Name of the local HTTP server that received the request
	AttrServerAddress server.AttrAddress `otel:"server.address"` // Port of the local HTTP server that received the request
	AttrServerPort    server.AttrPort    `otel:"server.port"`
}

func (a ServerActiveRequestsExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }
func (a ServerActiveRequestsExtra) ServerPort() server.AttrPort       { return a.AttrServerPort }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerActiveRequestsExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.active_requests",
        "Type": "ServerActiveRequests",
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
                "requirement_level": "required",
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
        ],
        "ctx": {
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
                    "requirement_level": "required",
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
            ],
            "brief": "Number of active HTTP server requests.",
            "events": [],
            "id": "metric.http.server.active_requests",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
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
            "metric_name": "http.server.active_requests",
            "name": none,
            "root_namespace": "http",
            "span_kind": none,
            "stability": "development",
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
