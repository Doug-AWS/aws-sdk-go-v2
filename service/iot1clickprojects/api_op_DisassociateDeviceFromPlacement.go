// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a physical device from a placement.
func (c *Client) DisassociateDeviceFromPlacement(ctx context.Context, params *DisassociateDeviceFromPlacementInput, optFns ...func(*Options)) (*DisassociateDeviceFromPlacementOutput, error) {
	stack := middleware.NewStack("DisassociateDeviceFromPlacement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateDeviceFromPlacementMiddlewares(stack)
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
	addOpDisassociateDeviceFromPlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDeviceFromPlacement(options.Region), middleware.Before)
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
			OperationName: "DisassociateDeviceFromPlacement",
			Err:           err,
		}
	}
	out := result.(*DisassociateDeviceFromPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDeviceFromPlacementInput struct {

	// The device ID that should be removed from the placement.
	//
	// This member is required.
	DeviceTemplateName *string

	// The name of the placement that the device should be removed from.
	//
	// This member is required.
	PlacementName *string

	// The name of the project that contains the placement.
	//
	// This member is required.
	ProjectName *string
}

type DisassociateDeviceFromPlacementOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateDeviceFromPlacementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateDeviceFromPlacement{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateDeviceFromPlacement{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateDeviceFromPlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "DisassociateDeviceFromPlacement",
	}
}
