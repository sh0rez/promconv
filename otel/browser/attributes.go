package browser

// Array of brand name and version separated by a space
// This value is intended to be taken from the [UA client hints API] (`navigator.userAgentData.brands`)
//
// [UA client hints API]: https://wicg.github.io/ua-client-hints/#interface
type AttrBrands string // browser.brands

func (AttrBrands) Development()    {}
func (AttrBrands) Recommended()    {}
func (AttrBrands) Key() string     { return "browser_brands" }
func (a AttrBrands) Value() string { return string(a) }

// Preferred language of the user using the browser
// This value is intended to be taken from the Navigator API `navigator.language`
type AttrLanguage string // browser.language

func (AttrLanguage) Development()    {}
func (AttrLanguage) Recommended()    {}
func (AttrLanguage) Key() string     { return "browser_language" }
func (a AttrLanguage) Value() string { return string(a) }

// A boolean that is true if the browser is running on a mobile device
// This value is intended to be taken from the [UA client hints API] (`navigator.userAgentData.mobile`). If unavailable, this attribute SHOULD be left unset
//
// [UA client hints API]: https://wicg.github.io/ua-client-hints/#interface
type AttrMobile string // browser.mobile

func (AttrMobile) Development()    {}
func (AttrMobile) Recommended()    {}
func (AttrMobile) Key() string     { return "browser_mobile" }
func (a AttrMobile) Value() string { return string(a) }

// The platform on which the browser is running
// This value is intended to be taken from the [UA client hints API] (`navigator.userAgentData.platform`). If unavailable, the legacy `navigator.platform` API SHOULD NOT be used instead and this attribute SHOULD be left unset in order for the values to be consistent.
// The list of possible values is defined in the [W3C User-Agent Client Hints specification]. Note that some (but not all) of these values can overlap with values in the [`os.type` and `os.name` attributes]. However, for consistency, the values in the `browser.platform` attribute should capture the exact value that the user agent provides
//
// [UA client hints API]: https://wicg.github.io/ua-client-hints/#interface
// [W3C User-Agent Client Hints specification]: https://wicg.github.io/ua-client-hints/#sec-ch-ua-platform
// [`os.type` and `os.name` attributes]: ./os.md
type AttrPlatform string // browser.platform

func (AttrPlatform) Development()    {}
func (AttrPlatform) Recommended()    {}
func (AttrPlatform) Key() string     { return "browser_platform" }
func (a AttrPlatform) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Array of brand name and version separated by a space",
                    "examples": [
                        [
                            " Not A;Brand 99",
                            "Chromium 99",
                            "Chrome 99",
                        ],
                    ],
                    "name": "browser.brands",
                    "note": "This value is intended to be taken from the [UA client hints API](https://wicg.github.io/ua-client-hints/#interface) (`navigator.userAgentData.brands`).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "browser",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Preferred language of the user using the browser",
                    "examples": [
                        "en",
                        "en-US",
                        "fr",
                        "fr-FR",
                    ],
                    "name": "browser.language",
                    "note": "This value is intended to be taken from the Navigator API `navigator.language`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "browser",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A boolean that is true if the browser is running on a mobile device",
                    "name": "browser.mobile",
                    "note": "This value is intended to be taken from the [UA client hints API](https://wicg.github.io/ua-client-hints/#interface) (`navigator.userAgentData.mobile`). If unavailable, this attribute SHOULD be left unset.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "browser",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The platform on which the browser is running",
                    "examples": [
                        "Windows",
                        "macOS",
                        "Android",
                    ],
                    "name": "browser.platform",
                    "note": "This value is intended to be taken from the [UA client hints API](https://wicg.github.io/ua-client-hints/#interface) (`navigator.userAgentData.platform`). If unavailable, the legacy `navigator.platform` API SHOULD NOT be used instead and this attribute SHOULD be left unset in order for the values to be consistent.\nThe list of possible values is defined in the [W3C User-Agent Client Hints specification](https://wicg.github.io/ua-client-hints/#sec-ch-ua-platform). Note that some (but not all) of these values can overlap with values in the [`os.type` and `os.name` attributes](./os.md). However, for consistency, the values in the `browser.platform` attribute should capture the exact value that the user agent provides.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "browser",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "browser",
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
