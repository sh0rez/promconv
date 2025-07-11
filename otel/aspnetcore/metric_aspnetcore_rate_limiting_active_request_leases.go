package aspnetcore

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of requests that are currently active on the server that hold a rate limiting lease.
type RateLimitingActiveRequestLeases struct {
	*prometheus.GaugeVec
	extra RateLimitingActiveRequestLeasesExtra
}

func NewRateLimitingActiveRequestLeases() RateLimitingActiveRequestLeases {
	labels := []string{AttrRateLimitingPolicy("").Key()}
	return RateLimitingActiveRequestLeases{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aspnetcore_rate_limiting_active_request_leases",
		Help: "Number of requests that are currently active on the server that hold a rate limiting lease.",
	}, labels)}
}

func (m RateLimitingActiveRequestLeases) With(extras ...interface{ AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy }) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.AspnetcoreRateLimitingPolicy().Value())
}

// Deprecated: Use [RateLimitingActiveRequestLeases.With] instead
func (m RateLimitingActiveRequestLeases) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a RateLimitingActiveRequestLeases) WithRateLimitingPolicy(attr interface{ AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy }) RateLimitingActiveRequestLeases {
	a.extra.AttrRateLimitingPolicy = attr.AspnetcoreRateLimitingPolicy()
	return a
}

type RateLimitingActiveRequestLeasesExtra struct {
	// Rate limiting policy name
	AttrRateLimitingPolicy AttrRateLimitingPolicy `otel:"aspnetcore.rate_limiting.policy"`
}

func (a RateLimitingActiveRequestLeasesExtra) AspnetcoreRateLimitingPolicy() AttrRateLimitingPolicy {
	return a.AttrRateLimitingPolicy
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "RateLimitingActiveRequestLeasesExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "rate_limiting.active_request_leases",
        "Type": "RateLimitingActiveRequestLeases",
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
            ],
            "brief": "Number of requests that are currently active on the server that hold a rate limiting lease.",
            "events": [],
            "id": "metric.aspnetcore.rate_limiting.active_request_leases",
            "instrument": "updowncounter",
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/aspnetcore/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "aspnetcore.rate_limiting.active_request_leases",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.RateLimiting`; Added in: ASP.NET Core 8.0\n",
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
