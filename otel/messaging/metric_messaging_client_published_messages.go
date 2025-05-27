package messaging

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/server"
)

// Deprecated. Use `messaging.client.sent.messages` instead.
type ClientPublishedMessages struct {
	*prometheus.CounterVec
}

func NewClientPublishedMessages() ClientPublishedMessages {
	labels := []string{"messaging_operation_name", "messaging_system", "error_type", "messaging_destination_name", "messaging_destination_template", "server_address", "messaging_destination_partition_id", "server_port"}
	return ClientPublishedMessages{CounterVec: prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "messaging",
		Name:      "client_published_messages",
		Help:      "Deprecated. Use `messaging.client.sent.messages` instead.",
	}, labels)}
}

func (m ClientPublishedMessages) With(operationName AttrOperationName, system AttrSystem, extra interface {
	AttrErrorType() error.AttrType
	AttrMessagingDestinationName() AttrDestinationName
	AttrMessagingDestinationTemplate() AttrDestinationTemplate
	AttrServerAddress() server.AttrAddress
	AttrMessagingDestinationPartitionId() AttrDestinationPartitionId
	AttrServerPort() server.AttrPort
}) prometheus.Counter {
	if extra == nil {
		extra = ClientPublishedMessagesExtra{}
	}
	return m.WithLabelValues(
		string(operationName),
		string(system),
		string(extra.AttrErrorType()),
		string(extra.AttrMessagingDestinationName()),
		string(extra.AttrMessagingDestinationTemplate()),
		string(extra.AttrServerAddress()),
		string(extra.AttrMessagingDestinationPartitionId()),
		string(extra.AttrServerPort()),
	)
}

type ClientPublishedMessagesExtra struct {
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// The message destination name
	MessagingDestinationName AttrDestinationName `otel:"messaging.destination.name"`
	// Low cardinality representation of the messaging destination name
	MessagingDestinationTemplate AttrDestinationTemplate `otel:"messaging.destination.template"`
	// Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
	MessagingDestinationPartitionId AttrDestinationPartitionId `otel:"messaging.destination.partition.id"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
}

func (a ClientPublishedMessagesExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a ClientPublishedMessagesExtra) AttrMessagingDestinationName() AttrDestinationName {
	return a.MessagingDestinationName
}
func (a ClientPublishedMessagesExtra) AttrMessagingDestinationTemplate() AttrDestinationTemplate {
	return a.MessagingDestinationTemplate
}
func (a ClientPublishedMessagesExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ClientPublishedMessagesExtra) AttrMessagingDestinationPartitionId() AttrDestinationPartitionId {
	return a.MessagingDestinationPartitionId
}
func (a ClientPublishedMessagesExtra) AttrServerPort() server.AttrPort { return a.ServerPort }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientPublishedMessagesExtra",
        "Instr": "Counter",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.published.messages",
        "Type": "ClientPublishedMessages",
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
                "brief": "The messaging system as identified by the client instrumentation.",
                "name": "messaging.system",
                "note": "The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
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
                    "brief": "The messaging system as identified by the client instrumentation.",
                    "name": "messaging.system",
                    "note": "The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
            "brief": "Deprecated. Use `messaging.client.sent.messages` instead.",
            "deprecated": {
                "note": "Replaced by `messaging.client.sent.messages`.",
                "reason": "renamed",
                "renamed_to": "messaging.client.sent.messages",
            },
            "events": [],
            "id": "metric.messaging.client.published.messages",
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
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
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
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/messaging/deprecated/metrics-deprecated.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "messaging.client.published.messages",
            "name": none,
            "root_namespace": "messaging",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{message}",
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
