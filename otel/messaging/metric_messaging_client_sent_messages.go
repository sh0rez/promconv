package messaging

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/server"
)

// Number of messages producer attempted to send to the broker.
type ClientSentMessages struct {
	*prometheus.CounterVec
	extra ClientSentMessagesExtra
}

func NewClientSentMessages() ClientSentMessages {
	labels := []string{AttrOperationName("").Key(), AttrSystem("").Key(), error.AttrType("").Key(), AttrDestinationName("").Key(), AttrDestinationTemplate("").Key(), server.AttrAddress("").Key(), AttrDestinationPartitionId("").Key(), server.AttrPort("").Key()}
	return ClientSentMessages{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "messaging_client_sent_messages",
		Help: "Number of messages producer attempted to send to the broker.",
	}, labels)}
}

func (m ClientSentMessages) With(operationName AttrOperationName, system AttrSystem, extras ...interface {
	ErrorType() error.AttrType
	MessagingDestinationName() AttrDestinationName
	MessagingDestinationTemplate() AttrDestinationTemplate
	ServerAddress() server.AttrAddress
	MessagingDestinationPartitionId() AttrDestinationPartitionId
	ServerPort() server.AttrPort
}) prometheus.Counter {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.CounterVec.WithLabelValues(operationName.Value(), system.Value(), extra.ErrorType().Value(), extra.MessagingDestinationName().Value(), extra.MessagingDestinationTemplate().Value(), extra.ServerAddress().Value(), extra.MessagingDestinationPartitionId().Value(), extra.ServerPort().Value())
}

// Deprecated: Use [ClientSentMessages.With] instead
func (m ClientSentMessages) WithLabelValues(lvs ...string) prometheus.Counter {
	return m.CounterVec.WithLabelValues(lvs...)
}

func (a ClientSentMessages) WithErrorType(attr interface{ ErrorType() error.AttrType }) ClientSentMessages {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a ClientSentMessages) WithDestinationName(attr interface{ MessagingDestinationName() AttrDestinationName }) ClientSentMessages {
	a.extra.AttrDestinationName = attr.MessagingDestinationName()
	return a
}
func (a ClientSentMessages) WithDestinationTemplate(attr interface {
	MessagingDestinationTemplate() AttrDestinationTemplate
}) ClientSentMessages {
	a.extra.AttrDestinationTemplate = attr.MessagingDestinationTemplate()
	return a
}
func (a ClientSentMessages) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) ClientSentMessages {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a ClientSentMessages) WithDestinationPartitionId(attr interface {
	MessagingDestinationPartitionId() AttrDestinationPartitionId
}) ClientSentMessages {
	a.extra.AttrDestinationPartitionId = attr.MessagingDestinationPartitionId()
	return a
}
func (a ClientSentMessages) WithServerPort(attr interface{ ServerPort() server.AttrPort }) ClientSentMessages {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}

type ClientSentMessagesExtra struct {
	// Describes a class of error the operation ended with
	AttrErrorType              error.AttrType             `otel:"error.type"`                         // The message destination name
	AttrDestinationName        AttrDestinationName        `otel:"messaging.destination.name"`         // Low cardinality representation of the messaging destination name
	AttrDestinationTemplate    AttrDestinationTemplate    `otel:"messaging.destination.template"`     // Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name
	AttrServerAddress          server.AttrAddress         `otel:"server.address"`                     // The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`
	AttrDestinationPartitionId AttrDestinationPartitionId `otel:"messaging.destination.partition.id"` // Server port number
	AttrServerPort             server.AttrPort            `otel:"server.port"`
}

func (a ClientSentMessagesExtra) ErrorType() error.AttrType { return a.AttrErrorType }
func (a ClientSentMessagesExtra) MessagingDestinationName() AttrDestinationName {
	return a.AttrDestinationName
}
func (a ClientSentMessagesExtra) MessagingDestinationTemplate() AttrDestinationTemplate {
	return a.AttrDestinationTemplate
}
func (a ClientSentMessagesExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }
func (a ClientSentMessagesExtra) MessagingDestinationPartitionId() AttrDestinationPartitionId {
	return a.AttrDestinationPartitionId
}
func (a ClientSentMessagesExtra) ServerPort() server.AttrPort { return a.AttrServerPort }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientSentMessagesExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.sent.messages",
        "Type": "ClientSentMessages",
        "attributes": [
            {
                "brief": "The system-specific name of the messaging operation.\n",
                "examples": [
                    "send",
                    "schedule",
                    "enqueue",
                ],
                "name": "messaging.operation.name",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The messaging system as identified by the client instrumentation.",
                "name": "messaging.system",
                "note": "The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Apache ActiveMQ",
                            "deprecated": none,
                            "id": "activemq",
                            "note": none,
                            "stability": "development",
                            "value": "activemq",
                        },
                        {
                            "brief": "Amazon Simple Queue Service (SQS)",
                            "deprecated": none,
                            "id": "aws_sqs",
                            "note": none,
                            "stability": "development",
                            "value": "aws_sqs",
                        },
                        {
                            "brief": "Azure Event Grid",
                            "deprecated": none,
                            "id": "eventgrid",
                            "note": none,
                            "stability": "development",
                            "value": "eventgrid",
                        },
                        {
                            "brief": "Azure Event Hubs",
                            "deprecated": none,
                            "id": "eventhubs",
                            "note": none,
                            "stability": "development",
                            "value": "eventhubs",
                        },
                        {
                            "brief": "Azure Service Bus",
                            "deprecated": none,
                            "id": "servicebus",
                            "note": none,
                            "stability": "development",
                            "value": "servicebus",
                        },
                        {
                            "brief": "Google Cloud Pub/Sub",
                            "deprecated": none,
                            "id": "gcp_pubsub",
                            "note": none,
                            "stability": "development",
                            "value": "gcp_pubsub",
                        },
                        {
                            "brief": "Java Message Service",
                            "deprecated": none,
                            "id": "jms",
                            "note": none,
                            "stability": "development",
                            "value": "jms",
                        },
                        {
                            "brief": "Apache Kafka",
                            "deprecated": none,
                            "id": "kafka",
                            "note": none,
                            "stability": "development",
                            "value": "kafka",
                        },
                        {
                            "brief": "RabbitMQ",
                            "deprecated": none,
                            "id": "rabbitmq",
                            "note": none,
                            "stability": "development",
                            "value": "rabbitmq",
                        },
                        {
                            "brief": "Apache RocketMQ",
                            "deprecated": none,
                            "id": "rocketmq",
                            "note": none,
                            "stability": "development",
                            "value": "rocketmq",
                        },
                        {
                            "brief": "Apache Pulsar",
                            "deprecated": none,
                            "id": "pulsar",
                            "note": none,
                            "stability": "development",
                            "value": "pulsar",
                        },
                    ],
                },
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
                "brief": "The message destination name",
                "examples": [
                    "MyQueue",
                    "MyTopic",
                ],
                "name": "messaging.destination.name",
                "note": "Destination name SHOULD uniquely identify a specific queue, topic or other entity within the broker. If\nthe broker doesn't have such notion, the destination name SHOULD uniquely identify the broker.\n",
                "requirement_level": {
                    "conditionally_required": "if and only if `messaging.destination.name` is known to have low cardinality. Otherwise, `messaging.destination.template` MAY be populated.",
                },
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "Low cardinality representation of the messaging destination name",
                "examples": [
                    "/customers/{customerId}",
                ],
                "name": "messaging.destination.template",
                "note": "Destination names could be constructed from templates. An example would be a destination name involving a user name or product id. Although the destination name in this case is of high cardinality, the underlying template is of low cardinality and can be effectively used for grouping and aggregation.\n",
                "requirement_level": {
                    "conditionally_required": "if available.",
                },
                "stability": "development",
                "type": "string",
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
                "brief": "The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.\n",
                "examples": "1",
                "name": "messaging.destination.partition.id",
                "requirement_level": "recommended",
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
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.\n",
                    "examples": "1",
                    "name": "messaging.destination.partition.id",
                    "requirement_level": "recommended",
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
                    "brief": "The message destination name",
                    "examples": [
                        "MyQueue",
                        "MyTopic",
                    ],
                    "name": "messaging.destination.name",
                    "note": "Destination name SHOULD uniquely identify a specific queue, topic or other entity within the broker. If\nthe broker doesn't have such notion, the destination name SHOULD uniquely identify the broker.\n",
                    "requirement_level": {
                        "conditionally_required": "if and only if `messaging.destination.name` is known to have low cardinality. Otherwise, `messaging.destination.template` MAY be populated.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Low cardinality representation of the messaging destination name",
                    "examples": [
                        "/customers/{customerId}",
                    ],
                    "name": "messaging.destination.template",
                    "note": "Destination names could be constructed from templates. An example would be a destination name involving a user name or product id. Although the destination name in this case is of high cardinality, the underlying template is of low cardinality and can be effectively used for grouping and aggregation.\n",
                    "requirement_level": {
                        "conditionally_required": "if available.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The messaging system as identified by the client instrumentation.",
                    "name": "messaging.system",
                    "note": "The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Apache ActiveMQ",
                                "deprecated": none,
                                "id": "activemq",
                                "note": none,
                                "stability": "development",
                                "value": "activemq",
                            },
                            {
                                "brief": "Amazon Simple Queue Service (SQS)",
                                "deprecated": none,
                                "id": "aws_sqs",
                                "note": none,
                                "stability": "development",
                                "value": "aws_sqs",
                            },
                            {
                                "brief": "Azure Event Grid",
                                "deprecated": none,
                                "id": "eventgrid",
                                "note": none,
                                "stability": "development",
                                "value": "eventgrid",
                            },
                            {
                                "brief": "Azure Event Hubs",
                                "deprecated": none,
                                "id": "eventhubs",
                                "note": none,
                                "stability": "development",
                                "value": "eventhubs",
                            },
                            {
                                "brief": "Azure Service Bus",
                                "deprecated": none,
                                "id": "servicebus",
                                "note": none,
                                "stability": "development",
                                "value": "servicebus",
                            },
                            {
                                "brief": "Google Cloud Pub/Sub",
                                "deprecated": none,
                                "id": "gcp_pubsub",
                                "note": none,
                                "stability": "development",
                                "value": "gcp_pubsub",
                            },
                            {
                                "brief": "Java Message Service",
                                "deprecated": none,
                                "id": "jms",
                                "note": none,
                                "stability": "development",
                                "value": "jms",
                            },
                            {
                                "brief": "Apache Kafka",
                                "deprecated": none,
                                "id": "kafka",
                                "note": none,
                                "stability": "development",
                                "value": "kafka",
                            },
                            {
                                "brief": "RabbitMQ",
                                "deprecated": none,
                                "id": "rabbitmq",
                                "note": none,
                                "stability": "development",
                                "value": "rabbitmq",
                            },
                            {
                                "brief": "Apache RocketMQ",
                                "deprecated": none,
                                "id": "rocketmq",
                                "note": none,
                                "stability": "development",
                                "value": "rocketmq",
                            },
                            {
                                "brief": "Apache Pulsar",
                                "deprecated": none,
                                "id": "pulsar",
                                "note": none,
                                "stability": "development",
                                "value": "pulsar",
                            },
                        ],
                    },
                },
                {
                    "brief": "The system-specific name of the messaging operation.\n",
                    "examples": [
                        "send",
                        "schedule",
                        "enqueue",
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
            "brief": "Number of messages producer attempted to send to the broker.",
            "events": [],
            "id": "metric.messaging.client.sent.messages",
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
                    "messaging.destination.name": {
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
                    "messaging.destination.partition.id": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "source_group": "registry.messaging",
                    },
                    "messaging.destination.template": {
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
                    "messaging.operation.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.messaging",
                    },
                    "messaging.system": {
                        "inherited_fields": [
                            "brief",
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/messaging/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "messaging.client.sent.messages",
            "name": none,
            "note": "This metric MUST NOT count messages that were created but haven't yet been sent.\n",
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
