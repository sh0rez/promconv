package process

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/disk"
)

// Disk bytes transferred.
type DiskIo struct {
	*prometheus.CounterVec
	extra DiskIoExtra
}

func NewDiskIo() DiskIo {
	labels := []string{"disk_io_direction"}
	return DiskIo{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "process",
		Name:      "disk_io",
		Help:      "Disk bytes transferred.",
	}, labels)}
}

func (m DiskIo) With(extra interface {
	AttrDiskIoDirection() disk.AttrIoDirection
}) prometheus.Counter {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrDiskIoDirection()),
	)
}

func (a DiskIo) WithDiskIoDirection(attr interface{ AttrDiskIoDirection() disk.AttrIoDirection }) DiskIo {
	a.extra.DiskIoDirection = attr.AttrDiskIoDirection()
	return a
}

type DiskIoExtra struct {
	// The disk IO operation direction.
	DiskIoDirection disk.AttrIoDirection `otel:"disk.io.direction"`
}

func (a DiskIoExtra) AttrDiskIoDirection() disk.AttrIoDirection { return a.DiskIoDirection }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "DiskIoExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "disk.io",
        "Type": "DiskIo",
        "attributes": [
            {
                "brief": "The disk IO operation direction.",
                "examples": [
                    "read",
                ],
                "name": "disk.io.direction",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "read",
                            "note": none,
                            "stability": "development",
                            "value": "read",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "write",
                            "note": none,
                            "stability": "development",
                            "value": "write",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The disk IO operation direction.",
                    "examples": [
                        "read",
                    ],
                    "name": "disk.io.direction",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "read",
                                "note": none,
                                "stability": "development",
                                "value": "read",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "write",
                                "note": none,
                                "stability": "development",
                                "value": "write",
                            },
                        ],
                    },
                },
            ],
            "brief": "Disk bytes transferred.",
            "entity_associations": [
                "process",
            ],
            "events": [],
            "id": "metric.process.disk.io",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "disk.io.direction": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.disk",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/process/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "process.disk.io",
            "name": none,
            "root_namespace": "process",
            "span_kind": none,
            "stability": "development",
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
