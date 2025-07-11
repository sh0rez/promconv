package otel

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/server"
)

// The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)
type SdkExporterMetricDataPointInflight struct {
	*prometheus.GaugeVec
	extra SdkExporterMetricDataPointInflightExtra
}

func NewSdkExporterMetricDataPointInflight() SdkExporterMetricDataPointInflight {
	labels := []string{AttrComponentName("").Key(), AttrComponentType("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key()}
	return SdkExporterMetricDataPointInflight{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "otel_sdk_exporter_metric_data_point_inflight",
		Help: "The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)",
	}, labels)}
}

func (m SdkExporterMetricDataPointInflight) With(extras ...interface {
	OtelComponentName() AttrComponentName
	OtelComponentType() AttrComponentType
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(extra.OtelComponentName().Value(), extra.OtelComponentType().Value(), extra.ServerAddress().Value(), extra.ServerPort().Value())
}

// Deprecated: Use [SdkExporterMetricDataPointInflight.With] instead
func (m SdkExporterMetricDataPointInflight) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a SdkExporterMetricDataPointInflight) WithComponentName(attr interface{ OtelComponentName() AttrComponentName }) SdkExporterMetricDataPointInflight {
	a.extra.AttrComponentName = attr.OtelComponentName()
	return a
}
func (a SdkExporterMetricDataPointInflight) WithComponentType(attr interface{ OtelComponentType() AttrComponentType }) SdkExporterMetricDataPointInflight {
	a.extra.AttrComponentType = attr.OtelComponentType()
	return a
}
func (a SdkExporterMetricDataPointInflight) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) SdkExporterMetricDataPointInflight {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a SdkExporterMetricDataPointInflight) WithServerPort(attr interface{ ServerPort() server.AttrPort }) SdkExporterMetricDataPointInflight {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}

type SdkExporterMetricDataPointInflightExtra struct {
	// A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance
	AttrComponentName AttrComponentName  `otel:"otel.component.name"` // A name identifying the type of the OpenTelemetry component
	AttrComponentType AttrComponentType  `otel:"otel.component.type"` // Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
	AttrServerAddress server.AttrAddress `otel:"server.address"`      // Server port number
	AttrServerPort    server.AttrPort    `otel:"server.port"`
}

func (a SdkExporterMetricDataPointInflightExtra) OtelComponentName() AttrComponentName {
	return a.AttrComponentName
}
func (a SdkExporterMetricDataPointInflightExtra) OtelComponentType() AttrComponentType {
	return a.AttrComponentType
}
func (a SdkExporterMetricDataPointInflightExtra) ServerAddress() server.AttrAddress {
	return a.AttrServerAddress
}
func (a SdkExporterMetricDataPointInflightExtra) ServerPort() server.AttrPort {
	return a.AttrServerPort
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "SdkExporterMetricDataPointInflightExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "sdk.exporter.metric_data_point.inflight",
        "Type": "SdkExporterMetricDataPointInflight",
        "attributes": [
            {
                "brief": "A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.\n",
                "examples": [
                    "otlp_grpc_span_exporter/0",
                    "custom-name",
                ],
                "name": "otel.component.name",
                "note": "Implementations SHOULD ensure a low cardinality for this attribute, even across application or SDK restarts.\nE.g. implementations MUST NOT use UUIDs as values for this attribute.\n\nImplementations MAY achieve these goals by following a `<otel.component.type>/<instance-counter>` pattern, e.g. `batching_span_processor/0`.\nHereby `otel.component.type` refers to the corresponding attribute value of the component.\n\nThe value of `instance-counter` MAY be automatically assigned by the component and uniqueness within the enclosing SDK instance MUST be guaranteed.\nFor example, `<instance-counter>` MAY be implemented by using a monotonically increasing counter (starting with `0`), which is incremented every time an\ninstance of the given component type is started.\n\nWith this implementation, for example the first Batching Span Processor would have `batching_span_processor/0`\nas `otel.component.name`, the second one `batching_span_processor/1` and so on.\nThese values will therefore be reused in the case of an application restart.\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "A name identifying the type of the OpenTelemetry component.\n",
                "examples": [
                    "batching_span_processor",
                    "com.example.MySpanExporter",
                ],
                "name": "otel.component.type",
                "note": "If none of the standardized values apply, implementations SHOULD use the language-defined name of the type.\nE.g. for Java the fully qualified classname SHOULD be used in this case.\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "The builtin SDK batching span processor\n",
                            "deprecated": none,
                            "id": "batching_span_processor",
                            "note": none,
                            "stability": "development",
                            "value": "batching_span_processor",
                        },
                        {
                            "brief": "The builtin SDK simple span processor\n",
                            "deprecated": none,
                            "id": "simple_span_processor",
                            "note": none,
                            "stability": "development",
                            "value": "simple_span_processor",
                        },
                        {
                            "brief": "The builtin SDK batching log record processor\n",
                            "deprecated": none,
                            "id": "batching_log_processor",
                            "note": none,
                            "stability": "development",
                            "value": "batching_log_processor",
                        },
                        {
                            "brief": "The builtin SDK simple log record processor\n",
                            "deprecated": none,
                            "id": "simple_log_processor",
                            "note": none,
                            "stability": "development",
                            "value": "simple_log_processor",
                        },
                        {
                            "brief": "OTLP span exporter over gRPC with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_grpc_span_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_grpc_span_exporter",
                        },
                        {
                            "brief": "OTLP span exporter over HTTP with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_span_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_span_exporter",
                        },
                        {
                            "brief": "OTLP span exporter over HTTP with JSON serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_json_span_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_json_span_exporter",
                        },
                        {
                            "brief": "OTLP log record exporter over gRPC with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_grpc_log_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_grpc_log_exporter",
                        },
                        {
                            "brief": "OTLP log record exporter over HTTP with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_log_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_log_exporter",
                        },
                        {
                            "brief": "OTLP log record exporter over HTTP with JSON serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_json_log_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_json_log_exporter",
                        },
                        {
                            "brief": "The builtin SDK periodically exporting metric reader\n",
                            "deprecated": none,
                            "id": "periodic_metric_reader",
                            "note": none,
                            "stability": "development",
                            "value": "periodic_metric_reader",
                        },
                        {
                            "brief": "OTLP metric exporter over gRPC with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_grpc_metric_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_grpc_metric_exporter",
                        },
                        {
                            "brief": "OTLP metric exporter over HTTP with protobuf serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_metric_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_metric_exporter",
                        },
                        {
                            "brief": "OTLP metric exporter over HTTP with JSON serialization\n",
                            "deprecated": none,
                            "id": "otlp_http_json_metric_exporter",
                            "note": none,
                            "stability": "development",
                            "value": "otlp_http_json_metric_exporter",
                        },
                    ],
                },
            },
            {
                "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": {
                    "recommended": "when applicable",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Server port number.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": {
                    "recommended": "when applicable",
                },
                "stability": "stable",
                "type": "int",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "A name identifying the type of the OpenTelemetry component.\n",
                    "examples": [
                        "batching_span_processor",
                        "com.example.MySpanExporter",
                    ],
                    "name": "otel.component.type",
                    "note": "If none of the standardized values apply, implementations SHOULD use the language-defined name of the type.\nE.g. for Java the fully qualified classname SHOULD be used in this case.\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The builtin SDK batching span processor\n",
                                "deprecated": none,
                                "id": "batching_span_processor",
                                "note": none,
                                "stability": "development",
                                "value": "batching_span_processor",
                            },
                            {
                                "brief": "The builtin SDK simple span processor\n",
                                "deprecated": none,
                                "id": "simple_span_processor",
                                "note": none,
                                "stability": "development",
                                "value": "simple_span_processor",
                            },
                            {
                                "brief": "The builtin SDK batching log record processor\n",
                                "deprecated": none,
                                "id": "batching_log_processor",
                                "note": none,
                                "stability": "development",
                                "value": "batching_log_processor",
                            },
                            {
                                "brief": "The builtin SDK simple log record processor\n",
                                "deprecated": none,
                                "id": "simple_log_processor",
                                "note": none,
                                "stability": "development",
                                "value": "simple_log_processor",
                            },
                            {
                                "brief": "OTLP span exporter over gRPC with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_grpc_span_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_grpc_span_exporter",
                            },
                            {
                                "brief": "OTLP span exporter over HTTP with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_span_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_span_exporter",
                            },
                            {
                                "brief": "OTLP span exporter over HTTP with JSON serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_json_span_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_json_span_exporter",
                            },
                            {
                                "brief": "OTLP log record exporter over gRPC with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_grpc_log_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_grpc_log_exporter",
                            },
                            {
                                "brief": "OTLP log record exporter over HTTP with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_log_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_log_exporter",
                            },
                            {
                                "brief": "OTLP log record exporter over HTTP with JSON serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_json_log_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_json_log_exporter",
                            },
                            {
                                "brief": "The builtin SDK periodically exporting metric reader\n",
                                "deprecated": none,
                                "id": "periodic_metric_reader",
                                "note": none,
                                "stability": "development",
                                "value": "periodic_metric_reader",
                            },
                            {
                                "brief": "OTLP metric exporter over gRPC with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_grpc_metric_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_grpc_metric_exporter",
                            },
                            {
                                "brief": "OTLP metric exporter over HTTP with protobuf serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_metric_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_metric_exporter",
                            },
                            {
                                "brief": "OTLP metric exporter over HTTP with JSON serialization\n",
                                "deprecated": none,
                                "id": "otlp_http_json_metric_exporter",
                                "note": none,
                                "stability": "development",
                                "value": "otlp_http_json_metric_exporter",
                            },
                        ],
                    },
                },
                {
                    "brief": "A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.\n",
                    "examples": [
                        "otlp_grpc_span_exporter/0",
                        "custom-name",
                    ],
                    "name": "otel.component.name",
                    "note": "Implementations SHOULD ensure a low cardinality for this attribute, even across application or SDK restarts.\nE.g. implementations MUST NOT use UUIDs as values for this attribute.\n\nImplementations MAY achieve these goals by following a `<otel.component.type>/<instance-counter>` pattern, e.g. `batching_span_processor/0`.\nHereby `otel.component.type` refers to the corresponding attribute value of the component.\n\nThe value of `instance-counter` MAY be automatically assigned by the component and uniqueness within the enclosing SDK instance MUST be guaranteed.\nFor example, `<instance-counter>` MAY be implemented by using a monotonically increasing counter (starting with `0`), which is incremented every time an\ninstance of the given component type is started.\n\nWith this implementation, for example the first Batching Span Processor would have `batching_span_processor/0`\nas `otel.component.name`, the second one `batching_span_processor/1` and so on.\nThese values will therefore be reused in the case of an application restart.\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": {
                        "recommended": "when applicable",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": {
                        "recommended": "when applicable",
                    },
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)",
            "events": [],
            "id": "metric.otel.sdk.exporter.metric_data_point.inflight",
            "instrument": "updowncounter",
            "lineage": {
                "attributes": {
                    "otel.component.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.otel.component",
                    },
                    "otel.component.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.otel.component",
                    },
                    "server.address": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/otel/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "otel.sdk.exporter.metric_data_point.inflight",
            "name": none,
            "note": "For successful exports, `error.type` MUST NOT be set. For failed exports, `error.type` MUST contain the failure cause.\n",
            "root_namespace": "otel",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{data_point}",
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
