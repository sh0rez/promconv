package rpc

// The [error codes] of the Connect request. Error codes are always string values
//
// [error codes]: https://connectrpc.com//docs/protocol/#error-codes
type AttrConnectRpcErrorCode string // rpc.connect_rpc.error_code

func (AttrConnectRpcErrorCode) Development() {}
func (AttrConnectRpcErrorCode) Recommended() {}

// Connect request metadata, `<key>` being the normalized Connect Metadata key (lowercase), the value being the metadata values.
//
// Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.
// Including all request metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// For example, a property `my-custom-key` with value `["1.2.3.4", "1.2.3.5"]` SHOULD be recorded as
// the `rpc.connect_rpc.request.metadata.my-custom-key` attribute with value `["1.2.3.4", "1.2.3.5"]`
type AttrConnectRpcRequestMetadata string // rpc.connect_rpc.request.metadata

func (AttrConnectRpcRequestMetadata) Development() {}
func (AttrConnectRpcRequestMetadata) Recommended() {}

// Connect response metadata, `<key>` being the normalized Connect Metadata key (lowercase), the value being the metadata values.
//
// Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.
// Including all response metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// For example, a property `my-custom-key` with value `"attribute_value"` SHOULD be recorded as
// the `rpc.connect_rpc.response.metadata.my-custom-key` attribute with value `["attribute_value"]`
type AttrConnectRpcResponseMetadata string // rpc.connect_rpc.response.metadata

func (AttrConnectRpcResponseMetadata) Development() {}
func (AttrConnectRpcResponseMetadata) Recommended() {}

// gRPC request metadata, `<key>` being the normalized gRPC Metadata key (lowercase), the value being the metadata values.
//
// Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.
// Including all request metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// For example, a property `my-custom-key` with value `["1.2.3.4", "1.2.3.5"]` SHOULD be recorded as
// `rpc.grpc.request.metadata.my-custom-key` attribute with value `["1.2.3.4", "1.2.3.5"]`
type AttrGrpcRequestMetadata string // rpc.grpc.request.metadata

func (AttrGrpcRequestMetadata) Development() {}
func (AttrGrpcRequestMetadata) Recommended() {}

// gRPC response metadata, `<key>` being the normalized gRPC Metadata key (lowercase), the value being the metadata values.
//
// Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.
// Including all response metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// For example, a property `my-custom-key` with value `["attribute_value"]` SHOULD be recorded as
// the `rpc.grpc.response.metadata.my-custom-key` attribute with value `["attribute_value"]`
type AttrGrpcResponseMetadata string // rpc.grpc.response.metadata

func (AttrGrpcResponseMetadata) Development() {}
func (AttrGrpcResponseMetadata) Recommended() {}

// The [numeric status code] of the gRPC request
//
// [numeric status code]: https://github.com/grpc/grpc/blob/v1.33.2/doc/statuscodes.md
type AttrGrpcStatusCode string // rpc.grpc.status_code

func (AttrGrpcStatusCode) Development() {}
func (AttrGrpcStatusCode) Recommended() {}

// `error.code` property of response if it is an error response
type AttrJsonrpcErrorCode string // rpc.jsonrpc.error_code

func (AttrJsonrpcErrorCode) Development() {}
func (AttrJsonrpcErrorCode) Recommended() {}

// `error.message` property of response if it is an error response
type AttrJsonrpcErrorMessage string // rpc.jsonrpc.error_message

func (AttrJsonrpcErrorMessage) Development() {}
func (AttrJsonrpcErrorMessage) Recommended() {}

// `id` property of request or response. Since protocol allows id to be int, string, `null` or missing (for notifications), value is expected to be cast to string for simplicity. Use empty string in case of `null` value. Omit entirely if this is a notification
type AttrJsonrpcRequestId string // rpc.jsonrpc.request_id

func (AttrJsonrpcRequestId) Development() {}
func (AttrJsonrpcRequestId) Recommended() {}

// Protocol version as in `jsonrpc` property of request/response. Since JSON-RPC 1.0 doesn't specify this, the value can be omitted
type AttrJsonrpcVersion string // rpc.jsonrpc.version

func (AttrJsonrpcVersion) Development() {}
func (AttrJsonrpcVersion) Recommended() {}

// Compressed size of the message in bytes
type AttrMessageCompressedSize string // rpc.message.compressed_size

func (AttrMessageCompressedSize) Development() {}
func (AttrMessageCompressedSize) Recommended() {}

// MUST be calculated as two different counters starting from `1` one for sent messages and one for received message.
// This way we guarantee that the values will be consistent between different implementations
type AttrMessageId string // rpc.message.id

func (AttrMessageId) Development() {}
func (AttrMessageId) Recommended() {}

// Whether this is a received or sent message
type AttrMessageType string // rpc.message.type

func (AttrMessageType) Development() {}
func (AttrMessageType) Recommended() {}

// Uncompressed size of the message in bytes
type AttrMessageUncompressedSize string // rpc.message.uncompressed_size

func (AttrMessageUncompressedSize) Development() {}
func (AttrMessageUncompressedSize) Recommended() {}

// The name of the (logical) method being called, must be equal to the $method part in the span name.
// This is the logical name of the method from the RPC interface perspective, which can be different from the name of any implementing method/function. The `code.function.name` attribute may be used to store the latter (e.g., method actually executing the call on the server side, RPC client stub method on the client side)
type AttrMethod string // rpc.method

func (AttrMethod) Development() {}
func (AttrMethod) Recommended() {}

// The full (logical) name of the service being called, including its package name, if applicable.
// This is the logical name of the service from the RPC interface perspective, which can be different from the name of any implementing class. The `code.namespace` attribute may be used to store the latter (despite the attribute name, it may include a class name; e.g., class with method actually executing the call on the server side, RPC client stub class on the client side)
type AttrService string // rpc.service

func (AttrService) Development() {}
func (AttrService) Recommended() {}

// A string identifying the remoting system. See below for a list of well-known identifiers
type AttrSystem string // rpc.system

func (AttrSystem) Development() {}
func (AttrSystem) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The [error codes](https://connectrpc.com//docs/protocol/#error-codes) of the Connect request. Error codes are always string values.",
                    "name": "rpc.connect_rpc.error_code",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "cancelled",
                                "note": none,
                                "stability": "development",
                                "value": "cancelled",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unknown",
                                "note": none,
                                "stability": "development",
                                "value": "unknown",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "invalid_argument",
                                "note": none,
                                "stability": "development",
                                "value": "invalid_argument",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "deadline_exceeded",
                                "note": none,
                                "stability": "development",
                                "value": "deadline_exceeded",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "not_found",
                                "note": none,
                                "stability": "development",
                                "value": "not_found",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "already_exists",
                                "note": none,
                                "stability": "development",
                                "value": "already_exists",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "permission_denied",
                                "note": none,
                                "stability": "development",
                                "value": "permission_denied",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "resource_exhausted",
                                "note": none,
                                "stability": "development",
                                "value": "resource_exhausted",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "failed_precondition",
                                "note": none,
                                "stability": "development",
                                "value": "failed_precondition",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "aborted",
                                "note": none,
                                "stability": "development",
                                "value": "aborted",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "out_of_range",
                                "note": none,
                                "stability": "development",
                                "value": "out_of_range",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unimplemented",
                                "note": none,
                                "stability": "development",
                                "value": "unimplemented",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "internal",
                                "note": none,
                                "stability": "development",
                                "value": "internal",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unavailable",
                                "note": none,
                                "stability": "development",
                                "value": "unavailable",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "data_loss",
                                "note": none,
                                "stability": "development",
                                "value": "data_loss",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "unauthenticated",
                                "note": none,
                                "stability": "development",
                                "value": "unauthenticated",
                            },
                        ],
                    },
                },
                {
                    "brief": "Connect request metadata, `<key>` being the normalized Connect Metadata key (lowercase), the value being the metadata values.\n",
                    "examples": [
                        [
                            "1.2.3.4",
                            "1.2.3.5",
                        ],
                    ],
                    "name": "rpc.connect_rpc.request.metadata",
                    "note": "Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.\nIncluding all request metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nFor example, a property `my-custom-key` with value `[\"1.2.3.4\", \"1.2.3.5\"]` SHOULD be recorded as\nthe `rpc.connect_rpc.request.metadata.my-custom-key` attribute with value `[\"1.2.3.4\", \"1.2.3.5\"]`\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "template[string[]]",
                },
                {
                    "brief": "Connect response metadata, `<key>` being the normalized Connect Metadata key (lowercase), the value being the metadata values.\n",
                    "examples": [
                        [
                            "attribute_value",
                        ],
                    ],
                    "name": "rpc.connect_rpc.response.metadata",
                    "note": "Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.\nIncluding all response metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nFor example, a property `my-custom-key` with value `\"attribute_value\"` SHOULD be recorded as\nthe `rpc.connect_rpc.response.metadata.my-custom-key` attribute with value `[\"attribute_value\"]`\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "template[string[]]",
                },
                {
                    "brief": "gRPC request metadata, `<key>` being the normalized gRPC Metadata key (lowercase), the value being the metadata values.\n",
                    "examples": [
                        [
                            "1.2.3.4",
                            "1.2.3.5",
                        ],
                    ],
                    "name": "rpc.grpc.request.metadata",
                    "note": "Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.\nIncluding all request metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nFor example, a property `my-custom-key` with value `[\"1.2.3.4\", \"1.2.3.5\"]` SHOULD be recorded as\n`rpc.grpc.request.metadata.my-custom-key` attribute with value `[\"1.2.3.4\", \"1.2.3.5\"]`\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "template[string[]]",
                },
                {
                    "brief": "gRPC response metadata, `<key>` being the normalized gRPC Metadata key (lowercase), the value being the metadata values.\n",
                    "examples": [
                        [
                            "attribute_value",
                        ],
                    ],
                    "name": "rpc.grpc.response.metadata",
                    "note": "Instrumentations SHOULD require an explicit configuration of which metadata values are to be captured.\nIncluding all response metadata values can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nFor example, a property `my-custom-key` with value `[\"attribute_value\"]` SHOULD be recorded as\nthe `rpc.grpc.response.metadata.my-custom-key` attribute with value `[\"attribute_value\"]`\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "template[string[]]",
                },
                {
                    "brief": "The [numeric status code](https://github.com/grpc/grpc/blob/v1.33.2/doc/statuscodes.md) of the gRPC request.",
                    "name": "rpc.grpc.status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
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
                    "brief": "`error.code` property of response if it is an error response.",
                    "examples": [
                        -32700,
                        100,
                    ],
                    "name": "rpc.jsonrpc.error_code",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "`error.message` property of response if it is an error response.",
                    "examples": [
                        "Parse error",
                        "User already exists",
                    ],
                    "name": "rpc.jsonrpc.error_message",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "`id` property of request or response. Since protocol allows id to be int, string, `null` or missing (for notifications), value is expected to be cast to string for simplicity. Use empty string in case of `null` value. Omit entirely if this is a notification.\n",
                    "examples": [
                        "10",
                        "request-7",
                        "",
                    ],
                    "name": "rpc.jsonrpc.request_id",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Protocol version as in `jsonrpc` property of request/response. Since JSON-RPC 1.0 doesn't specify this, the value can be omitted.",
                    "examples": [
                        "2.0",
                        "1.0",
                    ],
                    "name": "rpc.jsonrpc.version",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Compressed size of the message in bytes.",
                    "name": "rpc.message.compressed_size",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "MUST be calculated as two different counters starting from `1` one for sent messages and one for received message.",
                    "name": "rpc.message.id",
                    "note": "This way we guarantee that the values will be consistent between different implementations.",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Whether this is a received or sent message.",
                    "name": "rpc.message.type",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "sent",
                                "note": none,
                                "stability": "development",
                                "value": "SENT",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "received",
                                "note": none,
                                "stability": "development",
                                "value": "RECEIVED",
                            },
                        ],
                    },
                },
                {
                    "brief": "Uncompressed size of the message in bytes.",
                    "name": "rpc.message.uncompressed_size",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The name of the (logical) method being called, must be equal to the $method part in the span name.",
                    "examples": "exampleMethod",
                    "name": "rpc.method",
                    "note": "This is the logical name of the method from the RPC interface perspective, which can be different from the name of any implementing method/function. The `code.function.name` attribute may be used to store the latter (e.g., method actually executing the call on the server side, RPC client stub method on the client side).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full (logical) name of the service being called, including its package name, if applicable.",
                    "examples": "myservice.EchoService",
                    "name": "rpc.service",
                    "note": "This is the logical name of the service from the RPC interface perspective, which can be different from the name of any implementing class. The `code.namespace` attribute may be used to store the latter (despite the attribute name, it may include a class name; e.g., class with method actually executing the call on the server side, RPC client stub class on the client side).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A string identifying the remoting system. See below for a list of well-known identifiers.",
                    "name": "rpc.system",
                    "requirement_level": "recommended",
                    "root_namespace": "rpc",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "gRPC",
                                "deprecated": none,
                                "id": "grpc",
                                "note": none,
                                "stability": "development",
                                "value": "grpc",
                            },
                            {
                                "brief": "Java RMI",
                                "deprecated": none,
                                "id": "java_rmi",
                                "note": none,
                                "stability": "development",
                                "value": "java_rmi",
                            },
                            {
                                "brief": ".NET WCF",
                                "deprecated": none,
                                "id": "dotnet_wcf",
                                "note": none,
                                "stability": "development",
                                "value": "dotnet_wcf",
                            },
                            {
                                "brief": "Apache Dubbo",
                                "deprecated": none,
                                "id": "apache_dubbo",
                                "note": none,
                                "stability": "development",
                                "value": "apache_dubbo",
                            },
                            {
                                "brief": "Connect RPC",
                                "deprecated": none,
                                "id": "connect_rpc",
                                "note": none,
                                "stability": "development",
                                "value": "connect_rpc",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "rpc",
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
            "attr.go.j2",
        ],
    },
} */
