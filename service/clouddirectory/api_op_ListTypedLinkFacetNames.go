// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of TypedLink facet names for a particular schema. For
// more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListTypedLinkFacetNames(ctx context.Context, params *ListTypedLinkFacetNamesInput, optFns ...func(*Options)) (*ListTypedLinkFacetNamesOutput, error) {
	if params == nil {
		params = &ListTypedLinkFacetNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypedLinkFacetNames", params, optFns, addOperationListTypedLinkFacetNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypedLinkFacetNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypedLinkFacetNamesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListTypedLinkFacetNamesOutput struct {

	// The names of typed link facets that exist within the schema.
	FacetNames []*string

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTypedLinkFacetNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTypedLinkFacetNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypedLinkFacetNames{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListTypedLinkFacetNamesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypedLinkFacetNames(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListTypedLinkFacetNames(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListTypedLinkFacetNames",
	}
}
