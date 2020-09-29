// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the inputs for the change set and a list of changes that AWS
// CloudFormation will make if you execute the change set. For more information,
// see Updating Stacks Using Change Sets
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html)
// in the AWS CloudFormation User Guide.
func (c *Client) DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) {
	stack := middleware.NewStack("DescribeChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeChangeSetMiddlewares(stack)
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
	addOpDescribeChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSet(options.Region), middleware.Before)
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
			OperationName: "DescribeChangeSet",
			Err:           err,
		}
	}
	out := result.(*DescribeChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeChangeSet () action.
type DescribeChangeSetInput struct {
	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// describe.
	ChangeSetName *string
	// If you specified the name of a change set, specify the stack name or ID (ARN) of
	// the change set you want to describe.
	StackName *string
	// A string (provided by the DescribeChangeSet () response output) that identifies
	// the next page of information that you want to retrieve.
	NextToken *string
}

// The output for the DescribeChangeSet () action.
type DescribeChangeSetOutput struct {
	// A list of Parameter structures that describes the input parameters and their
	// values used to create the change set. For more information, see the Parameter
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []*types.Parameter
	// A list of Change structures that describes the resources AWS CloudFormation
	// changes if you execute the change set.
	Changes []*types.Change
	// The current status of the change set, such as CREATE_IN_PROGRESS,
	// CREATE_COMPLETE, or FAILED.
	Status types.ChangeSetStatus
	// The name of the stack that is associated with the change set.
	StackName *string
	// The ARN of the stack that is associated with the change set.
	StackId *string
	// The start time when the change set was created, in UTC.
	CreationTime *time.Time
	// Information about the change set.
	Description *string
	// If you execute the change set, the list of capabilities that were explicitly
	// acknowledged when the change set was created.
	Capabilities []types.Capability
	// If you execute the change set, the tags that will be associated with the stack.
	Tags []*types.Tag
	// The rollback triggers for AWS CloudFormation to monitor during stack creation
	// and updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration
	// The name of the change set.
	ChangeSetName *string
	// A description of the change set's status. For example, if your attempt to create
	// a change set failed, AWS CloudFormation shows the error message.
	StatusReason *string
	// If the output exceeds 1 MB, a string that identifies the next page of changes.
	// If there is no additional page, this value is null.
	NextToken *string
	// The ARNs of the Amazon Simple Notification Service (Amazon SNS) topics that will
	// be associated with the stack if you execute the change set.
	NotificationARNs []*string
	// The ARN of the change set.
	ChangeSetId *string
	// If the change set execution status is AVAILABLE, you can execute the change set.
	// If you can’t execute the change set, the status indicates why. For example, a
	// change set might be in an UNAVAILABLE state because AWS CloudFormation is still
	// creating it or in an OBSOLETE state because the stack was already updated.
	ExecutionStatus types.ExecutionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeChangeSet",
	}
}
