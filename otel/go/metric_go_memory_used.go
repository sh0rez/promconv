package golang

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Memory used by the Go runtime.
type MemoryUsed struct {
	*prometheus.GaugeVec
	extra MemoryUsedExtra
}

func NewMemoryUsed() MemoryUsed {
	labels := []string{AttrMemoryType("").Key()}
	return MemoryUsed{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "go_memory_used",
		Help: "Memory used by the Go runtime.",
	}, labels)}
}

func (m MemoryUsed) With(extras ...interface{ GoMemoryType() AttrMemoryType }) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.GoMemoryType().Value())
}

// Deprecated: Use [MemoryUsed.With] instead
func (m MemoryUsed) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a MemoryUsed) WithMemoryType(attr interface{ GoMemoryType() AttrMemoryType }) MemoryUsed {
	a.extra.AttrMemoryType = attr.GoMemoryType()
	return a
}

type MemoryUsedExtra struct {
	// The type of memory
	AttrMemoryType AttrMemoryType `otel:"go.memory.type"`
}

func (a MemoryUsedExtra) GoMemoryType() AttrMemoryType { return a.AttrMemoryType }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "MemoryUsedExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "memory.used",
        "Type": "MemoryUsed",
        "attributes": [
            {
                "brief": "The type of memory.",
                "examples": [
                    "other",
                    "stack",
                ],
                "name": "go.memory.type",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Memory allocated from the heap that is reserved for stack space, whether or not it is currently in-use.",
                            "deprecated": none,
                            "id": "stack",
                            "note": "Computed from `/memory/classes/heap/stacks:bytes`.\n",
                            "stability": "development",
                            "value": "stack",
                        },
                        {
                            "brief": "Memory used by the Go runtime, excluding other categories of memory usage described in this enumeration.",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "development",
                            "value": "other",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The type of memory.",
                    "examples": [
                        "other",
                        "stack",
                    ],
                    "name": "go.memory.type",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Memory allocated from the heap that is reserved for stack space, whether or not it is currently in-use.",
                                "deprecated": none,
                                "id": "stack",
                                "note": "Computed from `/memory/classes/heap/stacks:bytes`.\n",
                                "stability": "development",
                                "value": "stack",
                            },
                            {
                                "brief": "Memory used by the Go runtime, excluding other categories of memory usage described in this enumeration.",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "development",
                                "value": "other",
                            },
                        ],
                    },
                },
            ],
            "brief": "Memory used by the Go runtime.",
            "events": [],
            "id": "metric.go.memory.used",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "go.memory.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.go",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/go/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "go.memory.used",
            "name": none,
            "note": "Computed from `(/memory/classes/total:bytes - /memory/classes/heap/released:bytes)`.\n",
            "root_namespace": "go",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "By",
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
