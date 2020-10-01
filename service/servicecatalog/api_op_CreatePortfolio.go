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

// Creates a portfolio. A delegated admin is authorized to invoke this command.
func (c *Client) CreatePortfolio(ctx context.Context, params *CreatePortfolioInput, optFns ...func(*Options)) (*CreatePortfolioOutput, error) {
	stack := middleware.NewStack("CreatePortfolio", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePortfolioMiddlewares(stack)
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
	addIdempotencyToken_opCreatePortfolioMiddleware(stack, options)
	addOpCreatePortfolioValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePortfolio(options.Region), middleware.Before)
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
			OperationName: "CreatePortfolio",
			Err:           err,
		}
	}
	out := result.(*CreatePortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePortfolioInput struct {

	// The name to use for display purposes.
	//
	// This member is required.
	DisplayName *string

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of the portfolio provider.
	//
	// This member is required.
	ProviderName *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The description of the portfolio.
	Description *string

	// One or more tags.
	Tags []*types.Tag
}

type CreatePortfolioOutput struct {

	// Information about the portfolio.
	PortfolioDetail *types.PortfolioDetail

	// Information about the tags associated with the portfolio.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePortfolioMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePortfolio{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePortfolio{}, middleware.After)
}

type idempotencyToken_initializeOpCreatePortfolio struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePortfolio) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePortfolio) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePortfolioInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePortfolioInput ")
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
func addIdempotencyToken_opCreatePortfolioMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreatePortfolio{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePortfolio(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreatePortfolio",
	}
}
