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

// Describes the associations between virtual interface groups and local gateway
// route tables.
func (c *Client) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx context.Context, params *DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, optFns ...func(*Options)) (*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	stack := middleware.NewStack("DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(options.Region), middleware.Before)
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
			OperationName: "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
			Err:           err,
		}
	}
	out := result.(*DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput struct {
	// The IDs of the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationIds []*string
	// The token for the next page of results.
	NextToken *string
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// One or more filters.
	//
	//     * local-gateway-id - The ID of a local gateway.
	//
	//     *
	// local-gateway-route-table-id - The ID of the local gateway route table.
	//
	//     *
	// local-gateway-route-table-virtual-interface-group-association-id - The ID of the
	// association.
	//
	//     * local-gateway-route-table-virtual-interface-group-id - The
	// ID of the virtual interface group.
	//
	//     * state - The state of the association.
	Filters []*types.Filter
}

type DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput struct {
	// Information about the associations.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociations []*types.LocalGatewayRouteTableVirtualInterfaceGroupAssociation
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
	}
}
