// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The NotifyWorkers operation sends an email to one or more Workers that you
// specify with the Worker ID. You can specify up to 100 Worker IDs to send the
// same message with a single call to the NotifyWorkers operation. The
// NotifyWorkers operation will send a notification email to a Worker only if you
// have previously approved or rejected work from the Worker.
func (c *Client) NotifyWorkers(ctx context.Context, params *NotifyWorkersInput, optFns ...func(*Options)) (*NotifyWorkersOutput, error) {
	stack := middleware.NewStack("NotifyWorkers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpNotifyWorkersMiddlewares(stack)
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
	addOpNotifyWorkersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyWorkers(options.Region), middleware.Before)
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
			OperationName: "NotifyWorkers",
			Err:           err,
		}
	}
	out := result.(*NotifyWorkersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyWorkersInput struct {

	// The text of the email message to send. Can include up to 4,096 characters
	//
	// This member is required.
	MessageText *string

	// The subject line of the email message to send. Can include up to 200 characters.
	//
	// This member is required.
	Subject *string

	// A list of Worker IDs you wish to notify. You can notify upto 100 Workers at a
	// time.
	//
	// This member is required.
	WorkerIds []*string
}

type NotifyWorkersOutput struct {

	// When MTurk sends notifications to the list of Workers, it returns back any
	// failures it encounters in this list of NotifyWorkersFailureStatus objects.
	NotifyWorkersFailureStatuses []*types.NotifyWorkersFailureStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpNotifyWorkersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyWorkers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyWorkers{}, middleware.After)
}

func newServiceMetadataMiddleware_opNotifyWorkers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "NotifyWorkers",
	}
}
