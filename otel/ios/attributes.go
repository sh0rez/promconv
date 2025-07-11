package ios

// This attribute represents the state of the application.
//
// The iOS lifecycle states are defined in the [UIApplicationDelegate documentation], and from which the `OS terminology` column values are derived
//
// [UIApplicationDelegate documentation]: https://developer.apple.com/documentation/uikit/uiapplicationdelegate
type AttrAppState string // ios.app.state

func (AttrAppState) Development()    {}
func (AttrAppState) Recommended()    {}
func (AttrAppState) Key() string     { return "ios_app_state" }
func (a AttrAppState) Value() string { return string(a) }

const AppStateActive AttrAppState = "active"
const AppStateInactive AttrAppState = "inactive"
const AppStateBackground AttrAppState = "background"
const AppStateForeground AttrAppState = "foreground"
const AppStateTerminate AttrAppState = "terminate"

// The iOS lifecycle states are defined in the [UIApplicationDelegate documentation], and from which the `OS terminology` column values are derived
//
// [UIApplicationDelegate documentation]: https://developer.apple.com/documentation/uikit/uiapplicationdelegate
type AttrState string // ios.state

func (AttrState) Development()    {}
func (AttrState) Recommended()    {}
func (AttrState) Key() string     { return "ios_state" }
func (a AttrState) Value() string { return string(a) }

const StateActive AttrState = "active"
const StateInactive AttrState = "inactive"
const StateBackground AttrState = "background"
const StateForeground AttrState = "foreground"
const StateTerminate AttrState = "terminate"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "This attribute represents the state of the application.\n",
                    "name": "ios.app.state",
                    "note": "The iOS lifecycle states are defined in the [UIApplicationDelegate documentation](https://developer.apple.com/documentation/uikit/uiapplicationdelegate), and from which the `OS terminology` column values are derived.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "ios",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The app has become `active`. Associated with UIKit notification `applicationDidBecomeActive`.\n",
                                "deprecated": none,
                                "id": "active",
                                "note": none,
                                "stability": "development",
                                "value": "active",
                            },
                            {
                                "brief": "The app is now `inactive`. Associated with UIKit notification `applicationWillResignActive`.\n",
                                "deprecated": none,
                                "id": "inactive",
                                "note": none,
                                "stability": "development",
                                "value": "inactive",
                            },
                            {
                                "brief": "The app is now in the background. This value is associated with UIKit notification `applicationDidEnterBackground`.\n",
                                "deprecated": none,
                                "id": "background",
                                "note": none,
                                "stability": "development",
                                "value": "background",
                            },
                            {
                                "brief": "The app is now in the foreground. This value is associated with UIKit notification `applicationWillEnterForeground`.\n",
                                "deprecated": none,
                                "id": "foreground",
                                "note": none,
                                "stability": "development",
                                "value": "foreground",
                            },
                            {
                                "brief": "The app is about to terminate. Associated with UIKit notification `applicationWillTerminate`.\n",
                                "deprecated": none,
                                "id": "terminate",
                                "note": none,
                                "stability": "development",
                                "value": "terminate",
                            },
                        ],
                    },
                },
                {
                    "deprecated": {
                        "note": "Replaced by the `ios.app.state` event body field.",
                        "reason": "uncategorized",
                    },
                    "name": "ios.state",
                    "note": "The iOS lifecycle states are defined in the [UIApplicationDelegate documentation](https://developer.apple.com/documentation/uikit/uiapplicationdelegate), and from which the `OS terminology` column values are derived.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "ios",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "The app has become `active`. Associated with UIKit notification `applicationDidBecomeActive`.\n",
                                "deprecated": none,
                                "id": "active",
                                "note": none,
                                "stability": "development",
                                "value": "active",
                            },
                            {
                                "brief": "The app is now `inactive`. Associated with UIKit notification `applicationWillResignActive`.\n",
                                "deprecated": none,
                                "id": "inactive",
                                "note": none,
                                "stability": "development",
                                "value": "inactive",
                            },
                            {
                                "brief": "The app is now in the background. This value is associated with UIKit notification `applicationDidEnterBackground`.\n",
                                "deprecated": none,
                                "id": "background",
                                "note": none,
                                "stability": "development",
                                "value": "background",
                            },
                            {
                                "brief": "The app is now in the foreground. This value is associated with UIKit notification `applicationWillEnterForeground`.\n",
                                "deprecated": none,
                                "id": "foreground",
                                "note": none,
                                "stability": "development",
                                "value": "foreground",
                            },
                            {
                                "brief": "The app is about to terminate. Associated with UIKit notification `applicationWillTerminate`.\n",
                                "deprecated": none,
                                "id": "terminate",
                                "note": none,
                                "stability": "development",
                                "value": "terminate",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "ios",
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
