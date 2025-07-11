package cpython

import (
	"github.com/prometheus/client_golang/prometheus"
)

// The total number of objects collected inside a generation since interpreter start.
type GcCollectedObjects struct {
	*prometheus.CounterVec
	extra GcCollectedObjectsExtra
}

func NewGcCollectedObjects() GcCollectedObjects {
	labels := []string{AttrGcGeneration("").Key()}
	return GcCollectedObjects{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cpython_gc_collected_objects",
		Help: "The total number of objects collected inside a generation since interpreter start.",
	}, labels)}
}

func (m GcCollectedObjects) With(gcGeneration AttrGcGeneration, extras ...interface{}) prometheus.Counter {
	return m.CounterVec.WithLabelValues(gcGeneration.Value())
}

// Deprecated: Use [GcCollectedObjects.With] instead
func (m GcCollectedObjects) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

type GcCollectedObjectsExtra struct {
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "GcCollectedObjectsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "gc.collected_objects",
        "Type": "GcCollectedObjects",
        "attributes": [
            {
                "brief": "Value of the garbage collector collection generation.",
                "examples": [
                    0,
                    1,
                    2,
                ],
                "name": "cpython.gc.generation",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Generation 0",
                            "deprecated": none,
                            "id": "generation_0",
                            "note": none,
                            "stability": "development",
                            "value": 0,
                        },
                        {
                            "brief": "Generation 1",
                            "deprecated": none,
                            "id": "generation_1",
                            "note": none,
                            "stability": "development",
                            "value": 1,
                        },
                        {
                            "brief": "Generation 2",
                            "deprecated": none,
                            "id": "generation_2",
                            "note": none,
                            "stability": "development",
                            "value": 2,
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Value of the garbage collector collection generation.",
                    "examples": [
                        0,
                        1,
                        2,
                    ],
                    "name": "cpython.gc.generation",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Generation 0",
                                "deprecated": none,
                                "id": "generation_0",
                                "note": none,
                                "stability": "development",
                                "value": 0,
                            },
                            {
                                "brief": "Generation 1",
                                "deprecated": none,
                                "id": "generation_1",
                                "note": none,
                                "stability": "development",
                                "value": 1,
                            },
                            {
                                "brief": "Generation 2",
                                "deprecated": none,
                                "id": "generation_2",
                                "note": none,
                                "stability": "development",
                                "value": 2,
                            },
                        ],
                    },
                },
            ],
            "brief": "The total number of objects collected inside a generation since interpreter start.",
            "events": [],
            "id": "metric.cpython.gc.collected_objects",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "cpython.gc.generation": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.cpython",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/cpython/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "cpython.gc.collected_objects",
            "name": none,
            "note": "This metric reports data from [`gc.stats()`](https://docs.python.org/3/library/gc.html#gc.get_stats).\n",
            "root_namespace": "cpython",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{object}",
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
