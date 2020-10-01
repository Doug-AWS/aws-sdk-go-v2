// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Moves the specified instances into the standby state. If you choose to decrement
// the desired capacity of the Auto Scaling group, the instances can enter standby
// as long as the desired capacity of the Auto Scaling group after the instances
// are placed into standby is equal to or greater than the minimum capacity of the
// group. If you choose not to decrement the desired capacity of the Auto Scaling
// group, the Auto Scaling group launches new instances to replace the instances on
// standby. For more information, see Temporarily Removing Instances from Your Auto
// Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) EnterStandby(ctx context.Context, params *EnterStandbyInput, optFns ...func(*Options)) (*EnterStandbyOutput, error) {
	stack := middleware.NewStack("EnterStandby", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpEnterStandbyMiddlewares(stack)
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
	addOpEnterStandbyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnterStandby(options.Region), middleware.Before)
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
			OperationName: "EnterStandby",
			Err:           err,
		}
	}
	out := result.(*EnterStandbyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnterStandbyInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Indicates whether to decrement the desired capacity of the Auto Scaling group by
	// the number of instances moved to Standby mode.
	//
	// This member is required.
	ShouldDecrementDesiredCapacity *bool

	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []*string
}

type EnterStandbyOutput struct {

	// The activities related to moving instances into Standby mode.
	Activities []*types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpEnterStandbyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpEnterStandby{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpEnterStandby{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnterStandby(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "EnterStandby",
	}
}
