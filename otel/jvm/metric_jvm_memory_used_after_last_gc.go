package jvm

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Measure of memory used, as measured after the most recent garbage collection event on this pool.
type MemoryUsedAfterLastGc struct {
	*prometheus.GaugeVec
}

func NewMemoryUsedAfterLastGc() MemoryUsedAfterLastGc {
	labels := []string{"jvm_memory_pool_name", "jvm_memory_type"}
	return MemoryUsedAfterLastGc{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "jvm",
		Name:      "memory_used_after_last_gc",
		Help:      "Measure of memory used, as measured after the most recent garbage collection event on this pool.",
	}, labels)}
}

func (m MemoryUsedAfterLastGc) With(extra interface {
	AttrJvmMemoryPoolName() AttrMemoryPoolName
	AttrJvmMemoryType() AttrMemoryType
}) prometheus.Gauge {
	if extra == nil {
		extra = MemoryUsedAfterLastGcExtra{}
	}
	return m.WithLabelValues(
		string(extra.AttrJvmMemoryPoolName()),
		string(extra.AttrJvmMemoryType()),
	)
}

type MemoryUsedAfterLastGcExtra struct {
	// Name of the memory pool.
	JvmMemoryPoolName AttrMemoryPoolName `otel:"jvm.memory.pool.name"`
	// The type of memory.
	JvmMemoryType AttrMemoryType `otel:"jvm.memory.type"`
}

func (a MemoryUsedAfterLastGcExtra) AttrJvmMemoryPoolName() AttrMemoryPoolName {
	return a.JvmMemoryPoolName
}
func (a MemoryUsedAfterLastGcExtra) AttrJvmMemoryType() AttrMemoryType { return a.JvmMemoryType }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "MemoryUsedAfterLastGcExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "memory.used_after_last_gc",
        "Type": "MemoryUsedAfterLastGc",
        "attributes": [
            {
                "brief": "Name of the memory pool.",
                "examples": [
                    "G1 Old Gen",
                    "G1 Eden space",
                    "G1 Survivor Space",
                ],
                "name": "jvm.memory.pool.name",
                "note": "Pool names are generally obtained via [MemoryPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/MemoryPoolMXBean.html#getName()).\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The type of memory.",
                "examples": [
                    "heap",
                    "non_heap",
                ],
                "name": "jvm.memory.type",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "Heap memory.",
                            "deprecated": none,
                            "id": "heap",
                            "note": none,
                            "stability": "stable",
                            "value": "heap",
                        },
                        {
                            "brief": "Non-heap memory",
                            "deprecated": none,
                            "id": "non_heap",
                            "note": none,
                            "stability": "stable",
                            "value": "non_heap",
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
                        "heap",
                        "non_heap",
                    ],
                    "name": "jvm.memory.type",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Heap memory.",
                                "deprecated": none,
                                "id": "heap",
                                "note": none,
                                "stability": "stable",
                                "value": "heap",
                            },
                            {
                                "brief": "Non-heap memory",
                                "deprecated": none,
                                "id": "non_heap",
                                "note": none,
                                "stability": "stable",
                                "value": "non_heap",
                            },
                        ],
                    },
                },
                {
                    "brief": "Name of the memory pool.",
                    "examples": [
                        "G1 Old Gen",
                        "G1 Eden space",
                        "G1 Survivor Space",
                    ],
                    "name": "jvm.memory.pool.name",
                    "note": "Pool names are generally obtained via [MemoryPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/MemoryPoolMXBean.html#getName()).\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "brief": "Measure of memory used, as measured after the most recent garbage collection event on this pool.",
            "events": [],
            "id": "metric.jvm.memory.used_after_last_gc",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "jvm.memory.pool.name": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.jvm",
                    },
                    "jvm.memory.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.jvm",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/jvm/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "jvm.memory.used_after_last_gc",
            "name": none,
            "root_namespace": "jvm",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "By",
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
