package aspnetcore

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/http"
)

// Number of requests that were attempted to be matched to an endpoint.
type RoutingMatchAttempts struct {
	*prometheus.CounterVec
}

func NewRoutingMatchAttempts() RoutingMatchAttempts {
	labels := []string{"aspnetcore_routing_match_status", "aspnetcore_routing_is_fallback", "http_route"}
	return RoutingMatchAttempts{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "aspnetcore",
		Name:      "routing_match_attempts",
		Help:      "Number of requests that were attempted to be matched to an endpoint.",
	}, labels)}
}

func (m RoutingMatchAttempts) With(routingMatchStatus AttrRoutingMatchStatus, extra interface {
	AttrAspnetcoreRoutingIsFallback() AttrRoutingIsFallback
	AttrHttpRoute() http.AttrRoute
}) prometheus.Counter {
	if extra == nil {
		extra = RoutingMatchAttemptsExtra{}
	}
	return m.WithLabelValues(
		string(routingMatchStatus),
		string(extra.AttrAspnetcoreRoutingIsFallback()),
		string(extra.AttrHttpRoute()),
	)
}

type RoutingMatchAttemptsExtra struct {
	// A value that indicates whether the matched route is a fallback route.
	AspnetcoreRoutingIsFallback AttrRoutingIsFallback `otel:"aspnetcore.routing.is_fallback"`
	// The matched route, that is, the path template in the format used by the respective server framework.
	HttpRoute http.AttrRoute `otel:"http.route"`
}

func (a RoutingMatchAttemptsExtra) AttrAspnetcoreRoutingIsFallback() AttrRoutingIsFallback {
	return a.AspnetcoreRoutingIsFallback
}
func (a RoutingMatchAttemptsExtra) AttrHttpRoute() http.AttrRoute { return a.HttpRoute }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "RoutingMatchAttemptsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "routing.match_attempts",
        "Type": "RoutingMatchAttempts",
        "attributes": [
            {
                "brief": "Match result - success or failure",
                "examples": [
                    "success",
                    "failure",
                ],
                "name": "aspnetcore.routing.match_status",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "Match succeeded",
                            "deprecated": none,
                            "id": "success",
                            "note": none,
                            "stability": "stable",
                            "value": "success",
                        },
                        {
                            "brief": "Match failed",
                            "deprecated": none,
                            "id": "failure",
                            "note": none,
                            "stability": "stable",
                            "value": "failure",
                        },
                    ],
                },
            },
            {
                "brief": "A value that indicates whether the matched route is a fallback route.",
                "examples": [
                    true,
                ],
                "name": "aspnetcore.routing.is_fallback",
                "requirement_level": {
                    "conditionally_required": "if and only if a route was successfully matched.",
                },
                "stability": "stable",
                "type": "boolean",
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
                    "conditionally_required": "if and only if a route was successfully matched.",
                },
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                    "examples": [
                        "/users/:userID?",
                        "{controller}/{action}/{id?}",
                    ],
                    "name": "http.route",
                    "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                    "requirement_level": {
                        "conditionally_required": "if and only if a route was successfully matched.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "A value that indicates whether the matched route is a fallback route.",
                    "examples": [
                        true,
                    ],
                    "name": "aspnetcore.routing.is_fallback",
                    "requirement_level": {
                        "conditionally_required": "if and only if a route was successfully matched.",
                    },
                    "stability": "stable",
                    "type": "boolean",
                },
                {
                    "brief": "Match result - success or failure",
                    "examples": [
                        "success",
                        "failure",
                    ],
                    "name": "aspnetcore.routing.match_status",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Match succeeded",
                                "deprecated": none,
                                "id": "success",
                                "note": none,
                                "stability": "stable",
                                "value": "success",
                            },
                            {
                                "brief": "Match failed",
                                "deprecated": none,
                                "id": "failure",
                                "note": none,
                                "stability": "stable",
                                "value": "failure",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of requests that were attempted to be matched to an endpoint.",
            "events": [],
            "id": "metric.aspnetcore.routing.match_attempts",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "aspnetcore.routing.is_fallback": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.aspnetcore",
                    },
                    "aspnetcore.routing.match_status": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.aspnetcore",
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/aspnetcore/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "aspnetcore.routing.match_attempts",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Routing`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "aspnetcore",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "{match_attempt}",
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
