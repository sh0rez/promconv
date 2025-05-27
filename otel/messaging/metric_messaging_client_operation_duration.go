package messaging

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/server"
)

// Duration of messaging operation initiated by a producer or consumer client.
type ClientOperationDuration struct {
	*prometheus.HistogramVec
}

func NewClientOperationDuration() ClientOperationDuration {
	labels := []string{"messaging_operation_name", "messaging_system", "error_type", "messaging_consumer_group_name", "messaging_destination_name", "messaging_destination_subscription_name", "messaging_destination_template", "messaging_operation_type", "server_address", "messaging_destination_partition_id", "server_port"}
	return ClientOperationDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "messaging",
		Name:      "client_operation_duration",
		Help:      "Duration of messaging operation initiated by a producer or consumer client.",
	}, labels)}
}

func (m ClientOperationDuration) With(operationName AttrOperationName, system AttrSystem, extra interface {
	AttrErrorType() error.AttrType
	AttrMessagingConsumerGroupName() AttrConsumerGroupName
	AttrMessagingDestinationName() AttrDestinationName
	AttrMessagingDestinationSubscriptionName() AttrDestinationSubscriptionName
	AttrMessagingDestinationTemplate() AttrDestinationTemplate
	AttrMessagingOperationType() AttrOperationType
	AttrServerAddress() server.AttrAddress
	AttrMessagingDestinationPartitionId() AttrDestinationPartitionId
	AttrServerPort() server.AttrPort
}) prometheus.Observer {
	if extra == nil {
		extra = ClientOperationDurationExtra{}
	}
	return m.WithLabelValues(
		string(operationName),
		string(system),
		string(extra.AttrErrorType()),
		string(extra.AttrMessagingConsumerGroupName()),
		string(extra.AttrMessagingDestinationName()),
		string(extra.AttrMessagingDestinationSubscriptionName()),
		string(extra.AttrMessagingDestinationTemplate()),
		string(extra.AttrMessagingOperationType()),
		string(extra.AttrServerAddress()),
		string(extra.AttrMessagingDestinationPartitionId()),
		string(extra.AttrServerPort()),
	)
}

type ClientOperationDurationExtra struct {
	// Describes a class of error the operation ended with.
	ErrorType error.AttrType `otel:"error.type"`
	// The name of the consumer group with which a consumer is associated.
	MessagingConsumerGroupName AttrConsumerGroupName `otel:"messaging.consumer.group.name"`
	// The message destination name
	MessagingDestinationName AttrDestinationName `otel:"messaging.destination.name"`
	// The name of the destination subscription from which a message is consumed.
	MessagingDestinationSubscriptionName AttrDestinationSubscriptionName `otel:"messaging.destination.subscription.name"`
	// Low cardinality representation of the messaging destination name
	MessagingDestinationTemplate AttrDestinationTemplate `otel:"messaging.destination.template"`
	// A string identifying the type of the messaging operation.
	MessagingOperationType AttrOperationType `otel:"messaging.operation.type"`
	// Server domain name if available without reverse DNS lookup; otherwise, IP address or Unix domain socket name.
	ServerAddress server.AttrAddress `otel:"server.address"`
	// The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.
	MessagingDestinationPartitionId AttrDestinationPartitionId `otel:"messaging.destination.partition.id"`
	// Server port number.
	ServerPort server.AttrPort `otel:"server.port"`
}

func (a ClientOperationDurationExtra) AttrErrorType() error.AttrType { return a.ErrorType }
func (a ClientOperationDurationExtra) AttrMessagingConsumerGroupName() AttrConsumerGroupName {
	return a.MessagingConsumerGroupName
}
func (a ClientOperationDurationExtra) AttrMessagingDestinationName() AttrDestinationName {
	return a.MessagingDestinationName
}
func (a ClientOperationDurationExtra) AttrMessagingDestinationSubscriptionName() AttrDestinationSubscriptionName {
	return a.MessagingDestinationSubscriptionName
}
func (a ClientOperationDurationExtra) AttrMessagingDestinationTemplate() AttrDestinationTemplate {
	return a.MessagingDestinationTemplate
}
func (a ClientOperationDurationExtra) AttrMessagingOperationType() AttrOperationType {
	return a.MessagingOperationType
}
func (a ClientOperationDurationExtra) AttrServerAddress() server.AttrAddress { return a.ServerAddress }
func (a ClientOperationDurationExtra) AttrMessagingDestinationPartitionId() AttrDestinationPartitionId {
	return a.MessagingDestinationPartitionId
}
func (a ClientOperationDurationExtra) AttrServerPort() server.AttrPort { return a.ServerPort }

/*
State {
    name: "metric.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientOperationDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.operation.duration",
        "Type": "ClientOperationDuration",
        "attributes": [
            {
                "brief": "The system-specific name of the messaging operation.\n",
                "examples": [
                    "send",
                    "receive",
                    "ack",
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
                "brief": "The name of the consumer group with which a consumer is associated.\n",
                "examples": [
                    "my-group",
                    "indexer",
                ],
                "name": "messaging.consumer.group.name",
                "note": "Semantic conventions for individual messaging systems SHOULD document whether `messaging.consumer.group.name` is applicable and what it means in the context of that system.\n",
                "requirement_level": {
                    "conditionally_required": "if applicable.",
                },
                "stability": "development",
                "type": "string",
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
                "brief": "The name of the destination subscription from which a message is consumed.",
                "examples": [
                    "subscription-a",
                ],
                "name": "messaging.destination.subscription.name",
                "note": "Semantic conventions for individual messaging systems SHOULD document whether `messaging.destination.subscription.name` is applicable and what it means in the context of that system.\n",
                "requirement_level": {
                    "conditionally_required": "if applicable.",
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
                "brief": "A string identifying the type of the messaging operation.\n",
                "name": "messaging.operation.type",
                "note": "If a custom value is used, it MUST be of low cardinality.",
                "requirement_level": {
                    "conditionally_required": "If applicable.",
                },
                "stability": "development",
                "type": {
                    "allow_custom_values": none,
                    "members": [
                        {
                            "brief": "A message is created. \"Create\" spans always refer to a single message and are used to provide a unique creation context for messages in batch sending scenarios.\n",
                            "deprecated": none,
                            "id": "create",
                            "note": none,
                            "stability": "development",
                            "value": "create",
                        },
                        {
                            "brief": "One or more messages are provided for sending to an intermediary. If a single message is sent, the context of the \"Send\" span can be used as the creation context and no \"Create\" span needs to be created.\n",
                            "deprecated": none,
                            "id": "send",
                            "note": none,
                            "stability": "development",
                            "value": "send",
                        },
                        {
                            "brief": "One or more messages are requested by a consumer. This operation refers to pull-based scenarios, where consumers explicitly call methods of messaging SDKs to receive messages.\n",
                            "deprecated": none,
                            "id": "receive",
                            "note": none,
                            "stability": "development",
                            "value": "receive",
                        },
                        {
                            "brief": "One or more messages are processed by a consumer.\n",
                            "deprecated": none,
                            "id": "process",
                            "note": none,
                            "stability": "development",
                            "value": "process",
                        },
                        {
                            "brief": "One or more messages are settled.\n",
                            "deprecated": none,
                            "id": "settle",
                            "note": none,
                            "stability": "development",
                            "value": "settle",
                        },
                        {
                            "brief": "Deprecated. Use `process` instead.",
                            "deprecated": "Replaced by `process`.",
                            "id": "deliver",
                            "note": none,
                            "stability": "development",
                            "value": "deliver",
                        },
                        {
                            "brief": "Deprecated. Use `send` instead.",
                            "deprecated": "Replaced by `send`.",
                            "id": "publish",
                            "note": none,
                            "stability": "development",
                            "value": "publish",
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
                {
                    "brief": "The name of the consumer group with which a consumer is associated.\n",
                    "examples": [
                        "my-group",
                        "indexer",
                    ],
                    "name": "messaging.consumer.group.name",
                    "note": "Semantic conventions for individual messaging systems SHOULD document whether `messaging.consumer.group.name` is applicable and what it means in the context of that system.\n",
                    "requirement_level": {
                        "conditionally_required": "if applicable.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the destination subscription from which a message is consumed.",
                    "examples": [
                        "subscription-a",
                    ],
                    "name": "messaging.destination.subscription.name",
                    "note": "Semantic conventions for individual messaging systems SHOULD document whether `messaging.destination.subscription.name` is applicable and what it means in the context of that system.\n",
                    "requirement_level": {
                        "conditionally_required": "if applicable.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The system-specific name of the messaging operation.\n",
                    "examples": [
                        "send",
                        "receive",
                        "ack",
                    ],
                    "name": "messaging.operation.name",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A string identifying the type of the messaging operation.\n",
                    "name": "messaging.operation.type",
                    "note": "If a custom value is used, it MUST be of low cardinality.",
                    "requirement_level": {
                        "conditionally_required": "If applicable.",
                    },
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "A message is created. \"Create\" spans always refer to a single message and are used to provide a unique creation context for messages in batch sending scenarios.\n",
                                "deprecated": none,
                                "id": "create",
                                "note": none,
                                "stability": "development",
                                "value": "create",
                            },
                            {
                                "brief": "One or more messages are provided for sending to an intermediary. If a single message is sent, the context of the \"Send\" span can be used as the creation context and no \"Create\" span needs to be created.\n",
                                "deprecated": none,
                                "id": "send",
                                "note": none,
                                "stability": "development",
                                "value": "send",
                            },
                            {
                                "brief": "One or more messages are requested by a consumer. This operation refers to pull-based scenarios, where consumers explicitly call methods of messaging SDKs to receive messages.\n",
                                "deprecated": none,
                                "id": "receive",
                                "note": none,
                                "stability": "development",
                                "value": "receive",
                            },
                            {
                                "brief": "One or more messages are processed by a consumer.\n",
                                "deprecated": none,
                                "id": "process",
                                "note": none,
                                "stability": "development",
                                "value": "process",
                            },
                            {
                                "brief": "One or more messages are settled.\n",
                                "deprecated": none,
                                "id": "settle",
                                "note": none,
                                "stability": "development",
                                "value": "settle",
                            },
                            {
                                "brief": "Deprecated. Use `process` instead.",
                                "deprecated": "Replaced by `process`.",
                                "id": "deliver",
                                "note": none,
                                "stability": "development",
                                "value": "deliver",
                            },
                            {
                                "brief": "Deprecated. Use `send` instead.",
                                "deprecated": "Replaced by `send`.",
                                "id": "publish",
                                "note": none,
                                "stability": "development",
                                "value": "publish",
                            },
                        ],
                    },
                },
            ],
            "brief": "Duration of messaging operation initiated by a producer or consumer client.",
            "events": [],
            "id": "metric.messaging.client.operation.duration",
            "instrument": "histogram",
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
                    "messaging.consumer.group.name": {
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
                    "messaging.destination.subscription.name": {
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
                    "messaging.operation.type": {
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
            "metric_name": "messaging.client.operation.duration",
            "name": none,
            "note": "This metric SHOULD NOT be used to report processing duration - processing duration is reported in `messaging.process.duration` metric.\n",
            "root_namespace": "messaging",
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
