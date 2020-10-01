// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An object that represents the details of the consumer you registered.
type Consumer struct {

	// When you register a consumer, Kinesis Data Streams generates an ARN for it. You
	// need this ARN to be able to call SubscribeToShard (). If you delete a consumer
	// and then create a new one with the same name, it won't have the same ARN. That's
	// because consumer ARNs contain the creation timestamp. This is important to keep
	// in mind if you have IAM policies that reference consumer ARNs.
	//
	// This member is required.
	ConsumerARN *string

	//
	//
	// This member is required.
	ConsumerCreationTimestamp *time.Time

	// The name of the consumer is something you choose when you register the consumer.
	//
	// This member is required.
	ConsumerName *string

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// This member is required.
	ConsumerStatus ConsumerStatus
}

// An object that represents the details of a registered consumer.
type ConsumerDescription struct {

	// When you register a consumer, Kinesis Data Streams generates an ARN for it. You
	// need this ARN to be able to call SubscribeToShard (). If you delete a consumer
	// and then create a new one with the same name, it won't have the same ARN. That's
	// because consumer ARNs contain the creation timestamp. This is important to keep
	// in mind if you have IAM policies that reference consumer ARNs.
	//
	// This member is required.
	ConsumerARN *string

	//
	//
	// This member is required.
	ConsumerCreationTimestamp *time.Time

	// The name of the consumer is something you choose when you register the consumer.
	//
	// This member is required.
	ConsumerName *string

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// This member is required.
	ConsumerStatus ConsumerStatus

	// The ARN of the stream with which you registered the consumer.
	//
	// This member is required.
	StreamARN *string
}

// Represents enhanced metrics types.
type EnhancedMetrics struct {

	// List of shard-level metrics. The following are the valid shard-level metrics.
	// The value "ALL" enhances every metric.
	//
	//     * IncomingBytes
	//
	//     *
	// IncomingRecords
	//
	//     * OutgoingBytes
	//
	//     * OutgoingRecords
	//
	//     *
	// WriteProvisionedThroughputExceeded
	//
	//     * ReadProvisionedThroughputExceeded
	//
	//
	// * IteratorAgeMilliseconds
	//
	//     * ALL
	//
	// For more information, see Monitoring the
	// Amazon Kinesis Data Streams Service with Amazon CloudWatch
	// (https://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	ShardLevelMetrics []MetricsName
}

// The range of possible hash key values for the shard, which is a set of ordered
// contiguous positive integers.
type HashKeyRange struct {

	// The ending hash key of the hash key range.
	//
	// This member is required.
	EndingHashKey *string

	// The starting hash key of the hash key range.
	//
	// This member is required.
	StartingHashKey *string
}

// Represents the output for PutRecords.
type PutRecordsRequestEntry struct {

	// The data blob to put into the record, which is base64-encoded when the blob is
	// serialized. When the data blob (the payload before base64-encoding) is added to
	// the partition key size, the total size must not exceed the maximum record size
	// (1 MB).
	//
	// This member is required.
	Data []byte

	// Determines which shard in the stream the data record is assigned to. Partition
	// keys are Unicode strings with a maximum length limit of 256 characters for each
	// key. Amazon Kinesis Data Streams uses the partition key as input to a hash
	// function that maps the partition key and associated data to a specific shard.
	// Specifically, an MD5 hash function is used to map partition keys to 128-bit
	// integer values and to map associated data records to shards. As a result of this
	// hashing mechanism, all data records with the same partition key map to the same
	// shard within the stream.
	//
	// This member is required.
	PartitionKey *string

	// The hash value used to determine explicitly the shard that the data record is
	// assigned to by overriding the partition key hash.
	ExplicitHashKey *string
}

// Represents the result of an individual record from a PutRecords request. A
// record that is successfully added to a stream includes SequenceNumber and
// ShardId in the result. A record that fails to be added to the stream includes
// ErrorCode and ErrorMessage in the result.
type PutRecordsResultEntry struct {

	// The error code for an individual record result. ErrorCodes can be either
	// ProvisionedThroughputExceededException or InternalFailure.
	ErrorCode *string

	// The error message for an individual record result. An ErrorCode value of
	// ProvisionedThroughputExceededException has an error message that includes the
	// account ID, stream name, and shard ID. An ErrorCode value of InternalFailure has
	// the error message "Internal Service Failure".
	ErrorMessage *string

	// The sequence number for an individual record result.
	SequenceNumber *string

	// The shard ID for an individual record result.
	ShardId *string
}

// The unit of data of the Kinesis data stream, which is composed of a sequence
// number, a partition key, and a data blob.
type Record struct {

	// The data blob. The data in the blob is both opaque and immutable to Kinesis Data
	// Streams, which does not inspect, interpret, or change the data in the blob in
	// any way. When the data blob (the payload before base64-encoding) is added to the
	// partition key size, the total size must not exceed the maximum record size (1
	// MB).
	//
	// This member is required.
	Data []byte

	// Identifies which shard in the stream the data record is assigned to.
	//
	// This member is required.
	PartitionKey *string

	// The unique identifier of the record within its shard.
	//
	// This member is required.
	SequenceNumber *string

	// The approximate time that the record was inserted into the stream.
	ApproximateArrivalTimestamp *time.Time

	// The encryption type used on the record. This parameter can be one of the
	// following values:
	//
	//     * NONE: Do not encrypt the records in the stream.
	//
	//     *
	// KMS: Use server-side encryption on the records in the stream using a
	// customer-managed AWS KMS key.
	EncryptionType EncryptionType
}

// The range of possible sequence numbers for the shard.
type SequenceNumberRange struct {

	// The starting sequence number for the range.
	//
	// This member is required.
	StartingSequenceNumber *string

	// The ending sequence number for the range. Shards that are in the OPEN state have
	// an ending sequence number of null.
	EndingSequenceNumber *string
}

// A uniquely identified group of data records in a Kinesis data stream.
type Shard struct {

	// The range of possible hash key values for the shard, which is a set of ordered
	// contiguous positive integers.
	//
	// This member is required.
	HashKeyRange *HashKeyRange

	// The range of possible sequence numbers for the shard.
	//
	// This member is required.
	SequenceNumberRange *SequenceNumberRange

	// The unique identifier of the shard within the stream.
	//
	// This member is required.
	ShardId *string

	// The shard ID of the shard adjacent to the shard's parent.
	AdjacentParentShardId *string

	// The shard ID of the shard's parent.
	ParentShardId *string
}

type StartingPosition struct {
	Type ShardIteratorType

	SequenceNumber *string

	Timestamp *time.Time
}

// Represents the output for DescribeStream ().
type StreamDescription struct {

	// Represents the current enhanced monitoring settings of the stream.
	//
	// This member is required.
	EnhancedMonitoring []*EnhancedMetrics

	// If set to true, more shards in the stream are available to describe.
	//
	// This member is required.
	HasMoreShards *bool

	// The current retention period, in hours.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The shards that comprise the stream.
	//
	// This member is required.
	Shards []*Shard

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// This member is required.
	StreamARN *string

	// The approximate time that the stream was created.
	//
	// This member is required.
	StreamCreationTimestamp *time.Time

	// The name of the stream being described.
	//
	// This member is required.
	StreamName *string

	// The current status of the stream being described. The stream status is one of
	// the following states:
	//
	//     * CREATING - The stream is being created. Kinesis
	// Data Streams immediately returns and sets StreamStatus to CREATING.
	//
	//     *
	// DELETING - The stream is being deleted. The specified stream is in the DELETING
	// state until Kinesis Data Streams completes the deletion.
	//
	//     * ACTIVE - The
	// stream exists and is ready for read and write operations or deletion. You should
	// perform read and write operations only on an ACTIVE stream.
	//
	//     * UPDATING -
	// Shards in the stream are being merged or split. Read and write operations
	// continue to work while the stream is in the UPDATING state.
	//
	// This member is required.
	StreamStatus StreamStatus

	// The server-side encryption type used on the stream. This parameter can be one of
	// the following values:
	//
	//     * NONE: Do not encrypt the records in the stream.
	//
	//
	// * KMS: Use server-side encryption on the records in the stream using a
	// customer-managed AWS KMS key.
	EncryptionType EncryptionType

	// The GUID for the customer-managed AWS KMS key to use for encryption. This value
	// can be a globally unique identifier, a fully specified ARN to either an alias or
	// a key, or an alias name prefixed by "alias/".You can also use a master key owned
	// by Kinesis Data Streams by specifying the alias aws/kinesis.
	//
	//     * Key ARN
	// example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//
	// * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//     *
	// Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//     *
	// Alias name example: alias/MyAliasName
	//
	//     * Master key owned by Kinesis Data
	// Streams: alias/aws/kinesis
	KeyId *string
}

// Represents the output for DescribeStreamSummary ()
type StreamDescriptionSummary struct {

	// Represents the current enhanced monitoring settings of the stream.
	//
	// This member is required.
	EnhancedMonitoring []*EnhancedMetrics

	// The number of open shards in the stream.
	//
	// This member is required.
	OpenShardCount *int32

	// The current retention period, in hours.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// This member is required.
	StreamARN *string

	// The approximate time that the stream was created.
	//
	// This member is required.
	StreamCreationTimestamp *time.Time

	// The name of the stream being described.
	//
	// This member is required.
	StreamName *string

	// The current status of the stream being described. The stream status is one of
	// the following states:
	//
	//     * CREATING - The stream is being created. Kinesis
	// Data Streams immediately returns and sets StreamStatus to CREATING.
	//
	//     *
	// DELETING - The stream is being deleted. The specified stream is in the DELETING
	// state until Kinesis Data Streams completes the deletion.
	//
	//     * ACTIVE - The
	// stream exists and is ready for read and write operations or deletion. You should
	// perform read and write operations only on an ACTIVE stream.
	//
	//     * UPDATING -
	// Shards in the stream are being merged or split. Read and write operations
	// continue to work while the stream is in the UPDATING state.
	//
	// This member is required.
	StreamStatus StreamStatus

	// The number of enhanced fan-out consumers registered with the stream.
	ConsumerCount *int32

	// The encryption type used. This value is one of the following:
	//
	//     * KMS
	//
	//     *
	// NONE
	EncryptionType EncryptionType

	// The GUID for the customer-managed AWS KMS key to use for encryption. This value
	// can be a globally unique identifier, a fully specified ARN to either an alias or
	// a key, or an alias name prefixed by "alias/".You can also use a master key owned
	// by Kinesis Data Streams by specifying the alias aws/kinesis.
	//
	//     * Key ARN
	// example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//
	// * Alias ARN example:  arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//
	// * Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//     *
	// Alias name example: alias/MyAliasName
	//
	//     * Master key owned by Kinesis Data
	// Streams: alias/aws/kinesis
	KeyId *string
}

// After you call SubscribeToShard (), Kinesis Data Streams sends events of this
// type to your consumer.
type SubscribeToShardEvent struct {

	// Use this as StartingSequenceNumber in the next call to SubscribeToShard ().
	//
	// This member is required.
	ContinuationSequenceNumber *string

	// The number of milliseconds the read records are from the tip of the stream,
	// indicating how far behind current time the consumer is. A value of zero
	// indicates that record processing is caught up, and there are no new records to
	// process at this moment.
	//
	// This member is required.
	MillisBehindLatest *int64

	//
	//
	// This member is required.
	Records []*Record
}

type SubscribeToShardEventStream interface {
	isSubscribeToShardEventStream()
}

// The request was denied due to request throttling. For more information about
// throttling, see Limits
// (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#requests-per-second)
// in the AWS Key Management Service Developer Guide.
type SubscribeToShardEventStreamMemberKMSThrottlingException struct {
	Value *KMSThrottlingException
}

func (*SubscribeToShardEventStreamMemberKMSThrottlingException) isSubscribeToShardEventStream() {}

// The request was rejected because the specified customer master key (CMK) isn't
// enabled.
type SubscribeToShardEventStreamMemberKMSDisabledException struct {
	Value *KMSDisabledException
}

func (*SubscribeToShardEventStreamMemberKMSDisabledException) isSubscribeToShardEventStream() {}

// The ciphertext references a key that doesn't exist or that you don't have access
// to.
type SubscribeToShardEventStreamMemberKMSAccessDeniedException struct {
	Value *KMSAccessDeniedException
}

func (*SubscribeToShardEventStreamMemberKMSAccessDeniedException) isSubscribeToShardEventStream() {}

// The request was rejected because the state of the specified resource isn't valid
// for this request. For more information, see How Key State Affects Use of a
// Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
type SubscribeToShardEventStreamMemberKMSInvalidStateException struct {
	Value *KMSInvalidStateException
}

func (*SubscribeToShardEventStreamMemberKMSInvalidStateException) isSubscribeToShardEventStream() {}

// The request was rejected because the specified entity or resource can't be
// found.
type SubscribeToShardEventStreamMemberKMSNotFoundException struct {
	Value *KMSNotFoundException
}

func (*SubscribeToShardEventStreamMemberKMSNotFoundException) isSubscribeToShardEventStream() {}

// The requested resource could not be found. The stream might not be specified
// correctly.
type SubscribeToShardEventStreamMemberResourceNotFoundException struct {
	Value *ResourceNotFoundException
}

func (*SubscribeToShardEventStreamMemberResourceNotFoundException) isSubscribeToShardEventStream() {}

// After you call SubscribeToShard (), Kinesis Data Streams sends events of this
// type to your consumer.
type SubscribeToShardEventStreamMemberSubscribeToShardEvent struct {
	Value *SubscribeToShardEvent
}

func (*SubscribeToShardEventStreamMemberSubscribeToShardEvent) isSubscribeToShardEventStream() {}

// The resource is not available for this operation. For successful operation, the
// resource must be in the ACTIVE state.
type SubscribeToShardEventStreamMemberResourceInUseException struct {
	Value *ResourceInUseException
}

func (*SubscribeToShardEventStreamMemberResourceInUseException) isSubscribeToShardEventStream() {}

type SubscribeToShardEventStreamMemberInternalFailureException struct {
	Value *InternalFailureException
}

func (*SubscribeToShardEventStreamMemberInternalFailureException) isSubscribeToShardEventStream() {}

// The AWS access key ID needs a subscription for the service.
type SubscribeToShardEventStreamMemberKMSOptInRequired struct {
	Value *KMSOptInRequired
}

func (*SubscribeToShardEventStreamMemberKMSOptInRequired) isSubscribeToShardEventStream() {}

// Metadata assigned to the stream, consisting of a key-value pair.
type Tag struct {

	// A unique identifier for the tag. Maximum length: 128 characters. Valid
	// characters: Unicode letters, digits, white space, _ . / = + - % @
	//
	// This member is required.
	Key *string

	// An optional string, typically used to describe or define the tag. Maximum
	// length: 256 characters. Valid characters: Unicode letters, digits, white space,
	// _ . / = + - % @
	Value *string
}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isSubscribeToShardEventStream() {}
