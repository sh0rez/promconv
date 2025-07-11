package aspnetcore

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of requests that tried to acquire a rate limiting lease.
type RateLimitingRequests struct {
	*prometheus.CounterVec
	extra RateLimitingRequestsExtra
}

func NewRateLimitingRequests() RateLimitingRequests {
	labels := []string{AttrRateLimitingResult("").Key(), AttrRateLimitingPolicy("").Key()}
	return RateLimitingRequests{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "aspnetcore_rate_limiting_requests",
		Help: "Number of requests that tried to acquire a rate limiting lease.",
	}, labels)}
}

func (m RateLimitingRequests) With(rateLimitingResult AttrRateLimitingResult, extras ...interface{ AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy }) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(rateLimitingResult.Value(), extra.AspnetcoreRateLimitingPolicy().Value())
}

// Deprecated: Use [RateLimitingRequests.With] instead
func (m RateLimitingRequests) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a RateLimitingRequests) WithRateLimitingPolicy(attr interface{ AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy }) RateLimitingRequests {
	a.extra.AttrRateLimitingPolicy = attr.AspnetcoreRateLimitingPolicy()
	return a
}

type RateLimitingRequestsExtra struct {
	// Rate limiting policy name
	AttrRateLimitingPolicy AttrRateLimitingPolicy `otel:"aspnetcore.rate_limiting.policy"`
}

func (a RateLimitingRequestsExtra) AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy {
	return a.AttrRateLimitingPolicy
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "RateLimitingRequestsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "rate_limiting.requests",
        "Type": "RateLimitingRequests",
        "attributes": [
            {
                "brief": "Rate-limiting result, shows whether the lease was acquired or contains a rejection reason",
                "examples": [
                    "acquired",
                    "request_canceled",
                ],
                "name": "aspnetcore.rate_limiting.result",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "Lease was acquired",
                            "deprecated": none,
                            "id": "acquired",
                            "note": none,
                            "stability": "stable",
                            "value": "acquired",
                        },
                        {
                            "brief": "Lease request was rejected by the endpoint limiter",
                            "deprecated": none,
                            "id": "endpoint_limiter",
                            "note": none,
                            "stability": "stable",
                            "value": "endpoint_limiter",
                        },
                        {
                            "brief": "Lease request was rejected by the global limiter",
                            "deprecated": none,
                            "id": "global_limiter",
                            "note": none,
                            "stability": "stable",
                            "value": "global_limiter",
                        },
                        {
                            "brief": "Lease request was canceled",
                            "deprecated": none,
                            "id": "request_canceled",
                            "note": none,
                            "stability": "stable",
                            "value": "request_canceled",
                        },
                    ],
                },
            },
            {
                "brief": "Rate limiting policy name.",
                "examples": [
                    "fixed",
                    "sliding",
                    "token",
                ],
                "name": "aspnetcore.rate_limiting.policy",
                "requirement_level": {
                    "conditionally_required": "if the matched endpoint for the request had a rate-limiting policy.",
                },
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Rate limiting policy name.",
                    "examples": [
                        "fixed",
                        "sliding",
                        "token",
                    ],
                    "name": "aspnetcore.rate_limiting.policy",
                    "requirement_level": {
                        "conditionally_required": "if the matched endpoint for the request had a rate-limiting policy.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Rate-limiting result, shows whether the lease was acquired or contains a rejection reason",
                    "examples": [
                        "acquired",
                        "request_canceled",
                    ],
                    "name": "aspnetcore.rate_limiting.result",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "Lease was acquired",
                                "deprecated": none,
                                "id": "acquired",
                                "note": none,
                                "stability": "stable",
                                "value": "acquired",
                            },
                            {
                                "brief": "Lease request was rejected by the endpoint limiter",
                                "deprecated": none,
                                "id": "endpoint_limiter",
                                "note": none,
                                "stability": "stable",
                                "value": "endpoint_limiter",
                            },
                            {
                                "brief": "Lease request was rejected by the global limiter",
                                "deprecated": none,
                                "id": "global_limiter",
                                "note": none,
                                "stability": "stable",
                                "value": "global_limiter",
                            },
                            {
                                "brief": "Lease request was canceled",
                                "deprecated": none,
                                "id": "request_canceled",
                                "note": none,
                                "stability": "stable",
                                "value": "request_canceled",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of requests that tried to acquire a rate limiting lease.",
            "events": [],
            "id": "metric.aspnetcore.rate_limiting.requests",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "aspnetcore.rate_limiting.policy": {
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
                    "aspnetcore.rate_limiting.result": {
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/aspnetcore/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "aspnetcore.rate_limiting.requests",
            "name": none,
            "note": "Requests could be:\n\n* Rejected by global or endpoint rate limiting policies\n* Canceled while waiting for the lease.\n\nMeter name: `Microsoft.AspNetCore.RateLimiting`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "aspnetcore",
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
