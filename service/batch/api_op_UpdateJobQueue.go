// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a job queue.
func (c *Client) UpdateJobQueue(ctx context.Context, params *UpdateJobQueueInput, optFns ...func(*Options)) (*UpdateJobQueueOutput, error) {
	stack := middleware.NewStack("UpdateJobQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateJobQueueMiddlewares(stack)
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
	addOpUpdateJobQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobQueue(options.Region), middleware.Before)
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
			OperationName: "UpdateJobQueue",
			Err:           err,
		}
	}
	out := result.(*UpdateJobQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobQueueInput struct {

	// The name or the Amazon Resource Name (ARN) of the job queue.
	//
	// This member is required.
	JobQueue *string

	// Details the set of compute environments mapped to a job queue and their order
	// relative to each other. This is one of the parameters used by the job scheduler
	// to determine which compute environment should execute a given job.
	ComputeEnvironmentOrder []*types.ComputeEnvironmentOrder

	// The priority of the job queue. Job queues with a higher priority (or a higher
	// integer value for the priority parameter) are evaluated first when associated
	// with the same compute environment. Priority is determined in descending order,
	// for example, a job queue with a priority value of 10 is given scheduling
	// preference over a job queue with a priority value of 1.
	Priority *int32

	// Describes the queue's ability to accept new jobs.
	State types.JQState
}

type UpdateJobQueueOutput struct {

	// The Amazon Resource Name (ARN) of the job queue.
	JobQueueArn *string

	// The name of the job queue.
	JobQueueName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateJobQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJobQueue{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJobQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateJobQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "UpdateJobQueue",
	}
}
