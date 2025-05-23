package security_rule

// A categorization value keyword used by the entity using the rule for detection of this event
type AttrCategory string // security_rule.category

func (AttrCategory) Development() {}
func (AttrCategory) Recommended() {}

// The description of the rule generating the event
type AttrDescription string // security_rule.description

func (AttrDescription) Development() {}
func (AttrDescription) Recommended() {}

// Name of the license under which the rule used to generate this event is made available
type AttrLicense string // security_rule.license

func (AttrLicense) Development() {}
func (AttrLicense) Recommended() {}

// The name of the rule or signature generating the event
type AttrName string // security_rule.name

func (AttrName) Development() {}
func (AttrName) Recommended() {}

// Reference URL to additional information about the rule used to generate this event.
//
// The URL can point to the vendor’s documentation about the rule. If that’s not available, it can also be a link to a more general page describing this type of alert
type AttrReference string // security_rule.reference

func (AttrReference) Development() {}
func (AttrReference) Recommended() {}

// Name of the ruleset, policy, group, or parent category in which the rule used to generate this event is a member
type AttrRulesetName string // security_rule.ruleset.name

func (AttrRulesetName) Development() {}
func (AttrRulesetName) Recommended() {}

// A rule ID that is unique within the scope of a set or group of agents, observers, or other entities using the rule for detection of this event
type AttrUuid string // security_rule.uuid

func (AttrUuid) Development() {}
func (AttrUuid) Recommended() {}

// The version / revision of the rule being used for analysis
type AttrVersion string // security_rule.version

func (AttrVersion) Development() {}
func (AttrVersion) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "A categorization value keyword used by the entity using the rule for detection of this event\n",
                    "examples": [
                        "Attempted Information Leak",
                    ],
                    "name": "security_rule.category",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The description of the rule generating the event.\n",
                    "examples": [
                        "Block requests to public DNS over HTTPS / TLS protocols",
                    ],
                    "name": "security_rule.description",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the license under which the rule used to generate this event is made available.\n",
                    "examples": [
                        "Apache 2.0",
                    ],
                    "name": "security_rule.license",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the rule or signature generating the event.\n",
                    "examples": [
                        "BLOCK_DNS_over_TLS",
                    ],
                    "name": "security_rule.name",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Reference URL to additional information about the rule used to generate this event.\n",
                    "examples": [
                        "https://en.wikipedia.org/wiki/DNS_over_TLS",
                    ],
                    "name": "security_rule.reference",
                    "note": "The URL can point to the vendor’s documentation about the rule. If that’s not available, it can also be a link to a more general page describing this type of alert.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the ruleset, policy, group, or parent category in which the rule used to generate this event is a member.\n",
                    "examples": [
                        "Standard_Protocol_Filters",
                    ],
                    "name": "security_rule.ruleset.name",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A rule ID that is unique within the scope of a set or group of agents, observers, or other entities using the rule for detection of this event.\n",
                    "examples": [
                        "550e8400-e29b-41d4-a716-446655440000",
                        "1100110011",
                    ],
                    "name": "security_rule.uuid",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version / revision of the rule being used for analysis.\n",
                    "examples": [
                        "1.0.0",
                    ],
                    "name": "security_rule.version",
                    "requirement_level": "recommended",
                    "root_namespace": "security_rule",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "security_rule",
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
