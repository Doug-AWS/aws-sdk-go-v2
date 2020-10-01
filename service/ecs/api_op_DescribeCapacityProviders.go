// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your capacity providers.
func (c *Client) DescribeCapacityProviders(ctx context.Context, params *DescribeCapacityProvidersInput, optFns ...func(*Options)) (*DescribeCapacityProvidersOutput, error) {
	stack := middleware.NewStack("DescribeCapacityProviders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCapacityProvidersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCapacityProviders(options.Region), middleware.Before)
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
			OperationName: "DescribeCapacityProviders",
			Err:           err,
		}
	}
	out := result.(*DescribeCapacityProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCapacityProvidersInput struct {

	// The short name or full Amazon Resource Name (ARN) of one or more capacity
	// providers. Up to 100 capacity providers can be described in an action.
	CapacityProviders []*string

	// Specifies whether or not you want to see the resource tags for the capacity
	// provider. If TAGS is specified, the tags are included in the response. If this
	// field is omitted, tags are not included in the response.
	Include []types.CapacityProviderField

	// The maximum number of account setting results returned by
	// DescribeCapacityProviders in paginated output. When this parameter is used,
	// DescribeCapacityProviders only returns maxResults results in a single page along
	// with a nextToken response element. The remaining results of the initial request
	// can be seen by sending another DescribeCapacityProviders request with the
	// returned nextToken value. This value can be between 1 and 10. If this parameter
	// is not used, then DescribeCapacityProviders returns up to 10 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeCapacityProviders
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string
}

type DescribeCapacityProvidersOutput struct {

	// The list of capacity providers.
	CapacityProviders []*types.CapacityProvider

	// Any failures associated with the call.
	Failures []*types.Failure

	// The nextToken value to include in a future DescribeCapacityProviders request.
	// When the results of a DescribeCapacityProviders request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCapacityProvidersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCapacityProviders{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCapacityProviders{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCapacityProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeCapacityProviders",
	}
}
