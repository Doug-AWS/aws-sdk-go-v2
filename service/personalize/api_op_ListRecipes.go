// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of available recipes. The response provides the properties for
// each recipe, including the recipe's Amazon Resource Name (ARN).
func (c *Client) ListRecipes(ctx context.Context, params *ListRecipesInput, optFns ...func(*Options)) (*ListRecipesOutput, error) {
	stack := middleware.NewStack("ListRecipes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRecipesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRecipes(options.Region), middleware.Before)
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
			OperationName: "ListRecipes",
			Err:           err,
		}
	}
	out := result.(*ListRecipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecipesInput struct {

	// The maximum number of recipes to return.
	MaxResults *int32

	// A token returned from the previous call to ListRecipes for getting the next set
	// of recipes (if they exist).
	NextToken *string

	// The default is SERVICE.
	RecipeProvider types.RecipeProvider
}

type ListRecipesOutput struct {

	// A token for getting the next set of recipes.
	NextToken *string

	// The list of available recipes.
	Recipes []*types.RecipeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRecipesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRecipes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRecipes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRecipes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListRecipes",
	}
}
