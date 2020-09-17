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

// Deletes the specified transit gateway multicast domain.
func (c *Client) DeleteTransitGatewayMulticastDomain(ctx context.Context, params *DeleteTransitGatewayMulticastDomainInput, optFns ...func(*Options)) (*DeleteTransitGatewayMulticastDomainOutput, error) {
	stack := middleware.NewStack("DeleteTransitGatewayMulticastDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteTransitGatewayMulticastDomainMiddlewares(stack)
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
	addOpDeleteTransitGatewayMulticastDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTransitGatewayMulticastDomain(options.Region), middleware.Before)

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
			OperationName: "DeleteTransitGatewayMulticastDomain",
			Err:           err,
		}
	}
	out := result.(*DeleteTransitGatewayMulticastDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTransitGatewayMulticastDomainInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string
}

type DeleteTransitGatewayMulticastDomainOutput struct {
	// Information about the deleted transit gateway multicast domain.
	TransitGatewayMulticastDomain *types.TransitGatewayMulticastDomain

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteTransitGatewayMulticastDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteTransitGatewayMulticastDomain{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteTransitGatewayMulticastDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTransitGatewayMulticastDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTransitGatewayMulticastDomain",
	}
}