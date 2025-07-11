package aspnetcore

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
)

// Number of exceptions caught by exception handling middleware.
type DiagnosticsExceptions struct {
	*prometheus.CounterVec
	extra DiagnosticsExceptionsExtra
}

func NewDiagnosticsExceptions() DiagnosticsExceptions {
	labels := []string{AttrDiagnosticsExceptionResult("").Key(), error.AttrType("").Key(), AttrDiagnosticsHandlerType("").Key()}
	return DiagnosticsExceptions{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "aspnetcore_diagnostics_exceptions",
		Help: "Number of exceptions caught by exception handling middleware.",
	}, labels)}
}

func (m DiagnosticsExceptions) With(diagnosticsExceptionResult AttrDiagnosticsExceptionResult, kind error.AttrType, extras ...interface {
	AspnetcoreDiagnosticsHandlerType() AttrDiagnosticsHandlerType
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(diagnosticsExceptionResult.Value(), kind.Value(), extra.AspnetcoreDiagnosticsHandlerType().Value())
}

// Deprecated: Use [DiagnosticsExceptions.With] instead
func (m DiagnosticsExceptions) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a DiagnosticsExceptions) WithDiagnosticsHandlerType(attr interface {
	AspnetcoreDiagnosticsHandlerType() AttrDiagnosticsHandlerType
}) DiagnosticsExceptions {
	a.extra.AttrDiagnosticsHandlerType = attr.AspnetcoreDiagnosticsHandlerType()
	return a
}

type DiagnosticsExceptionsExtra struct {
	// Full type name of the [`IExceptionHandler`] implementation that handled the exception
	//
	// [`IExceptionHandler`]: https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler
	AttrDiagnosticsHandlerType AttrDiagnosticsHandlerType `otel:"aspnetcore.diagnostics.handler.type"`
}

func (a DiagnosticsExceptionsExtra) AspnetcoreDiagnosticsHandlerType() AttrDiagnosticsHandlerType {
	return a.AttrDiagnosticsHandlerType
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "DiagnosticsExceptionsExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "diagnostics.exceptions",
        "Type": "DiagnosticsExceptions",
        "attributes": [
            {
                "brief": "ASP.NET Core exception middleware handling result",
                "examples": [
                    "handled",
                    "unhandled",
                ],
                "name": "aspnetcore.diagnostics.exception.result",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "Exception was handled by the exception handling middleware.",
                            "deprecated": none,
                            "id": "handled",
                            "note": none,
                            "stability": "stable",
                            "value": "handled",
                        },
                        {
                            "brief": "Exception was not handled by the exception handling middleware.",
                            "deprecated": none,
                            "id": "unhandled",
                            "note": none,
                            "stability": "stable",
                            "value": "unhandled",
                        },
                        {
                            "brief": "Exception handling was skipped because the response had started.",
                            "deprecated": none,
                            "id": "skipped",
                            "note": none,
                            "stability": "stable",
                            "value": "skipped",
                        },
                        {
                            "brief": "Exception handling didn't run because the request was aborted.",
                            "deprecated": none,
                            "id": "aborted",
                            "note": none,
                            "stability": "stable",
                            "value": "aborted",
                        },
                    ],
                },
            },
            {
                "brief": "The full name of exception type.",
                "examples": [
                    "System.OperationCanceledException",
                    "Contoso.MyException",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
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
            {
                "brief": "Full type name of the [`IExceptionHandler`](https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler) implementation that handled the exception.",
                "examples": [
                    "Contoso.MyHandler",
                ],
                "name": "aspnetcore.diagnostics.handler.type",
                "requirement_level": {
                    "conditionally_required": "if and only if the exception was handled by this handler.",
                },
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The full name of exception type.",
                    "examples": [
                        "System.OperationCanceledException",
                        "Contoso.MyException",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
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
                {
                    "brief": "Full type name of the [`IExceptionHandler`](https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler) implementation that handled the exception.",
                    "examples": [
                        "Contoso.MyHandler",
                    ],
                    "name": "aspnetcore.diagnostics.handler.type",
                    "requirement_level": {
                        "conditionally_required": "if and only if the exception was handled by this handler.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "ASP.NET Core exception middleware handling result",
                    "examples": [
                        "handled",
                        "unhandled",
                    ],
                    "name": "aspnetcore.diagnostics.exception.result",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "Exception was handled by the exception handling middleware.",
                                "deprecated": none,
                                "id": "handled",
                                "note": none,
                                "stability": "stable",
                                "value": "handled",
                            },
                            {
                                "brief": "Exception was not handled by the exception handling middleware.",
                                "deprecated": none,
                                "id": "unhandled",
                                "note": none,
                                "stability": "stable",
                                "value": "unhandled",
                            },
                            {
                                "brief": "Exception handling was skipped because the response had started.",
                                "deprecated": none,
                                "id": "skipped",
                                "note": none,
                                "stability": "stable",
                                "value": "skipped",
                            },
                            {
                                "brief": "Exception handling didn't run because the request was aborted.",
                                "deprecated": none,
                                "id": "aborted",
                                "note": none,
                                "stability": "stable",
                                "value": "aborted",
                            },
                        ],
                    },
                },
            ],
            "brief": "Number of exceptions caught by exception handling middleware.",
            "events": [],
            "id": "metric.aspnetcore.diagnostics.exceptions",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "aspnetcore.diagnostics.exception.result": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.aspnetcore",
                    },
                    "aspnetcore.diagnostics.handler.type": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.aspnetcore",
                    },
                    "error.type": {
                        "inherited_fields": [
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/aspnetcore/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "aspnetcore.diagnostics.exceptions",
            "name": none,
            "note": "Meter name: `Microsoft.AspNetCore.Diagnostics`; Added in: ASP.NET Core 8.0\n",
            "root_namespace": "aspnetcore",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "{exception}",
        },
        "for_each_attr": <macro for_each_attr>,
        "module": "shorez.de/promconv/otel",
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
            "vec.go.j2",
        ],
    },
}
*/
