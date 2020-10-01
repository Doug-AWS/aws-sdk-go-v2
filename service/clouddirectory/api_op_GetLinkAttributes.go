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

// Retrieves attributes that are associated with a typed link.
func (c *Client) GetLinkAttributes(ctx context.Context, params *GetLinkAttributesInput, optFns ...func(*Options)) (*GetLinkAttributesOutput, error) {
	stack := middleware.NewStack("GetLinkAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetLinkAttributesMiddlewares(stack)
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
	addOpGetLinkAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLinkAttributes(options.Region), middleware.Before)
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
			OperationName: "GetLinkAttributes",
			Err:           err,
		}
	}
	out := result.(*GetLinkAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLinkAttributesInput struct {

	// A list of attribute names whose values will be retrieved.
	//
	// This member is required.
	AttributeNames []*string

	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// typed link resides. For more information, see arns () or Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	//
	// This member is required.
	DirectoryArn *string

	// Allows a typed link specifier to be accepted as input.
	//
	// This member is required.
	TypedLinkSpecifier *types.TypedLinkSpecifier

	// The consistency level at which to retrieve the attributes on a typed link.
	ConsistencyLevel types.ConsistencyLevel
}

type GetLinkAttributesOutput struct {

	// The attributes that are associated with the typed link.
	Attributes []*types.AttributeKeyAndValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetLinkAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetLinkAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLinkAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLinkAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetLinkAttributes",
	}
}
