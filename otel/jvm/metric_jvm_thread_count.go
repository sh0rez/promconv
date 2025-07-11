package jvm

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of executing platform threads.
type ThreadCount struct {
	*prometheus.GaugeVec
	extra ThreadCountExtra
}

func NewThreadCount() ThreadCount {
	labels := []string{AttrThreadDaemon("").Key(), AttrThreadState("").Key()}
	return ThreadCount{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "jvm_thread_count",
		Help: "Number of executing platform threads.",
	}, labels)}
}

func (m ThreadCount) With(extras ...interface {
	JvmThreadDaemon() AttrThreadDaemon
	JvmThreadState() AttrThreadState
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.JvmThreadDaemon().Value(), extra.JvmThreadState().Value())
}

// Deprecated: Use [ThreadCount.With] instead
func (m ThreadCount) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ThreadCount) WithThreadDaemon(attr interface{ JvmThreadDaemon() AttrThreadDaemon }) ThreadCount {
	a.extra.AttrThreadDaemon = attr.JvmThreadDaemon()
	return a
}
func (a ThreadCount) WithThreadState(attr interface{ JvmThreadState() AttrThreadState }) ThreadCount {
	a.extra.AttrThreadState = attr.JvmThreadState()
	return a
}

type ThreadCountExtra struct {
	// Whether the thread is daemon or not
	AttrThreadDaemon AttrThreadDaemon `otel:"jvm.thread.daemon"` // State of the thread
	AttrThreadState  AttrThreadState  `otel:"jvm.thread.state"`
}

func (a ThreadCountExtra) JvmThreadDaemon() AttrThreadDaemon { return a.AttrThreadDaemon }
func (a ThreadCountExtra) JvmThreadState() AttrThreadState   { return a.AttrThreadState }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ThreadCountExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "thread.count",
        "Type": "ThreadCount",
        "attributes": [
            {
                "brief": "Whether the thread is daemon or not.",
                "name": "jvm.thread.daemon",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "boolean",
            },
            {
                "brief": "State of the thread.",
                "examples": [
                    "runnable",
                    "blocked",
                ],
                "name": "jvm.thread.state",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "A thread that has not yet started is in this state.",
                            "deprecated": none,
                            "id": "new",
                            "note": none,
                            "stability": "stable",
                            "value": "new",
                        },
                        {
                            "brief": "A thread executing in the Java virtual machine is in this state.",
                            "deprecated": none,
                            "id": "runnable",
                            "note": none,
                            "stability": "stable",
                            "value": "runnable",
                        },
                        {
                            "brief": "A thread that is blocked waiting for a monitor lock is in this state.",
                            "deprecated": none,
                            "id": "blocked",
                            "note": none,
                            "stability": "stable",
                            "value": "blocked",
                        },
                        {
                            "brief": "A thread that is waiting indefinitely for another thread to perform a particular action is in this state.",
                            "deprecated": none,
                            "id": "waiting",
                            "note": none,
                            "stability": "stable",
                            "value": "waiting",
                        },
                        {
                            "brief": "A thread that is waiting for another thread to perform an action for up to a specified waiting time is in this state.",
                            "deprecated": none,
                            "id": "timed_waiting",
                            "note": none,
                            "stability": "stable",
                            "value": "timed_waiting",
                        },
                        {
                            "brief": "A thread that has exited is in this state.",
                            "deprecated": none,
                            "id": "terminated",
                            "note": none,
                            "stability": "stable",
                            "value": "terminated",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Whether the thread is daemon or not.",
                    "name": "jvm.thread.daemon",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "boolean",
                },
                {
                    "brief": "State of the thread.",
                    "examples": [
                        "runnable",
                        "blocked",
                    ],
                    "name": "jvm.thread.state",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "A thread that has not yet started is in this state.",
                                "deprecated": none,
                                "id": "new",
                                "note": none,
                                "stability": "stable",
                                "value": "new",
                            },
                            {
                                "brief": "A thread executing in the Java virtual machine is in this state.",
                                "deprecated": none,
                                "id": "runnable",
                                "note": none,
                                "stability": "stable",
                                "value": "runnable",
                            },
                            {
                                "brief": "A thread that is blocked waiting for a monitor lock is in this state.",
                                "deprecated": none,
                                "id": "blocked",
                                "note": none,
                                "stability": "stable",
                                "value": "blocked",
                            },
                            {
                                "brief": "A thread that is waiting indefinitely for another thread to perform a particular action is in this state.",
                                "deprecated": none,
                                "id": "waiting",
                                "note": none,
                                "stability": "stable",
                                "value": "waiting",
                            },
                            {
                                "brief": "A thread that is waiting for another thread to perform an action for up to a specified waiting time is in this state.",
                                "deprecated": none,
                                "id": "timed_waiting",
                                "note": none,
                                "stability": "stable",
                                "value": "timed_waiting",
                            },
                            {
                                "brief": "A thread that has exited is in this state.",
                                "deprecated": none,
                                "id": "terminated",
                                "note": none,
                                "stability": "stable",
                                "value": "terminated",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of executing platform threads.",
            "events": [],
            "id": "metric.jvm.thread.count",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "jvm.thread.daemon": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.jvm",
                    },
                    "jvm.thread.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.jvm",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/jvm/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "jvm.thread.count",
            "name": none,
            "root_namespace": "jvm",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "{thread}",
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
