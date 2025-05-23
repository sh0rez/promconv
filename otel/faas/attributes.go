package faas

// A boolean that is true if the serverless function is executed for the first time (aka cold-start)
type AttrColdstart string // faas.coldstart

func (AttrColdstart) Development() {}
func (AttrColdstart) Recommended() {}

// A string containing the schedule period as [Cron Expression]
//
// [Cron Expression]: https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expressions.htm
type AttrCron string // faas.cron

func (AttrCron) Development() {}
func (AttrCron) Recommended() {}

// The name of the source on which the triggering operation was performed. For example, in Cloud Storage or S3 corresponds to the bucket name, and in Cosmos DB to the database name
type AttrDocumentCollection string // faas.document.collection

func (AttrDocumentCollection) Development() {}
func (AttrDocumentCollection) Recommended() {}

// The document name/table subjected to the operation. For example, in Cloud Storage or S3 is the name of the file, and in Cosmos DB the table name
type AttrDocumentName string // faas.document.name

func (AttrDocumentName) Development() {}
func (AttrDocumentName) Recommended() {}

// Describes the type of the operation that was performed on the data
type AttrDocumentOperation string // faas.document.operation

func (AttrDocumentOperation) Development() {}
func (AttrDocumentOperation) Recommended() {}

// A string containing the time when the data was accessed in the [ISO 8601] format expressed in [UTC]
//
// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
// [UTC]: https://www.w3.org/TR/NOTE-datetime
type AttrDocumentTime string // faas.document.time

func (AttrDocumentTime) Development() {}
func (AttrDocumentTime) Recommended() {}

// The execution environment ID as a string, that will be potentially reused for other invocations to the same function/function version.
//
//   - **AWS Lambda:** Use the (full) log stream name
type AttrInstance string // faas.instance

func (AttrInstance) Development() {}
func (AttrInstance) Recommended() {}

// The invocation ID of the current function invocation
type AttrInvocationId string // faas.invocation_id

func (AttrInvocationId) Development() {}
func (AttrInvocationId) Recommended() {}

// The name of the invoked function.
//
// SHOULD be equal to the `faas.name` resource attribute of the invoked function
type AttrInvokedName string // faas.invoked_name

func (AttrInvokedName) Development() {}
func (AttrInvokedName) Recommended() {}

// The cloud provider of the invoked function.
//
// SHOULD be equal to the `cloud.provider` resource attribute of the invoked function
type AttrInvokedProvider string // faas.invoked_provider

func (AttrInvokedProvider) Development() {}
func (AttrInvokedProvider) Recommended() {}

// The cloud region of the invoked function.
//
// SHOULD be equal to the `cloud.region` resource attribute of the invoked function
type AttrInvokedRegion string // faas.invoked_region

func (AttrInvokedRegion) Development() {}
func (AttrInvokedRegion) Recommended() {}

// The amount of memory available to the serverless function converted to Bytes.
//
// It's recommended to set this attribute since e.g. too little memory can easily stop a Java AWS Lambda function from working correctly. On AWS Lambda, the environment variable `AWS_LAMBDA_FUNCTION_MEMORY_SIZE` provides this information (which must be multiplied by 1,048,576)
type AttrMaxMemory string // faas.max_memory

func (AttrMaxMemory) Development() {}
func (AttrMaxMemory) Recommended() {}

// The name of the single function that this runtime instance executes.
//
// This is the name of the function as configured/deployed on the FaaS
// platform and is usually different from the name of the callback
// function (which may be stored in the
// [`code.namespace`/`code.function.name`]
// span attributes).
//
// For some cloud providers, the above definition is ambiguous. The following
// definition of function name MUST be used for this attribute
// (and consequently the span name) for the listed cloud providers/products:
//
//   - **Azure:**  The full name `<FUNCAPP>/<FUNC>`, i.e., function app name
//     followed by a forward slash followed by the function name (this form
//     can also be seen in the resource JSON for the function).
//     This means that a span attribute MUST be used, as an Azure function
//     app can host multiple functions that would usually share
//     a TracerProvider (see also the `cloud.resource_id` attribute)
//
// [`code.namespace`/`code.function.name`]: /docs/general/attributes.md#source-code-attributes
type AttrName string // faas.name

func (AttrName) Development() {}
func (AttrName) Recommended() {}

// A string containing the function invocation time in the [ISO 8601] format expressed in [UTC]
//
// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
// [UTC]: https://www.w3.org/TR/NOTE-datetime
type AttrTime string // faas.time

func (AttrTime) Development() {}
func (AttrTime) Recommended() {}

// Type of the trigger which caused this function invocation
type AttrTrigger string // faas.trigger

func (AttrTrigger) Development() {}
func (AttrTrigger) Recommended() {}

// The immutable version of the function being executed.
// Depending on the cloud provider and platform, use:
//
//   - **AWS Lambda:** The [function version]
//     (an integer represented as a decimal string).
//   - **Google Cloud Run (Services):** The [revision]
//     (i.e., the function name plus the revision suffix).
//   - **Google Cloud Functions:** The value of the
//     [`K_REVISION` environment variable].
//   - **Azure Functions:** Not applicable. Do not set this attribute
//
// [function version]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html
// [revision]: https://cloud.google.com/run/docs/managing/revisions
// [`K_REVISION` environment variable]: https://cloud.google.com/functions/docs/env-var#runtime_environment_variables_set_automatically
type AttrVersion string // faas.version

func (AttrVersion) Development() {}
func (AttrVersion) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "A boolean that is true if the serverless function is executed for the first time (aka cold-start).\n",
                    "name": "faas.coldstart",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "A string containing the schedule period as [Cron Expression](https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expressions.htm).\n",
                    "examples": "0/5 * * * ? *",
                    "name": "faas.cron",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the source on which the triggering operation was performed. For example, in Cloud Storage or S3 corresponds to the bucket name, and in Cosmos DB to the database name.\n",
                    "examples": [
                        "myBucketName",
                        "myDbName",
                    ],
                    "name": "faas.document.collection",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The document name/table subjected to the operation. For example, in Cloud Storage or S3 is the name of the file, and in Cosmos DB the table name.\n",
                    "examples": [
                        "myFile.txt",
                        "myTableName",
                    ],
                    "name": "faas.document.name",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Describes the type of the operation that was performed on the data.",
                    "name": "faas.document.operation",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "When a new object is created.",
                                "deprecated": none,
                                "id": "insert",
                                "note": none,
                                "stability": "development",
                                "value": "insert",
                            },
                            {
                                "brief": "When an object is modified.",
                                "deprecated": none,
                                "id": "edit",
                                "note": none,
                                "stability": "development",
                                "value": "edit",
                            },
                            {
                                "brief": "When an object is deleted.",
                                "deprecated": none,
                                "id": "delete",
                                "note": none,
                                "stability": "development",
                                "value": "delete",
                            },
                        ],
                    },
                },
                {
                    "brief": "A string containing the time when the data was accessed in the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).\n",
                    "examples": "2020-01-23T13:47:06Z",
                    "name": "faas.document.time",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The execution environment ID as a string, that will be potentially reused for other invocations to the same function/function version.\n",
                    "examples": [
                        "2021/06/28/[$LATEST]2f399eb14537447da05ab2a2e39309de",
                    ],
                    "name": "faas.instance",
                    "note": "- **AWS Lambda:** Use the (full) log stream name.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The invocation ID of the current function invocation.\n",
                    "examples": "af9d5aa4-a685-4c5f-a22b-444f80b3cc28",
                    "name": "faas.invocation_id",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the invoked function.\n",
                    "examples": "my-function",
                    "name": "faas.invoked_name",
                    "note": "SHOULD be equal to the `faas.name` resource attribute of the invoked function.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The cloud provider of the invoked function.\n",
                    "name": "faas.invoked_provider",
                    "note": "SHOULD be equal to the `cloud.provider` resource attribute of the invoked function.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Alibaba Cloud",
                                "deprecated": none,
                                "id": "alibaba_cloud",
                                "note": none,
                                "stability": "development",
                                "value": "alibaba_cloud",
                            },
                            {
                                "brief": "Amazon Web Services",
                                "deprecated": none,
                                "id": "aws",
                                "note": none,
                                "stability": "development",
                                "value": "aws",
                            },
                            {
                                "brief": "Microsoft Azure",
                                "deprecated": none,
                                "id": "azure",
                                "note": none,
                                "stability": "development",
                                "value": "azure",
                            },
                            {
                                "brief": "Google Cloud Platform",
                                "deprecated": none,
                                "id": "gcp",
                                "note": none,
                                "stability": "development",
                                "value": "gcp",
                            },
                            {
                                "brief": "Tencent Cloud",
                                "deprecated": none,
                                "id": "tencent_cloud",
                                "note": none,
                                "stability": "development",
                                "value": "tencent_cloud",
                            },
                        ],
                    },
                },
                {
                    "brief": "The cloud region of the invoked function.\n",
                    "examples": "eu-central-1",
                    "name": "faas.invoked_region",
                    "note": "SHOULD be equal to the `cloud.region` resource attribute of the invoked function.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The amount of memory available to the serverless function converted to Bytes.\n",
                    "examples": 134217728,
                    "name": "faas.max_memory",
                    "note": "It's recommended to set this attribute since e.g. too little memory can easily stop a Java AWS Lambda function from working correctly. On AWS Lambda, the environment variable `AWS_LAMBDA_FUNCTION_MEMORY_SIZE` provides this information (which must be multiplied by 1,048,576).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The name of the single function that this runtime instance executes.\n",
                    "examples": [
                        "my-function",
                        "myazurefunctionapp/some-function-name",
                    ],
                    "name": "faas.name",
                    "note": "This is the name of the function as configured/deployed on the FaaS\nplatform and is usually different from the name of the callback\nfunction (which may be stored in the\n[`code.namespace`/`code.function.name`](/docs/general/attributes.md#source-code-attributes)\nspan attributes).\n\nFor some cloud providers, the above definition is ambiguous. The following\ndefinition of function name MUST be used for this attribute\n(and consequently the span name) for the listed cloud providers/products:\n\n- **Azure:**  The full name `<FUNCAPP>/<FUNC>`, i.e., function app name\n  followed by a forward slash followed by the function name (this form\n  can also be seen in the resource JSON for the function).\n  This means that a span attribute MUST be used, as an Azure function\n  app can host multiple functions that would usually share\n  a TracerProvider (see also the `cloud.resource_id` attribute).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A string containing the function invocation time in the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format expressed in [UTC](https://www.w3.org/TR/NOTE-datetime).\n",
                    "examples": "2020-01-23T13:47:06Z",
                    "name": "faas.time",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Type of the trigger which caused this function invocation.\n",
                    "name": "faas.trigger",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                {
                    "brief": "The immutable version of the function being executed.",
                    "examples": [
                        "26",
                        "pinkfroid-00002",
                    ],
                    "name": "faas.version",
                    "note": "Depending on the cloud provider and platform, use:\n\n- **AWS Lambda:** The [function version](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html)\n  (an integer represented as a decimal string).\n- **Google Cloud Run (Services):** The [revision](https://cloud.google.com/run/docs/managing/revisions)\n  (i.e., the function name plus the revision suffix).\n- **Google Cloud Functions:** The value of the\n  [`K_REVISION` environment variable](https://cloud.google.com/functions/docs/env-var#runtime_environment_variables_set_automatically).\n- **Azure Functions:** Not applicable. Do not set this attribute.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "faas",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "faas",
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
