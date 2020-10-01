// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays a list of flows that are associated with this account. This request
// returns a paginated result.
func (c *Client) ListFlows(ctx context.Context, params *ListFlowsInput, optFns ...func(*Options)) (*ListFlowsOutput, error) {
	stack := middleware.NewStack("ListFlows", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListFlowsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFlows(options.Region), middleware.Before)
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
			OperationName: "ListFlows",
			Err:           err,
		}
	}
	out := result.(*ListFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListFlows request with MaxResults set at 5. Although 20 items match your
	// request, the service returns no more than the first 5 items. (The service also
	// returns a NextToken value that you can use to fetch the next batch of results.)
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 10 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListFlows request with MaxResults set at 5. The service
	// returns the first batch of results (up to 5) and a NextToken value. To see the
	// next batch of results, you can submit the ListFlows request a second time and
	// specify the NextToken value.
	NextToken *string
}

type ListFlowsOutput struct {

	// A list of flow summaries.
	Flows []*types.ListedFlow

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListFlows request with MaxResults set at 5. The service
	// returns the first batch of results (up to 5) and a NextToken value. To see the
	// next batch of results, you can submit the ListFlows request a second time and
	// specify the NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListFlowsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListFlows{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListFlows{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFlows(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListFlows",
	}
}
