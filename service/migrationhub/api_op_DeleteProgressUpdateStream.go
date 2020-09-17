// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a progress update stream, including all of its tasks, which was
// previously created as an AWS resource used for access control. This API has the
// following traits:
//
//     * The only parameter needed for
// DeleteProgressUpdateStream is the stream name (same as a
// CreateProgressUpdateStream call).
//
//     * The call will return, and a background
// process will asynchronously delete the stream and all of its resources (tasks,
// associated resources, resource attributes, created artifacts).
//
//     * If the
// stream takes time to be deleted, it might still show up on a
// ListProgressUpdateStreams call.
//
//     * CreateProgressUpdateStream,
// ImportMigrationTask, NotifyMigrationTaskState, and all Associate[*] APIs related
// to the tasks belonging to the stream will throw "InvalidInputException" if the
// stream of the same name is in the process of being deleted.
//
//     * Once the
// stream and all of its resources are deleted, CreateProgressUpdateStream for a
// stream of the same name will succeed, and that stream will be an entirely new
// logical resource (without any resources associated with the old stream).
func (c *Client) DeleteProgressUpdateStream(ctx context.Context, params *DeleteProgressUpdateStreamInput, optFns ...func(*Options)) (*DeleteProgressUpdateStreamOutput, error) {
	stack := middleware.NewStack("DeleteProgressUpdateStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteProgressUpdateStreamMiddlewares(stack)
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
	addOpDeleteProgressUpdateStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProgressUpdateStream(options.Region), middleware.Before)

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
			OperationName: "DeleteProgressUpdateStream",
			Err:           err,
		}
	}
	out := result.(*DeleteProgressUpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProgressUpdateStreamInput struct {
	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
	// The name of the ProgressUpdateStream. Do not store personal data in this field.
	ProgressUpdateStreamName *string
}

type DeleteProgressUpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteProgressUpdateStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProgressUpdateStream{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProgressUpdateStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteProgressUpdateStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "DeleteProgressUpdateStream",
	}
}