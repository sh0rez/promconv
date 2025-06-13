package hw

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
)

// Number of errors encountered by the component
type Errors struct {
	*prometheus.CounterVec
	extra ErrorsExtra
}

func NewErrors() Errors {
	labels := []string{AttrId("").Key(), AttrType("").Key(), error.AttrType("").Key(), AttrName("").Key(), AttrParent("").Key()}
	return Errors{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hw_errors",
		Help: "Number of errors encountered by the component",
	}, labels)}
}

func (m Errors) With(id AttrId, kind AttrType, extras ...interface {
	ErrorType() error.AttrType
	HwName() AttrName
	HwParent() AttrParent
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(id.Value(), kind.Value(), extra.ErrorType().Value(), extra.HwName().Value(), extra.HwParent().Value())
}

// Deprecated: Use [Errors.With] instead
func (m Errors) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a Errors) WithErrorType(attr interface{ ErrorType() error.AttrType }) Errors {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a Errors) WithName(attr interface{ HwName() AttrName }) Errors {
	a.extra.AttrName = attr.HwName()
	return a
}
func (a Errors) WithParent(attr interface{ HwParent() AttrParent }) Errors {
	a.extra.AttrParent = attr.HwParent()
	return a
}

type ErrorsExtra struct {
	// The type of error encountered by the component
	AttrErrorType error.AttrType `otel:"error.type"` // An easily-recognizable name for the hardware component
	AttrName      AttrName       `otel:"hw.name"`    // Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
	AttrParent    AttrParent     `otel:"hw.parent"`
}

func (a ErrorsExtra) ErrorType() error.AttrType { return a.AttrErrorType }
func (a ErrorsExtra) HwName() AttrName          { return a.AttrName }
func (a ErrorsExtra) HwParent() AttrParent      { return a.AttrParent }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ErrorsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "errors",
        "Type": "Errors",
        "attributes": [
            {
                "brief": "An identifier for the hardware component, unique within the monitored host\n",
                "examples": [
                    "win32battery_battery_testsysa33_1",
                ],
                "name": "hw.id",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "Type of the component\n",
                "name": "hw.type",
                "note": "Describes the category of the hardware component for which `hw.state` is being reported. For example, `hw.type=temperature` along with `hw.state=degraded` would indicate that the temperature of the hardware component has been reported as `degraded`.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Battery",
                            "deprecated": none,
                            "id": "battery",
                            "note": none,
                            "stability": "development",
                            "value": "battery",
                        },
                        {
                            "brief": "CPU",
                            "deprecated": none,
                            "id": "cpu",
                            "note": none,
                            "stability": "development",
                            "value": "cpu",
                        },
                        {
                            "brief": "Disk controller",
                            "deprecated": none,
                            "id": "disk_controller",
                            "note": none,
                            "stability": "development",
                            "value": "disk_controller",
                        },
                        {
                            "brief": "Enclosure",
                            "deprecated": none,
                            "id": "enclosure",
                            "note": none,
                            "stability": "development",
                            "value": "enclosure",
                        },
                        {
                            "brief": "Fan",
                            "deprecated": none,
                            "id": "fan",
                            "note": none,
                            "stability": "development",
                            "value": "fan",
                        },
                        {
                            "brief": "GPU",
                            "deprecated": none,
                            "id": "gpu",
                            "note": none,
                            "stability": "development",
                            "value": "gpu",
                        },
                        {
                            "brief": "Logical disk",
                            "deprecated": none,
                            "id": "logical_disk",
                            "note": none,
                            "stability": "development",
                            "value": "logical_disk",
                        },
                        {
                            "brief": "Memory",
                            "deprecated": none,
                            "id": "memory",
                            "note": none,
                            "stability": "development",
                            "value": "memory",
                        },
                        {
                            "brief": "Network",
                            "deprecated": none,
                            "id": "network",
                            "note": none,
                            "stability": "development",
                            "value": "network",
                        },
                        {
                            "brief": "Physical disk",
                            "deprecated": none,
                            "id": "physical_disk",
                            "note": none,
                            "stability": "development",
                            "value": "physical_disk",
                        },
                        {
                            "brief": "Power supply",
                            "deprecated": none,
                            "id": "power_supply",
                            "note": none,
                            "stability": "development",
                            "value": "power_supply",
                        },
                        {
                            "brief": "Tape drive",
                            "deprecated": none,
                            "id": "tape_drive",
                            "note": none,
                            "stability": "development",
                            "value": "tape_drive",
                        },
                        {
                            "brief": "Temperature",
                            "deprecated": none,
                            "id": "temperature",
                            "note": none,
                            "stability": "development",
                            "value": "temperature",
                        },
                        {
                            "brief": "Voltage",
                            "deprecated": none,
                            "id": "voltage",
                            "note": none,
                            "stability": "development",
                            "value": "voltage",
                        },
                    ],
                },
            },
            {
                "brief": "The type of error encountered by the component",
                "examples": [
                    "uncorrected",
                    "zero_buffer_credit",
                    "crc",
                    "bad_sector",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD match the error code reported by the component, the canonical name of the error, or another low-cardinality error identifier. Instrumentations SHOULD document the list of errors they report.\n",
                "requirement_level": {
                    "conditionally_required": "if and only if an error has occurred",
                },
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "An easily-recognizable name for the hardware component\n",
                "examples": [
                    "eth0",
                ],
                "name": "hw.name",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)\n",
                "examples": [
                    "dellStorage_perc_0",
                ],
                "name": "hw.parent",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "An easily-recognizable name for the hardware component\n",
                    "examples": [
                        "eth0",
                    ],
                    "name": "hw.name",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)\n",
                    "examples": [
                        "dellStorage_perc_0",
                    ],
                    "name": "hw.parent",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "An identifier for the hardware component, unique within the monitored host\n",
                    "examples": [
                        "win32battery_battery_testsysa33_1",
                    ],
                    "name": "hw.id",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Type of the component\n",
                    "name": "hw.type",
                    "note": "Describes the category of the hardware component for which `hw.state` is being reported. For example, `hw.type=temperature` along with `hw.state=degraded` would indicate that the temperature of the hardware component has been reported as `degraded`.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Battery",
                                "deprecated": none,
                                "id": "battery",
                                "note": none,
                                "stability": "development",
                                "value": "battery",
                            },
                            {
                                "brief": "CPU",
                                "deprecated": none,
                                "id": "cpu",
                                "note": none,
                                "stability": "development",
                                "value": "cpu",
                            },
                            {
                                "brief": "Disk controller",
                                "deprecated": none,
                                "id": "disk_controller",
                                "note": none,
                                "stability": "development",
                                "value": "disk_controller",
                            },
                            {
                                "brief": "Enclosure",
                                "deprecated": none,
                                "id": "enclosure",
                                "note": none,
                                "stability": "development",
                                "value": "enclosure",
                            },
                            {
                                "brief": "Fan",
                                "deprecated": none,
                                "id": "fan",
                                "note": none,
                                "stability": "development",
                                "value": "fan",
                            },
                            {
                                "brief": "GPU",
                                "deprecated": none,
                                "id": "gpu",
                                "note": none,
                                "stability": "development",
                                "value": "gpu",
                            },
                            {
                                "brief": "Logical disk",
                                "deprecated": none,
                                "id": "logical_disk",
                                "note": none,
                                "stability": "development",
                                "value": "logical_disk",
                            },
                            {
                                "brief": "Memory",
                                "deprecated": none,
                                "id": "memory",
                                "note": none,
                                "stability": "development",
                                "value": "memory",
                            },
                            {
                                "brief": "Network",
                                "deprecated": none,
                                "id": "network",
                                "note": none,
                                "stability": "development",
                                "value": "network",
                            },
                            {
                                "brief": "Physical disk",
                                "deprecated": none,
                                "id": "physical_disk",
                                "note": none,
                                "stability": "development",
                                "value": "physical_disk",
                            },
                            {
                                "brief": "Power supply",
                                "deprecated": none,
                                "id": "power_supply",
                                "note": none,
                                "stability": "development",
                                "value": "power_supply",
                            },
                            {
                                "brief": "Tape drive",
                                "deprecated": none,
                                "id": "tape_drive",
                                "note": none,
                                "stability": "development",
                                "value": "tape_drive",
                            },
                            {
                                "brief": "Temperature",
                                "deprecated": none,
                                "id": "temperature",
                                "note": none,
                                "stability": "development",
                                "value": "temperature",
                            },
                            {
                                "brief": "Voltage",
                                "deprecated": none,
                                "id": "voltage",
                                "note": none,
                                "stability": "development",
                                "value": "voltage",
                            },
                        ],
                    },
                },
                {
                    "brief": "The type of error encountered by the component",
                    "examples": [
                        "uncorrected",
                        "zero_buffer_credit",
                        "crc",
                        "bad_sector",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD match the error code reported by the component, the canonical name of the error, or another low-cardinality error identifier. Instrumentations SHOULD document the list of errors they report.\n",
                    "requirement_level": {
                        "conditionally_required": "if and only if an error has occurred",
                    },
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of errors encountered by the component",
            "events": [],
            "id": "metric.hw.errors",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "hw.id": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.hardware",
                    },
                    "hw.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.hardware",
                    },
                    "hw.parent": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.hardware",
                    },
                    "hw.type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.hardware",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/hardware/common-metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "hw.errors",
            "name": none,
            "root_namespace": "hw",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{error}",
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
