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

// Describes your Elastic IP addresses that are being moved to the EC2-VPC
// platform, or that are being restored to the EC2-Classic platform. This request
// does not return information about any other Elastic IP addresses in your
// account.
func (c *Client) DescribeMovingAddresses(ctx context.Context, params *DescribeMovingAddressesInput, optFns ...func(*Options)) (*DescribeMovingAddressesOutput, error) {
	stack := middleware.NewStack("DescribeMovingAddresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeMovingAddressesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMovingAddresses(options.Region), middleware.Before)
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
			OperationName: "DescribeMovingAddresses",
			Err:           err,
		}
	}
	out := result.(*DescribeMovingAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMovingAddressesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * moving-status - The status of the Elastic IP address
	// (MovingToVpc | RestoringToClassic).
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1000; if
	// MaxResults is given a value outside of this range, an error is returned.
	// Default: If no value is provided, the default is 1000.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more Elastic IP addresses.
	PublicIps []*string
}

type DescribeMovingAddressesOutput struct {

	// The status for each Elastic IP address.
	MovingAddressStatuses []*types.MovingAddressStatus

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeMovingAddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeMovingAddresses{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeMovingAddresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMovingAddresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeMovingAddresses",
	}
}
