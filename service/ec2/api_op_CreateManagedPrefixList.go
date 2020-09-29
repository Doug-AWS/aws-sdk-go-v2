// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a managed prefix list. You can specify one or more entries for the
// prefix list. Each entry consists of a CIDR block and an optional description.
// You must specify the maximum number of entries for the prefix list. The maximum
// number of entries cannot be changed later.
func (c *Client) CreateManagedPrefixList(ctx context.Context, params *CreateManagedPrefixListInput, optFns ...func(*Options)) (*CreateManagedPrefixListOutput, error) {
	stack := middleware.NewStack("CreateManagedPrefixList", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateManagedPrefixListMiddlewares(stack)
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
	addIdempotencyToken_opCreateManagedPrefixListMiddleware(stack, options)
	addOpCreateManagedPrefixListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateManagedPrefixList(options.Region), middleware.Before)
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
			OperationName: "CreateManagedPrefixList",
			Err:           err,
		}
	}
	out := result.(*CreateManagedPrefixListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateManagedPrefixListInput struct {
	// The IP address type. Valid Values: IPv4 | IPv6
	AddressFamily *string
	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	// Constraints: Up to 255 UTF-8 characters in length.
	ClientToken *string
	// One or more entries for the prefix list.
	Entries []*types.AddPrefixListEntry
	// The maximum number of entries for the prefix list.
	MaxEntries *int32
	// The tags to apply to the prefix list during creation.
	TagSpecifications []*types.TagSpecification
	// A name for the prefix list. Constraints: Up to 255 characters in length. The
	// name cannot start with com.amazonaws.
	PrefixListName *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type CreateManagedPrefixListOutput struct {
	// Information about the prefix list.
	PrefixList *types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateManagedPrefixListMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateManagedPrefixList{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateManagedPrefixList{}, middleware.After)
}

type idempotencyToken_initializeOpCreateManagedPrefixList struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateManagedPrefixList) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateManagedPrefixList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateManagedPrefixListInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateManagedPrefixListInput ")
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
func addIdempotencyToken_opCreateManagedPrefixListMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateManagedPrefixList{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateManagedPrefixList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateManagedPrefixList",
	}
}
