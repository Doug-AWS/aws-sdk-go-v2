// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The Amazon S3 buckets or AWS Lambda functions that you specify in your event
// selectors for your trail to log data events. Data events provide information
// about the resource operations performed on or within a resource itself. These
// are also known as data plane operations. You can specify up to 250 data
// resources for a trail. The total number of allowed data resources is 250. This
// number can be distributed between 1 and 5 event selectors, but the total cannot
// exceed 250 across all selectors. The following example demonstrates how logging
// works when you configure logging of all data events for an S3 bucket named
// bucket-1. In this example, the CloudTrail user specified an empty prefix, and
// the option to log both Read and Write data events.
//
//     * A user uploads an
// image file to bucket-1.
//
//     * The PutObject API operation is an Amazon S3
// object-level API. It is recorded as a data event in CloudTrail. Because the
// CloudTrail user specified an S3 bucket with an empty prefix, events that occur
// on any object in that bucket are logged. The trail processes and logs the
// event.
//
//     * A user uploads an object to an Amazon S3 bucket named
// arn:aws:s3:::bucket-2.
//
//     * The PutObject API operation occurred for an object
// in an S3 bucket that the CloudTrail user didn't specify for the trail. The trail
// doesn’t log the event.
//
// The following example demonstrates how logging works
// when you configure logging of AWS Lambda data events for a Lambda function named
// MyLambdaFunction, but not for all AWS Lambda functions.
//
//     * A user runs a
// script that includes a call to the MyLambdaFunction function and the
// MyOtherLambdaFunction function.
//
//     * The Invoke API operation on
// MyLambdaFunction is an AWS Lambda API. It is recorded as a data event in
// CloudTrail. Because the CloudTrail user specified logging data events for
// MyLambdaFunction, any invocations of that function are logged. The trail
// processes and logs the event.
//
//     * The Invoke API operation on
// MyOtherLambdaFunction is an AWS Lambda API. Because the CloudTrail user did not
// specify logging data events for all Lambda functions, the Invoke operation for
// MyOtherLambdaFunction does not match the function specified for the trail. The
// trail doesn’t log the event.
type DataResource struct {

	// The resource type in which you want to log data events. You can specify
	// AWS::S3::Object or AWS::Lambda::Function resources.
	Type *string

	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the
	// specified objects.
	//
	//     * To log data events for all objects in all S3 buckets
	// in your AWS account, specify the prefix as arn:aws:s3:::. This will also enable
	// logging of data event activity performed by any user or role in your AWS
	// account, even if that activity is performed on a bucket that belongs to another
	// AWS account.
	//
	//     * To log data events for all objects in an S3 bucket, specify
	// the bucket and an empty object prefix such as arn:aws:s3:::bucket-1/. The trail
	// logs data events for all objects in this S3 bucket.
	//
	//     * To log data events
	// for specific objects, specify the S3 bucket and object prefix such as
	// arn:aws:s3:::bucket-1/example-images. The trail logs data events for objects in
	// this S3 bucket that match the prefix.
	//
	//     * To log data events for all
	// functions in your AWS account, specify the prefix as arn:aws:lambda. This will
	// also enable logging of Invoke activity performed by any user or role in your AWS
	// account, even if that activity is performed on a function that belongs to
	// another AWS account.
	//
	//     * To log data events for a specific Lambda function,
	// specify the function ARN. Lambda function ARNs are exact. For example, if you
	// specify a function ARN
	// arn:aws:lambda:us-west-2:111111111111:function:helloworld, data events will only
	// be logged for arn:aws:lambda:us-west-2:111111111111:function:helloworld. They
	// will not be logged for
	// arn:aws:lambda:us-west-2:111111111111:function:helloworld2.
	Values []*string
}

// Contains information about an event that was returned by a lookup request. The
// result includes a representation of a CloudTrail event.
type Event struct {

	// The AWS access key ID that was used to sign the request. If the request was made
	// with temporary security credentials, this is the access key ID of the temporary
	// credentials.
	AccessKeyId *string

	// A JSON string that contains a representation of the event returned.
	CloudTrailEvent *string

	// The CloudTrail ID of the event returned.
	EventId *string

	// The name of the event returned.
	EventName *string

	// The AWS service that the request was made to.
	EventSource *string

	// The date and time of the event returned.
	EventTime *time.Time

	// Information about whether the event is a write event or a read event.
	ReadOnly *string

	// A list of resources referenced by the event returned.
	Resources []*Resource

	// A user name or role name of the requester that called the API in the event
	// returned.
	Username *string
}

// Use event selectors to further specify the management and data event settings
// for your trail. By default, trails created without specific event selectors will
// be configured to log all read and write management events, and no data events.
// When an event occurs in your account, CloudTrail evaluates the event selector
// for all trails. For each trail, if the event matches any event selector, the
// trail processes and logs the event. If the event doesn't match any event
// selector, the trail doesn't log the event. You can configure up to five event
// selectors for a trail.
type EventSelector struct {

	// CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda
	// functions. You can specify up to 250 resources for an individual event selector,
	// but the total number of data resources cannot exceed 250 across all event
	// selectors in a trail. This limit does not apply if you configure resource
	// logging for all data events. For more information, see Data Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-data-events)
	// and Limits in AWS CloudTrail
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html)
	// in the AWS CloudTrail User Guide.
	DataResources []*DataResource

	// An optional list of service event sources from which you do not want management
	// events to be logged on your trail. In this release, the list can be empty
	// (disables the filter), or it can filter out AWS Key Management Service events by
	// containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is
	// empty, and AWS KMS events are included in events that are logged to your trail.
	ExcludeManagementEventSources []*string

	// Specify if you want your event selector to include management events for your
	// trail. For more information, see Management Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-management-events)
	// in the AWS CloudTrail User Guide.  <p>By default, the value is
	// <code>true</code>.</p>
	IncludeManagementEvents *bool

	// Specify if you want your trail to log read-only events, write-only events, or
	// all. For example, the EC2 GetConsoleOutput is a read-only API operation and
	// RunInstances is a write-only API operation. By default, the value is All.
	ReadWriteType ReadWriteType
}

// A JSON string that contains a list of insight types that are logged on a trail.
type InsightSelector struct {

	// The type of insights to log on a trail. In this release, only ApiCallRateInsight
	// is supported as an insight type.
	InsightType InsightType
}

// Specifies an attribute and value that filter the events returned.
type LookupAttribute struct {

	// Specifies an attribute on which to filter the events returned.
	//
	// This member is required.
	AttributeKey LookupAttributeKey

	// Specifies a value for the specified AttributeKey.
	//
	// This member is required.
	AttributeValue *string
}

// Contains information about a returned public key.
type PublicKey struct {

	// The fingerprint of the public key.
	Fingerprint *string

	// The ending time of validity of the public key.
	ValidityEndTime *time.Time

	// The starting time of validity of the public key.
	ValidityStartTime *time.Time

	// The DER encoded public key value in PKCS#1 format.
	Value []byte
}

// Specifies the type and name of a resource referenced by an event.
type Resource struct {

	// The name of the resource referenced by the event returned. These are
	// user-created names whose values will depend on the environment. For example, the
	// resource name might be "auto-scaling-test-group" for an Auto Scaling Group or
	// "i-1234567" for an EC2 Instance.
	ResourceName *string

	// The type of a resource referenced by the event returned. When the resource type
	// cannot be determined, null is returned. Some examples of resource types are:
	// Instance for EC2, Trail for CloudTrail, DBInstance for RDS, and AccessKey for
	// IAM. To learn more about how to look up and filter events by the resource types
	// supported for a service, see Filtering CloudTrail Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events-console.html#filtering-cloudtrail-events).
	ResourceType *string
}

// A resource tag.
type ResourceTag struct {

	// Specifies the ARN of the resource.
	ResourceId *string

	// A list of tags.
	TagsList []*Tag
}

// A custom key-value pair associated with a resource such as a CloudTrail trail.
type Tag struct {

	// The key in a key-value pair. The key must be must be no longer than 128 Unicode
	// characters. The key must be unique for the resource to which it applies.
	//
	// This member is required.
	Key *string

	// The value in a key-value pair of a tag. The value must be no longer than 256
	// Unicode characters.
	Value *string
}

// The settings for a trail.
type Trail struct {

	// Specifies an Amazon Resource Name (ARN), a unique identifier that represents the
	// log group to which CloudTrail logs will be delivered.
	CloudWatchLogsLogGroupArn *string

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a
	// user's log group.
	CloudWatchLogsRoleArn *string

	// Specifies if the trail has custom event selectors.
	HasCustomEventSelectors *bool

	// Specifies whether a trail has insight types specified in an InsightSelector
	// list.
	HasInsightSelectors *bool

	// The region in which the trail was created.
	HomeRegion *string

	// Set to True to include AWS API calls from AWS global services such as IAM.
	// Otherwise, False.
	IncludeGlobalServiceEvents *bool

	// Specifies whether the trail exists only in one region or exists in all regions.
	IsMultiRegionTrail *bool

	// Specifies whether the trail is an organization trail.
	IsOrganizationTrail *bool

	// Specifies the KMS key ID that encrypts the logs delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the format:
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies whether log file validation is enabled.
	LogFileValidationEnabled *bool

	// Name of the trail set by calling CreateTrail (). The maximum length is 128
	// characters.
	Name *string

	// Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.
	// See Amazon S3 Bucket Naming Requirements
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).
	S3BucketName *string

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you
	// have designated for log file delivery. For more information, see Finding Your
	// CloudTrail Log Files
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).The
	// maximum length is 200 characters.
	S3KeyPrefix *string

	// Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send
	// notifications when log files are delivered. The format of a topic ARN is:
	// arn:aws:sns:us-east-2:123456789012:MyTopic
	SnsTopicARN *string

	// This field is no longer in use. Use SnsTopicARN.
	SnsTopicName *string

	// Specifies the ARN of the trail. The format of a trail ARN is:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string
}

// Information about a CloudTrail trail, including the trail's name, home region,
// and Amazon Resource Name (ARN).
type TrailInfo struct {

	// The AWS region in which a trail was created.
	HomeRegion *string

	// The name of a trail.
	Name *string

	// The ARN of a trail.
	TrailARN *string
}
