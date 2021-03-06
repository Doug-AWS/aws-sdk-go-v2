// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Batch job queue. When you create a job queue, you associate one
// or more compute environments to the queue and assign an order of preference for
// the compute environments. You also set a priority to the job queue that
// determines the order in which the AWS Batch scheduler places jobs onto its
// associated compute environments. For example, if a compute environment is
// associated with more than one job queue, the job queue with a higher priority is
// given preference for scheduling jobs to that compute environment.
func (c *Client) CreateJobQueue(ctx context.Context, params *CreateJobQueueInput, optFns ...func(*Options)) (*CreateJobQueueOutput, error) {
	if params == nil {
		params = &CreateJobQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJobQueue", params, optFns, addOperationCreateJobQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobQueueInput struct {

	// The set of compute environments mapped to a job queue and their order relative
	// to each other. The job scheduler uses this parameter to determine which compute
	// environment should execute a given job. Compute environments must be in the
	// VALID state before you can associate them with a job queue. You can associate up
	// to three compute environments with a job queue.
	//
	// This member is required.
	ComputeEnvironmentOrder []*types.ComputeEnvironmentOrder

	// The name of the job queue.
	//
	// This member is required.
	JobQueueName *string

	// The priority of the job queue. Job queues with a higher priority (or a higher
	// integer value for the priority parameter) are evaluated first when associated
	// with the same compute environment. Priority is determined in descending order,
	// for example, a job queue with a priority value of 10 is given scheduling
	// preference over a job queue with a priority value of 1.
	//
	// This member is required.
	Priority *int32

	// The state of the job queue. If the job queue state is ENABLED, it is able to
	// accept jobs. If the job queue state is DISABLED, new jobs cannot be added to the
	// queue, but jobs already in the queue can finish.
	State types.JQState

	// The tags that you apply to the job queue to help you categorize and organize
	// your resources. Each tag consists of a key and an optional value. For more
	// information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in AWS General
	// Reference.
	Tags map[string]*string
}

type CreateJobQueueOutput struct {

	// The Amazon Resource Name (ARN) of the job queue.
	//
	// This member is required.
	JobQueueArn *string

	// The name of the job queue.
	//
	// This member is required.
	JobQueueName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateJobQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobQueue{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateJobQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobQueue(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateJobQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateJobQueue",
	}
}
