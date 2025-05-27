package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Unix swap or windows pagefile usage
type PagingUsage struct {
	*prometheus.GaugeVec
	extra PagingUsageExtra
}

func NewPagingUsage() PagingUsage {
	labels := []string{"system_device", "system_paging_state"}
	return PagingUsage{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "system",
		Name:      "paging_usage",
		Help:      "Unix swap or windows pagefile usage",
	}, labels)}
}

func (m PagingUsage) With(extra interface {
	AttrSystemDevice() AttrDevice
	AttrSystemPagingState() AttrPagingState
}) prometheus.Gauge {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrSystemDevice()),
		string(extra.AttrSystemPagingState()),
	)
}

func (a PagingUsage) WithSystemDevice(attr interface{ AttrSystemDevice() AttrDevice }) PagingUsage {
	a.extra.SystemDevice = attr.AttrSystemDevice()
	return a
}
func (a PagingUsage) WithSystemPagingState(attr interface{ AttrSystemPagingState() AttrPagingState }) PagingUsage {
	a.extra.SystemPagingState = attr.AttrSystemPagingState()
	return a
}

type PagingUsageExtra struct {
	// Unique identifier for the device responsible for managing paging operations.
	SystemDevice AttrDevice `otel:"system.device"`
	// The memory paging state
	SystemPagingState AttrPagingState `otel:"system.paging.state"`
}

func (a PagingUsageExtra) AttrSystemDevice() AttrDevice           { return a.SystemDevice }
func (a PagingUsageExtra) AttrSystemPagingState() AttrPagingState { return a.SystemPagingState }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "PagingUsageExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "paging.usage",
        "Type": "PagingUsage",
        "attributes": [
            {
                "brief": "Unique identifier for the device responsible for managing paging operations.",
                "examples": [
                    "/dev/dm-0",
                ],
                "name": "system.device",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The memory paging state",
                "examples": [
                    "free",
                ],
                "name": "system.paging.state",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
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
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The memory paging state",
                    "examples": [
                        "free",
                    ],
                    "name": "system.paging.state",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                        ],
                    },
                },
                {
                    "brief": "Unique identifier for the device responsible for managing paging operations.",
                    "examples": [
                        "/dev/dm-0",
                    ],
                    "name": "system.device",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Unix swap or windows pagefile usage",
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.paging.usage",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "system.device": {
                        "inherited_fields": [
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "examples",
                        ],
                        "source_group": "registry.system",
                    },
                    "system.paging.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.system.paging",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.paging.usage",
            "name": none,
            "root_namespace": "system",
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
