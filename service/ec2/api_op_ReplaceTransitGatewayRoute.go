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

// Replaces the specified route in the specified transit gateway route table.
func (c *Client) ReplaceTransitGatewayRoute(ctx context.Context, params *ReplaceTransitGatewayRouteInput, optFns ...func(*Options)) (*ReplaceTransitGatewayRouteOutput, error) {
	stack := middleware.NewStack("ReplaceTransitGatewayRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReplaceTransitGatewayRouteMiddlewares(stack)
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
	addOpReplaceTransitGatewayRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceTransitGatewayRoute(options.Region), middleware.Before)
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
			OperationName: "ReplaceTransitGatewayRoute",
			Err:           err,
		}
	}
	out := result.(*ReplaceTransitGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceTransitGatewayRouteInput struct {

	// The CIDR range used for the destination match. Routing decisions are based on
	// the most specific match.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Indicates whether traffic matching this route is to be dropped.
	Blackhole *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the attachment.
	TransitGatewayAttachmentId *string
}

type ReplaceTransitGatewayRouteOutput struct {

	// Information about the modified route.
	Route *types.TransitGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReplaceTransitGatewayRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReplaceTransitGatewayRoute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceTransitGatewayRoute{}, middleware.After)
}

func newServiceMetadataMiddleware_opReplaceTransitGatewayRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceTransitGatewayRoute",
	}
}
