package otel

import (
	"github.com/prometheus/client_golang/prometheus"
)

// The number of created spans for which the end operation was called
type SdkSpanEnded struct {
	*prometheus.CounterVec
	extra SdkSpanEndedExtra
}

func NewSdkSpanEnded() SdkSpanEnded {
	labels := []string{"otel_span_sampling_result"}
	return SdkSpanEnded{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "otel",
		Name:      "sdk_span_ended",
		Help:      "The number of created spans for which the end operation was called",
	}, labels)}
}

func (m SdkSpanEnded) With(extra interface {
	AttrOtelSpanSamplingResult() AttrSpanSamplingResult
}) prometheus.Counter {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrOtelSpanSamplingResult()),
	)
}

func (a SdkSpanEnded) WithOtelSpanSamplingResult(attr interface{ AttrOtelSpanSamplingResult() AttrSpanSamplingResult }) SdkSpanEnded {
	a.extra.OtelSpanSamplingResult = attr.AttrOtelSpanSamplingResult()
	return a
}

type SdkSpanEndedExtra struct {
	// The result value of the sampler for this span
	OtelSpanSamplingResult AttrSpanSamplingResult `otel:"otel.span.sampling_result"`
}

func (a SdkSpanEndedExtra) AttrOtelSpanSamplingResult() AttrSpanSamplingResult {
	return a.OtelSpanSamplingResult
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "SdkSpanEndedExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "sdk.span.ended",
        "Type": "SdkSpanEnded",
        "attributes": [
            {
                "brief": "The result value of the sampler for this span",
                "name": "otel.span.sampling_result",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "The span is not sampled and not recording",
                            "deprecated": none,
                            "id": "drop",
                            "note": none,
                            "stability": "development",
                            "value": "DROP",
                        },
                        {
                            "brief": "The span is not sampled, but recording",
                            "deprecated": none,
                            "id": "record_only",
                            "note": none,
                            "stability": "development",
                            "value": "RECORD_ONLY",
                        },
                        {
                            "brief": "The span is sampled and recording",
                            "deprecated": none,
                            "id": "record_and_sample",
                            "note": none,
                            "stability": "development",
                            "value": "RECORD_AND_SAMPLE",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The result value of the sampler for this span",
                    "name": "otel.span.sampling_result",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The span is not sampled and not recording",
                                "deprecated": none,
                                "id": "drop",
                                "note": none,
                                "stability": "development",
                                "value": "DROP",
                            },
                            {
                                "brief": "The span is not sampled, but recording",
                                "deprecated": none,
                                "id": "record_only",
                                "note": none,
                                "stability": "development",
                                "value": "RECORD_ONLY",
                            },
                            {
                                "brief": "The span is sampled and recording",
                                "deprecated": none,
                                "id": "record_and_sample",
                                "note": none,
                                "stability": "development",
                                "value": "RECORD_AND_SAMPLE",
                            },
                        ],
                    },
                },
            ],
            "brief": "The number of created spans for which the end operation was called",
            "events": [],
            "id": "metric.otel.sdk.span.ended",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "otel.span.sampling_result": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.otel",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/otel/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "otel.sdk.span.ended",
            "name": none,
            "note": "For spans with `recording=true`: Implementations MUST record both `otel.sdk.span.live` and `otel.sdk.span.ended`.\nFor spans with `recording=false`: If implementations decide to record this metric, they MUST also record `otel.sdk.span.live`.\n",
            "root_namespace": "otel",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{span}",
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
