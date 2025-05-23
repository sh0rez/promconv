package cicd

// The kind of action a pipeline run is performing
type AttrPipelineActionName string // cicd.pipeline.action.name

func (AttrPipelineActionName) Development() {}
func (AttrPipelineActionName) Recommended() {}

// The human readable name of the pipeline within a CI/CD system
type AttrPipelineName string // cicd.pipeline.name

func (AttrPipelineName) Development() {}
func (AttrPipelineName) Recommended() {}

// The result of a pipeline run
type AttrPipelineResult string // cicd.pipeline.result

func (AttrPipelineResult) Development() {}
func (AttrPipelineResult) Recommended() {}

// The unique identifier of a pipeline run within a CI/CD system
type AttrPipelineRunId string // cicd.pipeline.run.id

func (AttrPipelineRunId) Development() {}
func (AttrPipelineRunId) Recommended() {}

// The pipeline run goes through these states during its lifecycle
type AttrPipelineRunState string // cicd.pipeline.run.state

func (AttrPipelineRunState) Development() {}
func (AttrPipelineRunState) Recommended() {}

// The [URL] of the pipeline run, providing the complete address in order to locate and identify the pipeline run
//
// [URL]: https://wikipedia.org/wiki/URL
type AttrPipelineRunUrlFull string // cicd.pipeline.run.url.full

func (AttrPipelineRunUrlFull) Development() {}
func (AttrPipelineRunUrlFull) Recommended() {}

// The human readable name of a task within a pipeline. Task here most closely aligns with a [computing process] in a pipeline. Other terms for tasks include commands, steps, and procedures
//
// [computing process]: https://wikipedia.org/wiki/Pipeline_(computing)
type AttrPipelineTaskName string // cicd.pipeline.task.name

func (AttrPipelineTaskName) Development() {}
func (AttrPipelineTaskName) Recommended() {}

// The unique identifier of a task run within a pipeline
type AttrPipelineTaskRunId string // cicd.pipeline.task.run.id

func (AttrPipelineTaskRunId) Development() {}
func (AttrPipelineTaskRunId) Recommended() {}

// The result of a task run
type AttrPipelineTaskRunResult string // cicd.pipeline.task.run.result

func (AttrPipelineTaskRunResult) Development() {}
func (AttrPipelineTaskRunResult) Recommended() {}

// The [URL] of the pipeline task run, providing the complete address in order to locate and identify the pipeline task run
//
// [URL]: https://wikipedia.org/wiki/URL
type AttrPipelineTaskRunUrlFull string // cicd.pipeline.task.run.url.full

func (AttrPipelineTaskRunUrlFull) Development() {}
func (AttrPipelineTaskRunUrlFull) Recommended() {}

// The type of the task within a pipeline
type AttrPipelineTaskType string // cicd.pipeline.task.type

func (AttrPipelineTaskType) Development() {}
func (AttrPipelineTaskType) Recommended() {}

// The name of a component of the CICD system
type AttrSystemComponent string // cicd.system.component

func (AttrSystemComponent) Development() {}
func (AttrSystemComponent) Recommended() {}

// The unique identifier of a worker within a CICD system
type AttrWorkerId string // cicd.worker.id

func (AttrWorkerId) Development() {}
func (AttrWorkerId) Recommended() {}

// The name of a worker within a CICD system
type AttrWorkerName string // cicd.worker.name

func (AttrWorkerName) Development() {}
func (AttrWorkerName) Recommended() {}

// The state of a CICD worker / agent
type AttrWorkerState string // cicd.worker.state

func (AttrWorkerState) Development() {}
func (AttrWorkerState) Recommended() {}

// The [URL] of the worker, providing the complete address in order to locate and identify the worker
//
// [URL]: https://wikipedia.org/wiki/URL
type AttrWorkerUrlFull string // cicd.worker.url.full

func (AttrWorkerUrlFull) Development() {}
func (AttrWorkerUrlFull) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The kind of action a pipeline run is performing.\n",
                    "examples": [
                        "BUILD",
                        "RUN",
                        "SYNC",
                    ],
                    "name": "cicd.pipeline.action.name",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The pipeline run is executing a build.",
                                "deprecated": none,
                                "id": "build",
                                "note": none,
                                "stability": "development",
                                "value": "BUILD",
                            },
                            {
                                "brief": "The pipeline run is executing.",
                                "deprecated": none,
                                "id": "run",
                                "note": none,
                                "stability": "development",
                                "value": "RUN",
                            },
                            {
                                "brief": "The pipeline run is executing a sync.",
                                "deprecated": none,
                                "id": "sync",
                                "note": none,
                                "stability": "development",
                                "value": "SYNC",
                            },
                        ],
                    },
                },
                {
                    "brief": "The human readable name of the pipeline within a CI/CD system.\n",
                    "examples": [
                        "Build and Test",
                        "Lint",
                        "Deploy Go Project",
                        "deploy_to_environment",
                    ],
                    "name": "cicd.pipeline.name",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The result of a pipeline run.\n",
                    "examples": [
                        "success",
                        "failure",
                        "timeout",
                        "skipped",
                    ],
                    "name": "cicd.pipeline.result",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The pipeline run finished successfully.",
                                "deprecated": none,
                                "id": "success",
                                "note": none,
                                "stability": "development",
                                "value": "success",
                            },
                            {
                                "brief": "The pipeline run did not finish successfully, eg. due to a compile error or a failing test. Such failures are usually detected by non-zero exit codes of the tools executed in the pipeline run.",
                                "deprecated": none,
                                "id": "failure",
                                "note": none,
                                "stability": "development",
                                "value": "failure",
                            },
                            {
                                "brief": "The pipeline run failed due to an error in the CICD system, eg. due to the worker being killed.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "development",
                                "value": "error",
                            },
                            {
                                "brief": "A timeout caused the pipeline run to be interrupted.",
                                "deprecated": none,
                                "id": "timeout",
                                "note": none,
                                "stability": "development",
                                "value": "timeout",
                            },
                            {
                                "brief": "The pipeline run was cancelled, eg. by a user manually cancelling the pipeline run.",
                                "deprecated": none,
                                "id": "cancellation",
                                "note": none,
                                "stability": "development",
                                "value": "cancellation",
                            },
                            {
                                "brief": "The pipeline run was skipped, eg. due to a precondition not being met.",
                                "deprecated": none,
                                "id": "skip",
                                "note": none,
                                "stability": "development",
                                "value": "skip",
                            },
                        ],
                    },
                },
                {
                    "brief": "The unique identifier of a pipeline run within a CI/CD system.\n",
                    "examples": [
                        "120912",
                    ],
                    "name": "cicd.pipeline.run.id",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The pipeline run goes through these states during its lifecycle.\n",
                    "examples": [
                        "pending",
                        "executing",
                        "finalizing",
                    ],
                    "name": "cicd.pipeline.run.state",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The run pending state spans from the event triggering the pipeline run until the execution of the run starts (eg. time spent in a queue, provisioning agents, creating run resources).\n",
                                "deprecated": none,
                                "id": "pending",
                                "note": none,
                                "stability": "development",
                                "value": "pending",
                            },
                            {
                                "brief": "The executing state spans the execution of any run tasks (eg. build, test).",
                                "deprecated": none,
                                "id": "executing",
                                "note": none,
                                "stability": "development",
                                "value": "executing",
                            },
                            {
                                "brief": "The finalizing state spans from when the run has finished executing (eg. cleanup of run resources).",
                                "deprecated": none,
                                "id": "finalizing",
                                "note": none,
                                "stability": "development",
                                "value": "finalizing",
                            },
                        ],
                    },
                },
                {
                    "brief": "The [URL](https://wikipedia.org/wiki/URL) of the pipeline run, providing the complete address in order to locate and identify the pipeline run.\n",
                    "examples": [
                        "https://github.com/open-telemetry/semantic-conventions/actions/runs/9753949763?pr=1075",
                    ],
                    "name": "cicd.pipeline.run.url.full",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The human readable name of a task within a pipeline. Task here most closely aligns with a [computing process](https://wikipedia.org/wiki/Pipeline_(computing)) in a pipeline. Other terms for tasks include commands, steps, and procedures.\n",
                    "examples": [
                        "Run GoLang Linter",
                        "Go Build",
                        "go-test",
                        "deploy_binary",
                    ],
                    "name": "cicd.pipeline.task.name",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique identifier of a task run within a pipeline.\n",
                    "examples": [
                        "12097",
                    ],
                    "name": "cicd.pipeline.task.run.id",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The result of a task run.\n",
                    "examples": [
                        "success",
                        "failure",
                        "timeout",
                        "skipped",
                    ],
                    "name": "cicd.pipeline.task.run.result",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The task run finished successfully.",
                                "deprecated": none,
                                "id": "success",
                                "note": none,
                                "stability": "development",
                                "value": "success",
                            },
                            {
                                "brief": "The task run did not finish successfully, eg. due to a compile error or a failing test. Such failures are usually detected by non-zero exit codes of the tools executed in the task run.",
                                "deprecated": none,
                                "id": "failure",
                                "note": none,
                                "stability": "development",
                                "value": "failure",
                            },
                            {
                                "brief": "The task run failed due to an error in the CICD system, eg. due to the worker being killed.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "development",
                                "value": "error",
                            },
                            {
                                "brief": "A timeout caused the task run to be interrupted.",
                                "deprecated": none,
                                "id": "timeout",
                                "note": none,
                                "stability": "development",
                                "value": "timeout",
                            },
                            {
                                "brief": "The task run was cancelled, eg. by a user manually cancelling the task run.",
                                "deprecated": none,
                                "id": "cancellation",
                                "note": none,
                                "stability": "development",
                                "value": "cancellation",
                            },
                            {
                                "brief": "The task run was skipped, eg. due to a precondition not being met.",
                                "deprecated": none,
                                "id": "skip",
                                "note": none,
                                "stability": "development",
                                "value": "skip",
                            },
                        ],
                    },
                },
                {
                    "brief": "The [URL](https://wikipedia.org/wiki/URL) of the pipeline task run, providing the complete address in order to locate and identify the pipeline task run.\n",
                    "examples": [
                        "https://github.com/open-telemetry/semantic-conventions/actions/runs/9753949763/job/26920038674?pr=1075",
                    ],
                    "name": "cicd.pipeline.task.run.url.full",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The type of the task within a pipeline.\n",
                    "examples": [
                        "build",
                        "test",
                        "deploy",
                    ],
                    "name": "cicd.pipeline.task.type",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "build",
                                "deprecated": none,
                                "id": "build",
                                "note": none,
                                "stability": "development",
                                "value": "build",
                            },
                            {
                                "brief": "test",
                                "deprecated": none,
                                "id": "test",
                                "note": none,
                                "stability": "development",
                                "value": "test",
                            },
                            {
                                "brief": "deploy",
                                "deprecated": none,
                                "id": "deploy",
                                "note": none,
                                "stability": "development",
                                "value": "deploy",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of a component of the CICD system.",
                    "examples": [
                        "controller",
                        "scheduler",
                        "agent",
                    ],
                    "name": "cicd.system.component",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique identifier of a worker within a CICD system.",
                    "examples": [
                        "abc123",
                        "10.0.1.2",
                        "controller",
                    ],
                    "name": "cicd.worker.id",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of a worker within a CICD system.",
                    "examples": [
                        "agent-abc",
                        "controller",
                        "Ubuntu LTS",
                    ],
                    "name": "cicd.worker.name",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The state of a CICD worker / agent.\n",
                    "examples": [
                        "idle",
                        "busy",
                        "down",
                    ],
                    "name": "cicd.worker.state",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The worker is not performing work for the CICD system. It is available to the CICD system to perform work on (online / idle).",
                                "deprecated": none,
                                "id": "available",
                                "note": "Pipelines might have conditions on which workers they are able to run so not every worker might be available to every pipeline.",
                                "stability": "development",
                                "value": "available",
                            },
                            {
                                "brief": "The worker is performing work for the CICD system.",
                                "deprecated": none,
                                "id": "busy",
                                "note": none,
                                "stability": "development",
                                "value": "busy",
                            },
                            {
                                "brief": "The worker is not available to the CICD system (disconnected / down).",
                                "deprecated": none,
                                "id": "offline",
                                "note": none,
                                "stability": "development",
                                "value": "offline",
                            },
                        ],
                    },
                },
                {
                    "brief": "The [URL](https://wikipedia.org/wiki/URL) of the worker, providing the complete address in order to locate and identify the worker.",
                    "examples": [
                        "https://cicd.example.org/worker/abc123",
                    ],
                    "name": "cicd.worker.url.full",
                    "requirement_level": "recommended",
                    "root_namespace": "cicd",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "cicd",
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
