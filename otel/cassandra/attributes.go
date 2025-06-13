package cassandra

// The consistency level of the query. Based on consistency values from [CQL]
//
// [CQL]: https://docs.datastax.com/en/cassandra-oss/3.0/cassandra/dml/dmlConfigConsistency.html
type AttrConsistencyLevel string // cassandra.consistency.level

func (AttrConsistencyLevel) Development()    {}
func (AttrConsistencyLevel) Recommended()    {}
func (AttrConsistencyLevel) Key() string     { return "cassandra_consistency_level" }
func (a AttrConsistencyLevel) Value() string { return string(a) }

const ConsistencyLevelAll AttrConsistencyLevel = "all"
const ConsistencyLevelEachQuorum AttrConsistencyLevel = "each_quorum"
const ConsistencyLevelQuorum AttrConsistencyLevel = "quorum"
const ConsistencyLevelLocalQuorum AttrConsistencyLevel = "local_quorum"
const ConsistencyLevelOne AttrConsistencyLevel = "one"
const ConsistencyLevelTwo AttrConsistencyLevel = "two"
const ConsistencyLevelThree AttrConsistencyLevel = "three"
const ConsistencyLevelLocalOne AttrConsistencyLevel = "local_one"
const ConsistencyLevelAny AttrConsistencyLevel = "any"
const ConsistencyLevelSerial AttrConsistencyLevel = "serial"
const ConsistencyLevelLocalSerial AttrConsistencyLevel = "local_serial"

// The data center of the coordinating node for a query
type AttrCoordinatorDc string // cassandra.coordinator.dc

func (AttrCoordinatorDc) Development()    {}
func (AttrCoordinatorDc) Recommended()    {}
func (AttrCoordinatorDc) Key() string     { return "cassandra_coordinator_dc" }
func (a AttrCoordinatorDc) Value() string { return string(a) }

// The ID of the coordinating node for a query
type AttrCoordinatorId string // cassandra.coordinator.id

func (AttrCoordinatorId) Development()    {}
func (AttrCoordinatorId) Recommended()    {}
func (AttrCoordinatorId) Key() string     { return "cassandra_coordinator_id" }
func (a AttrCoordinatorId) Value() string { return string(a) }

// The fetch size used for paging, i.e. how many rows will be returned at once
type AttrPageSize string // cassandra.page.size

func (AttrPageSize) Development()    {}
func (AttrPageSize) Recommended()    {}
func (AttrPageSize) Key() string     { return "cassandra_page_size" }
func (a AttrPageSize) Value() string { return string(a) }

// Whether or not the query is idempotent
type AttrQueryIdempotent string // cassandra.query.idempotent

func (AttrQueryIdempotent) Development()    {}
func (AttrQueryIdempotent) Recommended()    {}
func (AttrQueryIdempotent) Key() string     { return "cassandra_query_idempotent" }
func (a AttrQueryIdempotent) Value() string { return string(a) }

// The number of times a query was speculatively executed. Not set or `0` if the query was not executed speculatively
type AttrSpeculativeExecutionCount string // cassandra.speculative_execution.count

func (AttrSpeculativeExecutionCount) Development()    {}
func (AttrSpeculativeExecutionCount) Recommended()    {}
func (AttrSpeculativeExecutionCount) Key() string     { return "cassandra_speculative_execution_count" }
func (a AttrSpeculativeExecutionCount) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The consistency level of the query. Based on consistency values from [CQL](https://docs.datastax.com/en/cassandra-oss/3.0/cassandra/dml/dmlConfigConsistency.html).\n",
                    "name": "cassandra.consistency.level",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "all",
                                "note": none,
                                "stability": "development",
                                "value": "all",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "each_quorum",
                                "note": none,
                                "stability": "development",
                                "value": "each_quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "quorum",
                                "note": none,
                                "stability": "development",
                                "value": "quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_quorum",
                                "note": none,
                                "stability": "development",
                                "value": "local_quorum",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "one",
                                "note": none,
                                "stability": "development",
                                "value": "one",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "two",
                                "note": none,
                                "stability": "development",
                                "value": "two",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "three",
                                "note": none,
                                "stability": "development",
                                "value": "three",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_one",
                                "note": none,
                                "stability": "development",
                                "value": "local_one",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "any",
                                "note": none,
                                "stability": "development",
                                "value": "any",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "serial",
                                "note": none,
                                "stability": "development",
                                "value": "serial",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "local_serial",
                                "note": none,
                                "stability": "development",
                                "value": "local_serial",
                            },
                        ],
                    },
                },
                {
                    "brief": "The data center of the coordinating node for a query.\n",
                    "examples": "us-west-2",
                    "name": "cassandra.coordinator.dc",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ID of the coordinating node for a query.\n",
                    "examples": "be13faa2-8574-4d71-926d-27f16cf8a7af",
                    "name": "cassandra.coordinator.id",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The fetch size used for paging, i.e. how many rows will be returned at once.\n",
                    "examples": [
                        5000,
                    ],
                    "name": "cassandra.page.size",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Whether or not the query is idempotent.\n",
                    "name": "cassandra.query.idempotent",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The number of times a query was speculatively executed. Not set or `0` if the query was not executed speculatively.\n",
                    "examples": [
                        0,
                        2,
                    ],
                    "name": "cassandra.speculative_execution.count",
                    "requirement_level": "recommended",
                    "root_namespace": "cassandra",
                    "stability": "development",
                    "type": "int",
                },
            ],
            "root_namespace": "cassandra",
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
