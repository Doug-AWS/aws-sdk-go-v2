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

// Copies the specified source product to the specified target product or a new
// product. You can copy a product to the same account or another account. You can
// copy a product to the same region or another region. This operation is performed
// asynchronously. To track the progress of the operation, use
// DescribeCopyProductStatus ().
func (c *Client) CopyProduct(ctx context.Context, params *CopyProductInput, optFns ...func(*Options)) (*CopyProductOutput, error) {
	stack := middleware.NewStack("CopyProduct", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCopyProductMiddlewares(stack)
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
	addIdempotencyToken_opCopyProductMiddleware(stack, options)
	addOpCopyProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyProduct(options.Region), middleware.Before)
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
			OperationName: "CopyProduct",
			Err:           err,
		}
	}
	out := result.(*CopyProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyProductInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The Amazon Resource Name (ARN) of the source product.
	//
	// This member is required.
	SourceProductArn *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The copy options. If the value is CopyTags, the tags from the source product are
	// copied to the target product.
	CopyOptions []types.CopyOption

	// The identifiers of the provisioning artifacts (also known as versions) of the
	// product to copy. By default, all provisioning artifacts are copied.
	SourceProvisioningArtifactIdentifiers []map[string]*string

	// The identifier of the target product. By default, a new product is created.
	TargetProductId *string

	// A name for the target product. The default is the name of the source product.
	TargetProductName *string
}

type CopyProductOutput struct {

	// The token to use to track the progress of the operation.
	CopyProductToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCopyProductMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCopyProduct{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopyProduct{}, middleware.After)
}

type idempotencyToken_initializeOpCopyProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCopyProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCopyProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CopyProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CopyProductInput ")
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
func addIdempotencyToken_opCopyProductMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCopyProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCopyProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CopyProduct",
	}
}
