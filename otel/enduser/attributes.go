package enduser

// Unique identifier of an end user in the system. It maybe a username, email address, or other identifier.
// Unique identifier of an end user in the system.
//
// > [!Warning]
// > This field contains sensitive (PII) information
type AttrId string // enduser.id

func (AttrId) Development()    {}
func (AttrId) Recommended()    {}
func (AttrId) Key() string     { return "enduser_id" }
func (a AttrId) Value() string { return string(a) }

// Pseudonymous identifier of an end user. This identifier should be a random value that is not directly linked or associated with the end user's actual identity.
//
// Pseudonymous identifier of an end user.
//
// > [!Warning]
// > This field contains sensitive (linkable PII) information
type AttrPseudoId string // enduser.pseudo.id

func (AttrPseudoId) Development()    {}
func (AttrPseudoId) Recommended()    {}
func (AttrPseudoId) Key() string     { return "enduser_pseudo_id" }
func (a AttrPseudoId) Value() string { return string(a) }

// Deprecated, use `user.roles` instead
type AttrRole string // enduser.role

func (AttrRole) Development()    {}
func (AttrRole) Recommended()    {}
func (AttrRole) Key() string     { return "enduser_role" }
func (a AttrRole) Value() string { return string(a) }

// Deprecated, no replacement at this time
type AttrScope string // enduser.scope

func (AttrScope) Development()    {}
func (AttrScope) Recommended()    {}
func (AttrScope) Key() string     { return "enduser_scope" }
func (a AttrScope) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Unique identifier of an end user in the system. It maybe a username, email address, or other identifier.",
                    "examples": [
                        "username",
                    ],
                    "name": "enduser.id",
                    "note": "Unique identifier of an end user in the system.\n\n> [!Warning]\n> This field contains sensitive (PII) information.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "enduser",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Pseudonymous identifier of an end user. This identifier should be a random value that is not directly linked or associated with the end user's actual identity.\n",
                    "examples": [
                        "QdH5CAWJgqVT4rOr0qtumf",
                    ],
                    "name": "enduser.pseudo.id",
                    "note": "Pseudonymous identifier of an end user.\n\n> [!Warning]\n> This field contains sensitive (linkable PII) information.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "enduser",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `user.roles` instead.",
                    "deprecated": {
                        "note": "Use `user.roles` attribute instead.",
                        "reason": "uncategorized",
                    },
                    "examples": "admin",
                    "name": "enduser.role",
                    "requirement_level": "recommended",
                    "root_namespace": "enduser",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": "read:message, write:files",
                    "name": "enduser.scope",
                    "requirement_level": "recommended",
                    "root_namespace": "enduser",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "enduser",
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
