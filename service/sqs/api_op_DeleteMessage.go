// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified message from the specified queue. To select the message to
// delete, use the ReceiptHandle of the message (not the MessageId which you
// receive when you send the message). Amazon SQS can delete a message from a queue
// even if a visibility timeout setting causes the message to be locked by another
// consumer. Amazon SQS automatically deletes messages left in a queue longer than
// the retention period configured for the queue. The ReceiptHandle is associated
// with a specific instance of receiving a message. If you receive a message more
// than once, the ReceiptHandle is different each time you receive a message. When
// you use the DeleteMessage action, you must provide the most recently received
// ReceiptHandle for the message (otherwise, the request succeeds, but the message
// might not be deleted). For standard queues, it is possible to receive a message
// even after you delete it. This might happen on rare occasions if one of the
// servers which stores a copy of the message is unavailable when you send the
// request to delete the message. The copy remains on the server and might be
// returned to you during a subsequent receive request. You should ensure that your
// application is idempotent, so that receiving a message more than once does not
// cause issues.
func (c *Client) DeleteMessage(ctx context.Context, params *DeleteMessageInput, optFns ...func(*Options)) (*DeleteMessageOutput, error) {
	stack := middleware.NewStack("DeleteMessage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteMessageMiddlewares(stack)
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
	addOpDeleteMessageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMessage(options.Region), middleware.Before)
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
			OperationName: "DeleteMessage",
			Err:           err,
		}
	}
	out := result.(*DeleteMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteMessageInput struct {
	// The receipt handle associated with the message to delete.
	ReceiptHandle *string
	// The URL of the Amazon SQS queue from which messages are deleted. Queue URLs and
	// names are case-sensitive.
	QueueUrl *string
}

type DeleteMessageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteMessageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteMessage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteMessage{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMessage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "DeleteMessage",
	}
}
