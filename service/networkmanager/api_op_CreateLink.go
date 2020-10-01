// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new link for a specified site.
func (c *Client) CreateLink(ctx context.Context, params *CreateLinkInput, optFns ...func(*Options)) (*CreateLinkOutput, error) {
	stack := middleware.NewStack("CreateLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateLinkMiddlewares(stack)
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
	addOpCreateLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLink(options.Region), middleware.Before)
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
			OperationName: "CreateLink",
			Err:           err,
		}
	}
	out := result.(*CreateLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLinkInput struct {

	// The upload speed and download speed in Mbps.
	//
	// This member is required.
	Bandwidth *types.Bandwidth

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The ID of the site.
	//
	// This member is required.
	SiteId *string

	// A description of the link. Length Constraints: Maximum length of 256 characters.
	Description *string

	// The provider of the link. Constraints: Cannot include the following characters:
	// | \ ^ Length Constraints: Maximum length of 128 characters.
	Provider *string

	// The tags to apply to the resource during creation.
	Tags []*types.Tag

	// The type of the link. Constraints: Cannot include the following characters: | \
	// ^ Length Constraints: Maximum length of 128 characters.
	Type *string
}

type CreateLinkOutput struct {

	// Information about the link.
	Link *types.Link

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateLink{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "CreateLink",
	}
}
