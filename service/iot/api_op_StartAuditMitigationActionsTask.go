// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a task that applies a set of mitigation actions to the specified target.
func (c *Client) StartAuditMitigationActionsTask(ctx context.Context, params *StartAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*StartAuditMitigationActionsTaskOutput, error) {
	stack := middleware.NewStack("StartAuditMitigationActionsTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartAuditMitigationActionsTaskMiddlewares(stack)
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
	addIdempotencyToken_opStartAuditMitigationActionsTaskMiddleware(stack, options)
	addOpStartAuditMitigationActionsTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartAuditMitigationActionsTask(options.Region), middleware.Before)
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
			OperationName: "StartAuditMitigationActionsTask",
			Err:           err,
		}
	}
	out := result.(*StartAuditMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAuditMitigationActionsTaskInput struct {

	// For an audit check, specifies which mitigation actions to apply. Those actions
	// must be defined in your AWS account.
	//
	// This member is required.
	AuditCheckToActionsMapping map[string][]*string

	// Each audit mitigation task must have a unique client request token. If you try
	// to start a new task with the same token as a task that already exists, an
	// exception occurs. If you omit this value, a unique client request token is
	// generated automatically.
	//
	// This member is required.
	ClientRequestToken *string

	// Specifies the audit findings to which the mitigation actions are applied. You
	// can apply them to a type of audit check, to all findings from an audit, or to a
	// speecific set of findings.
	//
	// This member is required.
	Target *types.AuditMitigationActionsTaskTarget

	// A unique identifier for the task. You can use this identifier to check the
	// status of the task or to cancel it.
	//
	// This member is required.
	TaskId *string
}

type StartAuditMitigationActionsTaskOutput struct {

	// The unique identifier for the audit mitigation task. This matches the taskId
	// that you specified in the request.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartAuditMitigationActionsTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartAuditMitigationActionsTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAuditMitigationActionsTask{}, middleware.After)
}

type idempotencyToken_initializeOpStartAuditMitigationActionsTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartAuditMitigationActionsTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartAuditMitigationActionsTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartAuditMitigationActionsTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartAuditMitigationActionsTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartAuditMitigationActionsTaskMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartAuditMitigationActionsTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartAuditMitigationActionsTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartAuditMitigationActionsTask",
	}
}
