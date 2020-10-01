// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more flow logs to capture information about IP traffic for a
// specific network interface, subnet, or VPC.  <p>Flow log data for a monitored
// network interface is recorded as flow log records, which are log events
// consisting of fields that describe the traffic flow. For more information, see
// <a
// href="https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records">Flow
// Log Records</a> in the <i>Amazon Virtual Private Cloud User Guide</i>.</p>
// <p>When publishing to CloudWatch Logs, flow log records are published to a log
// group, and each network interface has a unique log stream in the log group. When
// publishing to Amazon S3, flow log records for all of the monitored network
// interfaces are published to a single log file object that is stored in the
// specified bucket.</p> <p>For more information, see <a
// href="https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html">VPC Flow
// Logs</a> in the <i>Amazon Virtual Private Cloud User Guide</i>.</p>
func (c *Client) CreateFlowLogs(ctx context.Context, params *CreateFlowLogsInput, optFns ...func(*Options)) (*CreateFlowLogsOutput, error) {
	stack := middleware.NewStack("CreateFlowLogs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateFlowLogsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateFlowLogsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowLogs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateFlowLogs",
			Err:           err,
		}
	}
	out := result.(*CreateFlowLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowLogsInput struct {

	// The ID of the subnet, network interface, or VPC for which you want to create a
	// flow log. Constraints: Maximum of 1000 resources
	//
	// This member is required.
	ResourceIds []*string

	// The type of resource for which to create the flow log. For example, if you
	// specified a VPC ID for the ResourceId property, specify VPC for this property.
	//
	// This member is required.
	ResourceType types.FlowLogsResourceType

	// The type of traffic to log. You can log traffic that the resource accepts or
	// rejects, or all traffic.
	//
	// This member is required.
	TrafficType types.TrafficType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a
	// CloudWatch Logs log group in your account. If you specify LogDestinationType as
	// s3, do not specify DeliverLogsPermissionArn or LogGroupName.
	DeliverLogsPermissionArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Specifies the destination to which the flow log data is to be published. Flow
	// log data can be published to a CloudWatch Logs log group or an Amazon S3 bucket.
	// The value specified for this parameter depends on the value specified for
	// LogDestinationType. If LogDestinationType is not specified or cloud-watch-logs,
	// specify the Amazon Resource Name (ARN) of the CloudWatch Logs log group. For
	// example, to publish to a log group called my-logs, specify
	// arn:aws:logs:us-east-1:123456789012:log-group:my-logs. Alternatively, use
	// LogGroupName instead. If LogDestinationType is s3, specify the ARN of the Amazon
	// S3 bucket. You can also specify a subfolder in the bucket. To specify a
	// subfolder in the bucket, use the following ARN format:
	// bucket_ARN/subfolder_name/. For example, to specify a subfolder named my-logs in
	// a bucket named my-bucket, use the following ARN:
	// arn:aws:s3:::my-bucket/my-logs/. You cannot use AWSLogs as a subfolder name.
	// This is a reserved term.
	LogDestination *string

	// Specifies the type of destination to which the flow log data is to be published.
	// Flow log data can be published to CloudWatch Logs or Amazon S3. To publish flow
	// log data to CloudWatch Logs, specify cloud-watch-logs. To publish flow log data
	// to Amazon S3, specify s3. If you specify LogDestinationType as s3, do not
	// specify DeliverLogsPermissionArn or LogGroupName. Default: cloud-watch-logs
	LogDestinationType types.LogDestinationType

	// The fields to include in the flow log record, in the order in which they should
	// appear. For a list of available fields, see Flow Log Records
	// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records).
	// If you omit this parameter, the flow log is created using the default format. If
	// you specify this parameter, you must specify at least one field. Specify the
	// fields using the ${field-id} format, separated by spaces. For the AWS CLI, use
	// single quotation marks (' ') to surround the parameter value.
	LogFormat *string

	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs. If you specify LogDestinationType as s3, do not
	// specify DeliverLogsPermissionArn or LogGroupName.
	LogGroupName *string

	// The maximum interval of time during which a flow of packets is captured and
	// aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600
	// seconds (10 minutes). When a network interface is attached to a Nitro-based
	// instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances),
	// the aggregation interval is always 60 seconds or less, regardless of the value
	// that you specify. Default: 600
	MaxAggregationInterval *int32

	// The tags to apply to the flow logs.
	TagSpecifications []*types.TagSpecification
}

type CreateFlowLogsOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// The IDs of the flow logs.
	FlowLogIds []*string

	// Information about the flow logs that could not be created successfully.
	Unsuccessful []*types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateFlowLogsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateFlowLogs{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFlowLogs{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFlowLogs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateFlowLogs",
	}
}
