package faas

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of invocation timeouts
type Timeouts struct {
	*prometheus.CounterVec
	extra TimeoutsExtra
}

func NewTimeouts() Timeouts {
	labels := []string{AttrTrigger("").Key()}
	return Timeouts{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "faas_timeouts",
		Help: "Number of invocation timeouts",
	}, labels)}
}

func (m Timeouts) With(extras ...interface{ FaasTrigger() AttrTrigger }) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(extra.FaasTrigger().Value())
}

// Deprecated: Use [Timeouts.With] instead
func (m Timeouts) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a Timeouts) WithTrigger(attr interface{ FaasTrigger() AttrTrigger }) Timeouts {
	a.extra.AttrTrigger = attr.FaasTrigger()
	return a
}

type TimeoutsExtra struct {
	// Type of the trigger which caused this function invocation
	AttrTrigger AttrTrigger `otel:"faas.trigger"`
}

func (a TimeoutsExtra) FaasTrigger() AttrTrigger { return a.AttrTrigger }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "TimeoutsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "timeouts",
        "Type": "Timeouts",
        "attributes": [
            {
                "brief": "Type of the trigger which caused this function invocation.\n",
                "name": "faas.trigger",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "A response to some data source operation such as a database or filesystem read/write",
                            "deprecated": none,
                            "id": "datasource",
                            "note": none,
                            "stability": "development",
                            "value": "datasource",
                        },
                        {
                            "brief": "To provide an answer to an inbound HTTP request",
                            "deprecated": none,
                            "id": "http",
                            "note": none,
                            "stability": "development",
                            "value": "http",
                        },
                        {
                            "brief": "A function is set to be executed when messages are sent to a messaging system",
                            "deprecated": none,
                            "id": "pubsub",
                            "note": none,
                            "stability": "development",
                            "value": "pubsub",
                        },
                        {
                            "brief": "A function is scheduled to be executed regularly",
                            "deprecated": none,
                            "id": "timer",
                            "note": none,
                            "stability": "development",
                            "value": "timer",
                        },
                        {
                            "brief": "If none of the others apply",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "development",
                            "value": "other",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Type of the trigger which caused this function invocation.\n",
                    "name": "faas.trigger",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "A response to some data source operation such as a database or filesystem read/write",
                                "deprecated": none,
                                "id": "datasource",
                                "note": none,
                                "stability": "development",
                                "value": "datasource",
                            },
                            {
                                "brief": "To provide an answer to an inbound HTTP request",
                                "deprecated": none,
                                "id": "http",
                                "note": none,
                                "stability": "development",
                                "value": "http",
                            },
                            {
                                "brief": "A function is set to be executed when messages are sent to a messaging system",
                                "deprecated": none,
                                "id": "pubsub",
                                "note": none,
                                "stability": "development",
                                "value": "pubsub",
                            },
                            {
                                "brief": "A function is scheduled to be executed regularly",
                                "deprecated": none,
                                "id": "timer",
                                "note": none,
                                "stability": "development",
                                "value": "timer",
                            },
                            {
                                "brief": "If none of the others apply",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "development",
                                "value": "other",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of invocation timeouts",
            "events": [],
            "id": "metric.faas.timeouts",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "faas.trigger": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.faas",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/faas/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "faas.timeouts",
            "name": none,
            "root_namespace": "faas",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{timeout}",
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
