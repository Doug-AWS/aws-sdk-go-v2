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

// Updates the specified attributes of the Direct Connect gateway association. Add
// or remove prefixes from the association.
func (c *Client) UpdateDirectConnectGatewayAssociation(ctx context.Context, params *UpdateDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*UpdateDirectConnectGatewayAssociationOutput, error) {
	stack := middleware.NewStack("UpdateDirectConnectGatewayAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDirectConnectGatewayAssociationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(options.Region), middleware.Before)

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
			OperationName: "UpdateDirectConnectGatewayAssociation",
			Err:           err,
		}
	}
	out := result.(*UpdateDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDirectConnectGatewayAssociationInput struct {
	// The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.
	RemoveAllowedPrefixesToDirectConnectGateway []*types.RouteFilterPrefix
	// The ID of the Direct Connect gateway association.
	AssociationId *string
	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AddAllowedPrefixesToDirectConnectGateway []*types.RouteFilterPrefix
}

type UpdateDirectConnectGatewayAssociationOutput struct {
	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "UpdateDirectConnectGatewayAssociation",
	}
}