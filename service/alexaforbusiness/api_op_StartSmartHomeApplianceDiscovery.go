// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates the discovery of any smart home appliances associated with the room.
func (c *Client) StartSmartHomeApplianceDiscovery(ctx context.Context, params *StartSmartHomeApplianceDiscoveryInput, optFns ...func(*Options)) (*StartSmartHomeApplianceDiscoveryOutput, error) {
	stack := middleware.NewStack("StartSmartHomeApplianceDiscovery", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartSmartHomeApplianceDiscoveryMiddlewares(stack)
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
	addOpStartSmartHomeApplianceDiscoveryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSmartHomeApplianceDiscovery(options.Region), middleware.Before)
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
			OperationName: "StartSmartHomeApplianceDiscovery",
			Err:           err,
		}
	}
	out := result.(*StartSmartHomeApplianceDiscoveryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSmartHomeApplianceDiscoveryInput struct {
	// The room where smart home appliance discovery was initiated.
	RoomArn *string
}

type StartSmartHomeApplianceDiscoveryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartSmartHomeApplianceDiscoveryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartSmartHomeApplianceDiscovery{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSmartHomeApplianceDiscovery{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSmartHomeApplianceDiscovery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "StartSmartHomeApplianceDiscovery",
	}
}
