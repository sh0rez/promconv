package app

// A unique identifier representing the installation of an application on a specific device
//
// Its value SHOULD persist across launches of the same application installation, including through application upgrades.
// It SHOULD change if the application is uninstalled or if all applications of the vendor are uninstalled.
// Additionally, users might be able to reset this value (e.g. by clearing application data).
// If an app is installed multiple times on the same device (e.g. in different accounts on Android), each `app.installation.id` SHOULD have a different value.
// If multiple OpenTelemetry SDKs are used within the same application, they SHOULD use the same value for `app.installation.id`.
// Hardware IDs (e.g. serial number, IMEI, MAC address) MUST NOT be used as the `app.installation.id`.
//
// For iOS, this value SHOULD be equal to the [vendor identifier].
//
// For Android, examples of `app.installation.id` implementations include:
//
//   - [Firebase Installation ID].
//   - A globally unique UUID which is persisted across sessions in your application.
//   - [App set ID].
//   - [`Settings.getString(Settings.Secure.ANDROID_ID)`].
//
// More information about Android identifier best practices can be found [here]
//
// [vendor identifier]: https://developer.apple.com/documentation/uikit/uidevice/identifierforvendor
// [Firebase Installation ID]: https://firebase.google.com/docs/projects/manage-installations
// [App set ID]: https://developer.android.com/identity/app-set-id
// [`Settings.getString(Settings.Secure.ANDROID_ID)`]: https://developer.android.com/reference/android/provider/Settings.Secure#ANDROID_ID
// [here]: https://developer.android.com/training/articles/user-data-ids
type AttrInstallationId string // app.installation.id

func (AttrInstallationId) Development()    {}
func (AttrInstallationId) Recommended()    {}
func (AttrInstallationId) Key() string     { return "app_installation_id" }
func (a AttrInstallationId) Value() string { return string(a) }

// The x (horizontal) coordinate of a screen coordinate, in screen pixels
type AttrScreenCoordinateX string // app.screen.coordinate.x

func (AttrScreenCoordinateX) Development()    {}
func (AttrScreenCoordinateX) Recommended()    {}
func (AttrScreenCoordinateX) Key() string     { return "app_screen_coordinate_x" }
func (a AttrScreenCoordinateX) Value() string { return string(a) }

// The y (vertical) component of a screen coordinate, in screen pixels
type AttrScreenCoordinateY string // app.screen.coordinate.y

func (AttrScreenCoordinateY) Development()    {}
func (AttrScreenCoordinateY) Recommended()    {}
func (AttrScreenCoordinateY) Key() string     { return "app_screen_coordinate_y" }
func (a AttrScreenCoordinateY) Value() string { return string(a) }

// An identifier that uniquely differentiates this widget from other widgets in the same application.
//
// A widget is an application component, typically an on-screen visual GUI element
type AttrWidgetId string // app.widget.id

func (AttrWidgetId) Development()    {}
func (AttrWidgetId) Recommended()    {}
func (AttrWidgetId) Key() string     { return "app_widget_id" }
func (a AttrWidgetId) Value() string { return string(a) }

// The name of an application widget.
// A widget is an application component, typically an on-screen visual GUI element
type AttrWidgetName string // app.widget.name

func (AttrWidgetName) Development()    {}
func (AttrWidgetName) Recommended()    {}
func (AttrWidgetName) Key() string     { return "app_widget_name" }
func (a AttrWidgetName) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "A unique identifier representing the installation of an application on a specific device\n",
                    "examples": [
                        "2ab2916d-a51f-4ac8-80ee-45ac31a28092",
                    ],
                    "name": "app.installation.id",
                    "note": "Its value SHOULD persist across launches of the same application installation, including through application upgrades.\nIt SHOULD change if the application is uninstalled or if all applications of the vendor are uninstalled.\nAdditionally, users might be able to reset this value (e.g. by clearing application data).\nIf an app is installed multiple times on the same device (e.g. in different accounts on Android), each `app.installation.id` SHOULD have a different value.\nIf multiple OpenTelemetry SDKs are used within the same application, they SHOULD use the same value for `app.installation.id`.\nHardware IDs (e.g. serial number, IMEI, MAC address) MUST NOT be used as the `app.installation.id`.\n\nFor iOS, this value SHOULD be equal to the [vendor identifier](https://developer.apple.com/documentation/uikit/uidevice/identifierforvendor).\n\nFor Android, examples of `app.installation.id` implementations include:\n\n- [Firebase Installation ID](https://firebase.google.com/docs/projects/manage-installations).\n- A globally unique UUID which is persisted across sessions in your application.\n- [App set ID](https://developer.android.com/identity/app-set-id).\n- [`Settings.getString(Settings.Secure.ANDROID_ID)`](https://developer.android.com/reference/android/provider/Settings.Secure#ANDROID_ID).\n\nMore information about Android identifier best practices can be found [here](https://developer.android.com/training/articles/user-data-ids).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "app",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The x (horizontal) coordinate of a screen coordinate, in screen pixels.",
                    "examples": [
                        0,
                        131,
                    ],
                    "name": "app.screen.coordinate.x",
                    "requirement_level": "recommended",
                    "root_namespace": "app",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The y (vertical) component of a screen coordinate, in screen pixels.\n",
                    "examples": [
                        12,
                        99,
                    ],
                    "name": "app.screen.coordinate.y",
                    "requirement_level": "recommended",
                    "root_namespace": "app",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "An identifier that uniquely differentiates this widget from other widgets in the same application.\n",
                    "examples": [
                        "f9bc787d-ff05-48ad-90e1-fca1d46130b3",
                        "submit_order_1829",
                    ],
                    "name": "app.widget.id",
                    "note": "A widget is an application component, typically an on-screen visual GUI element.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "app",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of an application widget.",
                    "examples": [
                        "submit",
                        "attack",
                        "Clear Cart",
                    ],
                    "name": "app.widget.name",
                    "note": "A widget is an application component, typically an on-screen visual GUI element.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "app",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "app",
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
