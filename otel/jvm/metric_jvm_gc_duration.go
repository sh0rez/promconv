package jvm

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Duration of JVM garbage collection actions.
type GcDuration struct {
	*prometheus.HistogramVec
	extra GcDurationExtra
}

func NewGcDuration() GcDuration {
	labels := []string{AttrGcAction("").Key(), AttrGcName("").Key(), AttrGcCause("").Key()}
	return GcDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "jvm_gc_duration",
		Help: "Duration of JVM garbage collection actions.",
	}, labels)}
}

func (m GcDuration) With(extras ...interface {
	JvmGcAction() AttrGcAction
	JvmGcName() AttrGcName
	JvmGcCause() AttrGcCause
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(extra.JvmGcAction().Value(), extra.JvmGcName().Value(), extra.JvmGcCause().Value())
}

// Deprecated: Use [GcDuration.With] instead
func (m GcDuration) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a GcDuration) WithGcAction(attr interface{ JvmGcAction() AttrGcAction }) GcDuration {
	a.extra.AttrGcAction = attr.JvmGcAction()
	return a
}
func (a GcDuration) WithGcName(attr interface{ JvmGcName() AttrGcName }) GcDuration {
	a.extra.AttrGcName = attr.JvmGcName()
	return a
}
func (a GcDuration) WithGcCause(attr interface{ JvmGcCause() AttrGcCause }) GcDuration {
	a.extra.AttrGcCause = attr.JvmGcCause()
	return a
}

type GcDurationExtra struct {
	// Name of the garbage collector action
	AttrGcAction AttrGcAction `otel:"jvm.gc.action"` // Name of the garbage collector
	AttrGcName   AttrGcName   `otel:"jvm.gc.name"`   // Name of the garbage collector cause
	AttrGcCause  AttrGcCause  `otel:"jvm.gc.cause"`
}

func (a GcDurationExtra) JvmGcAction() AttrGcAction { return a.AttrGcAction }
func (a GcDurationExtra) JvmGcName() AttrGcName     { return a.AttrGcName }
func (a GcDurationExtra) JvmGcCause() AttrGcCause   { return a.AttrGcCause }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "GcDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "gc.duration",
        "Type": "GcDuration",
        "attributes": [
            {
                "brief": "Name of the garbage collector action.",
                "examples": [
                    "end of minor GC",
                    "end of major GC",
                ],
                "name": "jvm.gc.action",
                "note": "Garbage collector action is generally obtained via [GarbageCollectionNotificationInfo#getGcAction()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcAction()).\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Name of the garbage collector.",
                "examples": [
                    "G1 Young Generation",
                    "G1 Old Generation",
                ],
                "name": "jvm.gc.name",
                "note": "Garbage collector name is generally obtained via [GarbageCollectionNotificationInfo#getGcName()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcName()).\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Name of the garbage collector cause.",
                "examples": [
                    "System.gc()",
                    "Allocation Failure",
                ],
                "name": "jvm.gc.cause",
                "note": "Garbage collector cause is generally obtained via [GarbageCollectionNotificationInfo#getGcCause()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcCause()).\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Name of the garbage collector action.",
                    "examples": [
                        "end of minor GC",
                        "end of major GC",
                    ],
                    "name": "jvm.gc.action",
                    "note": "Garbage collector action is generally obtained via [GarbageCollectionNotificationInfo#getGcAction()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcAction()).\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Name of the garbage collector.",
                    "examples": [
                        "G1 Young Generation",
                        "G1 Old Generation",
                    ],
                    "name": "jvm.gc.name",
                    "note": "Garbage collector name is generally obtained via [GarbageCollectionNotificationInfo#getGcName()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcName()).\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Name of the garbage collector cause.",
                    "examples": [
                        "System.gc()",
                        "Allocation Failure",
                    ],
                    "name": "jvm.gc.cause",
                    "note": "Garbage collector cause is generally obtained via [GarbageCollectionNotificationInfo#getGcCause()](https://docs.oracle.com/en/java/javase/11/docs/api/jdk.management/com/sun/management/GarbageCollectionNotificationInfo.html#getGcCause()).\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Duration of JVM garbage collection actions.",
            "events": [],
            "id": "metric.jvm.gc.duration",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "jvm.gc.action": {
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
                    "jvm.gc.cause": {
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
                    "jvm.gc.name": {
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
            "metric_name": "jvm.gc.duration",
            "name": none,
            "root_namespace": "jvm",
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
