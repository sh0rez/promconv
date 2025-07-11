package hw

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Total energy consumed by the entire physical host, in joules
type HostEnergy struct {
	*prometheus.CounterVec
	extra HostEnergyExtra
}

func NewHostEnergy() HostEnergy {
	labels := []string{AttrId("").Key(), AttrName("").Key(), AttrParent("").Key()}
	return HostEnergy{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hw_host_energy",
		Help: "Total energy consumed by the entire physical host, in joules",
	}, labels)}
}

func (m HostEnergy) With(id AttrId, extras ...interface {
	HwName() AttrName
	HwParent() AttrParent
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(id.Value(), extra.HwName().Value(), extra.HwParent().Value())
}

// Deprecated: Use [HostEnergy.With] instead
func (m HostEnergy) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a HostEnergy) WithName(attr interface{ HwName() AttrName }) HostEnergy {
	a.extra.AttrName = attr.HwName()
	return a
}
func (a HostEnergy) WithParent(attr interface{ HwParent() AttrParent }) HostEnergy {
	a.extra.AttrParent = attr.HwParent()
	return a
}

type HostEnergyExtra struct {
	// An easily-recognizable name for the hardware component
	AttrName   AttrName   `otel:"hw.name"` // Unique identifier of the parent component (typically the `hw.id` attribute of the enclosure, or disk controller)
	AttrParent AttrParent `otel:"hw.parent"`
}

func (a HostEnergyExtra) HwName() AttrName     { return a.AttrName }
func (a HostEnergyExtra) HwParent() AttrParent { return a.AttrParent }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "HostEnergyExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "host.energy",
        "Type": "HostEnergy",
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
            ],
            "brief": "Total energy consumed by the entire physical host, in joules",
            "events": [],
            "id": "metric.hw.host.energy",
            "instrument": "counter",
            "lineage": {
                "attributes": {
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/hardware/host-metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "hw.host.energy",
            "name": none,
            "note": "The overall energy usage of a host MUST be reported using the specific `hw.host.energy` and `hw.host.power` metrics **only**, instead of the generic `hw.energy` and `hw.power` described in the previous section, to prevent summing up overlapping values.\n",
            "root_namespace": "hw",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "J",
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
