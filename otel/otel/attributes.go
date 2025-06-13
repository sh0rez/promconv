package otel

// A name uniquely identifying the instance of the OpenTelemetry component within its containing SDK instance.
//
// Implementations SHOULD ensure a low cardinality for this attribute, even across application or SDK restarts.
// E.g. implementations MUST NOT use UUIDs as values for this attribute.
//
// Implementations MAY achieve these goals by following a `<otel.component.type>/<instance-counter>` pattern, e.g. `batching_span_processor/0`.
// Hereby `otel.component.type` refers to the corresponding attribute value of the component.
//
// The value of `instance-counter` MAY be automatically assigned by the component and uniqueness within the enclosing SDK instance MUST be guaranteed.
// For example, `<instance-counter>` MAY be implemented by using a monotonically increasing counter (starting with `0`), which is incremented every time an
// instance of the given component type is started.
//
// With this implementation, for example the first Batching Span Processor would have `batching_span_processor/0`
// as `otel.component.name`, the second one `batching_span_processor/1` and so on.
// These values will therefore be reused in the case of an application restart
type AttrComponentName string // otel.component.name

func (AttrComponentName) Development()    {}
func (AttrComponentName) Recommended()    {}
func (AttrComponentName) Key() string     { return "otel_component_name" }
func (a AttrComponentName) Value() string { return string(a) }

// A name identifying the type of the OpenTelemetry component.
//
// If none of the standardized values apply, implementations SHOULD use the language-defined name of the type.
// E.g. for Java the fully qualified classname SHOULD be used in this case
type AttrComponentType string // otel.component.type

func (AttrComponentType) Development()    {}
func (AttrComponentType) Recommended()    {}
func (AttrComponentType) Key() string     { return "otel_component_type" }
func (a AttrComponentType) Value() string { return string(a) }

const ComponentTypeBatchingSpanProcessor AttrComponentType = "batching_span_processor"
const ComponentTypeSimpleSpanProcessor AttrComponentType = "simple_span_processor"
const ComponentTypeBatchingLogProcessor AttrComponentType = "batching_log_processor"
const ComponentTypeSimpleLogProcessor AttrComponentType = "simple_log_processor"
const ComponentTypeOtlpGrpcSpanExporter AttrComponentType = "otlp_grpc_span_exporter"
const ComponentTypeOtlpHttpSpanExporter AttrComponentType = "otlp_http_span_exporter"
const ComponentTypeOtlpHttpJsonSpanExporter AttrComponentType = "otlp_http_json_span_exporter"
const ComponentTypeOtlpGrpcLogExporter AttrComponentType = "otlp_grpc_log_exporter"
const ComponentTypeOtlpHttpLogExporter AttrComponentType = "otlp_http_log_exporter"
const ComponentTypeOtlpHttpJsonLogExporter AttrComponentType = "otlp_http_json_log_exporter"
const ComponentTypePeriodicMetricReader AttrComponentType = "periodic_metric_reader"
const ComponentTypeOtlpGrpcMetricExporter AttrComponentType = "otlp_grpc_metric_exporter"
const ComponentTypeOtlpHttpMetricExporter AttrComponentType = "otlp_http_metric_exporter"
const ComponentTypeOtlpHttpJsonMetricExporter AttrComponentType = "otlp_http_json_metric_exporter"

// Deprecated. Use the `otel.scope.name` attribute
type AttrLibraryName string // otel.library.name

func (AttrLibraryName) Development()    {}
func (AttrLibraryName) Recommended()    {}
func (AttrLibraryName) Key() string     { return "otel_library_name" }
func (a AttrLibraryName) Value() string { return string(a) }

// Deprecated. Use the `otel.scope.version` attribute
type AttrLibraryVersion string // otel.library.version

func (AttrLibraryVersion) Development()    {}
func (AttrLibraryVersion) Recommended()    {}
func (AttrLibraryVersion) Key() string     { return "otel_library_version" }
func (a AttrLibraryVersion) Value() string { return string(a) }

// The name of the instrumentation scope - (`InstrumentationScope.Name` in OTLP)
type AttrScopeName string // otel.scope.name

func (AttrScopeName) Stable()         {}
func (AttrScopeName) Recommended()    {}
func (AttrScopeName) Key() string     { return "otel_scope_name" }
func (a AttrScopeName) Value() string { return string(a) }

// The version of the instrumentation scope - (`InstrumentationScope.Version` in OTLP)
type AttrScopeVersion string // otel.scope.version

func (AttrScopeVersion) Stable()         {}
func (AttrScopeVersion) Recommended()    {}
func (AttrScopeVersion) Key() string     { return "otel_scope_version" }
func (a AttrScopeVersion) Value() string { return string(a) }

// The result value of the sampler for this span
type AttrSpanSamplingResult string // otel.span.sampling_result

func (AttrSpanSamplingResult) Development()    {}
func (AttrSpanSamplingResult) Recommended()    {}
func (AttrSpanSamplingResult) Key() string     { return "otel_span_sampling_result" }
func (a AttrSpanSamplingResult) Value() string { return string(a) }

const SpanSamplingResultDrop AttrSpanSamplingResult = "DROP"
const SpanSamplingResultRecordOnly AttrSpanSamplingResult = "RECORD_ONLY"
const SpanSamplingResultRecordAndSample AttrSpanSamplingResult = "RECORD_AND_SAMPLE"

// Name of the code, either "OK" or "ERROR". MUST NOT be set if the status code is UNSET
type AttrStatusCode string // otel.status_code

func (AttrStatusCode) Stable()         {}
func (AttrStatusCode) Recommended()    {}
func (AttrStatusCode) Key() string     { return "otel_status_code" }
func (a AttrStatusCode) Value() string { return string(a) }

const StatusCodeOk AttrStatusCode = "OK"
const StatusCodeError AttrStatusCode = "ERROR"

// Description of the Status if it has a value, otherwise not set
type AttrStatusDescription string // otel.status_description

func (AttrStatusDescription) Stable()         {}
func (AttrStatusDescription) Recommended()    {}
func (AttrStatusDescription) Key() string     { return "otel_status_description" }
func (a AttrStatusDescription) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
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
                    "root_namespace": "otel",
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
                    "root_namespace": "otel",
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
                    "brief": "Deprecated. Use the `otel.scope.name` attribute",
                    "deprecated": {
                        "note": "Replaced by `otel.scope.name`.",
                        "reason": "renamed",
                        "renamed_to": "otel.scope.name",
                    },
                    "examples": [
                        "io.opentelemetry.contrib.mongodb",
                    ],
                    "name": "otel.library.name",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated. Use the `otel.scope.version` attribute.",
                    "deprecated": {
                        "note": "Replaced by `otel.scope.version`.",
                        "reason": "renamed",
                        "renamed_to": "otel.scope.version",
                    },
                    "examples": [
                        "1.0.0",
                    ],
                    "name": "otel.library.version",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the instrumentation scope - (`InstrumentationScope.Name` in OTLP).",
                    "examples": [
                        "io.opentelemetry.contrib.mongodb",
                    ],
                    "name": "otel.scope.name",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The version of the instrumentation scope - (`InstrumentationScope.Version` in OTLP).",
                    "examples": [
                        "1.0.0",
                    ],
                    "name": "otel.scope.version",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The result value of the sampler for this span",
                    "name": "otel.span.sampling_result",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The span is not sampled and not recording",
                                "deprecated": none,
                                "id": "drop",
                                "note": none,
                                "stability": "development",
                                "value": "DROP",
                            },
                            {
                                "brief": "The span is not sampled, but recording",
                                "deprecated": none,
                                "id": "record_only",
                                "note": none,
                                "stability": "development",
                                "value": "RECORD_ONLY",
                            },
                            {
                                "brief": "The span is sampled and recording",
                                "deprecated": none,
                                "id": "record_and_sample",
                                "note": none,
                                "stability": "development",
                                "value": "RECORD_AND_SAMPLE",
                            },
                        ],
                    },
                },
                {
                    "brief": "Name of the code, either \"OK\" or \"ERROR\". MUST NOT be set if the status code is UNSET.",
                    "name": "otel.status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "The operation has been validated by an Application developer or Operator to have completed successfully.",
                                "deprecated": none,
                                "id": "ok",
                                "note": none,
                                "stability": "stable",
                                "value": "OK",
                            },
                            {
                                "brief": "The operation contains an error.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "stable",
                                "value": "ERROR",
                            },
                        ],
                    },
                },
                {
                    "brief": "Description of the Status if it has a value, otherwise not set.",
                    "examples": [
                        "resource not found",
                    ],
                    "name": "otel.status_description",
                    "requirement_level": "recommended",
                    "root_namespace": "otel",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "root_namespace": "otel",
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
            "attr.go.j2",
        ],
    },
} */
