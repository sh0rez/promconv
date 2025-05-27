package cicd

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
)

// Duration of a pipeline run grouped by pipeline, state and result.
type PipelineRunDuration struct {
	*prometheus.HistogramVec
	extra PipelineRunDurationExtra
}

func NewPipelineRunDuration() PipelineRunDuration {
	labels := []string{"cicd_pipeline_name", "cicd_pipeline_run_state", "cicd_pipeline_result", "error_type"}
	return PipelineRunDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "cicd",
		Name:      "pipeline_run_duration",
		Help:      "Duration of a pipeline run grouped by pipeline, state and result.",
	}, labels)}
}

func (m PipelineRunDuration) With(pipelineName AttrPipelineName, pipelineRunState AttrPipelineRunState, extra interface {
	AttrCicdPipelineResult() AttrPipelineResult
	AttrErrorType() error.AttrType
}) prometheus.Observer {
	if extra == nil {
		extra = m.extra
	}
	return m.WithLabelValues(
		string(pipelineName),
		string(pipelineRunState),
		string(extra.AttrCicdPipelineResult()),
		string(extra.AttrErrorType()),
	)
}

func (a PipelineRunDuration) WithCicdPipelineResult(attr interface{ AttrCicdPipelineResult() AttrPipelineResult }) PipelineRunDuration {
	a.extra.CicdPipelineResult = attr.AttrCicdPipelineResult()
	return a
}
func (a PipelineRunDuration) WithErrorType(attr interface{ AttrErrorType() error.AttrType }) PipelineRunDuration {
	a.extra.ErrorType = attr.AttrErrorType()
	return a
}

type PipelineRunDurationExtra struct {
	// The result of a pipeline run.
	CicdPipelineResult AttrPipelineResult `otel:"cicd.pipeline.result"`
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
}

func (a PipelineRunDurationExtra) AttrCicdPipelineResult() AttrPipelineResult {
	return a.CicdPipelineResult
}
func (a PipelineRunDurationExtra) AttrErrorType() error.AttrType { return a.ErrorType }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "PipelineRunDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "pipeline.run.duration",
        "Type": "PipelineRunDuration",
        "attributes": [
            {
                "brief": "The human readable name of the pipeline within a CI/CD system.\n",
                "examples": [
                    "Build and Test",
                    "Lint",
                    "Deploy Go Project",
                    "deploy_to_environment",
                ],
                "name": "cicd.pipeline.name",
                "requirement_level": "required",
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
                "requirement_level": "required",
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
                "brief": "The result of a pipeline run.\n",
                "examples": [
                    "success",
                    "failure",
                    "timeout",
                    "skipped",
                ],
                "name": "cicd.pipeline.result",
                "requirement_level": {
                    "conditionally_required": "If and only if the pipeline run result has been set during that state.",
                },
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
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "timeout",
                    "java.net.UnknownHostException",
                    "server_certificate_invalid",
                    "500",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if the pipeline run failed.",
                },
                "stability": "stable",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The human readable name of the pipeline within a CI/CD system.\n",
                    "examples": [
                        "Build and Test",
                        "Lint",
                        "Deploy Go Project",
                        "deploy_to_environment",
                    ],
                    "name": "cicd.pipeline.name",
                    "requirement_level": "required",
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
                    "requirement_level": "required",
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
                    "brief": "The result of a pipeline run.\n",
                    "examples": [
                        "success",
                        "failure",
                        "timeout",
                        "skipped",
                    ],
                    "name": "cicd.pipeline.result",
                    "requirement_level": {
                        "conditionally_required": "If and only if the pipeline run result has been set during that state.",
                    },
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
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "timeout",
                        "java.net.UnknownHostException",
                        "server_certificate_invalid",
                        "500",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                    "requirement_level": {
                        "conditionally_required": "If and only if the pipeline run failed.",
                    },
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
                },
            ],
            "brief": "Duration of a pipeline run grouped by pipeline, state and result.",
            "entity_associations": [
                "cicd.pipeline",
            ],
            "events": [],
            "id": "metric.cicd.pipeline.run.duration",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "cicd.pipeline.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.cicd.pipeline",
                    },
                    "cicd.pipeline.result": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.cicd.pipeline",
                    },
                    "cicd.pipeline.run.state": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.cicd.pipeline",
                    },
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/cicd/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "cicd.pipeline.run.duration",
            "name": none,
            "root_namespace": "cicd",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "s",
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
            "metric.go.j2",
        ],
    },
}
*/
