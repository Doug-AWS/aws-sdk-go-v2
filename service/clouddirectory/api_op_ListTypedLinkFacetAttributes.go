// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of all attribute definitions for a particular
// TypedLinkFacet (). For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListTypedLinkFacetAttributes(ctx context.Context, params *ListTypedLinkFacetAttributesInput, optFns ...func(*Options)) (*ListTypedLinkFacetAttributesOutput, error) {
	stack := middleware.NewStack("ListTypedLinkFacetAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListTypedLinkFacetAttributesMiddlewares(stack)
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
	addOpListTypedLinkFacetAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypedLinkFacetAttributes(options.Region), middleware.Before)
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
			OperationName: "ListTypedLinkFacetAttributes",
			Err:           err,
		}
	}
	out := result.(*ListTypedLinkFacetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypedLinkFacetAttributesInput struct {

	// The unique name of the typed link facet.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListTypedLinkFacetAttributesOutput struct {

	// An ordered set of attributes associate with the typed link.
	Attributes []*types.TypedLinkAttributeDefinition

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListTypedLinkFacetAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListTypedLinkFacetAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypedLinkFacetAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTypedLinkFacetAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListTypedLinkFacetAttributes",
	}
}
