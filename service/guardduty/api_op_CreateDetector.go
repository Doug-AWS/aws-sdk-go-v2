// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a single Amazon GuardDuty detector. A detector is a resource that
// represents the GuardDuty service. To start using GuardDuty, you must create a
// detector in each Region where you enable the service. You can have only one
// detector per account per Region. All data sources are enabled in a new detector
// by default.
func (c *Client) CreateDetector(ctx context.Context, params *CreateDetectorInput, optFns ...func(*Options)) (*CreateDetectorOutput, error) {
	stack := middleware.NewStack("CreateDetector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDetectorMiddlewares(stack)
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
	addIdempotencyToken_opCreateDetectorMiddleware(stack, options)
	addOpCreateDetectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDetector(options.Region), middleware.Before)
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
			OperationName: "CreateDetector",
			Err:           err,
		}
	}
	out := result.(*CreateDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDetectorInput struct {

	// A Boolean value that specifies whether the detector is to be enabled.
	//
	// This member is required.
	Enable *bool

	// The idempotency token for the create request.
	ClientToken *string

	// An object that describes which data sources will be enabled for the detector.
	DataSources *types.DataSourceConfigurations

	// An enum value that specifies how frequently updated findings are exported.
	FindingPublishingFrequency types.FindingPublishingFrequency

	// The tags to be added to a new detector resource.
	Tags map[string]*string
}

type CreateDetectorOutput struct {

	// The unique ID of the created detector.
	DetectorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDetectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDetector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDetector{}, middleware.After)
}

type idempotencyToken_initializeOpCreateDetector struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDetector) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDetectorInput ")
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
func addIdempotencyToken_opCreateDetectorMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDetector{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDetector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateDetector",
	}
}
