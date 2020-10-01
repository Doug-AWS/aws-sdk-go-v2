// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a transit virtual interface. A transit virtual interface should be used
// to access one or more transit gateways associated with Direct Connect gateways.
// A transit virtual interface enables the connection of multiple VPCs attached to
// a transit gateway to a Direct Connect gateway. If you associate your transit
// gateway with one or more Direct Connect gateways, the Autonomous System Number
// (ASN) used by the transit gateway and the Direct Connect gateway must be
// different. For example, if you use the default ASN 64512 for both your the
// transit gateway and Direct Connect gateway, the association request fails.
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections (). To
// check whether your virtual interface supports jumbo frames, call
// DescribeVirtualInterfaces ().
func (c *Client) CreateTransitVirtualInterface(ctx context.Context, params *CreateTransitVirtualInterfaceInput, optFns ...func(*Options)) (*CreateTransitVirtualInterfaceOutput, error) {
	stack := middleware.NewStack("CreateTransitVirtualInterface", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTransitVirtualInterfaceMiddlewares(stack)
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
	addOpCreateTransitVirtualInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitVirtualInterface(options.Region), middleware.Before)
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
			OperationName: "CreateTransitVirtualInterface",
			Err:           err,
		}
	}
	out := result.(*CreateTransitVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitVirtualInterfaceInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// Information about the transit virtual interface.
	//
	// This member is required.
	NewTransitVirtualInterface *types.NewTransitVirtualInterface
}

type CreateTransitVirtualInterfaceOutput struct {

	// Information about a virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTransitVirtualInterfaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTransitVirtualInterface{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTransitVirtualInterface{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTransitVirtualInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateTransitVirtualInterface",
	}
}
