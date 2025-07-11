package messaging

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/server"
)

// Deprecated. Use `messaging.client.consumed.messages` instead.
type ReceiveMessages struct {
	*prometheus.CounterVec
	extra ReceiveMessagesExtra
}

func NewReceiveMessages() ReceiveMessages {
	labels := []string{AttrOperationName("").Key(), error.AttrType("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key()}
	return ReceiveMessages{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "messaging_receive_messages",
		Help: "Deprecated. Use `messaging.client.consumed.messages` instead.",
	}, labels)}
}

func (m ReceiveMessages) With(operationName AttrOperationName, extras ...interface {
	ErrorType() error.AttrType
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(operationName.Value(), extra.ErrorType().Value(), extra.ServerAddress().Value(), extra.ServerPort().Value())
}

// Deprecated: Use [ReceiveMessages.With] instead
func (m ReceiveMessages) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a ReceiveMessages) WithErrorType(attr interface{ ErrorType() error.AttrType }) ReceiveMessages {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a ReceiveMessages) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) ReceiveMessages {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a ReceiveMessages) WithServerPort(attr interface{ ServerPort() server.AttrPort }) ReceiveMessages {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}

type ReceiveMessagesExtra struct {
	// Describes a class of error the operation ended with
	AttrErrorType     error.AttrType     `otel:"error.type"`     // Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
	AttrServerAddress server.AttrAddress `otel:"server.address"` // Server port number
	AttrServerPort    server.AttrPort    `otel:"server.port"`
}

func (a ReceiveMessagesExtra) ErrorType() error.AttrType         { return a.AttrErrorType }
func (a ReceiveMessagesExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }
func (a ReceiveMessagesExtra) ServerPort() server.AttrPort       { return a.AttrServerPort }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ReceiveMessagesExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "receive.messages",
        "Type": "ReceiveMessages",
        "attributes": [
            {
                "brief": "The system-specific name of the messaging operation.\n",
                "examples": [
                    "ack",
                    "nack",
                    "send",
                ],
                "name": "messaging.operation.name",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "amqp:decode-error",
                    "KAFKA_STORAGE_ERROR",
                    "channel-error",
                ],
                "name": "error.type",
                "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if the messaging operation has failed.",
                },
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
                "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "Server domain name of the broker if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.\n",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Server port number.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "int",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "amqp:decode-error",
                        "KAFKA_STORAGE_ERROR",
                        "channel-error",
                    ],
                    "name": "error.type",
                    "note": "The `error.type` SHOULD be predictable, and SHOULD have low cardinality.\n\nWhen `error.type` is set to a type (e.g., an exception type), its\ncanonical class name identifying the type within the artifact SHOULD be used.\n\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low.\nTelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time when no\nadditional filters are applied.\n\nIf the operation has completed successfully, instrumentations SHOULD NOT set `error.type`.\n\nIf a specific domain defines its own set of error identifiers (such as HTTP or gRPC status codes),\nit's RECOMMENDED to:\n\n- Use a domain-specific attribute\n- Set `error.type` to capture all errors, regardless of whether they are defined within the domain-specific set or not.\n",
                    "requirement_level": {
                        "conditionally_required": "If and only if the messaging operation has failed.",
                    },
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
                    "brief": "The system-specific name of the messaging operation.\n",
                    "examples": [
                        "ack",
                        "nack",
                        "send",
                    ],
                    "name": "messaging.operation.name",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Server port number.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "Server domain name of the broker if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.\n",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
            ],
            "brief": "Deprecated. Use `messaging.client.consumed.messages` instead.",
            "deprecated": {
                "note": "Replaced by `messaging.client.consumed.messages`.",
                "reason": "renamed",
                "renamed_to": "messaging.client.consumed.messages",
            },
            "events": [],
            "id": "metric.messaging.receive.messages",
            "instrument": "counter",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "messaging.operation.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.messaging",
                    },
                    "server.address": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.server",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/messaging/deprecated/metrics-deprecated.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "messaging.receive.messages",
            "name": none,
            "root_namespace": "messaging",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{message}",
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
