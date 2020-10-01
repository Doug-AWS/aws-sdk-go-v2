// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Add a tag to a resource.
func (c *Client) CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) {
	stack := middleware.NewStack("CreateTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateTagsMiddlewares(stack)
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
	addOpCreateTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTags(options.Region), middleware.Before)
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
			OperationName: "CreateTags",
			Err:           err,
		}
	}
	out := result.(*CreateTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A map of the key-value pairs for the resource tag.
type CreateTagsInput struct {

	// The Amazon Resource Name (ARN) of the resource tag.
	//
	// This member is required.
	ResourceArn *string

	// The key-value pair for the resource tag.
	Tags map[string]*string
}

type CreateTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateTags{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "CreateTags",
	}
}
