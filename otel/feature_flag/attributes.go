package feature_flag

// The unique identifier for the flag evaluation context. For example, the targeting key
type AttrContextId string // feature_flag.context.id

func (AttrContextId) ReleaseCandidate() {}
func (AttrContextId) Recommended()      {}

// Deprecated, use `error.message` instead
type AttrEvaluationErrorMessage string // feature_flag.evaluation.error.message

func (AttrEvaluationErrorMessage) Development() {}
func (AttrEvaluationErrorMessage) Recommended() {}

// Deprecated, use `feature_flag.result.reason` instead
type AttrEvaluationReason string // feature_flag.evaluation.reason

func (AttrEvaluationReason) Development() {}
func (AttrEvaluationReason) Recommended() {}

// The lookup key of the feature flag
type AttrKey string // feature_flag.key

func (AttrKey) ReleaseCandidate() {}
func (AttrKey) Recommended()      {}

// Identifies the feature flag provider
type AttrProviderName string // feature_flag.provider.name

func (AttrProviderName) ReleaseCandidate() {}
func (AttrProviderName) Recommended()      {}

// The reason code which shows how a feature flag value was determined
type AttrResultReason string // feature_flag.result.reason

func (AttrResultReason) ReleaseCandidate() {}
func (AttrResultReason) Recommended()      {}

// The evaluated value of the feature flag.
// With some feature flag providers, feature flag results can be quite large or contain private or sensitive details.
// Because of this, `feature_flag.result.variant` is often the preferred attribute if it is available.
//
// It may be desirable to redact or otherwise limit the size and scope of `feature_flag.result.value` if possible.
// Because the evaluated flag value is unstructured and may be any type, it is left to the instrumentation author to determine how best to achieve this
type AttrResultValue string // feature_flag.result.value

func (AttrResultValue) ReleaseCandidate() {}
func (AttrResultValue) Recommended()      {}

// A semantic identifier for an evaluated flag value.
//
// A semantic identifier, commonly referred to as a variant, provides a means
// for referring to a value without including the value itself. This can
// provide additional context for understanding the meaning behind a value.
// For example, the variant `red` maybe be used for the value `#c05543`
type AttrResultVariant string // feature_flag.result.variant

func (AttrResultVariant) ReleaseCandidate() {}
func (AttrResultVariant) Recommended()      {}

// The identifier of the [flag set] to which the feature flag belongs
//
// [flag set]: https://openfeature.dev/specification/glossary/#flag-set
type AttrSetId string // feature_flag.set.id

func (AttrSetId) ReleaseCandidate() {}
func (AttrSetId) Recommended()      {}

// Deprecated, use `feature_flag.result.variant` instead
type AttrVariant string // feature_flag.variant

func (AttrVariant) Development() {}
func (AttrVariant) Recommended() {}

// The version of the ruleset used during the evaluation. This may be any stable value which uniquely identifies the ruleset
type AttrVersion string // feature_flag.version

func (AttrVersion) ReleaseCandidate() {}
func (AttrVersion) Recommended()      {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The unique identifier for the flag evaluation context. For example, the targeting key.\n",
                    "examples": [
                        "5157782b-2203-4c80-a857-dbbd5e7761db",
                    ],
                    "name": "feature_flag.context.id",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `error.message` instead.",
                    "deprecated": {
                        "note": "Replaced by `error.message`.",
                        "reason": "renamed",
                        "renamed_to": "error.message",
                    },
                    "examples": [
                        "Flag `header-color` expected type `string` but found type `number`",
                    ],
                    "name": "feature_flag.evaluation.error.message",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `feature_flag.result.reason` instead.",
                    "deprecated": {
                        "note": "Replaced by `feature_flag.result.reason`.",
                        "reason": "renamed",
                        "renamed_to": "feature_flag.result.reason",
                    },
                    "examples": [
                        "static",
                        "targeting_match",
                        "error",
                        "default",
                    ],
                    "name": "feature_flag.evaluation.reason",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The resolved value is static (no dynamic evaluation).",
                                "deprecated": none,
                                "id": "static",
                                "note": none,
                                "stability": "development",
                                "value": "static",
                            },
                            {
                                "brief": "The resolved value fell back to a pre-configured value (no dynamic evaluation occurred or dynamic evaluation yielded no result).",
                                "deprecated": none,
                                "id": "default",
                                "note": none,
                                "stability": "development",
                                "value": "default",
                            },
                            {
                                "brief": "The resolved value was the result of a dynamic evaluation, such as a rule or specific user-targeting.",
                                "deprecated": none,
                                "id": "targeting_match",
                                "note": none,
                                "stability": "development",
                                "value": "targeting_match",
                            },
                            {
                                "brief": "The resolved value was the result of pseudorandom assignment.",
                                "deprecated": none,
                                "id": "split",
                                "note": none,
                                "stability": "development",
                                "value": "split",
                            },
                            {
                                "brief": "The resolved value was retrieved from cache.",
                                "deprecated": none,
                                "id": "cached",
                                "note": none,
                                "stability": "development",
                                "value": "cached",
                            },
                            {
                                "brief": "The resolved value was the result of the flag being disabled in the management system.",
                                "deprecated": none,
                                "id": "disabled",
                                "note": none,
                                "stability": "development",
                                "value": "disabled",
                            },
                            {
                                "brief": "The reason for the resolved value could not be determined.",
                                "deprecated": none,
                                "id": "unknown",
                                "note": none,
                                "stability": "development",
                                "value": "unknown",
                            },
                            {
                                "brief": "The resolved value is non-authoritative or possibly out of date",
                                "deprecated": none,
                                "id": "stale",
                                "note": none,
                                "stability": "development",
                                "value": "stale",
                            },
                            {
                                "brief": "The resolved value was the result of an error.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "development",
                                "value": "error",
                            },
                        ],
                    },
                },
                {
                    "brief": "The lookup key of the feature flag.",
                    "examples": [
                        "logo-color",
                    ],
                    "name": "feature_flag.key",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
                {
                    "brief": "Identifies the feature flag provider.",
                    "examples": [
                        "Flag Manager",
                    ],
                    "name": "feature_flag.provider.name",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
                {
                    "brief": "The reason code which shows how a feature flag value was determined.\n",
                    "examples": [
                        "static",
                        "targeting_match",
                        "error",
                        "default",
                    ],
                    "name": "feature_flag.result.reason",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The resolved value is static (no dynamic evaluation).",
                                "deprecated": none,
                                "id": "static",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "static",
                            },
                            {
                                "brief": "The resolved value fell back to a pre-configured value (no dynamic evaluation occurred or dynamic evaluation yielded no result).",
                                "deprecated": none,
                                "id": "default",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "default",
                            },
                            {
                                "brief": "The resolved value was the result of a dynamic evaluation, such as a rule or specific user-targeting.",
                                "deprecated": none,
                                "id": "targeting_match",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "targeting_match",
                            },
                            {
                                "brief": "The resolved value was the result of pseudorandom assignment.",
                                "deprecated": none,
                                "id": "split",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "split",
                            },
                            {
                                "brief": "The resolved value was retrieved from cache.",
                                "deprecated": none,
                                "id": "cached",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "cached",
                            },
                            {
                                "brief": "The resolved value was the result of the flag being disabled in the management system.",
                                "deprecated": none,
                                "id": "disabled",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "disabled",
                            },
                            {
                                "brief": "The reason for the resolved value could not be determined.",
                                "deprecated": none,
                                "id": "unknown",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "unknown",
                            },
                            {
                                "brief": "The resolved value is non-authoritative or possibly out of date",
                                "deprecated": none,
                                "id": "stale",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "stale",
                            },
                            {
                                "brief": "The resolved value was the result of an error.",
                                "deprecated": none,
                                "id": "error",
                                "note": none,
                                "stability": "release_candidate",
                                "value": "error",
                            },
                        ],
                    },
                },
                {
                    "brief": "The evaluated value of the feature flag.",
                    "examples": [
                        "#ff0000",
                        true,
                        3,
                    ],
                    "name": "feature_flag.result.value",
                    "note": "With some feature flag providers, feature flag results can be quite large or contain private or sensitive details.\nBecause of this, `feature_flag.result.variant` is often the preferred attribute if it is available.\n\nIt may be desirable to redact or otherwise limit the size and scope of `feature_flag.result.value` if possible.\nBecause the evaluated flag value is unstructured and may be any type, it is left to the instrumentation author to determine how best to achieve this.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "any",
                },
                {
                    "brief": "A semantic identifier for an evaluated flag value.\n",
                    "examples": [
                        "red",
                        "true",
                        "on",
                    ],
                    "name": "feature_flag.result.variant",
                    "note": "A semantic identifier, commonly referred to as a variant, provides a means\nfor referring to a value without including the value itself. This can\nprovide additional context for understanding the meaning behind a value.\nFor example, the variant `red` maybe be used for the value `#c05543`.",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
                {
                    "brief": "The identifier of the [flag set](https://openfeature.dev/specification/glossary/#flag-set) to which the feature flag belongs.\n",
                    "examples": [
                        "proj-1",
                        "ab98sgs",
                        "service1/dev",
                    ],
                    "name": "feature_flag.set.id",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `feature_flag.result.variant` instead.",
                    "deprecated": {
                        "note": "Replaced by `feature_flag.result.variant`.",
                        "reason": "renamed",
                        "renamed_to": "feature_flag.result.variant",
                    },
                    "examples": [
                        "red",
                        "true",
                        "on",
                    ],
                    "name": "feature_flag.variant",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version of the ruleset used during the evaluation. This may be any stable value which uniquely identifies the ruleset.\n",
                    "examples": [
                        "1",
                        "01ABCDEF",
                    ],
                    "name": "feature_flag.version",
                    "requirement_level": "recommended",
                    "root_namespace": "feature_flag",
                    "stability": "release_candidate",
                    "type": "string",
                },
            ],
            "root_namespace": "feature_flag",
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
