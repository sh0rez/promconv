package device

// A unique identifier representing the device
//
// Its value SHOULD be identical for all apps on a device and it SHOULD NOT change if an app is uninstalled and re-installed.
// However, it might be resettable by the user for all apps on a device.
// Hardware IDs (e.g. vendor-specific serial number, IMEI or MAC address) MAY be used as values.
//
// More information about Android identifier best practices can be found [here].
//
// > [!WARNING]> This attribute may contain sensitive (PII) information. Caution should be taken when storing personal data or anything which can identify a user. GDPR and data protection laws may apply,
// > ensure you do your own due diligence.> Due to these reasons, this identifier is not recommended for consumer applications and will likely result in rejection from both Google Play and App Store.
// > However, it may be appropriate for specific enterprise scenarios, such as kiosk devices or enterprise-managed devices, with appropriate compliance clearance.
// > Any instrumentation providing this identifier MUST implement it as an opt-in feature.> See [`app.installation.id`]>  for a more privacy-preserving alternative
//
// [here]: https://developer.android.com/training/articles/user-data-ids
// [`app.installation.id`]: /docs/registry/attributes/app.md#app-installation-id
type AttrId string // device.id

func (AttrId) Development()    {}
func (AttrId) Recommended()    {}
func (AttrId) Key() string     { return "device_id" }
func (a AttrId) Value() string { return string(a) }

// The name of the device manufacturer
//
// The Android OS provides this field via [Build]. iOS apps SHOULD hardcode the value `Apple`
//
// [Build]: https://developer.android.com/reference/android/os/Build#MANUFACTURER
type AttrManufacturer string // device.manufacturer

func (AttrManufacturer) Development()    {}
func (AttrManufacturer) Recommended()    {}
func (AttrManufacturer) Key() string     { return "device_manufacturer" }
func (a AttrManufacturer) Value() string { return string(a) }

// The model identifier for the device
//
// It's recommended this value represents a machine-readable version of the model identifier rather than the market or consumer-friendly name of the device
type AttrModelIdentifier string // device.model.identifier

func (AttrModelIdentifier) Development()    {}
func (AttrModelIdentifier) Recommended()    {}
func (AttrModelIdentifier) Key() string     { return "device_model_identifier" }
func (a AttrModelIdentifier) Value() string { return string(a) }

// The marketing name for the device model
//
// It's recommended this value represents a human-readable version of the device model rather than a machine-readable alternative
type AttrModelName string // device.model.name

func (AttrModelName) Development()    {}
func (AttrModelName) Recommended()    {}
func (AttrModelName) Key() string     { return "device_model_name" }
func (a AttrModelName) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "A unique identifier representing the device\n",
                    "examples": [
                        "123456789012345",
                        "01:23:45:67:89:AB",
                    ],
                    "name": "device.id",
                    "note": "Its value SHOULD be identical for all apps on a device and it SHOULD NOT change if an app is uninstalled and re-installed.\nHowever, it might be resettable by the user for all apps on a device.\nHardware IDs (e.g. vendor-specific serial number, IMEI or MAC address) MAY be used as values.\n\nMore information about Android identifier best practices can be found [here](https://developer.android.com/training/articles/user-data-ids).\n\n> [!WARNING]\n>\n> This attribute may contain sensitive (PII) information. Caution should be taken when storing personal data or anything which can identify a user. GDPR and data protection laws may apply,\n> ensure you do your own due diligence.\n>\n> Due to these reasons, this identifier is not recommended for consumer applications and will likely result in rejection from both Google Play and App Store.\n> However, it may be appropriate for specific enterprise scenarios, such as kiosk devices or enterprise-managed devices, with appropriate compliance clearance.\n> Any instrumentation providing this identifier MUST implement it as an opt-in feature.\n>\n> See [`app.installation.id`](/docs/registry/attributes/app.md#app-installation-id) for a more privacy-preserving alternative.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "device",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the device manufacturer\n",
                    "examples": [
                        "Apple",
                        "Samsung",
                    ],
                    "name": "device.manufacturer",
                    "note": "The Android OS provides this field via [Build](https://developer.android.com/reference/android/os/Build#MANUFACTURER). iOS apps SHOULD hardcode the value `Apple`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "device",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The model identifier for the device\n",
                    "examples": [
                        "iPhone3,4",
                        "SM-G920F",
                    ],
                    "name": "device.model.identifier",
                    "note": "It's recommended this value represents a machine-readable version of the model identifier rather than the market or consumer-friendly name of the device.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "device",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The marketing name for the device model\n",
                    "examples": [
                        "iPhone 6s Plus",
                        "Samsung Galaxy S6",
                    ],
                    "name": "device.model.name",
                    "note": "It's recommended this value represents a human-readable version of the device model rather than a machine-readable alternative.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "device",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "device",
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
