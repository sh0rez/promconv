package cloudevents

// The [event_id] uniquely identifies the event
//
// [event_id]: https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#id
type AttrEventId string // cloudevents.event_id

func (AttrEventId) Development()    {}
func (AttrEventId) Recommended()    {}
func (AttrEventId) Key() string     { return "cloudevents_event_id" }
func (a AttrEventId) Value() string { return string(a) }

// The [source] identifies the context in which an event happened
//
// [source]: https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#source-1
type AttrEventSource string // cloudevents.event_source

func (AttrEventSource) Development()    {}
func (AttrEventSource) Recommended()    {}
func (AttrEventSource) Key() string     { return "cloudevents_event_source" }
func (a AttrEventSource) Value() string { return string(a) }

// The [version of the CloudEvents specification] which the event uses
//
// [version of the CloudEvents specification]: https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#specversion
type AttrEventSpecVersion string // cloudevents.event_spec_version

func (AttrEventSpecVersion) Development()    {}
func (AttrEventSpecVersion) Recommended()    {}
func (AttrEventSpecVersion) Key() string     { return "cloudevents_event_spec_version" }
func (a AttrEventSpecVersion) Value() string { return string(a) }

// The [subject] of the event in the context of the event producer (identified by source)
//
// [subject]: https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#subject
type AttrEventSubject string // cloudevents.event_subject

func (AttrEventSubject) Development()    {}
func (AttrEventSubject) Recommended()    {}
func (AttrEventSubject) Key() string     { return "cloudevents_event_subject" }
func (a AttrEventSubject) Value() string { return string(a) }

// The [event_type] contains a value describing the type of event related to the originating occurrence
//
// [event_type]: https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#type
type AttrEventType string // cloudevents.event_type

func (AttrEventType) Development()    {}
func (AttrEventType) Recommended()    {}
func (AttrEventType) Key() string     { return "cloudevents_event_type" }
func (a AttrEventType) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The [event_id](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#id) uniquely identifies the event.\n",
                    "examples": [
                        "123e4567-e89b-12d3-a456-426614174000",
                        "0001",
                    ],
                    "name": "cloudevents.event_id",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudevents",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [source](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#source-1) identifies the context in which an event happened.\n",
                    "examples": [
                        "https://github.com/cloudevents",
                        "/cloudevents/spec/pull/123",
                        "my-service",
                    ],
                    "name": "cloudevents.event_source",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudevents",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [version of the CloudEvents specification](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#specversion) which the event uses.\n",
                    "examples": "1.0",
                    "name": "cloudevents.event_spec_version",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudevents",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [subject](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#subject) of the event in the context of the event producer (identified by source).\n",
                    "examples": "mynewfile.jpg",
                    "name": "cloudevents.event_subject",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudevents",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [event_type](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md#type) contains a value describing the type of event related to the originating occurrence.\n",
                    "examples": [
                        "com.github.pull_request.opened",
                        "com.example.object.deleted.v2",
                    ],
                    "name": "cloudevents.event_type",
                    "requirement_level": "recommended",
                    "root_namespace": "cloudevents",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "cloudevents",
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
