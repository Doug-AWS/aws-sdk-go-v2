// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the route table associated with a given subnet, internet gateway, or
// virtual private gateway in a VPC. After the operation completes, the subnet or
// gateway uses the routes in the new route table. For more information about route
// tables, see Route Tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
// Amazon Virtual Private Cloud User Guide. You can also use this operation to
// change which table is the main route table in the VPC. Specify the main route
// table's association ID and the route table ID of the new main route table.
func (c *Client) ReplaceRouteTableAssociation(ctx context.Context, params *ReplaceRouteTableAssociationInput, optFns ...func(*Options)) (*ReplaceRouteTableAssociationOutput, error) {
	stack := middleware.NewStack("ReplaceRouteTableAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReplaceRouteTableAssociationMiddlewares(stack)
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
	addOpReplaceRouteTableAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceRouteTableAssociation(options.Region), middleware.Before)
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
			OperationName: "ReplaceRouteTableAssociation",
			Err:           err,
		}
	}
	out := result.(*ReplaceRouteTableAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceRouteTableAssociationInput struct {

	// The association ID.
	//
	// This member is required.
	AssociationId *string

	// The ID of the new route table to associate with the subnet.
	//
	// This member is required.
	RouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ReplaceRouteTableAssociationOutput struct {

	// The state of the association.
	AssociationState *types.RouteTableAssociationState

	// The ID of the new association.
	NewAssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReplaceRouteTableAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReplaceRouteTableAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceRouteTableAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opReplaceRouteTableAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceRouteTableAssociation",
	}
}
