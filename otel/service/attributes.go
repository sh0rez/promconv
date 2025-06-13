package service

// The string ID of the service instance.
//
// MUST be unique for each instance of the same `service.namespace,service.name` pair (in other words
// `service.namespace,service.name,service.instance.id` triplet MUST be globally unique). The ID helps to
// distinguish instances of the same service that exist at the same time (e.g. instances of a horizontally scaled
// service).
//
// Implementations, such as SDKs, are recommended to generate a random Version 1 or Version 4 [RFC
// 4122] UUID, but are free to use an inherent unique ID as the source of
// this value if stability is desirable. In that case, the ID SHOULD be used as source of a UUID Version 5 and
// SHOULD use the following UUID as the namespace: `4d63009a-8d0f-11ee-aad7-4c796ed8e320`.
//
// UUIDs are typically recommended, as only an opaque value for the purposes of identifying a service instance is
// needed. Similar to what can be seen in the man page for the
// [`/etc/machine-id`] file, the underlying
// data, such as pod name and namespace should be treated as confidential, being the user's choice to expose it
// or not via another resource attribute.
//
// For applications running behind an application server (like unicorn), we do not recommend using one identifier
// for all processes participating in the application. Instead, it's recommended each division (e.g. a worker
// thread in unicorn) to have its own instance.id.
//
// It's not recommended for a Collector to set `service.instance.id` if it can't unambiguously determine the
// service instance that is generating that telemetry. For instance, creating an UUID based on `pod.name` will
// likely be wrong, as the Collector might not know from which container within that pod the telemetry originated.
// However, Collectors can set the `service.instance.id` if they can unambiguously determine the service instance
// for that telemetry. This is typically the case for scraping receivers, as they know the target address and
// port
//
// [RFC
// 4122]: https://www.ietf.org/rfc/rfc4122.txt
// [`/etc/machine-id`]: https://www.freedesktop.org/software/systemd/man/latest/machine-id.html
type AttrInstanceId string // service.instance.id

func (AttrInstanceId) Development()    {}
func (AttrInstanceId) Recommended()    {}
func (AttrInstanceId) Key() string     { return "service_instance_id" }
func (a AttrInstanceId) Value() string { return string(a) }

// Logical name of the service.
//
// MUST be the same for all instances of horizontally scaled services. If the value was not specified, SDKs MUST fallback to `unknown_service:` concatenated with [`process.executable.name`], e.g. `unknown_service:bash`. If `process.executable.name` is not available, the value MUST be set to `unknown_service`
//
// [`process.executable.name`]: process.md
type AttrName string // service.name

func (AttrName) Stable()         {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "service_name" }
func (a AttrName) Value() string { return string(a) }

// A namespace for `service.name`.
//
// A string value having a meaning that helps to distinguish a group of services, for example the team name that owns a group of services. `service.name` is expected to be unique within the same namespace. If `service.namespace` is not specified in the Resource then `service.name` is expected to be unique for all services that have no explicit namespace defined (so the empty/unspecified namespace is simply one more valid namespace). Zero-length namespace string is assumed equal to unspecified namespace
type AttrNamespace string // service.namespace

func (AttrNamespace) Development()    {}
func (AttrNamespace) Recommended()    {}
func (AttrNamespace) Key() string     { return "service_namespace" }
func (a AttrNamespace) Value() string { return string(a) }

// The version string of the service API or implementation. The format is not defined by these conventions
type AttrVersion string // service.version

func (AttrVersion) Stable()         {}
func (AttrVersion) Recommended()    {}
func (AttrVersion) Key() string     { return "service_version" }
func (a AttrVersion) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The string ID of the service instance.\n",
                    "examples": [
                        "627cc493-f310-47de-96bd-71410b7dec09",
                    ],
                    "name": "service.instance.id",
                    "note": "MUST be unique for each instance of the same `service.namespace,service.name` pair (in other words\n`service.namespace,service.name,service.instance.id` triplet MUST be globally unique). The ID helps to\ndistinguish instances of the same service that exist at the same time (e.g. instances of a horizontally scaled\nservice).\n\nImplementations, such as SDKs, are recommended to generate a random Version 1 or Version 4 [RFC\n4122](https://www.ietf.org/rfc/rfc4122.txt) UUID, but are free to use an inherent unique ID as the source of\nthis value if stability is desirable. In that case, the ID SHOULD be used as source of a UUID Version 5 and\nSHOULD use the following UUID as the namespace: `4d63009a-8d0f-11ee-aad7-4c796ed8e320`.\n\nUUIDs are typically recommended, as only an opaque value for the purposes of identifying a service instance is\nneeded. Similar to what can be seen in the man page for the\n[`/etc/machine-id`](https://www.freedesktop.org/software/systemd/man/latest/machine-id.html) file, the underlying\ndata, such as pod name and namespace should be treated as confidential, being the user's choice to expose it\nor not via another resource attribute.\n\nFor applications running behind an application server (like unicorn), we do not recommend using one identifier\nfor all processes participating in the application. Instead, it's recommended each division (e.g. a worker\nthread in unicorn) to have its own instance.id.\n\nIt's not recommended for a Collector to set `service.instance.id` if it can't unambiguously determine the\nservice instance that is generating that telemetry. For instance, creating an UUID based on `pod.name` will\nlikely be wrong, as the Collector might not know from which container within that pod the telemetry originated.\nHowever, Collectors can set the `service.instance.id` if they can unambiguously determine the service instance\nfor that telemetry. This is typically the case for scraping receivers, as they know the target address and\nport.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "service",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Logical name of the service.\n",
                    "examples": [
                        "shoppingcart",
                    ],
                    "name": "service.name",
                    "note": "MUST be the same for all instances of horizontally scaled services. If the value was not specified, SDKs MUST fallback to `unknown_service:` concatenated with [`process.executable.name`](process.md), e.g. `unknown_service:bash`. If `process.executable.name` is not available, the value MUST be set to `unknown_service`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "service",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "A namespace for `service.name`.\n",
                    "examples": [
                        "Shop",
                    ],
                    "name": "service.namespace",
                    "note": "A string value having a meaning that helps to distinguish a group of services, for example the team name that owns a group of services. `service.name` is expected to be unique within the same namespace. If `service.namespace` is not specified in the Resource then `service.name` is expected to be unique for all services that have no explicit namespace defined (so the empty/unspecified namespace is simply one more valid namespace). Zero-length namespace string is assumed equal to unspecified namespace.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "service",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version string of the service API or implementation. The format is not defined by these conventions.\n",
                    "examples": [
                        "2.0.0",
                        "a01dbef8a",
                    ],
                    "name": "service.version",
                    "requirement_level": "recommended",
                    "root_namespace": "service",
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "root_namespace": "service",
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
