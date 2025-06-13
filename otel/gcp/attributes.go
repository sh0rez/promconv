package gcp

// The container within GCP where the AppHub application is defined
type AttrApphubApplicationContainer string // gcp.apphub.application.container

func (AttrApphubApplicationContainer) Development()    {}
func (AttrApphubApplicationContainer) Recommended()    {}
func (AttrApphubApplicationContainer) Key() string     { return "gcp_apphub_application_container" }
func (a AttrApphubApplicationContainer) Value() string { return string(a) }

// The name of the application as configured in AppHub
type AttrApphubApplicationId string // gcp.apphub.application.id

func (AttrApphubApplicationId) Development()    {}
func (AttrApphubApplicationId) Recommended()    {}
func (AttrApphubApplicationId) Key() string     { return "gcp_apphub_application_id" }
func (a AttrApphubApplicationId) Value() string { return string(a) }

// The GCP zone or region where the application is defined
type AttrApphubApplicationLocation string // gcp.apphub.application.location

func (AttrApphubApplicationLocation) Development()    {}
func (AttrApphubApplicationLocation) Recommended()    {}
func (AttrApphubApplicationLocation) Key() string     { return "gcp_apphub_application_location" }
func (a AttrApphubApplicationLocation) Value() string { return string(a) }

// Criticality of a service indicates its importance to the business.
//
// [See AppHub type enum]
//
// [See AppHub type enum]: https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type
type AttrApphubServiceCriticalityType string // gcp.apphub.service.criticality_type

func (AttrApphubServiceCriticalityType) Development()    {}
func (AttrApphubServiceCriticalityType) Recommended()    {}
func (AttrApphubServiceCriticalityType) Key() string     { return "gcp_apphub_service_criticality_type" }
func (a AttrApphubServiceCriticalityType) Value() string { return string(a) }

const ApphubServiceCriticalityTypeMissionCritical AttrApphubServiceCriticalityType = "MISSION_CRITICAL"
const ApphubServiceCriticalityTypeHigh AttrApphubServiceCriticalityType = "HIGH"
const ApphubServiceCriticalityTypeMedium AttrApphubServiceCriticalityType = "MEDIUM"
const ApphubServiceCriticalityTypeLow AttrApphubServiceCriticalityType = "LOW"

// Environment of a service is the stage of a software lifecycle.
//
// [See AppHub environment type]
//
// [See AppHub environment type]: https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type_1
type AttrApphubServiceEnvironmentType string // gcp.apphub.service.environment_type

func (AttrApphubServiceEnvironmentType) Development()    {}
func (AttrApphubServiceEnvironmentType) Recommended()    {}
func (AttrApphubServiceEnvironmentType) Key() string     { return "gcp_apphub_service_environment_type" }
func (a AttrApphubServiceEnvironmentType) Value() string { return string(a) }

const ApphubServiceEnvironmentTypeProduction AttrApphubServiceEnvironmentType = "PRODUCTION"
const ApphubServiceEnvironmentTypeStaging AttrApphubServiceEnvironmentType = "STAGING"
const ApphubServiceEnvironmentTypeTest AttrApphubServiceEnvironmentType = "TEST"
const ApphubServiceEnvironmentTypeDevelopment AttrApphubServiceEnvironmentType = "DEVELOPMENT"

// The name of the service as configured in AppHub
type AttrApphubServiceId string // gcp.apphub.service.id

func (AttrApphubServiceId) Development()    {}
func (AttrApphubServiceId) Recommended()    {}
func (AttrApphubServiceId) Key() string     { return "gcp_apphub_service_id" }
func (a AttrApphubServiceId) Value() string { return string(a) }

// Criticality of a workload indicates its importance to the business.
//
// [See AppHub type enum]
//
// [See AppHub type enum]: https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type
type AttrApphubWorkloadCriticalityType string // gcp.apphub.workload.criticality_type

func (AttrApphubWorkloadCriticalityType) Development()    {}
func (AttrApphubWorkloadCriticalityType) Recommended()    {}
func (AttrApphubWorkloadCriticalityType) Key() string     { return "gcp_apphub_workload_criticality_type" }
func (a AttrApphubWorkloadCriticalityType) Value() string { return string(a) }

const ApphubWorkloadCriticalityTypeMissionCritical AttrApphubWorkloadCriticalityType = "MISSION_CRITICAL"
const ApphubWorkloadCriticalityTypeHigh AttrApphubWorkloadCriticalityType = "HIGH"
const ApphubWorkloadCriticalityTypeMedium AttrApphubWorkloadCriticalityType = "MEDIUM"
const ApphubWorkloadCriticalityTypeLow AttrApphubWorkloadCriticalityType = "LOW"

// Environment of a workload is the stage of a software lifecycle.
//
// [See AppHub environment type]
//
// [See AppHub environment type]: https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type_1
type AttrApphubWorkloadEnvironmentType string // gcp.apphub.workload.environment_type

func (AttrApphubWorkloadEnvironmentType) Development()    {}
func (AttrApphubWorkloadEnvironmentType) Recommended()    {}
func (AttrApphubWorkloadEnvironmentType) Key() string     { return "gcp_apphub_workload_environment_type" }
func (a AttrApphubWorkloadEnvironmentType) Value() string { return string(a) }

const ApphubWorkloadEnvironmentTypeProduction AttrApphubWorkloadEnvironmentType = "PRODUCTION"
const ApphubWorkloadEnvironmentTypeStaging AttrApphubWorkloadEnvironmentType = "STAGING"
const ApphubWorkloadEnvironmentTypeTest AttrApphubWorkloadEnvironmentType = "TEST"
const ApphubWorkloadEnvironmentTypeDevelopment AttrApphubWorkloadEnvironmentType = "DEVELOPMENT"

// The name of the workload as configured in AppHub
type AttrApphubWorkloadId string // gcp.apphub.workload.id

func (AttrApphubWorkloadId) Development()    {}
func (AttrApphubWorkloadId) Recommended()    {}
func (AttrApphubWorkloadId) Key() string     { return "gcp_apphub_workload_id" }
func (a AttrApphubWorkloadId) Value() string { return string(a) }

// Identifies the Google Cloud service for which the official client library is intended.
// Intended to be a stable identifier for Google Cloud client libraries that is uniform across implementation languages. The value should be derived from the canonical service domain for the service; for example, 'foo.googleapis.com' should result in a value of 'foo'
type AttrClientService string // gcp.client.service

func (AttrClientService) Development()    {}
func (AttrClientService) Recommended()    {}
func (AttrClientService) Key() string     { return "gcp_client_service" }
func (a AttrClientService) Value() string { return string(a) }

// The name of the Cloud Run [execution] being run for the Job, as set by the [`CLOUD_RUN_EXECUTION`] environment variable
//
// [execution]: https://cloud.google.com/run/docs/managing/job-executions
// [`CLOUD_RUN_EXECUTION`]: https://cloud.google.com/run/docs/container-contract#jobs-env-vars
type AttrCloudRunJobExecution string // gcp.cloud_run.job.execution

func (AttrCloudRunJobExecution) Development()    {}
func (AttrCloudRunJobExecution) Recommended()    {}
func (AttrCloudRunJobExecution) Key() string     { return "gcp_cloud_run_job_execution" }
func (a AttrCloudRunJobExecution) Value() string { return string(a) }

// The index for a task within an execution as provided by the [`CLOUD_RUN_TASK_INDEX`] environment variable
//
// [`CLOUD_RUN_TASK_INDEX`]: https://cloud.google.com/run/docs/container-contract#jobs-env-vars
type AttrCloudRunJobTaskIndex string // gcp.cloud_run.job.task_index

func (AttrCloudRunJobTaskIndex) Development()    {}
func (AttrCloudRunJobTaskIndex) Recommended()    {}
func (AttrCloudRunJobTaskIndex) Key() string     { return "gcp_cloud_run_job_task_index" }
func (a AttrCloudRunJobTaskIndex) Value() string { return string(a) }

// The hostname of a GCE instance. This is the full value of the default or [custom hostname]
//
// [custom hostname]: https://cloud.google.com/compute/docs/instances/custom-hostname-vm
type AttrGceInstanceHostname string // gcp.gce.instance.hostname

func (AttrGceInstanceHostname) Development()    {}
func (AttrGceInstanceHostname) Recommended()    {}
func (AttrGceInstanceHostname) Key() string     { return "gcp_gce_instance_hostname" }
func (a AttrGceInstanceHostname) Value() string { return string(a) }

// The instance name of a GCE instance. This is the value provided by `host.name`, the visible name of the instance in the Cloud Console UI, and the prefix for the default hostname of the instance as defined by the [default internal DNS name]
//
// [default internal DNS name]: https://cloud.google.com/compute/docs/internal-dns#instance-fully-qualified-domain-names
type AttrGceInstanceName string // gcp.gce.instance.name

func (AttrGceInstanceName) Development()    {}
func (AttrGceInstanceName) Recommended()    {}
func (AttrGceInstanceName) Key() string     { return "gcp_gce_instance_name" }
func (a AttrGceInstanceName) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The container within GCP where the AppHub application is defined.\n",
                    "examples": [
                        "projects/my-container-project",
                    ],
                    "name": "gcp.apphub.application.container",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the application as configured in AppHub.\n",
                    "examples": [
                        "my-application",
                    ],
                    "name": "gcp.apphub.application.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The GCP zone or region where the application is defined.\n",
                    "examples": [
                        "us-central1",
                    ],
                    "name": "gcp.apphub.application.location",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Criticality of a service indicates its importance to the business.\n",
                    "name": "gcp.apphub.service.criticality_type",
                    "note": "[See AppHub type enum](https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Mission critical service.",
                                "deprecated": none,
                                "id": "mission_critical",
                                "note": none,
                                "stability": "development",
                                "value": "MISSION_CRITICAL",
                            },
                            {
                                "brief": "High impact.",
                                "deprecated": none,
                                "id": "high",
                                "note": none,
                                "stability": "development",
                                "value": "HIGH",
                            },
                            {
                                "brief": "Medium impact.",
                                "deprecated": none,
                                "id": "medium",
                                "note": none,
                                "stability": "development",
                                "value": "MEDIUM",
                            },
                            {
                                "brief": "Low impact.",
                                "deprecated": none,
                                "id": "low",
                                "note": none,
                                "stability": "development",
                                "value": "LOW",
                            },
                        ],
                    },
                },
                {
                    "brief": "Environment of a service is the stage of a software lifecycle.\n",
                    "name": "gcp.apphub.service.environment_type",
                    "note": "[See AppHub environment type](https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type_1)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Production environment.",
                                "deprecated": none,
                                "id": "production",
                                "note": none,
                                "stability": "development",
                                "value": "PRODUCTION",
                            },
                            {
                                "brief": "Staging environment.",
                                "deprecated": none,
                                "id": "staging",
                                "note": none,
                                "stability": "development",
                                "value": "STAGING",
                            },
                            {
                                "brief": "Test environment.",
                                "deprecated": none,
                                "id": "test",
                                "note": none,
                                "stability": "development",
                                "value": "TEST",
                            },
                            {
                                "brief": "Development environment.",
                                "deprecated": none,
                                "id": "development",
                                "note": none,
                                "stability": "development",
                                "value": "DEVELOPMENT",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the service as configured in AppHub.\n",
                    "examples": [
                        "my-service",
                    ],
                    "name": "gcp.apphub.service.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Criticality of a workload indicates its importance to the business.\n",
                    "name": "gcp.apphub.workload.criticality_type",
                    "note": "[See AppHub type enum](https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Mission critical service.",
                                "deprecated": none,
                                "id": "mission_critical",
                                "note": none,
                                "stability": "development",
                                "value": "MISSION_CRITICAL",
                            },
                            {
                                "brief": "High impact.",
                                "deprecated": none,
                                "id": "high",
                                "note": none,
                                "stability": "development",
                                "value": "HIGH",
                            },
                            {
                                "brief": "Medium impact.",
                                "deprecated": none,
                                "id": "medium",
                                "note": none,
                                "stability": "development",
                                "value": "MEDIUM",
                            },
                            {
                                "brief": "Low impact.",
                                "deprecated": none,
                                "id": "low",
                                "note": none,
                                "stability": "development",
                                "value": "LOW",
                            },
                        ],
                    },
                },
                {
                    "brief": "Environment of a workload is the stage of a software lifecycle.\n",
                    "name": "gcp.apphub.workload.environment_type",
                    "note": "[See AppHub environment type](https://cloud.google.com/app-hub/docs/reference/rest/v1/Attributes#type_1)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Production environment.",
                                "deprecated": none,
                                "id": "production",
                                "note": none,
                                "stability": "development",
                                "value": "PRODUCTION",
                            },
                            {
                                "brief": "Staging environment.",
                                "deprecated": none,
                                "id": "staging",
                                "note": none,
                                "stability": "development",
                                "value": "STAGING",
                            },
                            {
                                "brief": "Test environment.",
                                "deprecated": none,
                                "id": "test",
                                "note": none,
                                "stability": "development",
                                "value": "TEST",
                            },
                            {
                                "brief": "Development environment.",
                                "deprecated": none,
                                "id": "development",
                                "note": none,
                                "stability": "development",
                                "value": "DEVELOPMENT",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the workload as configured in AppHub.\n",
                    "examples": [
                        "my-workload",
                    ],
                    "name": "gcp.apphub.workload.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Identifies the Google Cloud service for which the official client library is intended.",
                    "examples": [
                        "appengine",
                        "run",
                        "firestore",
                        "alloydb",
                        "spanner",
                    ],
                    "name": "gcp.client.service",
                    "note": "Intended to be a stable identifier for Google Cloud client libraries that is uniform across implementation languages. The value should be derived from the canonical service domain for the service; for example, 'foo.googleapis.com' should result in a value of 'foo'.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the Cloud Run [execution](https://cloud.google.com/run/docs/managing/job-executions) being run for the Job, as set by the [`CLOUD_RUN_EXECUTION`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars) environment variable.\n",
                    "examples": [
                        "job-name-xxxx",
                        "sample-job-mdw84",
                    ],
                    "name": "gcp.cloud_run.job.execution",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The index for a task within an execution as provided by the [`CLOUD_RUN_TASK_INDEX`](https://cloud.google.com/run/docs/container-contract#jobs-env-vars) environment variable.\n",
                    "examples": [
                        0,
                        1,
                    ],
                    "name": "gcp.cloud_run.job.task_index",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The hostname of a GCE instance. This is the full value of the default or [custom hostname](https://cloud.google.com/compute/docs/instances/custom-hostname-vm).\n",
                    "examples": [
                        "my-host1234.example.com",
                        "sample-vm.us-west1-b.c.my-project.internal",
                    ],
                    "name": "gcp.gce.instance.hostname",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The instance name of a GCE instance. This is the value provided by `host.name`, the visible name of the instance in the Cloud Console UI, and the prefix for the default hostname of the instance as defined by the [default internal DNS name](https://cloud.google.com/compute/docs/internal-dns#instance-fully-qualified-domain-names).\n",
                    "examples": [
                        "instance-1",
                        "my-vm-name",
                    ],
                    "name": "gcp.gce.instance.name",
                    "requirement_level": "recommended",
                    "root_namespace": "gcp",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "gcp",
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
