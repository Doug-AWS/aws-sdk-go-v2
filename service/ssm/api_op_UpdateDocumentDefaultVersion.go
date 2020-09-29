// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Set the default version of a document.
func (c *Client) UpdateDocumentDefaultVersion(ctx context.Context, params *UpdateDocumentDefaultVersionInput, optFns ...func(*Options)) (*UpdateDocumentDefaultVersionOutput, error) {
	stack := middleware.NewStack("UpdateDocumentDefaultVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDocumentDefaultVersionMiddlewares(stack)
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
	addOpUpdateDocumentDefaultVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDocumentDefaultVersion(options.Region), middleware.Before)
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
			OperationName: "UpdateDocumentDefaultVersion",
			Err:           err,
		}
	}
	out := result.(*UpdateDocumentDefaultVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDocumentDefaultVersionInput struct {
	// The version of a custom document that you want to set as the default version.
	DocumentVersion *string
	// The name of a custom document that you want to set as the default version.
	Name *string
}

type UpdateDocumentDefaultVersionOutput struct {
	// The description of a custom document that you want to set as the default
	// version.
	Description *types.DocumentDefaultVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDocumentDefaultVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDocumentDefaultVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDocumentDefaultVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDocumentDefaultVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateDocumentDefaultVersion",
	}
}
