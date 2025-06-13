package aspnetcore

// ASP.NET Core exception middleware handling result
type AttrDiagnosticsExceptionResult string // aspnetcore.diagnostics.exception.result

func (AttrDiagnosticsExceptionResult) Stable()         {}
func (AttrDiagnosticsExceptionResult) Recommended()    {}
func (AttrDiagnosticsExceptionResult) Key() string     { return "aspnetcore_diagnostics_exception_result" }
func (a AttrDiagnosticsExceptionResult) Value() string { return string(a) }

const DiagnosticsExceptionResultHandled AttrDiagnosticsExceptionResult = "handled"
const DiagnosticsExceptionResultUnhandled AttrDiagnosticsExceptionResult = "unhandled"
const DiagnosticsExceptionResultSkipped AttrDiagnosticsExceptionResult = "skipped"
const DiagnosticsExceptionResultAborted AttrDiagnosticsExceptionResult = "aborted"

// Full type name of the [`IExceptionHandler`] implementation that handled the exception
//
// [`IExceptionHandler`]: https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler
type AttrDiagnosticsHandlerType string // aspnetcore.diagnostics.handler.type

func (AttrDiagnosticsHandlerType) Stable()         {}
func (AttrDiagnosticsHandlerType) Recommended()    {}
func (AttrDiagnosticsHandlerType) Key() string     { return "aspnetcore_diagnostics_handler_type" }
func (a AttrDiagnosticsHandlerType) Value() string { return string(a) }

// Rate limiting policy name
type AttrRateLimitingPolicy string // aspnetcore.rate_limiting.policy

func (AttrRateLimitingPolicy) Stable()         {}
func (AttrRateLimitingPolicy) Recommended()    {}
func (AttrRateLimitingPolicy) Key() string     { return "aspnetcore_rate_limiting_policy" }
func (a AttrRateLimitingPolicy) Value() string { return string(a) }

// Rate-limiting result, shows whether the lease was acquired or contains a rejection reason
type AttrRateLimitingResult string // aspnetcore.rate_limiting.result

func (AttrRateLimitingResult) Stable()         {}
func (AttrRateLimitingResult) Recommended()    {}
func (AttrRateLimitingResult) Key() string     { return "aspnetcore_rate_limiting_result" }
func (a AttrRateLimitingResult) Value() string { return string(a) }

const RateLimitingResultAcquired AttrRateLimitingResult = "acquired"
const RateLimitingResultEndpointLimiter AttrRateLimitingResult = "endpoint_limiter"
const RateLimitingResultGlobalLimiter AttrRateLimitingResult = "global_limiter"
const RateLimitingResultRequestCanceled AttrRateLimitingResult = "request_canceled"

// Flag indicating if request was handled by the application pipeline
type AttrRequestIsUnhandled string // aspnetcore.request.is_unhandled

func (AttrRequestIsUnhandled) Stable()         {}
func (AttrRequestIsUnhandled) Recommended()    {}
func (AttrRequestIsUnhandled) Key() string     { return "aspnetcore_request_is_unhandled" }
func (a AttrRequestIsUnhandled) Value() string { return string(a) }

// A value that indicates whether the matched route is a fallback route
type AttrRoutingIsFallback string // aspnetcore.routing.is_fallback

func (AttrRoutingIsFallback) Stable()         {}
func (AttrRoutingIsFallback) Recommended()    {}
func (AttrRoutingIsFallback) Key() string     { return "aspnetcore_routing_is_fallback" }
func (a AttrRoutingIsFallback) Value() string { return string(a) }

// Match result - success or failure
type AttrRoutingMatchStatus string // aspnetcore.routing.match_status

func (AttrRoutingMatchStatus) Stable()         {}
func (AttrRoutingMatchStatus) Recommended()    {}
func (AttrRoutingMatchStatus) Key() string     { return "aspnetcore_routing_match_status" }
func (a AttrRoutingMatchStatus) Value() string { return string(a) }

const RoutingMatchStatusSuccess AttrRoutingMatchStatus = "success"
const RoutingMatchStatusFailure AttrRoutingMatchStatus = "failure"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "ASP.NET Core exception middleware handling result",
                    "examples": [
                        "handled",
                        "unhandled",
                    ],
                    "name": "aspnetcore.diagnostics.exception.result",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
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
                    "brief": "Full type name of the [`IExceptionHandler`](https://learn.microsoft.com/dotnet/api/microsoft.aspnetcore.diagnostics.iexceptionhandler) implementation that handled the exception.",
                    "examples": [
                        "Contoso.MyHandler",
                    ],
                    "name": "aspnetcore.diagnostics.handler.type",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Rate limiting policy name.",
                    "examples": [
                        "fixed",
                        "sliding",
                        "token",
                    ],
                    "name": "aspnetcore.rate_limiting.policy",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Rate-limiting result, shows whether the lease was acquired or contains a rejection reason",
                    "examples": [
                        "acquired",
                        "request_canceled",
                    ],
                    "name": "aspnetcore.rate_limiting.result",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "Lease was acquired",
                                "deprecated": none,
                                "id": "acquired",
                                "note": none,
                                "stability": "stable",
                                "value": "acquired",
                            },
                            {
                                "brief": "Lease request was rejected by the endpoint limiter",
                                "deprecated": none,
                                "id": "endpoint_limiter",
                                "note": none,
                                "stability": "stable",
                                "value": "endpoint_limiter",
                            },
                            {
                                "brief": "Lease request was rejected by the global limiter",
                                "deprecated": none,
                                "id": "global_limiter",
                                "note": none,
                                "stability": "stable",
                                "value": "global_limiter",
                            },
                            {
                                "brief": "Lease request was canceled",
                                "deprecated": none,
                                "id": "request_canceled",
                                "note": none,
                                "stability": "stable",
                                "value": "request_canceled",
                            },
                        ],
                    },
                },
                {
                    "brief": "Flag indicating if request was handled by the application pipeline.",
                    "examples": [
                        true,
                    ],
                    "name": "aspnetcore.request.is_unhandled",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": "boolean",
                },
                {
                    "brief": "A value that indicates whether the matched route is a fallback route.",
                    "examples": [
                        true,
                    ],
                    "name": "aspnetcore.routing.is_fallback",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": "boolean",
                },
                {
                    "brief": "Match result - success or failure",
                    "examples": [
                        "success",
                        "failure",
                    ],
                    "name": "aspnetcore.routing.match_status",
                    "requirement_level": "recommended",
                    "root_namespace": "aspnetcore",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "Match succeeded",
                                "deprecated": none,
                                "id": "success",
                                "note": none,
                                "stability": "stable",
                                "value": "success",
                            },
                            {
                                "brief": "Match failed",
                                "deprecated": none,
                                "id": "failure",
                                "note": none,
                                "stability": "stable",
                                "value": "failure",
                            },
                        ],
                    },
                },
            ],
            "root_namespace": "aspnetcore",
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
