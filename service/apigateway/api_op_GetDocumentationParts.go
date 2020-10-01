// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) GetDocumentationParts(ctx context.Context, params *GetDocumentationPartsInput, optFns ...func(*Options)) (*GetDocumentationPartsOutput, error) {
	stack := middleware.NewStack("GetDocumentationParts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDocumentationPartsMiddlewares(stack)
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
	addOpGetDocumentationPartsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentationParts(options.Region), middleware.Before)
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
			OperationName: "GetDocumentationParts",
			Err:           err,
		}
	}
	out := result.(*GetDocumentationPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the documentation parts of an API. The result may be filtered by the type,
// name, or path of API entities (targets).
type GetDocumentationPartsInput struct {

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The status of the API documentation parts to retrieve. Valid values are
	// DOCUMENTED for retrieving DocumentationPart () resources with content and
	// UNDOCUMENTED for DocumentationPart () resources without content.
	LocationStatus types.LocationStatusType

	Name *string

	// The name of API entities of the to-be-retrieved documentation parts.
	NameQuery *string

	// The path of API entities of the to-be-retrieved documentation parts.
	Path *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string

	// The type of API entities of the to-be-retrieved documentation parts.
	Type types.DocumentationPartType
}

// The collection of documentation parts of an API. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart ()
type GetDocumentationPartsOutput struct {

	// The current page of elements from this collection.
	Items []*types.DocumentationPart

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDocumentationPartsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentationParts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentationParts{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentationParts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDocumentationParts",
	}
}
