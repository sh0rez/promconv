package process

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of times the process has been context switched.
type ContextSwitches struct {
	*prometheus.CounterVec
	extra ContextSwitchesExtra
}

func NewContextSwitches() ContextSwitches {
	labels := []string{AttrContextSwitchType("").Key()}
	return ContextSwitches{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "process_context_switches",
		Help: "Number of times the process has been context switched.",
	}, labels)}
}

func (m ContextSwitches) With(extras ...interface{ ProcessContextSwitchType() AttrContextSwitchType }) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(extra.ProcessContextSwitchType().Value())
}

// Deprecated: Use [ContextSwitches.With] instead
func (m ContextSwitches) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a ContextSwitches) WithContextSwitchType(attr interface{ ProcessContextSwitchType() AttrContextSwitchType }) ContextSwitches {
	a.extra.AttrContextSwitchType = attr.ProcessContextSwitchType()
	return a
}

type ContextSwitchesExtra struct {
	// Specifies whether the context switches for this data point were voluntary or involuntary
	AttrContextSwitchType AttrContextSwitchType `otel:"process.context_switch_type"`
}

func (a ContextSwitchesExtra) ProcessContextSwitchType() AttrContextSwitchType {
	return a.AttrContextSwitchType
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ContextSwitchesExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "context_switches",
        "Type": "ContextSwitches",
        "attributes": [
            {
                "brief": "Specifies whether the context switches for this data point were voluntary or involuntary.",
                "name": "process.context_switch_type",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "voluntary",
                            "note": none,
                            "stability": "development",
                            "value": "voluntary",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "involuntary",
                            "note": none,
                            "stability": "development",
                            "value": "involuntary",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Specifies whether the context switches for this data point were voluntary or involuntary.",
                    "name": "process.context_switch_type",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "voluntary",
                                "note": none,
                                "stability": "development",
                                "value": "voluntary",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "involuntary",
                                "note": none,
                                "stability": "development",
                                "value": "involuntary",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of times the process has been context switched.",
            "entity_associations": [
                "process",
            ],
            "events": [],
            "id": "metric.process.context_switches",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "process.context_switch_type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.process",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/process/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "process.context_switches",
            "name": none,
            "root_namespace": "process",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{context_switch}",
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
