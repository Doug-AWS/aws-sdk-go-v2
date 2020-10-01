// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists and describes insights for the specified insight ARNs.
func (c *Client) GetInsights(ctx context.Context, params *GetInsightsInput, optFns ...func(*Options)) (*GetInsightsOutput, error) {
	stack := middleware.NewStack("GetInsights", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetInsightsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsights(options.Region), middleware.Before)
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
			OperationName: "GetInsights",
			Err:           err,
		}
	}
	out := result.(*GetInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightsInput struct {

	// The ARNs of the insights to describe. If you do not provide any insight ARNs,
	// then GetInsights returns all of your custom insights. It does not return any
	// managed insights.
	InsightArns []*string

	// The maximum number of items to return in the response.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the GetInsights
	// operation, set the value of this parameter to NULL. For subsequent calls to the
	// operation, to continue listing data, set the value of this parameter to the
	// value returned from the previous response.
	NextToken *string
}

type GetInsightsOutput struct {

	// The insights returned by the operation.
	//
	// This member is required.
	Insights []*types.Insight

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetInsightsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetInsights{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInsights{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInsights(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetInsights",
	}
}
