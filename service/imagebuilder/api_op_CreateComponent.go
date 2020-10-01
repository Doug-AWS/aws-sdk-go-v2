// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new component that can be used to build, validate, test, and assess
// your image.
func (c *Client) CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) {
	stack := middleware.NewStack("CreateComponent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateComponentMiddlewares(stack)
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
	addIdempotencyToken_opCreateComponentMiddleware(stack, options)
	addOpCreateComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponent(options.Region), middleware.Before)
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
			OperationName: "CreateComponent",
			Err:           err,
		}
	}
	out := result.(*CreateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentInput struct {

	// The idempotency token of the component.
	//
	// This member is required.
	ClientToken *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	// The platform of the component.
	//
	// This member is required.
	Platform types.Platform

	// The semantic version of the component. This version follows the semantic version
	// syntax. For example, major.minor.patch. This could be versioned like software
	// (2.0.1) or like a date (2019.12.01).
	//
	// This member is required.
	SemanticVersion *string

	// The change description of the component. Describes what change has been made in
	// this version, or what makes this version different from other versions of this
	// component.
	ChangeDescription *string

	// The data of the component. Used to specify the data inline. Either data or uri
	// can be used to specify the data within the component.
	Data *string

	// The description of the component. Describes the contents of the component.
	Description *string

	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []*string

	// The tags of the component.
	Tags map[string]*string

	// The uri of the component. Must be an S3 URL and the requester must have
	// permission to access the S3 bucket. If you use S3, you can specify component
	// content up to your service quota. Either data or uri can be used to specify the
	// data within the component.
	Uri *string
}

type CreateComponentOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the component that was created by this
	// request.
	ComponentBuildVersionArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateComponentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateComponent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComponent{}, middleware.After)
}

type idempotencyToken_initializeOpCreateComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateComponentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateComponentMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CreateComponent",
	}
}
