package jvm

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Number of buffers in the pool.
type BufferCount struct {
	*prometheus.GaugeVec
	extra BufferCountExtra
}

func NewBufferCount() BufferCount {
	labels := []string{AttrBufferPoolName("").Key()}
	return BufferCount{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "jvm_buffer_count",
		Help: "Number of buffers in the pool.",
	}, labels)}
}

func (m BufferCount) With(extras ...interface{ JvmBufferPoolName() AttrBufferPoolName }) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.JvmBufferPoolName().Value())
}

// Deprecated: Use [BufferCount.With] instead
func (m BufferCount) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a BufferCount) WithBufferPoolName(attr interface{ JvmBufferPoolName() AttrBufferPoolName }) BufferCount {
	a.extra.AttrBufferPoolName = attr.JvmBufferPoolName()
	return a
}

type BufferCountExtra struct {
	// Name of the buffer pool
	AttrBufferPoolName AttrBufferPoolName `otel:"jvm.buffer.pool.name"`
}

func (a BufferCountExtra) JvmBufferPoolName() AttrBufferPoolName { return a.AttrBufferPoolName }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "BufferCountExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "buffer.count",
        "Type": "BufferCount",
        "attributes": [
            {
                "brief": "Name of the buffer pool.",
                "examples": [
                    "mapped",
                    "direct",
                ],
                "name": "jvm.buffer.pool.name",
                "note": "Pool names are generally obtained via [BufferPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/BufferPoolMXBean.html#getName()).\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Name of the buffer pool.",
                    "examples": [
                        "mapped",
                        "direct",
                    ],
                    "name": "jvm.buffer.pool.name",
                    "note": "Pool names are generally obtained via [BufferPoolMXBean#getName()](https://docs.oracle.com/en/java/javase/11/docs/api/java.management/java/lang/management/BufferPoolMXBean.html#getName()).\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "Number of buffers in the pool.",
            "events": [],
            "id": "metric.jvm.buffer.count",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "jvm.buffer.pool.name": {
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/jvm/metrics-experimental.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "jvm.buffer.count",
            "name": none,
            "root_namespace": "jvm",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{buffer}",
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
