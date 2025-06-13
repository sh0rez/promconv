package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MemoryUtilization struct {
	*prometheus.GaugeVec
	extra MemoryUtilizationExtra
}

func NewMemoryUtilization() MemoryUtilization {
	labels := []string{AttrMemoryState("").Key()}
	return MemoryUtilization{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "system_memory_utilization",
		Help: "",
	}, labels)}
}

func (m MemoryUtilization) With(extras ...interface{ SystemMemoryState() AttrMemoryState }) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.SystemMemoryState().Value())
}

// Deprecated: Use [MemoryUtilization.With] instead
func (m MemoryUtilization) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a MemoryUtilization) WithMemoryState(attr interface{ SystemMemoryState() AttrMemoryState }) MemoryUtilization {
	a.extra.AttrMemoryState = attr.SystemMemoryState()
	return a
}

type MemoryUtilizationExtra struct {
	// The memory state
	AttrMemoryState AttrMemoryState `otel:"system.memory.state"`
}

func (a MemoryUtilizationExtra) SystemMemoryState() AttrMemoryState { return a.AttrMemoryState }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "MemoryUtilizationExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "memory.utilization",
        "Type": "MemoryUtilization",
        "attributes": [
            {
                "brief": "The memory state",
                "examples": [
                    "free",
                    "cached",
                ],
                "name": "system.memory.state",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "used",
                            "note": none,
                            "stability": "development",
                            "value": "used",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "free",
                            "note": none,
                            "stability": "development",
                            "value": "free",
                        },
                        {
                            "brief": none,
                            "deprecated": "Removed, report shared memory usage with `metric.system.memory.shared` metric",
                            "id": "shared",
                            "note": none,
                            "stability": "development",
                            "value": "shared",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "buffers",
                            "note": none,
                            "stability": "development",
                            "value": "buffers",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "cached",
                            "note": none,
                            "stability": "development",
                            "value": "cached",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The memory state",
                    "examples": [
                        "free",
                        "cached",
                    ],
                    "name": "system.memory.state",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "used",
                                "note": none,
                                "stability": "development",
                                "value": "used",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "free",
                                "note": none,
                                "stability": "development",
                                "value": "free",
                            },
                            {
                                "brief": none,
                                "deprecated": "Removed, report shared memory usage with `metric.system.memory.shared` metric",
                                "id": "shared",
                                "note": none,
                                "stability": "development",
                                "value": "shared",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "buffers",
                                "note": none,
                                "stability": "development",
                                "value": "buffers",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "cached",
                                "note": none,
                                "stability": "development",
                                "value": "cached",
                            },
                        ],
                    },
                },
            ],
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.memory.utilization",
            "instrument": "gauge",
            "lineage": {
                "attributes": {
                    "system.memory.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.memory",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.memory.utilization",
            "name": none,
            "root_namespace": "system",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "1",
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
