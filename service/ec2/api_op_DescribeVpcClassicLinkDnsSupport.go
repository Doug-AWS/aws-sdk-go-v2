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

// Describes the ClassicLink DNS support status of one or more VPCs. If enabled,
// the DNS hostname of a linked EC2-Classic instance resolves to its private IP
// address when addressed from an instance in the VPC to which it's linked.
// Similarly, the DNS hostname of an instance in a VPC resolves to its private IP
// address when addressed from a linked EC2-Classic instance. For more information,
// see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVpcClassicLinkDnsSupport(ctx context.Context, params *DescribeVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*DescribeVpcClassicLinkDnsSupportOutput, error) {
	stack := middleware.NewStack("DescribeVpcClassicLinkDnsSupport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcClassicLinkDnsSupportMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcClassicLinkDnsSupport",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcClassicLinkDnsSupportInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more VPC IDs.
	VpcIds []*string
}

type DescribeVpcClassicLinkDnsSupportOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the ClassicLink DNS support status of the VPCs.
	Vpcs []*types.ClassicLinkDnsSupport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcClassicLinkDnsSupportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcClassicLinkDnsSupport",
	}
}
