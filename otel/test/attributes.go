package test

// The fully qualified human readable name of the [test case]
//
// [test case]: https://wikipedia.org/wiki/Test_case
type AttrCaseName string // test.case.name

func (AttrCaseName) Development()    {}
func (AttrCaseName) Recommended()    {}
func (AttrCaseName) Key() string     { return "test_case_name" }
func (a AttrCaseName) Value() string { return string(a) }

// The status of the actual test case result from test execution
type AttrCaseResultStatus string // test.case.result.status

func (AttrCaseResultStatus) Development()    {}
func (AttrCaseResultStatus) Recommended()    {}
func (AttrCaseResultStatus) Key() string     { return "test_case_result_status" }
func (a AttrCaseResultStatus) Value() string { return string(a) }

const CaseResultStatusPass AttrCaseResultStatus = "pass"
const CaseResultStatusFail AttrCaseResultStatus = "fail"

// The human readable name of a [test suite]
//
// [test suite]: https://wikipedia.org/wiki/Test_suite
type AttrSuiteName string // test.suite.name

func (AttrSuiteName) Development()    {}
func (AttrSuiteName) Recommended()    {}
func (AttrSuiteName) Key() string     { return "test_suite_name" }
func (a AttrSuiteName) Value() string { return string(a) }

// The status of the test suite run
type AttrSuiteRunStatus string // test.suite.run.status

func (AttrSuiteRunStatus) Development()    {}
func (AttrSuiteRunStatus) Recommended()    {}
func (AttrSuiteRunStatus) Key() string     { return "test_suite_run_status" }
func (a AttrSuiteRunStatus) Value() string { return string(a) }

const SuiteRunStatusSuccess AttrSuiteRunStatus = "success"
const SuiteRunStatusFailure AttrSuiteRunStatus = "failure"
const SuiteRunStatusSkipped AttrSuiteRunStatus = "skipped"
const SuiteRunStatusAborted AttrSuiteRunStatus = "aborted"
const SuiteRunStatusTimedOut AttrSuiteRunStatus = "timed_out"
const SuiteRunStatusInProgress AttrSuiteRunStatus = "in_progress"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The fully qualified human readable name of the [test case](https://wikipedia.org/wiki/Test_case).\n",
                    "examples": [
                        "org.example.TestCase1.test1",
                        "example/tests/TestCase1.test1",
                        "ExampleTestCase1_test1",
                    ],
                    "name": "test.case.name",
                    "requirement_level": "recommended",
                    "root_namespace": "test",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The status of the actual test case result from test execution.\n",
                    "examples": [
                        "pass",
                        "fail",
                    ],
                    "name": "test.case.result.status",
                    "requirement_level": "recommended",
                    "root_namespace": "test",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "pass",
                                "deprecated": none,
                                "id": "pass",
                                "note": none,
                                "stability": "development",
                                "value": "pass",
                            },
                            {
                                "brief": "fail",
                                "deprecated": none,
                                "id": "fail",
                                "note": none,
                                "stability": "development",
                                "value": "fail",
                            },
                        ],
                    },
                },
                {
                    "brief": "The human readable name of a [test suite](https://wikipedia.org/wiki/Test_suite).\n",
                    "examples": [
                        "TestSuite1",
                    ],
                    "name": "test.suite.name",
                    "requirement_level": "recommended",
                    "root_namespace": "test",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The status of the test suite run.\n",
                    "examples": [
                        "success",
                        "failure",
                        "skipped",
                        "aborted",
                        "timed_out",
                        "in_progress",
                    ],
                    "name": "test.suite.run.status",
                    "requirement_level": "recommended",
                    "root_namespace": "test",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "success",
                                "deprecated": none,
                                "id": "success",
                                "note": none,
                                "stability": "development",
                                "value": "success",
                            },
                            {
                                "brief": "failure",
                                "deprecated": none,
                                "id": "failure",
                                "note": none,
                                "stability": "development",
                                "value": "failure",
                            },
                            {
                                "brief": "skipped",
                                "deprecated": none,
                                "id": "skipped",
                                "note": none,
                                "stability": "development",
                                "value": "skipped",
                            },
                            {
                                "brief": "aborted",
                                "deprecated": none,
                                "id": "aborted",
                                "note": none,
                                "stability": "development",
                                "value": "aborted",
                            },
                            {
                                "brief": "timed_out",
                                "deprecated": none,
                                "id": "timed_out",
                                "note": none,
                                "stability": "development",
                                "value": "timed_out",
                            },
                            {
                                "brief": "in_progress",
                                "deprecated": none,
                                "id": "in_progress",
                                "note": none,
                                "stability": "development",
                                "value": "in_progress",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "test",
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
