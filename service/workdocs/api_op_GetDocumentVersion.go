// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves version metadata for the specified document.
func (c *Client) GetDocumentVersion(ctx context.Context, params *GetDocumentVersionInput, optFns ...func(*Options)) (*GetDocumentVersionOutput, error) {
	stack := middleware.NewStack("GetDocumentVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDocumentVersionMiddlewares(stack)
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
	addOpGetDocumentVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentVersion(options.Region), middleware.Before)
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
			OperationName: "GetDocumentVersion",
			Err:           err,
		}
	}
	out := result.(*GetDocumentVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentVersionInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// The version ID of the document.
	//
	// This member is required.
	VersionId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// A comma-separated list of values. Specify "SOURCE" to include a URL for the
	// source document.
	Fields *string

	// Set this to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool
}

type GetDocumentVersionOutput struct {

	// The custom metadata on the document version.
	CustomMetadata map[string]*string

	// The version metadata.
	Metadata *types.DocumentVersionMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDocumentVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "GetDocumentVersion",
	}
}
