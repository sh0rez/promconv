package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/cpu"
)

// For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time.
type CpuUtilization struct {
	*prometheus.GaugeVec
	extra CpuUtilizationExtra
}

func NewCpuUtilization() CpuUtilization {
	labels := []string{cpu.AttrLogicalNumber("").Key(), cpu.AttrMode("").Key()}
	return CpuUtilization{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "system_cpu_utilization",
		Help: "For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time.",
	}, labels)}
}

func (m CpuUtilization) With(extras ...interface {
	CpuLogicalNumber() cpu.AttrLogicalNumber
	CpuMode() cpu.AttrMode
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.CpuLogicalNumber().Value(), extra.CpuMode().Value())
}

// Deprecated: Use [CpuUtilization.With] instead
func (m CpuUtilization) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a CpuUtilization) WithCpuLogicalNumber(attr interface{ CpuLogicalNumber() cpu.AttrLogicalNumber }) CpuUtilization {
	a.extra.AttrCpuLogicalNumber = attr.CpuLogicalNumber()
	return a
}
func (a CpuUtilization) WithCpuMode(attr interface{ CpuMode() cpu.AttrMode }) CpuUtilization {
	a.extra.AttrCpuMode = attr.CpuMode()
	return a
}

type CpuUtilizationExtra struct {
	// The logical CPU number [0..n-1]
	AttrCpuLogicalNumber cpu.AttrLogicalNumber `otel:"cpu.logical_number"` // The mode of the CPU
	AttrCpuMode          cpu.AttrMode          `otel:"cpu.mode"`
}

func (a CpuUtilizationExtra) CpuLogicalNumber() cpu.AttrLogicalNumber { return a.AttrCpuLogicalNumber }
func (a CpuUtilizationExtra) CpuMode() cpu.AttrMode                   { return a.AttrCpuMode }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "CpuUtilizationExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "cpu.utilization",
        "Type": "CpuUtilization",
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
                "note": "Following modes SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
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
                    "note": "Following modes SHOULD be used: `user`, `system`, `nice`, `idle`, `iowait`, `interrupt`, `steal`",
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
            "brief": "For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time.",
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.cpu.utilization",
            "instrument": "gauge",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.cpu.utilization",
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
