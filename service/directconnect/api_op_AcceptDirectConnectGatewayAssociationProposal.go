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

// Accepts a proposal request to attach a virtual private gateway or transit
// gateway to a Direct Connect gateway.
func (c *Client) AcceptDirectConnectGatewayAssociationProposal(ctx context.Context, params *AcceptDirectConnectGatewayAssociationProposalInput, optFns ...func(*Options)) (*AcceptDirectConnectGatewayAssociationProposalOutput, error) {
	stack := middleware.NewStack("AcceptDirectConnectGatewayAssociationProposal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAcceptDirectConnectGatewayAssociationProposalMiddlewares(stack)
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
	addOpAcceptDirectConnectGatewayAssociationProposalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptDirectConnectGatewayAssociationProposal(options.Region), middleware.Before)
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
			OperationName: "AcceptDirectConnectGatewayAssociationProposal",
			Err:           err,
		}
	}
	out := result.(*AcceptDirectConnectGatewayAssociationProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptDirectConnectGatewayAssociationProposalInput struct {

	// The ID of the AWS account that owns the virtual private gateway or transit
	// gateway.
	//
	// This member is required.
	AssociatedGatewayOwnerAccount *string

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The ID of the request proposal.
	//
	// This member is required.
	ProposalId *string

	// Overrides the Amazon VPC prefixes advertised to the Direct Connect gateway. For
	// information about how to set the prefixes, see Allowed Prefixes
	// (https://docs.aws.amazon.com/directconnect/latest/UserGuide/multi-account-associate-vgw.html#allowed-prefixes)
	// in the AWS Direct Connect User Guide.
	OverrideAllowedPrefixesToDirectConnectGateway []*types.RouteFilterPrefix
}

type AcceptDirectConnectGatewayAssociationProposalOutput struct {

	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAcceptDirectConnectGatewayAssociationProposalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptDirectConnectGatewayAssociationProposal{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptDirectConnectGatewayAssociationProposal{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptDirectConnectGatewayAssociationProposal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AcceptDirectConnectGatewayAssociationProposal",
	}
}
