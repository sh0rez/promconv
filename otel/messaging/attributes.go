package messaging

// The number of messages sent, received, or processed in the scope of the batching operation.
// Instrumentations SHOULD NOT set `messaging.batch.message_count` on spans that operate with a single message. When a messaging client library supports both batch and single-message API for the same operation, instrumentations SHOULD use `messaging.batch.message_count` for batching APIs and SHOULD NOT use it for single-message APIs
type AttrBatchMessageCount string // messaging.batch.message_count

func (AttrBatchMessageCount) Development()    {}
func (AttrBatchMessageCount) Recommended()    {}
func (AttrBatchMessageCount) Key() string     { return "messaging_batch_message_count" }
func (a AttrBatchMessageCount) Value() string { return string(a) }

// A unique identifier for the client that consumes or produces a message
type AttrClientId string // messaging.client.id

func (AttrClientId) Development()    {}
func (AttrClientId) Recommended()    {}
func (AttrClientId) Key() string     { return "messaging_client_id" }
func (a AttrClientId) Value() string { return string(a) }

// The name of the consumer group with which a consumer is associated.
//
// Semantic conventions for individual messaging systems SHOULD document whether `messaging.consumer.group.name` is applicable and what it means in the context of that system
type AttrConsumerGroupName string // messaging.consumer.group.name

func (AttrConsumerGroupName) Development()    {}
func (AttrConsumerGroupName) Recommended()    {}
func (AttrConsumerGroupName) Key() string     { return "messaging_consumer_group_name" }
func (a AttrConsumerGroupName) Value() string { return string(a) }

// A boolean that is true if the message destination is anonymous (could be unnamed or have auto-generated name)
type AttrDestinationAnonymous string // messaging.destination.anonymous

func (AttrDestinationAnonymous) Development()    {}
func (AttrDestinationAnonymous) Recommended()    {}
func (AttrDestinationAnonymous) Key() string     { return "messaging_destination_anonymous" }
func (a AttrDestinationAnonymous) Value() string { return string(a) }

// The message destination name
// Destination name SHOULD uniquely identify a specific queue, topic or other entity within the broker. If
// the broker doesn't have such notion, the destination name SHOULD uniquely identify the broker
type AttrDestinationName string // messaging.destination.name

func (AttrDestinationName) Development()    {}
func (AttrDestinationName) Recommended()    {}
func (AttrDestinationName) Key() string     { return "messaging_destination_name" }
func (a AttrDestinationName) Value() string { return string(a) }

// The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`
type AttrDestinationPartitionId string // messaging.destination.partition.id

func (AttrDestinationPartitionId) Development()    {}
func (AttrDestinationPartitionId) Recommended()    {}
func (AttrDestinationPartitionId) Key() string     { return "messaging_destination_partition_id" }
func (a AttrDestinationPartitionId) Value() string { return string(a) }

// The name of the destination subscription from which a message is consumed.
// Semantic conventions for individual messaging systems SHOULD document whether `messaging.destination.subscription.name` is applicable and what it means in the context of that system
type AttrDestinationSubscriptionName string // messaging.destination.subscription.name

func (AttrDestinationSubscriptionName) Development()    {}
func (AttrDestinationSubscriptionName) Recommended()    {}
func (AttrDestinationSubscriptionName) Key() string     { return "messaging_destination_subscription_name" }
func (a AttrDestinationSubscriptionName) Value() string { return string(a) }

// Low cardinality representation of the messaging destination name
// Destination names could be constructed from templates. An example would be a destination name involving a user name or product id. Although the destination name in this case is of high cardinality, the underlying template is of low cardinality and can be effectively used for grouping and aggregation
type AttrDestinationTemplate string // messaging.destination.template

func (AttrDestinationTemplate) Development()    {}
func (AttrDestinationTemplate) Recommended()    {}
func (AttrDestinationTemplate) Key() string     { return "messaging_destination_template" }
func (a AttrDestinationTemplate) Value() string { return string(a) }

// A boolean that is true if the message destination is temporary and might not exist anymore after messages are processed
type AttrDestinationTemporary string // messaging.destination.temporary

func (AttrDestinationTemporary) Development()    {}
func (AttrDestinationTemporary) Recommended()    {}
func (AttrDestinationTemporary) Key() string     { return "messaging_destination_temporary" }
func (a AttrDestinationTemporary) Value() string { return string(a) }

// Deprecated, no replacement at this time
type AttrDestinationPublishAnonymous string // messaging.destination_publish.anonymous

func (AttrDestinationPublishAnonymous) Development()    {}
func (AttrDestinationPublishAnonymous) Recommended()    {}
func (AttrDestinationPublishAnonymous) Key() string     { return "messaging_destination_publish_anonymous" }
func (a AttrDestinationPublishAnonymous) Value() string { return string(a) }

// Deprecated, no replacement at this time
type AttrDestinationPublishName string // messaging.destination_publish.name

func (AttrDestinationPublishName) Development()    {}
func (AttrDestinationPublishName) Recommended()    {}
func (AttrDestinationPublishName) Key() string     { return "messaging_destination_publish_name" }
func (a AttrDestinationPublishName) Value() string { return string(a) }

// Deprecated, use `messaging.consumer.group.name` instead
type AttrEventhubsConsumerGroup string // messaging.eventhubs.consumer.group

func (AttrEventhubsConsumerGroup) Development()    {}
func (AttrEventhubsConsumerGroup) Recommended()    {}
func (AttrEventhubsConsumerGroup) Key() string     { return "messaging_eventhubs_consumer_group" }
func (a AttrEventhubsConsumerGroup) Value() string { return string(a) }

// The UTC epoch seconds at which the message has been accepted and stored in the entity
type AttrEventhubsMessageEnqueuedTime string // messaging.eventhubs.message.enqueued_time

func (AttrEventhubsMessageEnqueuedTime) Development() {}
func (AttrEventhubsMessageEnqueuedTime) Recommended() {}
func (AttrEventhubsMessageEnqueuedTime) Key() string {
	return "messaging_eventhubs_message_enqueued_time"
}
func (a AttrEventhubsMessageEnqueuedTime) Value() string { return string(a) }

// The ack deadline in seconds set for the modify ack deadline request
type AttrGcpPubsubMessageAckDeadline string // messaging.gcp_pubsub.message.ack_deadline

func (AttrGcpPubsubMessageAckDeadline) Development() {}
func (AttrGcpPubsubMessageAckDeadline) Recommended() {}
func (AttrGcpPubsubMessageAckDeadline) Key() string {
	return "messaging_gcp_pubsub_message_ack_deadline"
}
func (a AttrGcpPubsubMessageAckDeadline) Value() string { return string(a) }

// The ack id for a given message
type AttrGcpPubsubMessageAckId string // messaging.gcp_pubsub.message.ack_id

func (AttrGcpPubsubMessageAckId) Development()    {}
func (AttrGcpPubsubMessageAckId) Recommended()    {}
func (AttrGcpPubsubMessageAckId) Key() string     { return "messaging_gcp_pubsub_message_ack_id" }
func (a AttrGcpPubsubMessageAckId) Value() string { return string(a) }

// The delivery attempt for a given message
type AttrGcpPubsubMessageDeliveryAttempt string // messaging.gcp_pubsub.message.delivery_attempt

func (AttrGcpPubsubMessageDeliveryAttempt) Development() {}
func (AttrGcpPubsubMessageDeliveryAttempt) Recommended() {}
func (AttrGcpPubsubMessageDeliveryAttempt) Key() string {
	return "messaging_gcp_pubsub_message_delivery_attempt"
}
func (a AttrGcpPubsubMessageDeliveryAttempt) Value() string { return string(a) }

// The ordering key for a given message. If the attribute is not present, the message does not have an ordering key
type AttrGcpPubsubMessageOrderingKey string // messaging.gcp_pubsub.message.ordering_key

func (AttrGcpPubsubMessageOrderingKey) Development() {}
func (AttrGcpPubsubMessageOrderingKey) Recommended() {}
func (AttrGcpPubsubMessageOrderingKey) Key() string {
	return "messaging_gcp_pubsub_message_ordering_key"
}
func (a AttrGcpPubsubMessageOrderingKey) Value() string { return string(a) }

// Deprecated, use `messaging.consumer.group.name` instead
type AttrKafkaConsumerGroup string // messaging.kafka.consumer.group

func (AttrKafkaConsumerGroup) Development()    {}
func (AttrKafkaConsumerGroup) Recommended()    {}
func (AttrKafkaConsumerGroup) Key() string     { return "messaging_kafka_consumer_group" }
func (a AttrKafkaConsumerGroup) Value() string { return string(a) }

// Deprecated, use `messaging.destination.partition.id` instead
type AttrKafkaDestinationPartition string // messaging.kafka.destination.partition

func (AttrKafkaDestinationPartition) Development()    {}
func (AttrKafkaDestinationPartition) Recommended()    {}
func (AttrKafkaDestinationPartition) Key() string     { return "messaging_kafka_destination_partition" }
func (a AttrKafkaDestinationPartition) Value() string { return string(a) }

// Message keys in Kafka are used for grouping alike messages to ensure they're processed on the same partition. They differ from `messaging.message.id` in that they're not unique. If the key is `null`, the attribute MUST NOT be set.
//
// If the key type is not string, it's string representation has to be supplied for the attribute. If the key has no unambiguous, canonical string form, don't include its value
type AttrKafkaMessageKey string // messaging.kafka.message.key

func (AttrKafkaMessageKey) Development()    {}
func (AttrKafkaMessageKey) Recommended()    {}
func (AttrKafkaMessageKey) Key() string     { return "messaging_kafka_message_key" }
func (a AttrKafkaMessageKey) Value() string { return string(a) }

// Deprecated, use `messaging.kafka.offset` instead
type AttrKafkaMessageOffset string // messaging.kafka.message.offset

func (AttrKafkaMessageOffset) Development()    {}
func (AttrKafkaMessageOffset) Recommended()    {}
func (AttrKafkaMessageOffset) Key() string     { return "messaging_kafka_message_offset" }
func (a AttrKafkaMessageOffset) Value() string { return string(a) }

// A boolean that is true if the message is a tombstone
type AttrKafkaMessageTombstone string // messaging.kafka.message.tombstone

func (AttrKafkaMessageTombstone) Development()    {}
func (AttrKafkaMessageTombstone) Recommended()    {}
func (AttrKafkaMessageTombstone) Key() string     { return "messaging_kafka_message_tombstone" }
func (a AttrKafkaMessageTombstone) Value() string { return string(a) }

// The offset of a record in the corresponding Kafka partition
type AttrKafkaOffset string // messaging.kafka.offset

func (AttrKafkaOffset) Development()    {}
func (AttrKafkaOffset) Recommended()    {}
func (AttrKafkaOffset) Key() string     { return "messaging_kafka_offset" }
func (a AttrKafkaOffset) Value() string { return string(a) }

// The size of the message body in bytes.
//
// This can refer to both the compressed or uncompressed body size. If both sizes are known, the uncompressed
// body size should be used
type AttrMessageBodySize string // messaging.message.body.size

func (AttrMessageBodySize) Development()    {}
func (AttrMessageBodySize) Recommended()    {}
func (AttrMessageBodySize) Key() string     { return "messaging_message_body_size" }
func (a AttrMessageBodySize) Value() string { return string(a) }

// The conversation ID identifying the conversation to which the message belongs, represented as a string. Sometimes called "Correlation ID"
type AttrMessageConversationId string // messaging.message.conversation_id

func (AttrMessageConversationId) Development()    {}
func (AttrMessageConversationId) Recommended()    {}
func (AttrMessageConversationId) Key() string     { return "messaging_message_conversation_id" }
func (a AttrMessageConversationId) Value() string { return string(a) }

// The size of the message body and metadata in bytes.
//
// This can refer to both the compressed or uncompressed size. If both sizes are known, the uncompressed
// size should be used
type AttrMessageEnvelopeSize string // messaging.message.envelope.size

func (AttrMessageEnvelopeSize) Development()    {}
func (AttrMessageEnvelopeSize) Recommended()    {}
func (AttrMessageEnvelopeSize) Key() string     { return "messaging_message_envelope_size" }
func (a AttrMessageEnvelopeSize) Value() string { return string(a) }

// A value used by the messaging system as an identifier for the message, represented as a string
type AttrMessageId string // messaging.message.id

func (AttrMessageId) Development()    {}
func (AttrMessageId) Recommended()    {}
func (AttrMessageId) Key() string     { return "messaging_message_id" }
func (a AttrMessageId) Value() string { return string(a) }

// Deprecated, use `messaging.operation.type` instead
type AttrOperation string // messaging.operation

func (AttrOperation) Development()    {}
func (AttrOperation) Recommended()    {}
func (AttrOperation) Key() string     { return "messaging_operation" }
func (a AttrOperation) Value() string { return string(a) }

// The system-specific name of the messaging operation
type AttrOperationName string // messaging.operation.name

func (AttrOperationName) Development()    {}
func (AttrOperationName) Recommended()    {}
func (AttrOperationName) Key() string     { return "messaging_operation_name" }
func (a AttrOperationName) Value() string { return string(a) }

// A string identifying the type of the messaging operation.
//
// If a custom value is used, it MUST be of low cardinality
type AttrOperationType string // messaging.operation.type

func (AttrOperationType) Development()    {}
func (AttrOperationType) Recommended()    {}
func (AttrOperationType) Key() string     { return "messaging_operation_type" }
func (a AttrOperationType) Value() string { return string(a) }

const OperationTypeCreate AttrOperationType = "create"
const OperationTypeSend AttrOperationType = "send"
const OperationTypeReceive AttrOperationType = "receive"
const OperationTypeProcess AttrOperationType = "process"
const OperationTypeSettle AttrOperationType = "settle"
const OperationTypeDeliver AttrOperationType = "deliver"
const OperationTypePublish AttrOperationType = "publish"

// RabbitMQ message routing key
type AttrRabbitmqDestinationRoutingKey string // messaging.rabbitmq.destination.routing_key

func (AttrRabbitmqDestinationRoutingKey) Development() {}
func (AttrRabbitmqDestinationRoutingKey) Recommended() {}
func (AttrRabbitmqDestinationRoutingKey) Key() string {
	return "messaging_rabbitmq_destination_routing_key"
}
func (a AttrRabbitmqDestinationRoutingKey) Value() string { return string(a) }

// RabbitMQ message delivery tag
type AttrRabbitmqMessageDeliveryTag string // messaging.rabbitmq.message.delivery_tag

func (AttrRabbitmqMessageDeliveryTag) Development()    {}
func (AttrRabbitmqMessageDeliveryTag) Recommended()    {}
func (AttrRabbitmqMessageDeliveryTag) Key() string     { return "messaging_rabbitmq_message_delivery_tag" }
func (a AttrRabbitmqMessageDeliveryTag) Value() string { return string(a) }

// Deprecated, use `messaging.consumer.group.name` instead
type AttrRocketmqClientGroup string // messaging.rocketmq.client_group

func (AttrRocketmqClientGroup) Development()    {}
func (AttrRocketmqClientGroup) Recommended()    {}
func (AttrRocketmqClientGroup) Key() string     { return "messaging_rocketmq_client_group" }
func (a AttrRocketmqClientGroup) Value() string { return string(a) }

// Model of message consumption. This only applies to consumer spans
type AttrRocketmqConsumptionModel string // messaging.rocketmq.consumption_model

func (AttrRocketmqConsumptionModel) Development()    {}
func (AttrRocketmqConsumptionModel) Recommended()    {}
func (AttrRocketmqConsumptionModel) Key() string     { return "messaging_rocketmq_consumption_model" }
func (a AttrRocketmqConsumptionModel) Value() string { return string(a) }

const RocketmqConsumptionModelClustering AttrRocketmqConsumptionModel = "clustering"
const RocketmqConsumptionModelBroadcasting AttrRocketmqConsumptionModel = "broadcasting"

// The delay time level for delay message, which determines the message delay time
type AttrRocketmqMessageDelayTimeLevel string // messaging.rocketmq.message.delay_time_level

func (AttrRocketmqMessageDelayTimeLevel) Development() {}
func (AttrRocketmqMessageDelayTimeLevel) Recommended() {}
func (AttrRocketmqMessageDelayTimeLevel) Key() string {
	return "messaging_rocketmq_message_delay_time_level"
}
func (a AttrRocketmqMessageDelayTimeLevel) Value() string { return string(a) }

// The timestamp in milliseconds that the delay message is expected to be delivered to consumer
type AttrRocketmqMessageDeliveryTimestamp string // messaging.rocketmq.message.delivery_timestamp

func (AttrRocketmqMessageDeliveryTimestamp) Development() {}
func (AttrRocketmqMessageDeliveryTimestamp) Recommended() {}
func (AttrRocketmqMessageDeliveryTimestamp) Key() string {
	return "messaging_rocketmq_message_delivery_timestamp"
}
func (a AttrRocketmqMessageDeliveryTimestamp) Value() string { return string(a) }

// It is essential for FIFO message. Messages that belong to the same message group are always processed one by one within the same consumer group
type AttrRocketmqMessageGroup string // messaging.rocketmq.message.group

func (AttrRocketmqMessageGroup) Development()    {}
func (AttrRocketmqMessageGroup) Recommended()    {}
func (AttrRocketmqMessageGroup) Key() string     { return "messaging_rocketmq_message_group" }
func (a AttrRocketmqMessageGroup) Value() string { return string(a) }

// Key(s) of message, another way to mark message besides message id
type AttrRocketmqMessageKeys string // messaging.rocketmq.message.keys

func (AttrRocketmqMessageKeys) Development()    {}
func (AttrRocketmqMessageKeys) Recommended()    {}
func (AttrRocketmqMessageKeys) Key() string     { return "messaging_rocketmq_message_keys" }
func (a AttrRocketmqMessageKeys) Value() string { return string(a) }

// The secondary classifier of message besides topic
type AttrRocketmqMessageTag string // messaging.rocketmq.message.tag

func (AttrRocketmqMessageTag) Development()    {}
func (AttrRocketmqMessageTag) Recommended()    {}
func (AttrRocketmqMessageTag) Key() string     { return "messaging_rocketmq_message_tag" }
func (a AttrRocketmqMessageTag) Value() string { return string(a) }

// Type of message
type AttrRocketmqMessageType string // messaging.rocketmq.message.type

func (AttrRocketmqMessageType) Development()    {}
func (AttrRocketmqMessageType) Recommended()    {}
func (AttrRocketmqMessageType) Key() string     { return "messaging_rocketmq_message_type" }
func (a AttrRocketmqMessageType) Value() string { return string(a) }

const RocketmqMessageTypeNormal AttrRocketmqMessageType = "normal"
const RocketmqMessageTypeFifo AttrRocketmqMessageType = "fifo"
const RocketmqMessageTypeDelay AttrRocketmqMessageType = "delay"
const RocketmqMessageTypeTransaction AttrRocketmqMessageType = "transaction"

// Namespace of RocketMQ resources, resources in different namespaces are individual
type AttrRocketmqNamespace string // messaging.rocketmq.namespace

func (AttrRocketmqNamespace) Development()    {}
func (AttrRocketmqNamespace) Recommended()    {}
func (AttrRocketmqNamespace) Key() string     { return "messaging_rocketmq_namespace" }
func (a AttrRocketmqNamespace) Value() string { return string(a) }

// Deprecated, use `messaging.destination.subscription.name` instead
type AttrServicebusDestinationSubscriptionName string // messaging.servicebus.destination.subscription_name

func (AttrServicebusDestinationSubscriptionName) Development() {}
func (AttrServicebusDestinationSubscriptionName) Recommended() {}
func (AttrServicebusDestinationSubscriptionName) Key() string {
	return "messaging_servicebus_destination_subscription_name"
}
func (a AttrServicebusDestinationSubscriptionName) Value() string { return string(a) }

// Describes the [settlement type]
//
// [settlement type]: https://learn.microsoft.com/azure/service-bus-messaging/message-transfers-locks-settlement#peeklock
type AttrServicebusDispositionStatus string // messaging.servicebus.disposition_status

func (AttrServicebusDispositionStatus) Development()    {}
func (AttrServicebusDispositionStatus) Recommended()    {}
func (AttrServicebusDispositionStatus) Key() string     { return "messaging_servicebus_disposition_status" }
func (a AttrServicebusDispositionStatus) Value() string { return string(a) }

const ServicebusDispositionStatusComplete AttrServicebusDispositionStatus = "complete"
const ServicebusDispositionStatusAbandon AttrServicebusDispositionStatus = "abandon"
const ServicebusDispositionStatusDeadLetter AttrServicebusDispositionStatus = "dead_letter"
const ServicebusDispositionStatusDefer AttrServicebusDispositionStatus = "defer"

// Number of deliveries that have been attempted for this message
type AttrServicebusMessageDeliveryCount string // messaging.servicebus.message.delivery_count

func (AttrServicebusMessageDeliveryCount) Development() {}
func (AttrServicebusMessageDeliveryCount) Recommended() {}
func (AttrServicebusMessageDeliveryCount) Key() string {
	return "messaging_servicebus_message_delivery_count"
}
func (a AttrServicebusMessageDeliveryCount) Value() string { return string(a) }

// The UTC epoch seconds at which the message has been accepted and stored in the entity
type AttrServicebusMessageEnqueuedTime string // messaging.servicebus.message.enqueued_time

func (AttrServicebusMessageEnqueuedTime) Development() {}
func (AttrServicebusMessageEnqueuedTime) Recommended() {}
func (AttrServicebusMessageEnqueuedTime) Key() string {
	return "messaging_servicebus_message_enqueued_time"
}
func (a AttrServicebusMessageEnqueuedTime) Value() string { return string(a) }

// The messaging system as identified by the client instrumentation.
// The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge
type AttrSystem string // messaging.system

func (AttrSystem) Development()    {}
func (AttrSystem) Recommended()    {}
func (AttrSystem) Key() string     { return "messaging_system" }
func (a AttrSystem) Value() string { return string(a) }

const SystemActivemq AttrSystem = "activemq"
const SystemAwsSqs AttrSystem = "aws_sqs"
const SystemEventgrid AttrSystem = "eventgrid"
const SystemEventhubs AttrSystem = "eventhubs"
const SystemServicebus AttrSystem = "servicebus"
const SystemGcpPubsub AttrSystem = "gcp_pubsub"
const SystemJms AttrSystem = "jms"
const SystemKafka AttrSystem = "kafka"
const SystemRabbitmq AttrSystem = "rabbitmq"
const SystemRocketmq AttrSystem = "rocketmq"
const SystemPulsar AttrSystem = "pulsar"

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The number of messages sent, received, or processed in the scope of the batching operation.",
                    "examples": [
                        0,
                        1,
                        2,
                    ],
                    "name": "messaging.batch.message_count",
                    "note": "Instrumentations SHOULD NOT set `messaging.batch.message_count` on spans that operate with a single message. When a messaging client library supports both batch and single-message API for the same operation, instrumentations SHOULD use `messaging.batch.message_count` for batching APIs and SHOULD NOT use it for single-message APIs.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "A unique identifier for the client that consumes or produces a message.\n",
                    "examples": [
                        "client-5",
                        "myhost@8742@s8083jm",
                    ],
                    "name": "messaging.client.id",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
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
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A boolean that is true if the message destination is anonymous (could be unnamed or have auto-generated name).",
                    "name": "messaging.destination.anonymous",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The message destination name",
                    "examples": [
                        "MyQueue",
                        "MyTopic",
                    ],
                    "name": "messaging.destination.name",
                    "note": "Destination name SHOULD uniquely identify a specific queue, topic or other entity within the broker. If\nthe broker doesn't have such notion, the destination name SHOULD uniquely identify the broker.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The identifier of the partition messages are sent to or received from, unique within the `messaging.destination.name`.\n",
                    "examples": "1",
                    "name": "messaging.destination.partition.id",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
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
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
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
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A boolean that is true if the message destination is temporary and might not exist anymore after messages are processed.",
                    "name": "messaging.destination.temporary",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed. No replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "name": "messaging.destination_publish.anonymous",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "Deprecated, no replacement at this time.",
                    "deprecated": {
                        "note": "Removed. No replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "MyQueue",
                        "MyTopic",
                    ],
                    "name": "messaging.destination_publish.name",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.consumer.group.name` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.consumer.group.name`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.consumer.group.name",
                    },
                    "examples": "$Default",
                    "name": "messaging.eventhubs.consumer.group",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UTC epoch seconds at which the message has been accepted and stored in the entity.\n",
                    "examples": 1701393730,
                    "name": "messaging.eventhubs.message.enqueued_time",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The ack deadline in seconds set for the modify ack deadline request.\n",
                    "examples": 10,
                    "name": "messaging.gcp_pubsub.message.ack_deadline",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The ack id for a given message.\n",
                    "examples": "ack_id",
                    "name": "messaging.gcp_pubsub.message.ack_id",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The delivery attempt for a given message.\n",
                    "examples": 2,
                    "name": "messaging.gcp_pubsub.message.delivery_attempt",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The ordering key for a given message. If the attribute is not present, the message does not have an ordering key.\n",
                    "examples": "ordering_key",
                    "name": "messaging.gcp_pubsub.message.ordering_key",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.consumer.group.name` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.consumer.group.name`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.consumer.group.name",
                    },
                    "examples": "my-group",
                    "name": "messaging.kafka.consumer.group",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.destination.partition.id` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.destination.partition.id`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.destination.partition.id",
                    },
                    "examples": 2,
                    "name": "messaging.kafka.destination.partition",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Message keys in Kafka are used for grouping alike messages to ensure they're processed on the same partition. They differ from `messaging.message.id` in that they're not unique. If the key is `null`, the attribute MUST NOT be set.\n",
                    "examples": "myKey",
                    "name": "messaging.kafka.message.key",
                    "note": "If the key type is not string, it's string representation has to be supplied for the attribute. If the key has no unambiguous, canonical string form, don't include its value.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.kafka.offset` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.kafka.offset`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.kafka.offset",
                    },
                    "examples": 42,
                    "name": "messaging.kafka.message.offset",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "A boolean that is true if the message is a tombstone.",
                    "name": "messaging.kafka.message.tombstone",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The offset of a record in the corresponding Kafka partition.\n",
                    "examples": 42,
                    "name": "messaging.kafka.offset",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The size of the message body in bytes.\n",
                    "examples": 1439,
                    "name": "messaging.message.body.size",
                    "note": "This can refer to both the compressed or uncompressed body size. If both sizes are known, the uncompressed\nbody size should be used.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The conversation ID identifying the conversation to which the message belongs, represented as a string. Sometimes called \"Correlation ID\".\n",
                    "examples": "MyConversationId",
                    "name": "messaging.message.conversation_id",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The size of the message body and metadata in bytes.\n",
                    "examples": 2738,
                    "name": "messaging.message.envelope.size",
                    "note": "This can refer to both the compressed or uncompressed size. If both sizes are known, the uncompressed\nsize should be used.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "A value used by the messaging system as an identifier for the message, represented as a string.",
                    "examples": "452a7c7c7c7048c2f887f61572b18fc2",
                    "name": "messaging.message.id",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.operation.type` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.operation.type`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.operation.type",
                    },
                    "examples": [
                        "publish",
                        "create",
                        "process",
                    ],
                    "name": "messaging.operation",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
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
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A string identifying the type of the messaging operation.\n",
                    "name": "messaging.operation.type",
                    "note": "If a custom value is used, it MUST be of low cardinality.",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": {
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
                    "brief": "RabbitMQ message routing key.\n",
                    "examples": "myKey",
                    "name": "messaging.rabbitmq.destination.routing_key",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "RabbitMQ message delivery tag\n",
                    "examples": 123,
                    "name": "messaging.rabbitmq.message.delivery_tag",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `messaging.consumer.group.name` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.consumer.group.name` on the consumer spans. No replacement for producer spans.\n",
                        "reason": "uncategorized",
                    },
                    "examples": "myConsumerGroup",
                    "name": "messaging.rocketmq.client_group",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Model of message consumption. This only applies to consumer spans.\n",
                    "name": "messaging.rocketmq.consumption_model",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Clustering consumption model",
                                "deprecated": none,
                                "id": "clustering",
                                "note": none,
                                "stability": "development",
                                "value": "clustering",
                            },
                            {
                                "brief": "Broadcasting consumption model",
                                "deprecated": none,
                                "id": "broadcasting",
                                "note": none,
                                "stability": "development",
                                "value": "broadcasting",
                            },
                        ],
                    },
                },
                {
                    "brief": "The delay time level for delay message, which determines the message delay time.\n",
                    "examples": 3,
                    "name": "messaging.rocketmq.message.delay_time_level",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The timestamp in milliseconds that the delay message is expected to be delivered to consumer.\n",
                    "examples": 1665987217045,
                    "name": "messaging.rocketmq.message.delivery_timestamp",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "It is essential for FIFO message. Messages that belong to the same message group are always processed one by one within the same consumer group.\n",
                    "examples": "myMessageGroup",
                    "name": "messaging.rocketmq.message.group",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Key(s) of message, another way to mark message besides message id.\n",
                    "examples": [
                        [
                            "keyA",
                            "keyB",
                        ],
                    ],
                    "name": "messaging.rocketmq.message.keys",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The secondary classifier of message besides topic.\n",
                    "examples": "tagA",
                    "name": "messaging.rocketmq.message.tag",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Type of message.\n",
                    "name": "messaging.rocketmq.message.type",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Normal message",
                                "deprecated": none,
                                "id": "normal",
                                "note": none,
                                "stability": "development",
                                "value": "normal",
                            },
                            {
                                "brief": "FIFO message",
                                "deprecated": none,
                                "id": "fifo",
                                "note": none,
                                "stability": "development",
                                "value": "fifo",
                            },
                            {
                                "brief": "Delay message",
                                "deprecated": none,
                                "id": "delay",
                                "note": none,
                                "stability": "development",
                                "value": "delay",
                            },
                            {
                                "brief": "Transaction message",
                                "deprecated": none,
                                "id": "transaction",
                                "note": none,
                                "stability": "development",
                                "value": "transaction",
                            },
                        ],
                    },
                },
                {
                    "brief": "Namespace of RocketMQ resources, resources in different namespaces are individual.\n",
                    "examples": "myNamespace",
                    "name": "messaging.rocketmq.namespace",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `messaging.destination.subscription.name` instead.\n",
                    "deprecated": {
                        "note": "Replaced by `messaging.destination.subscription.name`.",
                        "reason": "renamed",
                        "renamed_to": "messaging.destination.subscription.name",
                    },
                    "examples": "subscription-a",
                    "name": "messaging.servicebus.destination.subscription_name",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Describes the [settlement type](https://learn.microsoft.com/azure/service-bus-messaging/message-transfers-locks-settlement#peeklock).\n",
                    "name": "messaging.servicebus.disposition_status",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Message is completed",
                                "deprecated": none,
                                "id": "complete",
                                "note": none,
                                "stability": "development",
                                "value": "complete",
                            },
                            {
                                "brief": "Message is abandoned",
                                "deprecated": none,
                                "id": "abandon",
                                "note": none,
                                "stability": "development",
                                "value": "abandon",
                            },
                            {
                                "brief": "Message is sent to dead letter queue",
                                "deprecated": none,
                                "id": "dead_letter",
                                "note": none,
                                "stability": "development",
                                "value": "dead_letter",
                            },
                            {
                                "brief": "Message is deferred",
                                "deprecated": none,
                                "id": "defer",
                                "note": none,
                                "stability": "development",
                                "value": "defer",
                            },
                        ],
                    },
                },
                {
                    "brief": "Number of deliveries that have been attempted for this message.\n",
                    "examples": 2,
                    "name": "messaging.servicebus.message.delivery_count",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The UTC epoch seconds at which the message has been accepted and stored in the entity.\n",
                    "examples": 1701393730,
                    "name": "messaging.servicebus.message.enqueued_time",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The messaging system as identified by the client instrumentation.",
                    "name": "messaging.system",
                    "note": "The actual messaging system may differ from the one known by the client. For example, when using Kafka client libraries to communicate with Azure Event Hubs, the `messaging.system` is set to `kafka` based on the instrumentation's best knowledge.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "messaging",
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
            ],
            "root_namespace": "messaging",
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
