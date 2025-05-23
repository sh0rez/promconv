package cloudfoundry

// The guid of the application.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.application_id`. This is the same value as
// reported by `cf app <app-name> --guid`
type AttrAppId string // cloudfoundry.app.id

func (AttrAppId) Development() {}
func (AttrAppId) Recommended() {}

// The index of the application instance. 0 when just one instance is active.
//
// CloudFoundry defines the `instance_id` in the [Loggregator v2 envelope].
// It is used for logs and metrics emitted by CloudFoundry. It is
// supposed to contain the application instance index for applications
// deployed on the runtime.
//
// Application instrumentation should use the value from environment
// variable `CF_INSTANCE_INDEX`
//
// [Loggregator v2 envelope]: https://github.com/cloudfoundry/loggregator-api#v2-envelope
type AttrAppInstanceId string // cloudfoundry.app.instance.id

func (AttrAppInstanceId) Development() {}
func (AttrAppInstanceId) Recommended() {}

// The name of the application.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.application_name`. This is the same value
// as reported by `cf apps`
type AttrAppName string // cloudfoundry.app.name

func (AttrAppName) Development() {}
func (AttrAppName) Recommended() {}

// The guid of the CloudFoundry org the application is running in.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.org_id`. This is the same value as
// reported by `cf org <org-name> --guid`
type AttrOrgId string // cloudfoundry.org.id

func (AttrOrgId) Development() {}
func (AttrOrgId) Recommended() {}

// The name of the CloudFoundry organization the app is running in.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.org_name`. This is the same value as
// reported by `cf orgs`
type AttrOrgName string // cloudfoundry.org.name

func (AttrOrgName) Development() {}
func (AttrOrgName) Recommended() {}

// The UID identifying the process.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.process_id`. It is supposed to be equal to
// `VCAP_APPLICATION.app_id` for applications deployed to the runtime.
// For system components, this could be the actual PID
type AttrProcessId string // cloudfoundry.process.id

func (AttrProcessId) Development() {}
func (AttrProcessId) Recommended() {}

// The type of process.
//
// CloudFoundry applications can consist of multiple jobs. Usually the
// main process will be of type `web`. There can be additional background
// tasks or side-cars with different process types
type AttrProcessType string // cloudfoundry.process.type

func (AttrProcessType) Development() {}
func (AttrProcessType) Recommended() {}

// The guid of the CloudFoundry space the application is running in.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.space_id`. This is the same value as
// reported by `cf space <space-name> --guid`
type AttrSpaceId string // cloudfoundry.space.id

func (AttrSpaceId) Development() {}
func (AttrSpaceId) Recommended() {}

// The name of the CloudFoundry space the application is running in.
//
// Application instrumentation should use the value from environment
// variable `VCAP_APPLICATION.space_name`. This is the same value as
// reported by `cf spaces`
type AttrSpaceName string // cloudfoundry.space.name

func (AttrSpaceName) Development() {}
func (AttrSpaceName) Recommended() {}

// A guid or another name describing the event source.
//
// CloudFoundry defines the `source_id` in the [Loggregator v2 envelope].
// It is used for logs and metrics emitted by CloudFoundry. It is
// supposed to contain the component name, e.g. "gorouter", for
// CloudFoundry components.
//
// When system components are instrumented, values from the
// [Bosh spec]
// should be used. The `system.id` should be set to
// `spec.deployment/spec.name`
//
// [Loggregator v2 envelope]: https://github.com/cloudfoundry/loggregator-api#v2-envelope
// [Bosh spec]: https://bosh.io/docs/jobs/#properties-spec
type AttrSystemId string // cloudfoundry.system.id

func (AttrSystemId) Development() {}
func (AttrSystemId) Recommended() {}

// A guid describing the concrete instance of the event source.
//
// CloudFoundry defines the `instance_id` in the [Loggregator v2 envelope].
// It is used for logs and metrics emitted by CloudFoundry. It is
// supposed to contain the vm id for CloudFoundry components.
//
// When system components are instrumented, values from the
// [Bosh spec]
// should be used. The `system.instance.id` should be set to `spec.id`
//
// [Loggregator v2 envelope]: https://github.com/cloudfoundry/loggregator-api#v2-envelope
// [Bosh spec]: https://bosh.io/docs/jobs/#properties-spec
type AttrSystemInstanceId string // cloudfoundry.system.instance.id

func (AttrSystemInstanceId) Development() {}
func (AttrSystemInstanceId) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The guid of the application.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "cloudfoundry.app.id",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.application_id`. This is the same value as\nreported by `cf app <app-name> --guid`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The index of the application instance. 0 when just one instance is active.\n",
                    "examples": [
                        "0",
                        "1",
                    ],
                    "name": "cloudfoundry.app.instance.id",
                    "note": "CloudFoundry defines the `instance_id` in the [Loggregator v2 envelope](https://github.com/cloudfoundry/loggregator-api#v2-envelope).\nIt is used for logs and metrics emitted by CloudFoundry. It is\nsupposed to contain the application instance index for applications\ndeployed on the runtime.\n\nApplication instrumentation should use the value from environment\nvariable `CF_INSTANCE_INDEX`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the application.\n",
                    "examples": [
                        "my-app-name",
                    ],
                    "name": "cloudfoundry.app.name",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.application_name`. This is the same value\nas reported by `cf apps`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The guid of the CloudFoundry org the application is running in.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "cloudfoundry.org.id",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.org_id`. This is the same value as\nreported by `cf org <org-name> --guid`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the CloudFoundry organization the app is running in.\n",
                    "examples": [
                        "my-org-name",
                    ],
                    "name": "cloudfoundry.org.name",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.org_name`. This is the same value as\nreported by `cf orgs`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UID identifying the process.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "cloudfoundry.process.id",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.process_id`. It is supposed to be equal to\n`VCAP_APPLICATION.app_id` for applications deployed to the runtime.\nFor system components, this could be the actual PID.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The type of process.\n",
                    "examples": [
                        "web",
                    ],
                    "name": "cloudfoundry.process.type",
                    "note": "CloudFoundry applications can consist of multiple jobs. Usually the\nmain process will be of type `web`. There can be additional background\ntasks or side-cars with different process types.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The guid of the CloudFoundry space the application is running in.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "cloudfoundry.space.id",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.space_id`. This is the same value as\nreported by `cf space <space-name> --guid`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the CloudFoundry space the application is running in.\n",
                    "examples": [
                        "my-space-name",
                    ],
                    "name": "cloudfoundry.space.name",
                    "note": "Application instrumentation should use the value from environment\nvariable `VCAP_APPLICATION.space_name`. This is the same value as\nreported by `cf spaces`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A guid or another name describing the event source.\n",
                    "examples": [
                        "cf/gorouter",
                    ],
                    "name": "cloudfoundry.system.id",
                    "note": "CloudFoundry defines the `source_id` in the [Loggregator v2 envelope](https://github.com/cloudfoundry/loggregator-api#v2-envelope).\nIt is used for logs and metrics emitted by CloudFoundry. It is\nsupposed to contain the component name, e.g. \"gorouter\", for\nCloudFoundry components.\n\nWhen system components are instrumented, values from the\n[Bosh spec](https://bosh.io/docs/jobs/#properties-spec)\nshould be used. The `system.id` should be set to\n`spec.deployment/spec.name`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A guid describing the concrete instance of the event source.\n",
                    "examples": [
                        "218fc5a9-a5f1-4b54-aa05-46717d0ab26d",
                    ],
                    "name": "cloudfoundry.system.instance.id",
                    "note": "CloudFoundry defines the `instance_id` in the [Loggregator v2 envelope](https://github.com/cloudfoundry/loggregator-api#v2-envelope).\nIt is used for logs and metrics emitted by CloudFoundry. It is\nsupposed to contain the vm id for CloudFoundry components.\n\nWhen system components are instrumented, values from the\n[Bosh spec](https://bosh.io/docs/jobs/#properties-spec)\nshould be used. The `system.instance.id` should be set to `spec.id`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudfoundry",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "cloudfoundry",
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
