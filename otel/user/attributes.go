package user

// User email address
type AttrEmail string // user.email

func (AttrEmail) Development()    {}
func (AttrEmail) Recommended()    {}
func (AttrEmail) Key() string     { return "user_email" }
func (a AttrEmail) Value() string { return string(a) }

// User's full name
type AttrFullName string // user.full_name

func (AttrFullName) Development()    {}
func (AttrFullName) Recommended()    {}
func (AttrFullName) Key() string     { return "user_full_name" }
func (a AttrFullName) Value() string { return string(a) }

// Unique user hash to correlate information for a user in anonymized form.
//
// Useful if `user.id` or `user.name` contain confidential information and cannot be used
type AttrHash string // user.hash

func (AttrHash) Development()    {}
func (AttrHash) Recommended()    {}
func (AttrHash) Key() string     { return "user_hash" }
func (a AttrHash) Value() string { return string(a) }

// Unique identifier of the user
type AttrId string // user.id

func (AttrId) Development()    {}
func (AttrId) Recommended()    {}
func (AttrId) Key() string     { return "user_id" }
func (a AttrId) Value() string { return string(a) }

// Short name or login/username of the user
type AttrName string // user.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "user_name" }
func (a AttrName) Value() string { return string(a) }

// Array of user roles at the time of the event
type AttrRoles string // user.roles

func (AttrRoles) Development()    {}
func (AttrRoles) Recommended()    {}
func (AttrRoles) Key() string     { return "user_roles" }
func (a AttrRoles) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "User email address.\n",
                    "examples": [
                        "a.einstein@example.com",
                    ],
                    "name": "user.email",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "User's full name\n",
                    "examples": [
                        "Albert Einstein",
                    ],
                    "name": "user.full_name",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Unique user hash to correlate information for a user in anonymized form.\n",
                    "examples": [
                        "364fc68eaf4c8acec74a4e52d7d1feaa",
                    ],
                    "name": "user.hash",
                    "note": "Useful if `user.id` or `user.name` contain confidential information and cannot be used.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Unique identifier of the user.\n",
                    "examples": [
                        "S-1-5-21-202424912787-2692429404-2351956786-1000",
                    ],
                    "name": "user.id",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Short name or login/username of the user.\n",
                    "examples": [
                        "a.einstein",
                    ],
                    "name": "user.name",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Array of user roles at the time of the event.\n",
                    "examples": [
                        [
                            "admin",
                            "reporting_user",
                        ],
                    ],
                    "name": "user.roles",
                    "requirement_level": "recommended",
                    "root_namespace": "user",
                    "stability": "development",
                    "type": "string[]",
                },
            ],
            "root_namespace": "user",
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
