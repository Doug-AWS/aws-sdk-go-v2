// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates capacity settings for a fleet. Use this action to specify the number of
// EC2 instances (hosts) that you want this fleet to contain. Before calling this
// action, you may want to call DescribeEC2InstanceLimits () to get the maximum
// capacity based on the fleet's EC2 instance type. Specify minimum and maximum
// number of instances. Amazon GameLift will not change fleet capacity to values
// fall outside of this range. This is particularly important when using
// auto-scaling (see PutScalingPolicy ()) to allow capacity to adjust based on
// player demand while imposing limits on automatic adjustments. To update fleet
// capacity, specify the fleet ID and the number of instances you want the fleet to
// host. If successful, Amazon GameLift starts or terminates instances so that the
// fleet's active instance count matches the desired instance count. You can view a
// fleet's current capacity information by calling DescribeFleetCapacity (). If the
// desired instance count is higher than the instance type's limit, the "Limit
// Exceeded" exception occurs. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * DescribeFleetAttributes ()
//
//     * Update fleets:
//
//         *
// UpdateFleetAttributes ()
//
//         * UpdateFleetCapacity ()
//
//         *
// UpdateFleetPortSettings ()
//
//         * UpdateRuntimeConfiguration ()
//
//     *
// StartFleetActions () or StopFleetActions ()
func (c *Client) UpdateFleetCapacity(ctx context.Context, params *UpdateFleetCapacityInput, optFns ...func(*Options)) (*UpdateFleetCapacityOutput, error) {
	stack := middleware.NewStack("UpdateFleetCapacity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateFleetCapacityMiddlewares(stack)
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
	addOpUpdateFleetCapacityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetCapacity(options.Region), middleware.Before)
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
			OperationName: "UpdateFleetCapacity",
			Err:           err,
		}
	}
	out := result.(*UpdateFleetCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateFleetCapacityInput struct {

	// A unique identifier for a fleet to update capacity for. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Number of EC2 instances you want this fleet to host.
	DesiredInstances *int32

	// The maximum value allowed for the fleet's instance count. Default if not set is
	// 1.
	MaxSize *int32

	// The minimum value allowed for the fleet's instance count. Default if not set is
	// 0.
	MinSize *int32
}

// Represents the returned data in response to a request action.
type UpdateFleetCapacityOutput struct {

	// A unique identifier for a fleet that was updated.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateFleetCapacityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetCapacity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetCapacity{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFleetCapacity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateFleetCapacity",
	}
}
