package azure

// The unique identifier of the client instance
type AttrClientId string // azure.client.id

func (AttrClientId) Development() {}
func (AttrClientId) Recommended() {}

// Cosmos client connection mode
type AttrCosmosdbConnectionMode string // azure.cosmosdb.connection.mode

func (AttrCosmosdbConnectionMode) Development() {}
func (AttrCosmosdbConnectionMode) Recommended() {}

// Account or request [consistency level]
//
// [consistency level]: https://learn.microsoft.com/azure/cosmos-db/consistency-levels
type AttrCosmosdbConsistencyLevel string // azure.cosmosdb.consistency.level

func (AttrCosmosdbConsistencyLevel) Development() {}
func (AttrCosmosdbConsistencyLevel) Recommended() {}

// List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.
//
// Region name matches the format of `displayName` in [Azure Location API]
//
// [Azure Location API]: https://learn.microsoft.com/rest/api/subscription/subscriptions/list-locations?view=rest-subscription-2021-10-01&tabs=HTTP#location
type AttrCosmosdbOperationContactedRegions string // azure.cosmosdb.operation.contacted_regions

func (AttrCosmosdbOperationContactedRegions) Development() {}
func (AttrCosmosdbOperationContactedRegions) Recommended() {}

// The number of request units consumed by the operation
type AttrCosmosdbOperationRequestCharge string // azure.cosmosdb.operation.request_charge

func (AttrCosmosdbOperationRequestCharge) Development() {}
func (AttrCosmosdbOperationRequestCharge) Recommended() {}

// Request payload size in bytes
type AttrCosmosdbRequestBodySize string // azure.cosmosdb.request.body.size

func (AttrCosmosdbRequestBodySize) Development() {}
func (AttrCosmosdbRequestBodySize) Recommended() {}

// Cosmos DB sub status code
type AttrCosmosdbResponseSubStatusCode string // azure.cosmosdb.response.sub_status_code

func (AttrCosmosdbResponseSubStatusCode) Development() {}
func (AttrCosmosdbResponseSubStatusCode) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The unique identifier of the client instance.",
                    "examples": [
                        "3ba4827d-4422-483f-b59f-85b74211c11d",
                        "storage-client-1",
                    ],
                    "name": "azure.client.id",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Cosmos client connection mode.",
                    "name": "azure.cosmosdb.connection.mode",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Gateway (HTTP) connection.",
                                "deprecated": none,
                                "id": "gateway",
                                "note": none,
                                "stability": "development",
                                "value": "gateway",
                            },
                            {
                                "brief": "Direct connection.",
                                "deprecated": none,
                                "id": "direct",
                                "note": none,
                                "stability": "development",
                                "value": "direct",
                            },
                        ],
                    },
                },
                {
                    "brief": "Account or request [consistency level](https://learn.microsoft.com/azure/cosmos-db/consistency-levels).",
                    "examples": [
                        "Eventual",
                        "ConsistentPrefix",
                        "BoundedStaleness",
                        "Strong",
                        "Session",
                    ],
                    "name": "azure.cosmosdb.consistency.level",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "strong",
                                "note": none,
                                "stability": "development",
                                "value": "Strong",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "bounded_staleness",
                                "note": none,
                                "stability": "development",
                                "value": "BoundedStaleness",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "session",
                                "note": none,
                                "stability": "development",
                                "value": "Session",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "eventual",
                                "note": none,
                                "stability": "development",
                                "value": "Eventual",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "consistent_prefix",
                                "note": none,
                                "stability": "development",
                                "value": "ConsistentPrefix",
                            },
                        ],
                    },
                },
                {
                    "brief": "List of regions contacted during operation in the order that they were contacted. If there is more than one region listed, it indicates that the operation was performed on multiple regions i.e. cross-regional call.\n",
                    "examples": [
                        [
                            "North Central US",
                            "Australia East",
                            "Australia Southeast",
                        ],
                    ],
                    "name": "azure.cosmosdb.operation.contacted_regions",
                    "note": "Region name matches the format of `displayName` in [Azure Location API](https://learn.microsoft.com/rest/api/subscription/subscriptions/list-locations?view=rest-subscription-2021-10-01&tabs=HTTP#location)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The number of request units consumed by the operation.\n",
                    "examples": [
                        46.18,
                        1.0,
                    ],
                    "name": "azure.cosmosdb.operation.request_charge",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Request payload size in bytes.",
                    "name": "azure.cosmosdb.request.body.size",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Cosmos DB sub status code.",
                    "examples": [
                        1000,
                        1002,
                    ],
                    "name": "azure.cosmosdb.response.sub_status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "azure",
                    "stability": "development",
                    "type": "int",
                },
            ],
            "root_namespace": "azure",
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
