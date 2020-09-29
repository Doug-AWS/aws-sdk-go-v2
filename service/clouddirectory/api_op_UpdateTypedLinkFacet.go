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

// Updates a TypedLinkFacet (). For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) UpdateTypedLinkFacet(ctx context.Context, params *UpdateTypedLinkFacetInput, optFns ...func(*Options)) (*UpdateTypedLinkFacetOutput, error) {
	stack := middleware.NewStack("UpdateTypedLinkFacet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateTypedLinkFacetMiddlewares(stack)
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
	addOpUpdateTypedLinkFacetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTypedLinkFacet(options.Region), middleware.Before)
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
			OperationName: "UpdateTypedLinkFacet",
			Err:           err,
		}
	}
	out := result.(*UpdateTypedLinkFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTypedLinkFacetInput struct {
	// The unique name of the typed link facet.
	Name *string
	// Attributes update structure.
	AttributeUpdates []*types.TypedLinkFacetAttributeUpdate
	// The order of identity attributes for the facet, from most significant to least
	// significant. The ability to filter typed links considers the order that the
	// attributes are defined on the typed link facet. When providing ranges to a typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls. For more information
	// about identity attributes, see Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	IdentityAttributeOrder []*string
	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns ().
	SchemaArn *string
}

type UpdateTypedLinkFacetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateTypedLinkFacetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTypedLinkFacet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTypedLinkFacet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTypedLinkFacet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpdateTypedLinkFacet",
	}
}
