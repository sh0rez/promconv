package android

// This attribute represents the state of the application.
//
// The Android lifecycle states are defined in [Activity lifecycle callbacks], and from which the `OS identifiers` are derived
//
// [Activity lifecycle callbacks]: https://developer.android.com/guide/components/activities/activity-lifecycle#lc
type AttrAppState string // android.app.state

func (AttrAppState) Development()    {}
func (AttrAppState) Recommended()    {}
func (AttrAppState) Key() string     { return "android_app_state" }
func (a AttrAppState) Value() string { return string(a) }

const AppStateCreated AttrAppState = "created"
const AppStateBackground AttrAppState = "background"
const AppStateForeground AttrAppState = "foreground"

// Uniquely identifies the framework API revision offered by a version (`os.version`) of the android operating system. More information can be found [here]
//
// [here]: https://developer.android.com/guide/topics/manifest/uses-sdk-element#ApiLevels
type AttrOsApiLevel string // android.os.api_level

func (AttrOsApiLevel) Development()    {}
func (AttrOsApiLevel) Recommended()    {}
func (AttrOsApiLevel) Key() string     { return "android_os_api_level" }
func (a AttrOsApiLevel) Value() string { return string(a) }

// Deprecated. Use `android.app.state` body field instead
type AttrState string // android.state

func (AttrState) Development()    {}
func (AttrState) Recommended()    {}
func (AttrState) Key() string     { return "android_state" }
func (a AttrState) Value() string { return string(a) }

const StateCreated AttrState = "created"
const StateBackground AttrState = "background"
const StateForeground AttrState = "foreground"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "This attribute represents the state of the application.\n",
                    "examples": [
                        "created",
                    ],
                    "name": "android.app.state",
                    "note": "The Android lifecycle states are defined in [Activity lifecycle callbacks](https://developer.android.com/guide/components/activities/activity-lifecycle#lc), and from which the `OS identifiers` are derived.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "android",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Any time before Activity.onResume() or, if the app has no Activity, Context.startService() has been called in the app for the first time.\n",
                                "deprecated": none,
                                "id": "created",
                                "note": none,
                                "stability": "development",
                                "value": "created",
                            },
                            {
                                "brief": "Any time after Activity.onPause() or, if the app has no Activity, Context.stopService() has been called when the app was in the foreground state.\n",
                                "deprecated": none,
                                "id": "background",
                                "note": none,
                                "stability": "development",
                                "value": "background",
                            },
                            {
                                "brief": "Any time after Activity.onResume() or, if the app has no Activity, Context.startService() has been called when the app was in either the created or background states.\n",
                                "deprecated": none,
                                "id": "foreground",
                                "note": none,
                                "stability": "development",
                                "value": "foreground",
                            },
                        ],
                    },
                },
                {
                    "brief": "Uniquely identifies the framework API revision offered by a version (`os.version`) of the android operating system. More information can be found [here](https://developer.android.com/guide/topics/manifest/uses-sdk-element#ApiLevels).\n",
                    "examples": [
                        "33",
                        "32",
                    ],
                    "name": "android.os.api_level",
                    "requirement_level": "recommended",
                    "root_namespace": "android",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated. Use `android.app.state` body field instead.",
                    "deprecated": {
                        "note": "Use `android.app.state` body field instead.",
                        "reason": "uncategorized",
                    },
                    "name": "android.state",
                    "requirement_level": "recommended",
                    "root_namespace": "android",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Any time before Activity.onResume() or, if the app has no Activity, Context.startService() has been called in the app for the first time.\n",
                                "deprecated": none,
                                "id": "created",
                                "note": none,
                                "stability": "development",
                                "value": "created",
                            },
                            {
                                "brief": "Any time after Activity.onPause() or, if the app has no Activity, Context.stopService() has been called when the app was in the foreground state.\n",
                                "deprecated": none,
                                "id": "background",
                                "note": none,
                                "stability": "development",
                                "value": "background",
                            },
                            {
                                "brief": "Any time after Activity.onResume() or, if the app has no Activity, Context.startService() has been called when the app was in either the created or background states.\n",
                                "deprecated": none,
                                "id": "foreground",
                                "note": none,
                                "stability": "development",
                                "value": "foreground",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "android",
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
