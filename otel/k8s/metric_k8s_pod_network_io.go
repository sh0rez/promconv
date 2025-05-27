package k8s

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/network"
)

// Network bytes for the Pod
type PodNetworkIo struct {
	*prometheus.CounterVec
	extra PodNetworkIoExtra
}

func NewPodNetworkIo() PodNetworkIo {
	labels := []string{"network_interface_name", "network_io_direction"}
	return PodNetworkIo{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "k8s",
		Name:      "pod_network_io",
		Help:      "Network bytes for the Pod",
	}, labels)}
}

func (m PodNetworkIo) With(extra interface {
	AttrNetworkInterfaceName() network.AttrInterfaceName
	AttrNetworkIoDirection() network.AttrIoDirection
}) prometheus.Counter {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrNetworkInterfaceName()),
		string(extra.AttrNetworkIoDirection()),
	)
}

func (a PodNetworkIo) WithNetworkInterfaceName(attr interface {
	AttrNetworkInterfaceName() network.AttrInterfaceName
}) PodNetworkIo {
	a.extra.NetworkInterfaceName = attr.AttrNetworkInterfaceName()
	return a
}
func (a PodNetworkIo) WithNetworkIoDirection(attr interface {
	AttrNetworkIoDirection() network.AttrIoDirection
}) PodNetworkIo {
	a.extra.NetworkIoDirection = attr.AttrNetworkIoDirection()
	return a
}

type PodNetworkIoExtra struct {
	// The network interface name.
	NetworkInterfaceName network.AttrInterfaceName `otel:"network.interface.name"`
	// The network IO operation direction.
	NetworkIoDirection network.AttrIoDirection `otel:"network.io.direction"`
}

func (a PodNetworkIoExtra) AttrNetworkInterfaceName() network.AttrInterfaceName {
	return a.NetworkInterfaceName
}
func (a PodNetworkIoExtra) AttrNetworkIoDirection() network.AttrIoDirection {
	return a.NetworkIoDirection
}

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "PodNetworkIoExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "pod.network.io",
        "Type": "PodNetworkIo",
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
                    "allow_custom_values": none,
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
                        "allow_custom_values": none,
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
            "brief": "Network bytes for the Pod",
            "events": [],
            "id": "metric.k8s.pod.network.io",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/k8s/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "k8s.pod.network.io",
            "name": none,
            "root_namespace": "k8s",
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
