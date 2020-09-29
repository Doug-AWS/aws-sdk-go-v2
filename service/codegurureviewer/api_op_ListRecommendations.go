// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of all recommendations for a completed code review.
func (c *Client) ListRecommendations(ctx context.Context, params *ListRecommendationsInput, optFns ...func(*Options)) (*ListRecommendationsOutput, error) {
	stack := middleware.NewStack("ListRecommendations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListRecommendationsMiddlewares(stack)
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
	addOpListRecommendationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendations(options.Region), middleware.Before)
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
			OperationName: "ListRecommendations",
			Err:           err,
		}
	}
	out := result.(*ListRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationsInput struct {
	// The maximum number of results that are returned per call. The default is 100.
	MaxResults *int32
	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	CodeReviewArn *string
	// Pagination token.
	NextToken *string
}

type ListRecommendationsOutput struct {
	// List of recommendations for the requested code review.
	RecommendationSummaries []*types.RecommendationSummary
	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListRecommendationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListRecommendations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecommendations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRecommendations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListRecommendations",
	}
}
