package deployment

// 'Deprecated, use `deployment.environment.name` instead.'
type AttrEnvironment string // deployment.environment

func (AttrEnvironment) Development()    {}
func (AttrEnvironment) Recommended()    {}
func (AttrEnvironment) Key() string     { return "deployment_environment" }
func (a AttrEnvironment) Value() string { return string(a) }

// Name of the [deployment environment] (aka deployment tier).
//
// `deployment.environment.name` does not affect the uniqueness constraints defined through
// the `service.namespace`, `service.name` and `service.instance.id` resource attributes.
// This implies that resources carrying the following attribute combinations MUST be
// considered to be identifying the same service:
//
//   - `service.name=frontend`, `deployment.environment.name=production`
//   - `service.name=frontend`, `deployment.environment.name=staging`
//
// [deployment environment]: https://wikipedia.org/wiki/Deployment_environment
type AttrEnvironmentName string // deployment.environment.name

func (AttrEnvironmentName) Development()    {}
func (AttrEnvironmentName) Recommended()    {}
func (AttrEnvironmentName) Key() string     { return "deployment_environment_name" }
func (a AttrEnvironmentName) Value() string { return string(a) }

// The id of the deployment
type AttrId string // deployment.id

func (AttrId) Development()    {}
func (AttrId) Recommended()    {}
func (AttrId) Key() string     { return "deployment_id" }
func (a AttrId) Value() string { return string(a) }

// The name of the deployment
type AttrName string // deployment.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "deployment_name" }
func (a AttrName) Value() string { return string(a) }

// The status of the deployment
type AttrStatus string // deployment.status

func (AttrStatus) Development()    {}
func (AttrStatus) Recommended()    {}
func (AttrStatus) Key() string     { return "deployment_status" }
func (a AttrStatus) Value() string { return string(a) }

const StatusFailed AttrStatus = "failed"
const StatusSucceeded AttrStatus = "succeeded"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "'Deprecated, use `deployment.environment.name` instead.'\n",
                    "deprecated": {
                        "note": "Replaced by `deployment.environment.name`.",
                        "reason": "renamed",
                        "renamed_to": "deployment.environment.name",
                    },
                    "examples": [
                        "staging",
                        "production",
                    ],
                    "name": "deployment.environment",
                    "requirement_level": "recommended",
                    "root_namespace": "deployment",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the [deployment environment](https://wikipedia.org/wiki/Deployment_environment) (aka deployment tier).\n",
                    "examples": [
                        "staging",
                        "production",
                    ],
                    "name": "deployment.environment.name",
                    "note": "`deployment.environment.name` does not affect the uniqueness constraints defined through\nthe `service.namespace`, `service.name` and `service.instance.id` resource attributes.\nThis implies that resources carrying the following attribute combinations MUST be\nconsidered to be identifying the same service:\n\n- `service.name=frontend`, `deployment.environment.name=production`\n- `service.name=frontend`, `deployment.environment.name=staging`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "deployment",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The id of the deployment.\n",
                    "examples": [
                        "1208",
                    ],
                    "name": "deployment.id",
                    "requirement_level": "recommended",
                    "root_namespace": "deployment",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the deployment.\n",
                    "examples": [
                        "deploy my app",
                        "deploy-frontend",
                    ],
                    "name": "deployment.name",
                    "requirement_level": "recommended",
                    "root_namespace": "deployment",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The status of the deployment.\n",
                    "name": "deployment.status",
                    "requirement_level": "recommended",
                    "root_namespace": "deployment",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "failed",
                                "deprecated": none,
                                "id": "failed",
                                "note": none,
                                "stability": "development",
                                "value": "failed",
                            },
                            {
                                "brief": "succeeded",
                                "deprecated": none,
                                "id": "succeeded",
                                "note": none,
                                "stability": "development",
                                "value": "succeeded",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "deployment",
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
