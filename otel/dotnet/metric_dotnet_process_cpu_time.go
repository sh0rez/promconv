package dotnet

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/cpu"
)

// CPU time used by the process.
type ProcessCpuTime struct {
	*prometheus.CounterVec
	extra ProcessCpuTimeExtra
}

func NewProcessCpuTime() ProcessCpuTime {
	labels := []string{cpu.AttrMode("").Key()}
	return ProcessCpuTime{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dotnet_process_cpu_time",
		Help: "CPU time used by the process.",
	}, labels)}
}

func (m ProcessCpuTime) With(mode cpu.AttrMode, extras ...interface{}) prometheus.Counter {
	return m.CounterVec.WithLabelValues(mode.Value())
}

// Deprecated: Use [ProcessCpuTime.With] instead
func (m ProcessCpuTime) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

type ProcessCpuTimeExtra struct {
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ProcessCpuTimeExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "process.cpu.time",
        "Type": "ProcessCpuTime",
        "attributes": [
            {
                "brief": "The mode of the CPU",
                "examples": [
                    "user",
                    "system",
                ],
                "name": "cpu.mode",
                "requirement_level": "required",
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
                    "brief": "The mode of the CPU",
                    "examples": [
                        "user",
                        "system",
                    ],
                    "name": "cpu.mode",
                    "requirement_level": "required",
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
            "brief": "CPU time used by the process.",
            "events": [],
            "id": "metric.dotnet.process.cpu.time",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "cpu.mode": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.cpu",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/dotnet/runtime-metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "dotnet.process.cpu.time",
            "name": none,
            "note": "Meter name: `System.Runtime`; Added in: .NET 9.0.\nThis metric reports the same values as accessing the corresponding processor time properties on [`System.Diagnostics.Process`](https://learn.microsoft.com/dotnet/api/system.diagnostics.process).\n",
            "root_namespace": "dotnet",
            "span_kind": none,
            "stability": "stable",
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
