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

// Describes the VPC endpoint service configurations in your account (your
// services).
func (c *Client) DescribeVpcEndpointServiceConfigurations(ctx context.Context, params *DescribeVpcEndpointServiceConfigurationsInput, optFns ...func(*Options)) (*DescribeVpcEndpointServiceConfigurationsOutput, error) {
	stack := middleware.NewStack("DescribeVpcEndpointServiceConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcEndpointServiceConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointServiceConfigurations(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcEndpointServiceConfigurations",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcEndpointServiceConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointServiceConfigurationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * service-name - The name of the service.
	//
	//     *
	// service-id - The ID of the service.
	//
	//     * service-state - The state of the
	// service (Pending | Available | Deleting | Deleted | Failed).
	//
	//     * tag: - The
	// key/value combination of a tag assigned to the resource. Use the tag key in the
	// filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	//     * tag-key -
	// The key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	// The IDs of one or more services.
	ServiceIds []*string
}

type DescribeVpcEndpointServiceConfigurationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more services.
	ServiceConfigurations []*types.ServiceConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcEndpointServiceConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointServiceConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointServiceConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcEndpointServiceConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointServiceConfigurations",
	}
}
