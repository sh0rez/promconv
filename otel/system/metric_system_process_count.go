package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Total number of processes in each state
type ProcessCount struct {
	*prometheus.GaugeVec
	extra ProcessCountExtra
}

func NewProcessCount() ProcessCount {
	labels := []string{AttrProcessStatus("").Key()}
	return ProcessCount{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "system_process_count",
		Help: "Total number of processes in each state",
	}, labels)}
}

func (m ProcessCount) With(extras ...interface{ SystemProcessStatus() AttrProcessStatus }) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.SystemProcessStatus().Value())
}

// Deprecated: Use [ProcessCount.With] instead
func (m ProcessCount) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ProcessCount) WithProcessStatus(attr interface{ SystemProcessStatus() AttrProcessStatus }) ProcessCount {
	a.extra.AttrProcessStatus = attr.SystemProcessStatus()
	return a
}

type ProcessCountExtra struct {
	// The process state, e.g., [Linux Process State Codes]
	//
	// [Linux Process State Codes]: https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES
	AttrProcessStatus AttrProcessStatus `otel:"system.process.status"`
}

func (a ProcessCountExtra) SystemProcessStatus() AttrProcessStatus { return a.AttrProcessStatus }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ProcessCountExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "process.count",
        "Type": "ProcessCount",
        "attributes": [
            {
                "brief": "The process state, e.g., [Linux Process State Codes](https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES)\n",
                "examples": [
                    "running",
                ],
                "name": "system.process.status",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "running",
                            "note": none,
                            "stability": "development",
                            "value": "running",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "sleeping",
                            "note": none,
                            "stability": "development",
                            "value": "sleeping",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "stopped",
                            "note": none,
                            "stability": "development",
                            "value": "stopped",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "defunct",
                            "note": none,
                            "stability": "development",
                            "value": "defunct",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The process state, e.g., [Linux Process State Codes](https://man7.org/linux/man-pages/man1/ps.1.html#PROCESS_STATE_CODES)\n",
                    "examples": [
                        "running",
                    ],
                    "name": "system.process.status",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "running",
                                "note": none,
                                "stability": "development",
                                "value": "running",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "sleeping",
                                "note": none,
                                "stability": "development",
                                "value": "sleeping",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "stopped",
                                "note": none,
                                "stability": "development",
                                "value": "stopped",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "defunct",
                                "note": none,
                                "stability": "development",
                                "value": "defunct",
                            },
                        ],
                    },
                },
            ],
            "brief": "Total number of processes in each state",
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.process.count",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "system.process.status": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.process",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.process.count",
            "name": none,
            "root_namespace": "system",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{process}",
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
