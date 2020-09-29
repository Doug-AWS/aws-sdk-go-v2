// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops an asynchronous batch translation job that is in progress. If the job's
// state is IN_PROGRESS, the job will be marked for termination and put into the
// STOP_REQUESTED state. If the job completes before it can be stopped, it is put
// into the COMPLETED state. Otherwise, the job is put into the STOPPED state.
// Asynchronous batch translation jobs are started with the StartTextTranslationJob
// () operation. You can use the DescribeTextTranslationJob () or
// ListTextTranslationJobs () operations to get a batch translation job's JobId.
func (c *Client) StopTextTranslationJob(ctx context.Context, params *StopTextTranslationJobInput, optFns ...func(*Options)) (*StopTextTranslationJobOutput, error) {
	stack := middleware.NewStack("StopTextTranslationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopTextTranslationJobMiddlewares(stack)
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
	addOpStopTextTranslationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopTextTranslationJob(options.Region), middleware.Before)
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
			OperationName: "StopTextTranslationJob",
			Err:           err,
		}
	}
	out := result.(*StopTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTextTranslationJobInput struct {
	// The job ID of the job to be stopped.
	JobId *string
}

type StopTextTranslationJobOutput struct {
	// The status of the designated job. Upon successful completion, the job's status
	// will be STOPPED.
	JobStatus types.JobStatus
	// The job ID of the stopped batch translation job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopTextTranslationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopTextTranslationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTextTranslationJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopTextTranslationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "StopTextTranslationJob",
	}
}
