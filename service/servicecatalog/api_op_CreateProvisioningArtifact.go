// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a provisioning artifact (also known as a version) for the specified
// product. You cannot create a provisioning artifact for a product that was shared
// with you.
func (c *Client) CreateProvisioningArtifact(ctx context.Context, params *CreateProvisioningArtifactInput, optFns ...func(*Options)) (*CreateProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("CreateProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProvisioningArtifactMiddlewares(stack)
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
	addIdempotencyToken_opCreateProvisioningArtifactMiddleware(stack, options)
	addOpCreateProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningArtifact(options.Region), middleware.Before)
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
			OperationName: "CreateProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*CreateProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningArtifactInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The configuration for the provisioning artifact.
	//
	// This member is required.
	Parameters *types.ProvisioningArtifactProperties

	// The product identifier.
	//
	// This member is required.
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
}

type CreateProvisioningArtifactOutput struct {

	// The URL of the CloudFormation template in Amazon S3, in JSON format.
	Info map[string]*string

	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *types.ProvisioningArtifactDetail

	// The status of the current request.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProvisioningArtifact{}, middleware.After)
}

type idempotencyToken_initializeOpCreateProvisioningArtifact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProvisioningArtifact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProvisioningArtifact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProvisioningArtifactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProvisioningArtifactInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProvisioningArtifactMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateProvisioningArtifact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateProvisioningArtifact",
	}
}
