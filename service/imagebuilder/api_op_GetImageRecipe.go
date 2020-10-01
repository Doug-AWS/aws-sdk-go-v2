// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an image recipe.
func (c *Client) GetImageRecipe(ctx context.Context, params *GetImageRecipeInput, optFns ...func(*Options)) (*GetImageRecipeOutput, error) {
	stack := middleware.NewStack("GetImageRecipe", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetImageRecipeMiddlewares(stack)
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
	addOpGetImageRecipeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetImageRecipe(options.Region), middleware.Before)
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
			OperationName: "GetImageRecipe",
			Err:           err,
		}
	}
	out := result.(*GetImageRecipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImageRecipeInput struct {

	// The Amazon Resource Name (ARN) of the image recipe that you want to retrieve.
	//
	// This member is required.
	ImageRecipeArn *string
}

type GetImageRecipeOutput struct {

	// The image recipe object.
	ImageRecipe *types.ImageRecipe

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetImageRecipeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetImageRecipe{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImageRecipe{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetImageRecipe(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "GetImageRecipe",
	}
}
