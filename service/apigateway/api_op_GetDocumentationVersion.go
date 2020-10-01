// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) GetDocumentationVersion(ctx context.Context, params *GetDocumentationVersionInput, optFns ...func(*Options)) (*GetDocumentationVersionOutput, error) {
	stack := middleware.NewStack("GetDocumentationVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDocumentationVersionMiddlewares(stack)
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
	addOpGetDocumentationVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentationVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetDocumentationVersion",
			Err:           err,
		}
	}
	out := result.(*GetDocumentationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a documentation snapshot of an API.
type GetDocumentationVersionInput struct {

	// [Required] The version identifier of the to-be-retrieved documentation snapshot.
	//
	// This member is required.
	DocumentationVersion *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A snapshot of the documentation of an API. Publishing API documentation involves
// creating a documentation version associated with an API stage and exporting the
// versioned documentation to an external (e.g., OpenAPI) file. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart (), DocumentationVersions ()
type GetDocumentationVersionOutput struct {

	// The date when the API documentation snapshot is created.
	CreatedDate *time.Time

	// The description of the API documentation snapshot.
	Description *string

	// The version identifier of the API documentation snapshot.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDocumentationVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentationVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentationVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDocumentationVersion",
	}
}
