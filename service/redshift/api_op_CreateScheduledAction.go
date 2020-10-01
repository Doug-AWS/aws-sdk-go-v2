// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a scheduled action. A scheduled action contains a schedule and an Amazon
// Redshift API action. For example, you can create a schedule of when to run the
// ResizeCluster API operation.
func (c *Client) CreateScheduledAction(ctx context.Context, params *CreateScheduledActionInput, optFns ...func(*Options)) (*CreateScheduledActionOutput, error) {
	stack := middleware.NewStack("CreateScheduledAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateScheduledActionMiddlewares(stack)
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
	addOpCreateScheduledActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScheduledAction(options.Region), middleware.Before)
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
			OperationName: "CreateScheduledAction",
			Err:           err,
		}
	}
	out := result.(*CreateScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduledActionInput struct {

	// The IAM role to assume to run the target action. For more information about this
	// parameter, see ScheduledAction ().
	//
	// This member is required.
	IamRole *string

	// The schedule in at( ) or cron( ) format. For more information about this
	// parameter, see ScheduledAction ().
	//
	// This member is required.
	Schedule *string

	// The name of the scheduled action. The name must be unique within an account. For
	// more information about this parameter, see ScheduledAction ().
	//
	// This member is required.
	ScheduledActionName *string

	// A JSON format string of the Amazon Redshift API operation with input parameters.
	// For more information about this parameter, see ScheduledAction ().
	//
	// This member is required.
	TargetAction *types.ScheduledActionType

	// If true, the schedule is enabled. If false, the scheduled action does not
	// trigger. For more information about state of the scheduled action, see
	// ScheduledAction ().
	Enable *bool

	// The end time in UTC of the scheduled action. After this time, the scheduled
	// action does not trigger. For more information about this parameter, see
	// ScheduledAction ().
	EndTime *time.Time

	// The description of the scheduled action.
	ScheduledActionDescription *string

	// The start time in UTC of the scheduled action. Before this time, the scheduled
	// action does not trigger. For more information about this parameter, see
	// ScheduledAction ().
	StartTime *time.Time
}

// Describes a scheduled action. You can use a scheduled action to trigger some
// Amazon Redshift API operations on a schedule. For information about which API
// operations can be scheduled, see ScheduledActionType ().
type CreateScheduledActionOutput struct {

	// The end time in UTC when the schedule is no longer active. After this time, the
	// scheduled action does not trigger.
	EndTime *time.Time

	// The IAM role to assume to run the scheduled action. This IAM role must have
	// permission to run the Amazon Redshift API operation in the scheduled action.
	// This IAM role must allow the Amazon Redshift scheduler (Principal
	// scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more
	// information about the IAM role to use with the Amazon Redshift scheduler, see
	// Using Identity-Based Policies for Amazon Redshift
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html)
	// in the Amazon Redshift Cluster Management Guide.
	IamRole *string

	// List of times when the scheduled action will run.
	NextInvocations []*time.Time

	// The schedule for a one-time (at format) or recurring (cron format) scheduled
	// action. Schedule invocations must be separated by at least one hour. Format of
	// at expressions is "at(yyyy-mm-ddThh:mm:ss)". For example,
	// "at(2016-03-04T17:27:00)". Format of cron expressions is "cron(Minutes Hours
	// Day-of-month Month Day-of-week Year)". For example, "cron(0 10 ? * MON *)". For
	// more information, see Cron Expressions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide.
	Schedule *string

	// The description of the scheduled action.
	ScheduledActionDescription *string

	// The name of the scheduled action.
	ScheduledActionName *string

	// The start time in UTC when the schedule is active. Before this time, the
	// scheduled action does not trigger.
	StartTime *time.Time

	// The state of the scheduled action. For example, DISABLED.
	State types.ScheduledActionState

	// A JSON format string of the Amazon Redshift API operation with input parameters.
	// "{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}".
	TargetAction *types.ScheduledActionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateScheduledActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateScheduledAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateScheduledAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateScheduledAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateScheduledAction",
	}
}
