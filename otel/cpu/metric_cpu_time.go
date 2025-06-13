package cpu

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Deprecated. Use `system.cpu.time` instead.
type Time struct {
	*prometheus.CounterVec
	extra TimeExtra
}

func NewTime() Time {
	labels := []string{AttrLogicalNumber("").Key(), AttrMode("").Key()}
	return Time{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cpu_time",
		Help: "Deprecated. Use `system.cpu.time` instead.",
	}, labels)}
}

func (m Time) With(extras ...interface {
	CpuLogicalNumber() AttrLogicalNumber
	CpuMode() AttrMode
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(extra.CpuLogicalNumber().Value(), extra.CpuMode().Value())
}

// Deprecated: Use [Time.With] instead
func (m Time) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a Time) WithLogicalNumber(attr interface{ CpuLogicalNumber() AttrLogicalNumber }) Time {
	a.extra.AttrLogicalNumber = attr.CpuLogicalNumber()
	return a
}
func (a Time) WithMode(attr interface{ CpuMode() AttrMode }) Time {
	a.extra.AttrMode = attr.CpuMode()
	return a
}

type TimeExtra struct {
	// The logical CPU number [0..n-1]
	AttrLogicalNumber AttrLogicalNumber `otel:"cpu.logical_number"` // The mode of the CPU
	AttrMode          AttrMode          `otel:"cpu.mode"`
}

func (a TimeExtra) CpuLogicalNumber() AttrLogicalNumber { return a.AttrLogicalNumber }
func (a TimeExtra) CpuMode() AttrMode                   { return a.AttrMode }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "TimeExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "time",
        "Type": "Time",
        "attributes": [
            {
                "brief": "The logical CPU number [0..n-1]",
                "examples": [
                    1,
                ],
                "name": "cpu.logical_number",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "int",
            },
            {
                "brief": "The mode of the CPU",
                "examples": [
                    "user",
                    "system",
                ],
                "name": "cpu.mode",
                "note": "Following states SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "user",
                            "note": none,
                            "stability": "development",
                            "value": "user",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "system",
                            "note": none,
                            "stability": "development",
                            "value": "system",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "nice",
                            "note": none,
                            "stability": "development",
                            "value": "nice",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "idle",
                            "note": none,
                            "stability": "development",
                            "value": "idle",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "iowait",
                            "note": none,
                            "stability": "development",
                            "value": "iowait",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "interrupt",
                            "note": none,
                            "stability": "development",
                            "value": "interrupt",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "steal",
                            "note": none,
                            "stability": "development",
                            "value": "steal",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "kernel",
                            "note": none,
                            "stability": "development",
                            "value": "kernel",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The logical CPU number [0..n-1]",
                    "examples": [
                        1,
                    ],
                    "name": "cpu.logical_number",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The mode of the CPU",
                    "examples": [
                        "user",
                        "system",
                    ],
                    "name": "cpu.mode",
                    "note": "Following states SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "user",
                                "note": none,
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "system",
                                "note": none,
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "nice",
                                "note": none,
                                "stability": "development",
                                "value": "nice",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "iowait",
                                "note": none,
                                "stability": "development",
                                "value": "iowait",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "interrupt",
                                "note": none,
                                "stability": "development",
                                "value": "interrupt",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "steal",
                                "note": none,
                                "stability": "development",
                                "value": "steal",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "kernel",
                                "note": none,
                                "stability": "development",
                                "value": "kernel",
                            },
                        ],
                    },
                },
            ],
            "brief": "Deprecated. Use `system.cpu.time` instead.",
            "deprecated": {
                "note": "Replaced by `system.cpu.time`.",
                "reason": "renamed",
                "renamed_to": "system.cpu.time",
            },
            "events": [],
            "id": "metric.cpu.time",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "cpu.logical_number": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.cpu",
                    },
                    "cpu.mode": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                        ],
                        "source_group": "registry.cpu",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/cpu/deprecated.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "cpu.time",
            "name": none,
            "root_namespace": "cpu",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "s",
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
