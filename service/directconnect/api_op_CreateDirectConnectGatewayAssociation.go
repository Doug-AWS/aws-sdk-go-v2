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

// Creates an association between a Direct Connect gateway and a virtual private
// gateway. The virtual private gateway must be attached to a VPC and must not be
// associated with another Direct Connect gateway.
func (c *Client) CreateDirectConnectGatewayAssociation(ctx context.Context, params *CreateDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*CreateDirectConnectGatewayAssociationOutput, error) {
	stack := middleware.NewStack("CreateDirectConnectGatewayAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDirectConnectGatewayAssociationMiddlewares(stack)
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
	addOpCreateDirectConnectGatewayAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociation(options.Region), middleware.Before)
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
			OperationName: "CreateDirectConnectGatewayAssociation",
			Err:           err,
		}
	}
	out := result.(*CreateDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectConnectGatewayAssociationInput struct {

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway This
	// parameter is required when you create an association to a transit gateway. For
	// information about how to set the prefixes, see Allowed Prefixes
	// (https://docs.aws.amazon.com/directconnect/latest/UserGuide/multi-account-associate-vgw.html#allowed-prefixes)
	// in the AWS Direct Connect User Guide.
	AddAllowedPrefixesToDirectConnectGateway []*types.RouteFilterPrefix

	// The ID of the virtual private gateway or transit gateway.
	GatewayId *string

	// The ID of the virtual private gateway.
	VirtualGatewayId *string
}

type CreateDirectConnectGatewayAssociationOutput struct {

	// The association to be created.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectConnectGatewayAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectConnectGatewayAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateDirectConnectGatewayAssociation",
	}
}
