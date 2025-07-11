package system

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/network"
)

// Count of network errors detected
type NetworkErrors struct {
	*prometheus.CounterVec
	extra NetworkErrorsExtra
}

func NewNetworkErrors() NetworkErrors {
	labels := []string{network.AttrInterfaceName("").Key(), network.AttrIoDirection("").Key()}
	return NetworkErrors{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "system_network_errors",
		Help: "Count of network errors detected",
	}, labels)}
}

func (m NetworkErrors) With(extras ...interface {
	NetworkInterfaceName() network.AttrInterfaceName
	NetworkIoDirection() network.AttrIoDirection
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(extra.NetworkInterfaceName().Value(), extra.NetworkIoDirection().Value())
}

// Deprecated: Use [NetworkErrors.With] instead
func (m NetworkErrors) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a NetworkErrors) WithNetworkInterfaceName(attr interface {
	NetworkInterfaceName() network.AttrInterfaceName
}) NetworkErrors {
	a.extra.AttrNetworkInterfaceName = attr.NetworkInterfaceName()
	return a
}
func (a NetworkErrors) WithNetworkIoDirection(attr interface {
	NetworkIoDirection() network.AttrIoDirection
}) NetworkErrors {
	a.extra.AttrNetworkIoDirection = attr.NetworkIoDirection()
	return a
}

type NetworkErrorsExtra struct {
	// The network interface name
	AttrNetworkInterfaceName network.AttrInterfaceName `otel:"network.interface.name"` // The network IO operation direction
	AttrNetworkIoDirection   network.AttrIoDirection   `otel:"network.io.direction"`
}

func (a NetworkErrorsExtra) NetworkInterfaceName() network.AttrInterfaceName {
	return a.AttrNetworkInterfaceName
}
func (a NetworkErrorsExtra) NetworkIoDirection() network.AttrIoDirection {
	return a.AttrNetworkIoDirection
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "NetworkErrorsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "network.errors",
        "Type": "NetworkErrors",
        "attributes": [
            {
                "brief": "The network interface name.",
                "examples": [
                    "lo",
                    "eth0",
                ],
                "name": "network.interface.name",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The network IO operation direction.",
                "examples": [
                    "transmit",
                ],
                "name": "network.io.direction",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "transmit",
                            "note": none,
                            "stability": "development",
                            "value": "transmit",
                        },
                        {
                            "brief": none,
                            "deprecated": none,
                            "id": "receive",
                            "note": none,
                            "stability": "development",
                            "value": "receive",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The network IO operation direction.",
                    "examples": [
                        "transmit",
                    ],
                    "name": "network.io.direction",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "transmit",
                                "note": none,
                                "stability": "development",
                                "value": "transmit",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "receive",
                                "note": none,
                                "stability": "development",
                                "value": "receive",
                            },
                        ],
                    },
                },
                {
                    "brief": "The network interface name.",
                    "examples": [
                        "lo",
                        "eth0",
                    ],
                    "name": "network.interface.name",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Count of network errors detected",
            "entity_associations": [
                "host",
            ],
            "events": [],
            "id": "metric.system.network.errors",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "network.interface.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.io.direction": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.network",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/system/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "system.network.errors",
            "name": none,
            "note": "Measured as:\n\n- Linux: the `errs` column in `/proc/dev/net` ([source](https://web.archive.org/web/20180321091318/http://www.onlamp.com/pub/a/linux/2000/11/16/LinuxAdmin.html)).\n- Windows: [`InErrors`/`OutErrors`](https://docs.microsoft.com/windows/win32/api/netioapi/ns-netioapi-mib_if_row2)\n  from [`GetIfEntry2`](https://docs.microsoft.com/windows/win32/api/netioapi/nf-netioapi-getifentry2).\n",
            "root_namespace": "system",
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
