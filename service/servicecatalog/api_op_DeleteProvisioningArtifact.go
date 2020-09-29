// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified provisioning artifact (also known as a version) for the
// specified product. You cannot delete a provisioning artifact associated with a
// product that was shared with you. You cannot delete the last provisioning
// artifact for a product, because a product must have at least one provisioning
// artifact.
func (c *Client) DeleteProvisioningArtifact(ctx context.Context, params *DeleteProvisioningArtifactInput, optFns ...func(*Options)) (*DeleteProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("DeleteProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteProvisioningArtifactMiddlewares(stack)
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
	addOpDeleteProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProvisioningArtifact(options.Region), middleware.Before)
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
			OperationName: "DeleteProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*DeleteProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProvisioningArtifactInput struct {
	// The product identifier.
	ProductId *string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
}

type DeleteProvisioningArtifactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DeleteProvisioningArtifact",
	}
}
