// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the name, description, and rules in a device pool given the attributes
// and the pool ARN. Rule updates are all-or-nothing, meaning they can only be
// updated as a whole (or not at all).
func (c *Client) UpdateDevicePool(ctx context.Context, params *UpdateDevicePoolInput, optFns ...func(*Options)) (*UpdateDevicePoolOutput, error) {
	stack := middleware.NewStack("UpdateDevicePool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDevicePoolMiddlewares(stack)
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
	addOpUpdateDevicePoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevicePool(options.Region), middleware.Before)

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
			OperationName: "UpdateDevicePool",
			Err:           err,
		}
	}
	out := result.(*UpdateDevicePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the update device pool operation.
type UpdateDevicePoolInput struct {
	// The number of devices that Device Farm can add to your device pool. Device Farm
	// adds devices that are available and that meet the criteria that you assign for
	// the rules parameter. Depending on how many devices meet these constraints, your
	// device pool might contain fewer devices than the value for this parameter. By
	// specifying the maximum number of devices, you can control the costs that you
	// incur by running tests. If you use this parameter in your request, you cannot
	// use the clearMaxDevices parameter in the same request.
	MaxDevices *int32
	// The Amazon Resource Name (ARN) of the Device Farm device pool to update.
	Arn *string
	// Represents the rules to modify for the device pool. Updating rules is optional.
	// If you update rules for your request, the update replaces the existing rules.
	Rules []*types.Rule
	// Sets whether the maxDevices parameter applies to your device pool. If you set
	// this parameter to true, the maxDevices parameter does not apply, and Device Farm
	// does not limit the number of devices that it adds to your device pool. In this
	// case, Device Farm adds all available devices that meet the criteria specified in
	// the rules parameter. If you use this parameter in your request, you cannot use
	// the maxDevices parameter in the same request.
	ClearMaxDevices *bool
	// A string that represents the name of the device pool to update.
	Name *string
	// A description of the device pool to update.
	Description *string
}

// Represents the result of an update device pool request.
type UpdateDevicePoolOutput struct {
	// The device pool you just updated.
	DevicePool *types.DevicePool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDevicePoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDevicePool{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDevicePool{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDevicePool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateDevicePool",
	}
}