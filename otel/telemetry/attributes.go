package telemetry

// The name of the auto instrumentation agent or distribution, if used.
//
// Official auto instrumentation agents and distributions SHOULD set the `telemetry.distro.name` attribute to
// a string starting with `opentelemetry-`, e.g. `opentelemetry-java-instrumentation`
type AttrDistroName string // telemetry.distro.name

func (AttrDistroName) Development()    {}
func (AttrDistroName) Recommended()    {}
func (AttrDistroName) Key() string     { return "telemetry_distro_name" }
func (a AttrDistroName) Value() string { return string(a) }

// The version string of the auto instrumentation agent or distribution, if used
type AttrDistroVersion string // telemetry.distro.version

func (AttrDistroVersion) Development()    {}
func (AttrDistroVersion) Recommended()    {}
func (AttrDistroVersion) Key() string     { return "telemetry_distro_version" }
func (a AttrDistroVersion) Value() string { return string(a) }

// The language of the telemetry SDK
type AttrSdkLanguage string // telemetry.sdk.language

func (AttrSdkLanguage) Stable()         {}
func (AttrSdkLanguage) Recommended()    {}
func (AttrSdkLanguage) Key() string     { return "telemetry_sdk_language" }
func (a AttrSdkLanguage) Value() string { return string(a) }

const SdkLanguageCpp AttrSdkLanguage = "cpp"
const SdkLanguageDotnet AttrSdkLanguage = "dotnet"
const SdkLanguageErlang AttrSdkLanguage = "erlang"
const SdkLanguageGo AttrSdkLanguage = "go"
const SdkLanguageJava AttrSdkLanguage = "java"
const SdkLanguageNodejs AttrSdkLanguage = "nodejs"
const SdkLanguagePhp AttrSdkLanguage = "php"
const SdkLanguagePython AttrSdkLanguage = "python"
const SdkLanguageRuby AttrSdkLanguage = "ruby"
const SdkLanguageRust AttrSdkLanguage = "rust"
const SdkLanguageSwift AttrSdkLanguage = "swift"
const SdkLanguageWebjs AttrSdkLanguage = "webjs"

// The name of the telemetry SDK as defined above.
//
// The OpenTelemetry SDK MUST set the `telemetry.sdk.name` attribute to `opentelemetry`.
// If another SDK, like a fork or a vendor-provided implementation, is used, this SDK MUST set the
// `telemetry.sdk.name` attribute to the fully-qualified class or module name of this SDK's main entry point
// or another suitable identifier depending on the language.
// The identifier `opentelemetry` is reserved and MUST NOT be used in this case.
// All custom identifiers SHOULD be stable across different versions of an implementation
type AttrSdkName string // telemetry.sdk.name

func (AttrSdkName) Stable()         {}
func (AttrSdkName) Recommended()    {}
func (AttrSdkName) Key() string     { return "telemetry_sdk_name" }
func (a AttrSdkName) Value() string { return string(a) }

// The version string of the telemetry SDK
type AttrSdkVersion string // telemetry.sdk.version

func (AttrSdkVersion) Stable()         {}
func (AttrSdkVersion) Recommended()    {}
func (AttrSdkVersion) Key() string     { return "telemetry_sdk_version" }
func (a AttrSdkVersion) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The name of the auto instrumentation agent or distribution, if used.\n",
                    "examples": [
                        "parts-unlimited-java",
                    ],
                    "name": "telemetry.distro.name",
                    "note": "Official auto instrumentation agents and distributions SHOULD set the `telemetry.distro.name` attribute to\na string starting with `opentelemetry-`, e.g. `opentelemetry-java-instrumentation`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "telemetry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version string of the auto instrumentation agent or distribution, if used.\n",
                    "examples": [
                        "1.2.3",
                    ],
                    "name": "telemetry.distro.version",
                    "requirement_level": "recommended",
                    "root_namespace": "telemetry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The language of the telemetry SDK.\n",
                    "name": "telemetry.sdk.language",
                    "requirement_level": "recommended",
                    "root_namespace": "telemetry",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "cpp",
                                "note": none,
                                "stability": "stable",
                                "value": "cpp",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "dotnet",
                                "note": none,
                                "stability": "stable",
                                "value": "dotnet",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "erlang",
                                "note": none,
                                "stability": "stable",
                                "value": "erlang",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "go",
                                "note": none,
                                "stability": "stable",
                                "value": "go",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "java",
                                "note": none,
                                "stability": "stable",
                                "value": "java",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "nodejs",
                                "note": none,
                                "stability": "stable",
                                "value": "nodejs",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "php",
                                "note": none,
                                "stability": "stable",
                                "value": "php",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "python",
                                "note": none,
                                "stability": "stable",
                                "value": "python",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ruby",
                                "note": none,
                                "stability": "stable",
                                "value": "ruby",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "rust",
                                "note": none,
                                "stability": "stable",
                                "value": "rust",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "swift",
                                "note": none,
                                "stability": "stable",
                                "value": "swift",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "webjs",
                                "note": none,
                                "stability": "stable",
                                "value": "webjs",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the telemetry SDK as defined above.\n",
                    "examples": [
                        "opentelemetry",
                    ],
                    "name": "telemetry.sdk.name",
                    "note": "The OpenTelemetry SDK MUST set the `telemetry.sdk.name` attribute to `opentelemetry`.\nIf another SDK, like a fork or a vendor-provided implementation, is used, this SDK MUST set the\n`telemetry.sdk.name` attribute to the fully-qualified class or module name of this SDK's main entry point\nor another suitable identifier depending on the language.\nThe identifier `opentelemetry` is reserved and MUST NOT be used in this case.\nAll custom identifiers SHOULD be stable across different versions of an implementation.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "telemetry",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The version string of the telemetry SDK.\n",
                    "examples": [
                        "1.2.3",
                    ],
                    "name": "telemetry.sdk.version",
                    "requirement_level": "recommended",
                    "root_namespace": "telemetry",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "root_namespace": "telemetry",
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
