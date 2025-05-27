package otel

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/http"
	"shorez.de/promconv/otel/rpc"
	"shorez.de/promconv/otel/server"
)

// The duration of exporting a batch of telemetry records.
type SdkExporterOperationDuration struct {
	*prometheus.HistogramVec
	extra SdkExporterOperationDurationExtra
}

func NewSdkExporterOperationDuration() SdkExporterOperationDuration {
	labels := []string{"error_type", "http_response_status_code", "otel_component_name", "otel_component_type", "rpc_grpc_status_code", "server_address", "server_port"}
	return SdkExporterOperationDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "otel",
		Name:      "sdk_exporter_operation_duration",
		Help:      "The duration of exporting a batch of telemetry records.",
	}, labels)}
}

func (m SdkExporterOperationDuration) With(extra interface {
	AttrErrorType() error.AttrType
	AttrHttpResponseStatusCode() http.AttrResponseStatusCode
	AttrOtelComponentName() AttrComponentName
	AttrOtelComponentType() AttrComponentType
	AttrRpcGrpcStatusCode() rpc.AttrGrpcStatusCode
	AttrServerAddress() server.AttrAddress
	AttrServerPort() server.AttrPort
}) prometheus.Observer {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(extra.AttrErrorType()),
		string(extra.AttrHttpResponseStatusCode()),
		string(extra.AttrOtelComponentName()),
		string(extra.AttrOtelComponentType()),
		string(extra.AttrRpcGrpcStatusCode()),
		string(extra.AttrServerAddress()),
		string(extra.AttrServerPort()),
	)
}

func (a SdkExporterOperationDuration) WithErrorType(attr interface{ AttrErrorType() error.AttrType }) SdkExporterOperationDuration {
	a.extra.ErrorType = attr.AttrErrorType()
	return a
}
func (a SdkExporterOperationDuration) WithHttpResponseStatusCode(attr interface {
	AttrHttpResponseStatusCode() http.AttrResponseStatusCode
}) SdkExporterOperationDuration {
	a.extra.HttpResponseStatusCode = attr.AttrHttpResponseStatusCode()
	return a
}
func (a SdkExporterOperationDuration) WithOtelComponentName(attr interface{ AttrOtelComponentName() AttrComponentName }) SdkExporterOperationDuration {
	a.extra.OtelComponentName = attr.AttrOtelComponentName()
	return a
}
func (a SdkExporterOperationDuration) WithOtelComponentType(attr interface{ AttrOtelComponentType() AttrComponentType }) SdkExporterOperationDuration {
	a.extra.OtelComponentType = attr.AttrOtelComponentType()
	return a
}
func (a SdkExporterOperationDuration) WithRpcGrpcStatusCode(attr interface{ AttrRpcGrpcStatusCode() rpc.AttrGrpcStatusCode }) SdkExporterOperationDuration {
	a.extra.RpcGrpcStatusCode = attr.AttrRpcGrpcStatusCode()
	return a
}
func (a SdkExporterOperationDuration) WithServerAddress(attr interface{ AttrServerAddress() server.AttrAddress }) SdkExporterOperationDuration {
	a.extra.ServerAddress = attr.AttrServerAddress()
	return a
}
func (a SdkExporterOperationDuration) WithServerPort(attr interface{ AttrServerPort() server.AttrPort }) SdkExporterOperationDuration {
	a.extra.ServerPort = attr.AttrServerPort()
	return a
}

type SdkExporterOperationDurationExtra struct {
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// The HTTP status code of the last HTTP request performed in scope of this export call.
	HttpResponseStatusCode http.AttrResponseStatusCode `otel:"http.response.status_code"`
	// A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
	OtelComponentName AttrComponentName `otel:"otel.component.name"`
	// A name identifying the type of the OpenTelemetry component.
	OtelComponentType AttrComponentType `otel:"otel.component.type"`
	// The gRPC status code of the last gRPC requests performed in scope of this export call.
	RpcGrpcStatusCode rpc.AttrGrpcStatusCode `otel:"rpc.grpc.status_code"`
	// Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
}

func (a SdkExporterOperationDurationExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a SdkExporterOperationDurationExtra) AttrHttpResponseStatusCode() http.AttrResponseStatusCode {
	return a.HttpResponseStatusCode
}
func (a SdkExporterOperationDurationExtra) AttrOtelComponentName() AttrComponentName {
	return a.OtelComponentName
}
func (a SdkExporterOperationDurationExtra) AttrOtelComponentType() AttrComponentType {
	return a.OtelComponentType
}
func (a SdkExporterOperationDurationExtra) AttrRpcGrpcStatusCode() rpc.AttrGrpcStatusCode {
	return a.RpcGrpcStatusCode
}
func (a SdkExporterOperationDurationExtra) AttrServerAddress() server.AttrAddress {
	return a.ServerAddress
}
func (a SdkExporterOperationDurationExtra) AttrServerPort() server.AttrPort { return a.ServerPort }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "SdkExporterOperationDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "sdk.exporter.operation.duration",
        "Type": "SdkExporterOperationDuration",
        "attributes": [
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "rejected",
                    "timeout",
                    "500",
                    "java.net.UnknownHostException",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                "requirement_level": {
                    "conditionally_required": "If operation has ended with an error",
                },
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
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
                "brief": "The HTTP status code of the last HTTP request performed in scope of this export call.",
                "examples": [
                    200,
                ],
                "name": "http.response.status_code",
                "requirement_level": {
                    "recommended": "when applicable",
                },
                "stability": "stable",
                "type": "int",
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
                "brief": "A name identifying the type of the OpenTelemetry component.\n",
                "examples": [
                    "otlp_grpc_span_exporter",
                    "com.example.MySpanExporter",
                ],
                "name": "otel.component.type",
                "note": "If none of the standardized values apply, implementations SHOULD use the language-defined name of the type.\nE.g. for Java the fully qualified classname SHOULD be used in this case.\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
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
                "brief": "The gRPC status code of the last gRPC requests performed in scope of this export call.",
                "name": "rpc.grpc.status_code",
                "requirement_level": {
                    "recommended": "when applicable",
                },
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "OK",
                            "deprecated": none,
                            "id": "ok",
                            "note": none,
                            "stability": "development",
                            "value": 0,
                        },
                        {
                            "brief": "CANCELLED",
                            "deprecated": none,
                            "id": "cancelled",
                            "note": none,
                            "stability": "development",
                            "value": 1,
                        },
                        {
                            "brief": "UNKNOWN",
                            "deprecated": none,
                            "id": "unknown",
                            "note": none,
                            "stability": "development",
                            "value": 2,
                        },
                        {
                            "brief": "INVALID_ARGUMENT",
                            "deprecated": none,
                            "id": "invalid_argument",
                            "note": none,
                            "stability": "development",
                            "value": 3,
                        },
                        {
                            "brief": "DEADLINE_EXCEEDED",
                            "deprecated": none,
                            "id": "deadline_exceeded",
                            "note": none,
                            "stability": "development",
                            "value": 4,
                        },
                        {
                            "brief": "NOT_FOUND",
                            "deprecated": none,
                            "id": "not_found",
                            "note": none,
                            "stability": "development",
                            "value": 5,
                        },
                        {
                            "brief": "ALREADY_EXISTS",
                            "deprecated": none,
                            "id": "already_exists",
                            "note": none,
                            "stability": "development",
                            "value": 6,
                        },
                        {
                            "brief": "PERMISSION_DENIED",
                            "deprecated": none,
                            "id": "permission_denied",
                            "note": none,
                            "stability": "development",
                            "value": 7,
                        },
                        {
                            "brief": "RESOURCE_EXHAUSTED",
                            "deprecated": none,
                            "id": "resource_exhausted",
                            "note": none,
                            "stability": "development",
                            "value": 8,
                        },
                        {
                            "brief": "FAILED_PRECONDITION",
                            "deprecated": none,
                            "id": "failed_precondition",
                            "note": none,
                            "stability": "development",
                            "value": 9,
                        },
                        {
                            "brief": "ABORTED",
                            "deprecated": none,
                            "id": "aborted",
                            "note": none,
                            "stability": "development",
                            "value": 10,
                        },
                        {
                            "brief": "OUT_OF_RANGE",
                            "deprecated": none,
                            "id": "out_of_range",
                            "note": none,
                            "stability": "development",
                            "value": 11,
                        },
                        {
                            "brief": "UNIMPLEMENTED",
                            "deprecated": none,
                            "id": "unimplemented",
                            "note": none,
                            "stability": "development",
                            "value": 12,
                        },
                        {
                            "brief": "INTERNAL",
                            "deprecated": none,
                            "id": "internal",
                            "note": none,
                            "stability": "development",
                            "value": 13,
                        },
                        {
                            "brief": "UNAVAILABLE",
                            "deprecated": none,
                            "id": "unavailable",
                            "note": none,
                            "stability": "development",
                            "value": 14,
                        },
                        {
                            "brief": "DATA_LOSS",
                            "deprecated": none,
                            "id": "data_loss",
                            "note": none,
                            "stability": "development",
                            "value": 15,
                        },
                        {
                            "brief": "UNAUTHENTICATED",
                            "deprecated": none,
                            "id": "unauthenticated",
                            "note": none,
                            "stability": "development",
                            "value": 16,
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
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "rejected",
                        "timeout",
                        "500",
                        "java.net.UnknownHostException",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                    "requirement_level": {
                        "conditionally_required": "If operation has ended with an error",
                    },
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "The HTTP status code of the last HTTP request performed in scope of this export call.",
                    "examples": [
                        200,
                    ],
                    "name": "http.response.status_code",
                    "requirement_level": {
                        "recommended": "when applicable",
                    },
                    "stability": "stable",
                    "type": "int",
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
                {
                    "brief": "A name identifying the type of the OpenTelemetry component.\n",
                    "examples": [
                        "otlp_grpc_span_exporter",
                        "com.example.MySpanExporter",
                    ],
                    "name": "otel.component.type",
                    "note": "If none of the standardized values apply, implementations SHOULD use the language-defined name of the type.\nE.g. for Java the fully qualified classname SHOULD be used in this case.\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "The gRPC status code of the last gRPC requests performed in scope of this export call.",
                    "name": "rpc.grpc.status_code",
                    "requirement_level": {
                        "recommended": "when applicable",
                    },
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "OK",
                                "deprecated": none,
                                "id": "ok",
                                "note": none,
                                "stability": "development",
                                "value": 0,
                            },
                            {
                                "brief": "CANCELLED",
                                "deprecated": none,
                                "id": "cancelled",
                                "note": none,
                                "stability": "development",
                                "value": 1,
                            },
                            {
                                "brief": "UNKNOWN",
                                "deprecated": none,
                                "id": "unknown",
                                "note": none,
                                "stability": "development",
                                "value": 2,
                            },
                            {
                                "brief": "INVALID_ARGUMENT",
                                "deprecated": none,
                                "id": "invalid_argument",
                                "note": none,
                                "stability": "development",
                                "value": 3,
                            },
                            {
                                "brief": "DEADLINE_EXCEEDED",
                                "deprecated": none,
                                "id": "deadline_exceeded",
                                "note": none,
                                "stability": "development",
                                "value": 4,
                            },
                            {
                                "brief": "NOT_FOUND",
                                "deprecated": none,
                                "id": "not_found",
                                "note": none,
                                "stability": "development",
                                "value": 5,
                            },
                            {
                                "brief": "ALREADY_EXISTS",
                                "deprecated": none,
                                "id": "already_exists",
                                "note": none,
                                "stability": "development",
                                "value": 6,
                            },
                            {
                                "brief": "PERMISSION_DENIED",
                                "deprecated": none,
                                "id": "permission_denied",
                                "note": none,
                                "stability": "development",
                                "value": 7,
                            },
                            {
                                "brief": "RESOURCE_EXHAUSTED",
                                "deprecated": none,
                                "id": "resource_exhausted",
                                "note": none,
                                "stability": "development",
                                "value": 8,
                            },
                            {
                                "brief": "FAILED_PRECONDITION",
                                "deprecated": none,
                                "id": "failed_precondition",
                                "note": none,
                                "stability": "development",
                                "value": 9,
                            },
                            {
                                "brief": "ABORTED",
                                "deprecated": none,
                                "id": "aborted",
                                "note": none,
                                "stability": "development",
                                "value": 10,
                            },
                            {
                                "brief": "OUT_OF_RANGE",
                                "deprecated": none,
                                "id": "out_of_range",
                                "note": none,
                                "stability": "development",
                                "value": 11,
                            },
                            {
                                "brief": "UNIMPLEMENTED",
                                "deprecated": none,
                                "id": "unimplemented",
                                "note": none,
                                "stability": "development",
                                "value": 12,
                            },
                            {
                                "brief": "INTERNAL",
                                "deprecated": none,
                                "id": "internal",
                                "note": none,
                                "stability": "development",
                                "value": 13,
                            },
                            {
                                "brief": "UNAVAILABLE",
                                "deprecated": none,
                                "id": "unavailable",
                                "note": none,
                                "stability": "development",
                                "value": 14,
                            },
                            {
                                "brief": "DATA_LOSS",
                                "deprecated": none,
                                "id": "data_loss",
                                "note": none,
                                "stability": "development",
                                "value": 15,
                            },
                            {
                                "brief": "UNAUTHENTICATED",
                                "deprecated": none,
                                "id": "unauthenticated",
                                "note": none,
                                "stability": "development",
                                "value": 16,
                            },
                        ],
                    },
                },
            ],
            "brief": "The duration of exporting a batch of telemetry records.",
            "events": [],
            "id": "metric.otel.sdk.exporter.operation.duration",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "http.response.status_code": {
                        "inherited_fields": [
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
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
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                        ],
                        "source_group": "registry.otel.component",
                    },
                    "rpc.grpc.status_code": {
                        "inherited_fields": [
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.rpc",
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
            "metric_name": "otel.sdk.exporter.operation.duration",
            "name": none,
            "note": "This metric defines successful operations using the full success definitions for [http](https://github.com/open-telemetry/opentelemetry-proto/blob/v1.5.0/docs/specification.md#full-success-1)\nand [grpc](https://github.com/open-telemetry/opentelemetry-proto/blob/v1.5.0/docs/specification.md#full-success). Anything else is defined as an unsuccessful operation. For successful\noperations, `error.type` MUST NOT be set. For unsuccessful export operations, `error.type` MUST contain a relevant failure cause.\n",
            "root_namespace": "otel",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "s",
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
