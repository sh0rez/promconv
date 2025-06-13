package aws

// The unique identifier of the AWS Bedrock Guardrail. A [guardrail] helps safeguard and prevent unwanted behavior from model responses or user messages
//
// [guardrail]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html
type AttrBedrockGuardrailId string // aws.bedrock.guardrail.id

func (AttrBedrockGuardrailId) Development()    {}
func (AttrBedrockGuardrailId) Recommended()    {}
func (AttrBedrockGuardrailId) Key() string     { return "aws_bedrock_guardrail_id" }
func (a AttrBedrockGuardrailId) Value() string { return string(a) }

// The unique identifier of the AWS Bedrock Knowledge base. A [knowledge base] is a bank of information that can be queried by models to generate more relevant responses and augment prompts
//
// [knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html
type AttrBedrockKnowledgeBaseId string // aws.bedrock.knowledge_base.id

func (AttrBedrockKnowledgeBaseId) Development()    {}
func (AttrBedrockKnowledgeBaseId) Recommended()    {}
func (AttrBedrockKnowledgeBaseId) Key() string     { return "aws_bedrock_knowledge_base_id" }
func (a AttrBedrockKnowledgeBaseId) Value() string { return string(a) }

// The JSON-serialized value of each item in the `AttributeDefinitions` request field
type AttrDynamodbAttributeDefinitions string // aws.dynamodb.attribute_definitions

func (AttrDynamodbAttributeDefinitions) Development()    {}
func (AttrDynamodbAttributeDefinitions) Recommended()    {}
func (AttrDynamodbAttributeDefinitions) Key() string     { return "aws_dynamodb_attribute_definitions" }
func (a AttrDynamodbAttributeDefinitions) Value() string { return string(a) }

// The value of the `AttributesToGet` request parameter
type AttrDynamodbAttributesToGet string // aws.dynamodb.attributes_to_get

func (AttrDynamodbAttributesToGet) Development()    {}
func (AttrDynamodbAttributesToGet) Recommended()    {}
func (AttrDynamodbAttributesToGet) Key() string     { return "aws_dynamodb_attributes_to_get" }
func (a AttrDynamodbAttributesToGet) Value() string { return string(a) }

// The value of the `ConsistentRead` request parameter
type AttrDynamodbConsistentRead string // aws.dynamodb.consistent_read

func (AttrDynamodbConsistentRead) Development()    {}
func (AttrDynamodbConsistentRead) Recommended()    {}
func (AttrDynamodbConsistentRead) Key() string     { return "aws_dynamodb_consistent_read" }
func (a AttrDynamodbConsistentRead) Value() string { return string(a) }

// The JSON-serialized value of each item in the `ConsumedCapacity` response field
type AttrDynamodbConsumedCapacity string // aws.dynamodb.consumed_capacity

func (AttrDynamodbConsumedCapacity) Development()    {}
func (AttrDynamodbConsumedCapacity) Recommended()    {}
func (AttrDynamodbConsumedCapacity) Key() string     { return "aws_dynamodb_consumed_capacity" }
func (a AttrDynamodbConsumedCapacity) Value() string { return string(a) }

// The value of the `Count` response parameter
type AttrDynamodbCount string // aws.dynamodb.count

func (AttrDynamodbCount) Development()    {}
func (AttrDynamodbCount) Recommended()    {}
func (AttrDynamodbCount) Key() string     { return "aws_dynamodb_count" }
func (a AttrDynamodbCount) Value() string { return string(a) }

// The value of the `ExclusiveStartTableName` request parameter
type AttrDynamodbExclusiveStartTable string // aws.dynamodb.exclusive_start_table

func (AttrDynamodbExclusiveStartTable) Development()    {}
func (AttrDynamodbExclusiveStartTable) Recommended()    {}
func (AttrDynamodbExclusiveStartTable) Key() string     { return "aws_dynamodb_exclusive_start_table" }
func (a AttrDynamodbExclusiveStartTable) Value() string { return string(a) }

// The JSON-serialized value of each item in the `GlobalSecondaryIndexUpdates` request field
type AttrDynamodbGlobalSecondaryIndexUpdates string // aws.dynamodb.global_secondary_index_updates

func (AttrDynamodbGlobalSecondaryIndexUpdates) Development() {}
func (AttrDynamodbGlobalSecondaryIndexUpdates) Recommended() {}
func (AttrDynamodbGlobalSecondaryIndexUpdates) Key() string {
	return "aws_dynamodb_global_secondary_index_updates"
}
func (a AttrDynamodbGlobalSecondaryIndexUpdates) Value() string { return string(a) }

// The JSON-serialized value of each item of the `GlobalSecondaryIndexes` request field
type AttrDynamodbGlobalSecondaryIndexes string // aws.dynamodb.global_secondary_indexes

func (AttrDynamodbGlobalSecondaryIndexes) Development() {}
func (AttrDynamodbGlobalSecondaryIndexes) Recommended() {}
func (AttrDynamodbGlobalSecondaryIndexes) Key() string {
	return "aws_dynamodb_global_secondary_indexes"
}
func (a AttrDynamodbGlobalSecondaryIndexes) Value() string { return string(a) }

// The value of the `IndexName` request parameter
type AttrDynamodbIndexName string // aws.dynamodb.index_name

func (AttrDynamodbIndexName) Development()    {}
func (AttrDynamodbIndexName) Recommended()    {}
func (AttrDynamodbIndexName) Key() string     { return "aws_dynamodb_index_name" }
func (a AttrDynamodbIndexName) Value() string { return string(a) }

// The JSON-serialized value of the `ItemCollectionMetrics` response field
type AttrDynamodbItemCollectionMetrics string // aws.dynamodb.item_collection_metrics

func (AttrDynamodbItemCollectionMetrics) Development()    {}
func (AttrDynamodbItemCollectionMetrics) Recommended()    {}
func (AttrDynamodbItemCollectionMetrics) Key() string     { return "aws_dynamodb_item_collection_metrics" }
func (a AttrDynamodbItemCollectionMetrics) Value() string { return string(a) }

// The value of the `Limit` request parameter
type AttrDynamodbLimit string // aws.dynamodb.limit

func (AttrDynamodbLimit) Development()    {}
func (AttrDynamodbLimit) Recommended()    {}
func (AttrDynamodbLimit) Key() string     { return "aws_dynamodb_limit" }
func (a AttrDynamodbLimit) Value() string { return string(a) }

// The JSON-serialized value of each item of the `LocalSecondaryIndexes` request field
type AttrDynamodbLocalSecondaryIndexes string // aws.dynamodb.local_secondary_indexes

func (AttrDynamodbLocalSecondaryIndexes) Development()    {}
func (AttrDynamodbLocalSecondaryIndexes) Recommended()    {}
func (AttrDynamodbLocalSecondaryIndexes) Key() string     { return "aws_dynamodb_local_secondary_indexes" }
func (a AttrDynamodbLocalSecondaryIndexes) Value() string { return string(a) }

// The value of the `ProjectionExpression` request parameter
type AttrDynamodbProjection string // aws.dynamodb.projection

func (AttrDynamodbProjection) Development()    {}
func (AttrDynamodbProjection) Recommended()    {}
func (AttrDynamodbProjection) Key() string     { return "aws_dynamodb_projection" }
func (a AttrDynamodbProjection) Value() string { return string(a) }

// The value of the `ProvisionedThroughput.ReadCapacityUnits` request parameter
type AttrDynamodbProvisionedReadCapacity string // aws.dynamodb.provisioned_read_capacity

func (AttrDynamodbProvisionedReadCapacity) Development() {}
func (AttrDynamodbProvisionedReadCapacity) Recommended() {}
func (AttrDynamodbProvisionedReadCapacity) Key() string {
	return "aws_dynamodb_provisioned_read_capacity"
}
func (a AttrDynamodbProvisionedReadCapacity) Value() string { return string(a) }

// The value of the `ProvisionedThroughput.WriteCapacityUnits` request parameter
type AttrDynamodbProvisionedWriteCapacity string // aws.dynamodb.provisioned_write_capacity

func (AttrDynamodbProvisionedWriteCapacity) Development() {}
func (AttrDynamodbProvisionedWriteCapacity) Recommended() {}
func (AttrDynamodbProvisionedWriteCapacity) Key() string {
	return "aws_dynamodb_provisioned_write_capacity"
}
func (a AttrDynamodbProvisionedWriteCapacity) Value() string { return string(a) }

// The value of the `ScanIndexForward` request parameter
type AttrDynamodbScanForward string // aws.dynamodb.scan_forward

func (AttrDynamodbScanForward) Development()    {}
func (AttrDynamodbScanForward) Recommended()    {}
func (AttrDynamodbScanForward) Key() string     { return "aws_dynamodb_scan_forward" }
func (a AttrDynamodbScanForward) Value() string { return string(a) }

// The value of the `ScannedCount` response parameter
type AttrDynamodbScannedCount string // aws.dynamodb.scanned_count

func (AttrDynamodbScannedCount) Development()    {}
func (AttrDynamodbScannedCount) Recommended()    {}
func (AttrDynamodbScannedCount) Key() string     { return "aws_dynamodb_scanned_count" }
func (a AttrDynamodbScannedCount) Value() string { return string(a) }

// The value of the `Segment` request parameter
type AttrDynamodbSegment string // aws.dynamodb.segment

func (AttrDynamodbSegment) Development()    {}
func (AttrDynamodbSegment) Recommended()    {}
func (AttrDynamodbSegment) Key() string     { return "aws_dynamodb_segment" }
func (a AttrDynamodbSegment) Value() string { return string(a) }

// The value of the `Select` request parameter
type AttrDynamodbSelect string // aws.dynamodb.select

func (AttrDynamodbSelect) Development()    {}
func (AttrDynamodbSelect) Recommended()    {}
func (AttrDynamodbSelect) Key() string     { return "aws_dynamodb_select" }
func (a AttrDynamodbSelect) Value() string { return string(a) }

// The number of items in the `TableNames` response parameter
type AttrDynamodbTableCount string // aws.dynamodb.table_count

func (AttrDynamodbTableCount) Development()    {}
func (AttrDynamodbTableCount) Recommended()    {}
func (AttrDynamodbTableCount) Key() string     { return "aws_dynamodb_table_count" }
func (a AttrDynamodbTableCount) Value() string { return string(a) }

// The keys in the `RequestItems` object field
type AttrDynamodbTableNames string // aws.dynamodb.table_names

func (AttrDynamodbTableNames) Development()    {}
func (AttrDynamodbTableNames) Recommended()    {}
func (AttrDynamodbTableNames) Key() string     { return "aws_dynamodb_table_names" }
func (a AttrDynamodbTableNames) Value() string { return string(a) }

// The value of the `TotalSegments` request parameter
type AttrDynamodbTotalSegments string // aws.dynamodb.total_segments

func (AttrDynamodbTotalSegments) Development()    {}
func (AttrDynamodbTotalSegments) Recommended()    {}
func (AttrDynamodbTotalSegments) Key() string     { return "aws_dynamodb_total_segments" }
func (a AttrDynamodbTotalSegments) Value() string { return string(a) }

// The ARN of an [ECS cluster]
//
// [ECS cluster]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html
type AttrEcsClusterArn string // aws.ecs.cluster.arn

func (AttrEcsClusterArn) Development()    {}
func (AttrEcsClusterArn) Recommended()    {}
func (AttrEcsClusterArn) Key() string     { return "aws_ecs_cluster_arn" }
func (a AttrEcsClusterArn) Value() string { return string(a) }

// The Amazon Resource Name (ARN) of an [ECS container instance]
//
// [ECS container instance]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html
type AttrEcsContainerArn string // aws.ecs.container.arn

func (AttrEcsContainerArn) Development()    {}
func (AttrEcsContainerArn) Recommended()    {}
func (AttrEcsContainerArn) Key() string     { return "aws_ecs_container_arn" }
func (a AttrEcsContainerArn) Value() string { return string(a) }

// The [launch type] for an ECS task
//
// [launch type]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html
type AttrEcsLaunchtype string // aws.ecs.launchtype

func (AttrEcsLaunchtype) Development()    {}
func (AttrEcsLaunchtype) Recommended()    {}
func (AttrEcsLaunchtype) Key() string     { return "aws_ecs_launchtype" }
func (a AttrEcsLaunchtype) Value() string { return string(a) }

const EcsLaunchtypeEc2 AttrEcsLaunchtype = "ec2"
const EcsLaunchtypeFargate AttrEcsLaunchtype = "fargate"

// The ARN of a running [ECS task]
//
// [ECS task]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
type AttrEcsTaskArn string // aws.ecs.task.arn

func (AttrEcsTaskArn) Development()    {}
func (AttrEcsTaskArn) Recommended()    {}
func (AttrEcsTaskArn) Key() string     { return "aws_ecs_task_arn" }
func (a AttrEcsTaskArn) Value() string { return string(a) }

// The family name of the [ECS task definition] used to create the ECS task
//
// [ECS task definition]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html
type AttrEcsTaskFamily string // aws.ecs.task.family

func (AttrEcsTaskFamily) Development()    {}
func (AttrEcsTaskFamily) Recommended()    {}
func (AttrEcsTaskFamily) Key() string     { return "aws_ecs_task_family" }
func (a AttrEcsTaskFamily) Value() string { return string(a) }

// The ID of a running ECS task. The ID MUST be extracted from `task.arn`
type AttrEcsTaskId string // aws.ecs.task.id

func (AttrEcsTaskId) Development()    {}
func (AttrEcsTaskId) Recommended()    {}
func (AttrEcsTaskId) Key() string     { return "aws_ecs_task_id" }
func (a AttrEcsTaskId) Value() string { return string(a) }

// The revision for the task definition used to create the ECS task
type AttrEcsTaskRevision string // aws.ecs.task.revision

func (AttrEcsTaskRevision) Development()    {}
func (AttrEcsTaskRevision) Recommended()    {}
func (AttrEcsTaskRevision) Key() string     { return "aws_ecs_task_revision" }
func (a AttrEcsTaskRevision) Value() string { return string(a) }

// The ARN of an EKS cluster
type AttrEksClusterArn string // aws.eks.cluster.arn

func (AttrEksClusterArn) Development()    {}
func (AttrEksClusterArn) Recommended()    {}
func (AttrEksClusterArn) Key() string     { return "aws_eks_cluster_arn" }
func (a AttrEksClusterArn) Value() string { return string(a) }

// The AWS extended request ID as returned in the response header `x-amz-id-2`
type AttrExtendedRequestId string // aws.extended_request_id

func (AttrExtendedRequestId) Development()    {}
func (AttrExtendedRequestId) Recommended()    {}
func (AttrExtendedRequestId) Key() string     { return "aws_extended_request_id" }
func (a AttrExtendedRequestId) Value() string { return string(a) }

// The name of the AWS Kinesis [stream] the request refers to. Corresponds to the `--stream-name` parameter of the Kinesis [describe-stream] operation
//
// [stream]: https://docs.aws.amazon.com/streams/latest/dev/introduction.html
// [describe-stream]: https://docs.aws.amazon.com/cli/latest/reference/kinesis/describe-stream.html
type AttrKinesisStreamName string // aws.kinesis.stream_name

func (AttrKinesisStreamName) Development()    {}
func (AttrKinesisStreamName) Recommended()    {}
func (AttrKinesisStreamName) Key() string     { return "aws_kinesis_stream_name" }
func (a AttrKinesisStreamName) Value() string { return string(a) }

// The full invoked ARN as provided on the `Context` passed to the function (`Lambda-Runtime-Invoked-Function-Arn` header on the `/runtime/invocation/next` applicable).
//
// This may be different from `cloud.resource_id` if an alias is involved
type AttrLambdaInvokedArn string // aws.lambda.invoked_arn

func (AttrLambdaInvokedArn) Development()    {}
func (AttrLambdaInvokedArn) Recommended()    {}
func (AttrLambdaInvokedArn) Key() string     { return "aws_lambda_invoked_arn" }
func (a AttrLambdaInvokedArn) Value() string { return string(a) }

// The UUID of the [AWS Lambda EvenSource Mapping]. An event source is mapped to a lambda function. It's contents are read by Lambda and used to trigger a function. This isn't available in the lambda execution context or the lambda runtime environtment. This is going to be populated by the AWS SDK for each language when that UUID is present. Some of these operations are Create/Delete/Get/List/Update EventSourceMapping
//
// [AWS Lambda EvenSource Mapping]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html
type AttrLambdaResourceMappingId string // aws.lambda.resource_mapping.id

func (AttrLambdaResourceMappingId) Development()    {}
func (AttrLambdaResourceMappingId) Recommended()    {}
func (AttrLambdaResourceMappingId) Key() string     { return "aws_lambda_resource_mapping_id" }
func (a AttrLambdaResourceMappingId) Value() string { return string(a) }

// The Amazon Resource Name(s) (ARN) of the AWS log group(s).
//
// See the [log group ARN format documentation]
//
// [log group ARN format documentation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format
type AttrLogGroupArns string // aws.log.group.arns

func (AttrLogGroupArns) Development()    {}
func (AttrLogGroupArns) Recommended()    {}
func (AttrLogGroupArns) Key() string     { return "aws_log_group_arns" }
func (a AttrLogGroupArns) Value() string { return string(a) }

// The name(s) of the AWS log group(s) an application is writing to.
//
// Multiple log groups must be supported for cases like multi-container applications, where a single application has sidecar containers, and each write to their own log group
type AttrLogGroupNames string // aws.log.group.names

func (AttrLogGroupNames) Development()    {}
func (AttrLogGroupNames) Recommended()    {}
func (AttrLogGroupNames) Key() string     { return "aws_log_group_names" }
func (a AttrLogGroupNames) Value() string { return string(a) }

// The ARN(s) of the AWS log stream(s).
//
// See the [log stream ARN format documentation]. One log group can contain several log streams, so these ARNs necessarily identify both a log group and a log stream
//
// [log stream ARN format documentation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format
type AttrLogStreamArns string // aws.log.stream.arns

func (AttrLogStreamArns) Development()    {}
func (AttrLogStreamArns) Recommended()    {}
func (AttrLogStreamArns) Key() string     { return "aws_log_stream_arns" }
func (a AttrLogStreamArns) Value() string { return string(a) }

// The name(s) of the AWS log stream(s) an application is writing to
type AttrLogStreamNames string // aws.log.stream.names

func (AttrLogStreamNames) Development()    {}
func (AttrLogStreamNames) Recommended()    {}
func (AttrLogStreamNames) Key() string     { return "aws_log_stream_names" }
func (a AttrLogStreamNames) Value() string { return string(a) }

// The AWS request ID as returned in the response headers `x-amzn-requestid`, `x-amzn-request-id` or `x-amz-request-id`
type AttrRequestId string // aws.request_id

func (AttrRequestId) Development()    {}
func (AttrRequestId) Recommended()    {}
func (AttrRequestId) Key() string     { return "aws_request_id" }
func (a AttrRequestId) Value() string { return string(a) }

// The S3 bucket name the request refers to. Corresponds to the `--bucket` parameter of the [S3 API] operations.
// The `bucket` attribute is applicable to all S3 operations that reference a bucket, i.e. that require the bucket name as a mandatory parameter.
// This applies to almost all S3 operations except `list-buckets`
//
// [S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html
type AttrS3Bucket string // aws.s3.bucket

func (AttrS3Bucket) Development()    {}
func (AttrS3Bucket) Recommended()    {}
func (AttrS3Bucket) Key() string     { return "aws_s3_bucket" }
func (a AttrS3Bucket) Value() string { return string(a) }

// The source object (in the form `bucket`/`key`) for the copy operation.
// The `copy_source` attribute applies to S3 copy operations and corresponds to the `--copy-source` parameter
// of the [copy-object operation within the S3 API].
// This applies in particular to the following operations:
//
//   - [copy-object]
//   - [upload-part-copy]
//
// [copy-object operation within the S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html
// [copy-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html
// [upload-part-copy]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html
type AttrS3CopySource string // aws.s3.copy_source

func (AttrS3CopySource) Development()    {}
func (AttrS3CopySource) Recommended()    {}
func (AttrS3CopySource) Key() string     { return "aws_s3_copy_source" }
func (a AttrS3CopySource) Value() string { return string(a) }

// The delete request container that specifies the objects to be deleted.
// The `delete` attribute is only applicable to the [delete-object] operation.
// The `delete` attribute corresponds to the `--delete` parameter of the
// [delete-objects operation within the S3 API]
//
// [delete-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-object.html
// [delete-objects operation within the S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-objects.html
type AttrS3Delete string // aws.s3.delete

func (AttrS3Delete) Development()    {}
func (AttrS3Delete) Recommended()    {}
func (AttrS3Delete) Key() string     { return "aws_s3_delete" }
func (a AttrS3Delete) Value() string { return string(a) }

// The S3 object key the request refers to. Corresponds to the `--key` parameter of the [S3 API] operations.
// The `key` attribute is applicable to all object-related S3 operations, i.e. that require the object key as a mandatory parameter.
// This applies in particular to the following operations:
//
//   - [copy-object]
//   - [delete-object]
//   - [get-object]
//   - [head-object]
//   - [put-object]
//   - [restore-object]
//   - [select-object-content]
//   - [abort-multipart-upload]
//   - [complete-multipart-upload]
//   - [create-multipart-upload]
//   - [list-parts]
//   - [upload-part]
//   - [upload-part-copy]
//
// [S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html
// [copy-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html
// [delete-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-object.html
// [get-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/get-object.html
// [head-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/head-object.html
// [put-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/put-object.html
// [restore-object]: https://docs.aws.amazon.com/cli/latest/reference/s3api/restore-object.html
// [select-object-content]: https://docs.aws.amazon.com/cli/latest/reference/s3api/select-object-content.html
// [abort-multipart-upload]: https://docs.aws.amazon.com/cli/latest/reference/s3api/abort-multipart-upload.html
// [complete-multipart-upload]: https://docs.aws.amazon.com/cli/latest/reference/s3api/complete-multipart-upload.html
// [create-multipart-upload]: https://docs.aws.amazon.com/cli/latest/reference/s3api/create-multipart-upload.html
// [list-parts]: https://docs.aws.amazon.com/cli/latest/reference/s3api/list-parts.html
// [upload-part]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html
// [upload-part-copy]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html
type AttrS3Key string // aws.s3.key

func (AttrS3Key) Development()    {}
func (AttrS3Key) Recommended()    {}
func (AttrS3Key) Key() string     { return "aws_s3_key" }
func (a AttrS3Key) Value() string { return string(a) }

// The part number of the part being uploaded in a multipart-upload operation. This is a positive integer between 1 and 10,000.
// The `part_number` attribute is only applicable to the [upload-part]
// and [upload-part-copy] operations.
// The `part_number` attribute corresponds to the `--part-number` parameter of the
// [upload-part operation within the S3 API]
//
// [upload-part]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html
// [upload-part-copy]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html
// [upload-part operation within the S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html
type AttrS3PartNumber string // aws.s3.part_number

func (AttrS3PartNumber) Development()    {}
func (AttrS3PartNumber) Recommended()    {}
func (AttrS3PartNumber) Key() string     { return "aws_s3_part_number" }
func (a AttrS3PartNumber) Value() string { return string(a) }

// Upload ID that identifies the multipart upload.
// The `upload_id` attribute applies to S3 multipart-upload operations and corresponds to the `--upload-id` parameter
// of the [S3 API] multipart operations.
// This applies in particular to the following operations:
//
//   - [abort-multipart-upload]
//   - [complete-multipart-upload]
//   - [list-parts]
//   - [upload-part]
//   - [upload-part-copy]
//
// [S3 API]: https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html
// [abort-multipart-upload]: https://docs.aws.amazon.com/cli/latest/reference/s3api/abort-multipart-upload.html
// [complete-multipart-upload]: https://docs.aws.amazon.com/cli/latest/reference/s3api/complete-multipart-upload.html
// [list-parts]: https://docs.aws.amazon.com/cli/latest/reference/s3api/list-parts.html
// [upload-part]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html
// [upload-part-copy]: https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html
type AttrS3UploadId string // aws.s3.upload_id

func (AttrS3UploadId) Development()    {}
func (AttrS3UploadId) Recommended()    {}
func (AttrS3UploadId) Key() string     { return "aws_s3_upload_id" }
func (a AttrS3UploadId) Value() string { return string(a) }

// The ARN of the Secret stored in the Secrets Mangger
type AttrSecretsmanagerSecretArn string // aws.secretsmanager.secret.arn

func (AttrSecretsmanagerSecretArn) Development()    {}
func (AttrSecretsmanagerSecretArn) Recommended()    {}
func (AttrSecretsmanagerSecretArn) Key() string     { return "aws_secretsmanager_secret_arn" }
func (a AttrSecretsmanagerSecretArn) Value() string { return string(a) }

// The ARN of the AWS SNS Topic. An Amazon SNS [topic] is a logical access point that acts as a communication channel
//
// [topic]: https://docs.aws.amazon.com/sns/latest/dg/sns-create-topic.html
type AttrSnsTopicArn string // aws.sns.topic.arn

func (AttrSnsTopicArn) Development()    {}
func (AttrSnsTopicArn) Recommended()    {}
func (AttrSnsTopicArn) Key() string     { return "aws_sns_topic_arn" }
func (a AttrSnsTopicArn) Value() string { return string(a) }

// The URL of the AWS SQS Queue. It's a unique identifier for a queue in Amazon Simple Queue Service (SQS) and is used to access the queue and perform actions on it
type AttrSqsQueueUrl string // aws.sqs.queue.url

func (AttrSqsQueueUrl) Development()    {}
func (AttrSqsQueueUrl) Recommended()    {}
func (AttrSqsQueueUrl) Key() string     { return "aws_sqs_queue_url" }
func (a AttrSqsQueueUrl) Value() string { return string(a) }

// The ARN of the AWS Step Functions Activity
type AttrStepFunctionsActivityArn string // aws.step_functions.activity.arn

func (AttrStepFunctionsActivityArn) Development()    {}
func (AttrStepFunctionsActivityArn) Recommended()    {}
func (AttrStepFunctionsActivityArn) Key() string     { return "aws_step_functions_activity_arn" }
func (a AttrStepFunctionsActivityArn) Value() string { return string(a) }

// The ARN of the AWS Step Functions State Machine
type AttrStepFunctionsStateMachineArn string // aws.step_functions.state_machine.arn

func (AttrStepFunctionsStateMachineArn) Development()    {}
func (AttrStepFunctionsStateMachineArn) Recommended()    {}
func (AttrStepFunctionsStateMachineArn) Key() string     { return "aws_step_functions_state_machine_arn" }
func (a AttrStepFunctionsStateMachineArn) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The unique identifier of the AWS Bedrock Guardrail. A [guardrail](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html) helps safeguard and prevent unwanted behavior from model responses or user messages.\n",
                    "examples": [
                        "sgi5gkybzqak",
                    ],
                    "name": "aws.bedrock.guardrail.id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique identifier of the AWS Bedrock Knowledge base. A [knowledge base](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html) is a bank of information that can be queried by models to generate more relevant responses and augment prompts.\n",
                    "examples": [
                        "XFWUPB9PAW",
                    ],
                    "name": "aws.bedrock.knowledge_base.id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The JSON-serialized value of each item in the `AttributeDefinitions` request field.",
                    "examples": [
                        [
                            "{ \"AttributeName\": \"string\", \"AttributeType\": \"string\" }",
                        ],
                    ],
                    "name": "aws.dynamodb.attribute_definitions",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `AttributesToGet` request parameter.",
                    "examples": [
                        [
                            "lives",
                            "id",
                        ],
                    ],
                    "name": "aws.dynamodb.attributes_to_get",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `ConsistentRead` request parameter.",
                    "name": "aws.dynamodb.consistent_read",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The JSON-serialized value of each item in the `ConsumedCapacity` response field.",
                    "examples": [
                        [
                            "{ \"CapacityUnits\": number, \"GlobalSecondaryIndexes\": { \"string\" : { \"CapacityUnits\": number, \"ReadCapacityUnits\": number, \"WriteCapacityUnits\": number } }, \"LocalSecondaryIndexes\": { \"string\" : { \"CapacityUnits\": number, \"ReadCapacityUnits\": number, \"WriteCapacityUnits\": number } }, \"ReadCapacityUnits\": number, \"Table\": { \"CapacityUnits\": number, \"ReadCapacityUnits\": number, \"WriteCapacityUnits\": number }, \"TableName\": \"string\", \"WriteCapacityUnits\": number }",
                        ],
                    ],
                    "name": "aws.dynamodb.consumed_capacity",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `Count` response parameter.",
                    "examples": [
                        10,
                    ],
                    "name": "aws.dynamodb.count",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The value of the `ExclusiveStartTableName` request parameter.",
                    "examples": [
                        "Users",
                        "CatsTable",
                    ],
                    "name": "aws.dynamodb.exclusive_start_table",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The JSON-serialized value of each item in the `GlobalSecondaryIndexUpdates` request field.",
                    "examples": [
                        [
                            "{ \"Create\": { \"IndexName\": \"string\", \"KeySchema\": [ { \"AttributeName\": \"string\", \"KeyType\": \"string\" } ], \"Projection\": { \"NonKeyAttributes\": [ \"string\" ], \"ProjectionType\": \"string\" }, \"ProvisionedThroughput\": { \"ReadCapacityUnits\": number, \"WriteCapacityUnits\": number } }",
                        ],
                    ],
                    "name": "aws.dynamodb.global_secondary_index_updates",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The JSON-serialized value of each item of the `GlobalSecondaryIndexes` request field",
                    "examples": [
                        [
                            "{ \"IndexName\": \"string\", \"KeySchema\": [ { \"AttributeName\": \"string\", \"KeyType\": \"string\" } ], \"Projection\": { \"NonKeyAttributes\": [ \"string\" ], \"ProjectionType\": \"string\" }, \"ProvisionedThroughput\": { \"ReadCapacityUnits\": number, \"WriteCapacityUnits\": number } }",
                        ],
                    ],
                    "name": "aws.dynamodb.global_secondary_indexes",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `IndexName` request parameter.",
                    "examples": [
                        "name_to_group",
                    ],
                    "name": "aws.dynamodb.index_name",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The JSON-serialized value of the `ItemCollectionMetrics` response field.",
                    "examples": [
                        "{ \"string\" : [ { \"ItemCollectionKey\": { \"string\" : { \"B\": blob, \"BOOL\": boolean, \"BS\": [ blob ], \"L\": [ \"AttributeValue\" ], \"M\": { \"string\" : \"AttributeValue\" }, \"N\": \"string\", \"NS\": [ \"string\" ], \"NULL\": boolean, \"S\": \"string\", \"SS\": [ \"string\" ] } }, \"SizeEstimateRangeGB\": [ number ] } ] }",
                    ],
                    "name": "aws.dynamodb.item_collection_metrics",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The value of the `Limit` request parameter.",
                    "examples": [
                        10,
                    ],
                    "name": "aws.dynamodb.limit",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The JSON-serialized value of each item of the `LocalSecondaryIndexes` request field.",
                    "examples": [
                        [
                            "{ \"IndexArn\": \"string\", \"IndexName\": \"string\", \"IndexSizeBytes\": number, \"ItemCount\": number, \"KeySchema\": [ { \"AttributeName\": \"string\", \"KeyType\": \"string\" } ], \"Projection\": { \"NonKeyAttributes\": [ \"string\" ], \"ProjectionType\": \"string\" } }",
                        ],
                    ],
                    "name": "aws.dynamodb.local_secondary_indexes",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `ProjectionExpression` request parameter.",
                    "examples": [
                        "Title",
                        "Title, Price, Color",
                        "Title, Description, RelatedItems, ProductReviews",
                    ],
                    "name": "aws.dynamodb.projection",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The value of the `ProvisionedThroughput.ReadCapacityUnits` request parameter.",
                    "examples": [
                        1.0,
                        2.0,
                    ],
                    "name": "aws.dynamodb.provisioned_read_capacity",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "The value of the `ProvisionedThroughput.WriteCapacityUnits` request parameter.",
                    "examples": [
                        1.0,
                        2.0,
                    ],
                    "name": "aws.dynamodb.provisioned_write_capacity",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "The value of the `ScanIndexForward` request parameter.",
                    "name": "aws.dynamodb.scan_forward",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The value of the `ScannedCount` response parameter.",
                    "examples": [
                        50,
                    ],
                    "name": "aws.dynamodb.scanned_count",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The value of the `Segment` request parameter.",
                    "examples": [
                        10,
                    ],
                    "name": "aws.dynamodb.segment",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The value of the `Select` request parameter.",
                    "examples": [
                        "ALL_ATTRIBUTES",
                        "COUNT",
                    ],
                    "name": "aws.dynamodb.select",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The number of items in the `TableNames` response parameter.",
                    "examples": [
                        20,
                    ],
                    "name": "aws.dynamodb.table_count",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The keys in the `RequestItems` object field.",
                    "examples": [
                        [
                            "Users",
                            "Cats",
                        ],
                    ],
                    "name": "aws.dynamodb.table_names",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The value of the `TotalSegments` request parameter.",
                    "examples": [
                        100,
                    ],
                    "name": "aws.dynamodb.total_segments",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The ARN of an [ECS cluster](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/clusters.html).\n",
                    "examples": [
                        "arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster",
                    ],
                    "name": "aws.ecs.cluster.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The Amazon Resource Name (ARN) of an [ECS container instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_instances.html).\n",
                    "examples": [
                        "arn:aws:ecs:us-west-1:123456789123:container/32624152-9086-4f0e-acae-1a75b14fe4d9",
                    ],
                    "name": "aws.ecs.container.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [launch type](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) for an ECS task.\n",
                    "name": "aws.ecs.launchtype",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ec2",
                                "note": none,
                                "stability": "development",
                                "value": "ec2",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "fargate",
                                "note": none,
                                "stability": "development",
                                "value": "fargate",
                            },
                        ],
                    },
                },
                {
                    "brief": "The ARN of a running [ECS task](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids).\n",
                    "examples": [
                        "arn:aws:ecs:us-west-1:123456789123:task/10838bed-421f-43ef-870a-f43feacbbb5b",
                        "arn:aws:ecs:us-west-1:123456789123:task/my-cluster/task-id/23ebb8ac-c18f-46c6-8bbe-d55d0e37cfbd",
                    ],
                    "name": "aws.ecs.task.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The family name of the [ECS task definition](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html) used to create the ECS task.\n",
                    "examples": [
                        "opentelemetry-family",
                    ],
                    "name": "aws.ecs.task.family",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ID of a running ECS task. The ID MUST be extracted from `task.arn`.\n",
                    "examples": [
                        "10838bed-421f-43ef-870a-f43feacbbb5b",
                        "23ebb8ac-c18f-46c6-8bbe-d55d0e37cfbd",
                    ],
                    "name": "aws.ecs.task.id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The revision for the task definition used to create the ECS task.\n",
                    "examples": [
                        "8",
                        "26",
                    ],
                    "name": "aws.ecs.task.revision",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ARN of an EKS cluster.\n",
                    "examples": [
                        "arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster",
                    ],
                    "name": "aws.eks.cluster.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The AWS extended request ID as returned in the response header `x-amz-id-2`.",
                    "examples": [
                        "wzHcyEWfmOGDIE5QOhTAqFDoDWP3y8IUvpNINCwL9N4TEHbUw0/gZJ+VZTmCNCWR7fezEN3eCiQ=",
                    ],
                    "name": "aws.extended_request_id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the AWS Kinesis [stream](https://docs.aws.amazon.com/streams/latest/dev/introduction.html) the request refers to. Corresponds to the `--stream-name` parameter of the Kinesis [describe-stream](https://docs.aws.amazon.com/cli/latest/reference/kinesis/describe-stream.html) operation.\n",
                    "examples": [
                        "some-stream-name",
                    ],
                    "name": "aws.kinesis.stream_name",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full invoked ARN as provided on the `Context` passed to the function (`Lambda-Runtime-Invoked-Function-Arn` header on the `/runtime/invocation/next` applicable).\n",
                    "examples": [
                        "arn:aws:lambda:us-east-1:123456:function:myfunction:myalias",
                    ],
                    "name": "aws.lambda.invoked_arn",
                    "note": "This may be different from `cloud.resource_id` if an alias is involved.",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The UUID of the [AWS Lambda EvenSource Mapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html). An event source is mapped to a lambda function. It's contents are read by Lambda and used to trigger a function. This isn't available in the lambda execution context or the lambda runtime environtment. This is going to be populated by the AWS SDK for each language when that UUID is present. Some of these operations are Create/Delete/Get/List/Update EventSourceMapping.\n",
                    "examples": [
                        "587ad24b-03b9-4413-8202-bbd56b36e5b7",
                    ],
                    "name": "aws.lambda.resource_mapping.id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The Amazon Resource Name(s) (ARN) of the AWS log group(s).\n",
                    "examples": [
                        [
                            "arn:aws:logs:us-west-1:123456789012:log-group:/aws/my/group:*",
                        ],
                    ],
                    "name": "aws.log.group.arns",
                    "note": "See the [log group ARN format documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The name(s) of the AWS log group(s) an application is writing to.\n",
                    "examples": [
                        [
                            "/aws/lambda/my-function",
                            "opentelemetry-service",
                        ],
                    ],
                    "name": "aws.log.group.names",
                    "note": "Multiple log groups must be supported for cases like multi-container applications, where a single application has sidecar containers, and each write to their own log group.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The ARN(s) of the AWS log stream(s).\n",
                    "examples": [
                        [
                            "arn:aws:logs:us-west-1:123456789012:log-group:/aws/my/group:log-stream:logs/main/10838bed-421f-43ef-870a-f43feacbbb5b",
                        ],
                    ],
                    "name": "aws.log.stream.arns",
                    "note": "See the [log stream ARN format documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format). One log group can contain several log streams, so these ARNs necessarily identify both a log group and a log stream.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The name(s) of the AWS log stream(s) an application is writing to.\n",
                    "examples": [
                        [
                            "logs/main/10838bed-421f-43ef-870a-f43feacbbb5b",
                        ],
                    ],
                    "name": "aws.log.stream.names",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The AWS request ID as returned in the response headers `x-amzn-requestid`, `x-amzn-request-id` or `x-amz-request-id`.",
                    "examples": [
                        "79b9da39-b7ae-508a-a6bc-864b2829c622",
                        "C9ER4AJX75574TDJ",
                    ],
                    "name": "aws.request_id",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The S3 bucket name the request refers to. Corresponds to the `--bucket` parameter of the [S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html) operations.",
                    "examples": [
                        "some-bucket-name",
                    ],
                    "name": "aws.s3.bucket",
                    "note": "The `bucket` attribute is applicable to all S3 operations that reference a bucket, i.e. that require the bucket name as a mandatory parameter.\nThis applies to almost all S3 operations except `list-buckets`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The source object (in the form `bucket`/`key`) for the copy operation.",
                    "examples": [
                        "someFile.yml",
                    ],
                    "name": "aws.s3.copy_source",
                    "note": "The `copy_source` attribute applies to S3 copy operations and corresponds to the `--copy-source` parameter\nof the [copy-object operation within the S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html).\nThis applies in particular to the following operations:\n\n- [copy-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html)\n- [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The delete request container that specifies the objects to be deleted.",
                    "examples": [
                        "Objects=[{Key=string,VersionId=string},{Key=string,VersionId=string}],Quiet=boolean",
                    ],
                    "name": "aws.s3.delete",
                    "note": "The `delete` attribute is only applicable to the [delete-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-object.html) operation.\nThe `delete` attribute corresponds to the `--delete` parameter of the\n[delete-objects operation within the S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-objects.html).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The S3 object key the request refers to. Corresponds to the `--key` parameter of the [S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html) operations.",
                    "examples": [
                        "someFile.yml",
                    ],
                    "name": "aws.s3.key",
                    "note": "The `key` attribute is applicable to all object-related S3 operations, i.e. that require the object key as a mandatory parameter.\nThis applies in particular to the following operations:\n\n- [copy-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/copy-object.html)\n- [delete-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/delete-object.html)\n- [get-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/get-object.html)\n- [head-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/head-object.html)\n- [put-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/put-object.html)\n- [restore-object](https://docs.aws.amazon.com/cli/latest/reference/s3api/restore-object.html)\n- [select-object-content](https://docs.aws.amazon.com/cli/latest/reference/s3api/select-object-content.html)\n- [abort-multipart-upload](https://docs.aws.amazon.com/cli/latest/reference/s3api/abort-multipart-upload.html)\n- [complete-multipart-upload](https://docs.aws.amazon.com/cli/latest/reference/s3api/complete-multipart-upload.html)\n- [create-multipart-upload](https://docs.aws.amazon.com/cli/latest/reference/s3api/create-multipart-upload.html)\n- [list-parts](https://docs.aws.amazon.com/cli/latest/reference/s3api/list-parts.html)\n- [upload-part](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html)\n- [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The part number of the part being uploaded in a multipart-upload operation. This is a positive integer between 1 and 10,000.",
                    "examples": [
                        3456,
                    ],
                    "name": "aws.s3.part_number",
                    "note": "The `part_number` attribute is only applicable to the [upload-part](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html)\nand [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html) operations.\nThe `part_number` attribute corresponds to the `--part-number` parameter of the\n[upload-part operation within the S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Upload ID that identifies the multipart upload.",
                    "examples": [
                        "dfRtDYWFbkRONycy.Yxwh66Yjlx.cph0gtNBtJ",
                    ],
                    "name": "aws.s3.upload_id",
                    "note": "The `upload_id` attribute applies to S3 multipart-upload operations and corresponds to the `--upload-id` parameter\nof the [S3 API](https://docs.aws.amazon.com/cli/latest/reference/s3api/index.html) multipart operations.\nThis applies in particular to the following operations:\n\n- [abort-multipart-upload](https://docs.aws.amazon.com/cli/latest/reference/s3api/abort-multipart-upload.html)\n- [complete-multipart-upload](https://docs.aws.amazon.com/cli/latest/reference/s3api/complete-multipart-upload.html)\n- [list-parts](https://docs.aws.amazon.com/cli/latest/reference/s3api/list-parts.html)\n- [upload-part](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part.html)\n- [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html)\n",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ARN of the Secret stored in the Secrets Mangger\n",
                    "examples": [
                        "arn:aws:secretsmanager:us-east-1:123456789012:secret:SecretName-6RandomCharacters",
                    ],
                    "name": "aws.secretsmanager.secret.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ARN of the AWS SNS Topic. An Amazon SNS [topic](https://docs.aws.amazon.com/sns/latest/dg/sns-create-topic.html) is a logical access point that acts as a communication channel.\n",
                    "examples": [
                        "arn:aws:sns:us-east-1:123456789012:mystack-mytopic-NZJ5JSMVGFIE",
                    ],
                    "name": "aws.sns.topic.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The URL of the AWS SQS Queue. It's a unique identifier for a queue in Amazon Simple Queue Service (SQS) and is used to access the queue and perform actions on it.\n",
                    "examples": [
                        "https://sqs.us-east-1.amazonaws.com/123456789012/MyQueue",
                    ],
                    "name": "aws.sqs.queue.url",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ARN of the AWS Step Functions Activity.\n",
                    "examples": [
                        "arn:aws:states:us-east-1:123456789012:activity:get-greeting",
                    ],
                    "name": "aws.step_functions.activity.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The ARN of the AWS Step Functions State Machine.\n",
                    "examples": [
                        "arn:aws:states:us-east-1:123456789012:stateMachine:myStateMachine:1",
                    ],
                    "name": "aws.step_functions.state_machine.arn",
                    "requirement_level": "recommended",
                    "root_namespace": "aws",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "aws",
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
