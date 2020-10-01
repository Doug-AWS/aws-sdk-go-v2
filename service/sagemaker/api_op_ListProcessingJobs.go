// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists processing jobs that satisfy various filters.
func (c *Client) ListProcessingJobs(ctx context.Context, params *ListProcessingJobsInput, optFns ...func(*Options)) (*ListProcessingJobsOutput, error) {
	stack := middleware.NewStack("ListProcessingJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListProcessingJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProcessingJobs(options.Region), middleware.Before)
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
			OperationName: "ListProcessingJobs",
			Err:           err,
		}
	}
	out := result.(*ListProcessingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProcessingJobsInput struct {

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only processing jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only processing jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of processing jobs to return in the response.
	MaxResults *int32

	// A string in the processing job name. This filter returns only processing jobs
	// whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListProcessingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of processing jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves only processing jobs with a specific status.
	StatusEquals types.ProcessingJobStatus
}

type ListProcessingJobsOutput struct {

	// An array of ProcessingJobSummary objects, each listing a processing job.
	//
	// This member is required.
	ProcessingJobSummaries []*types.ProcessingJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of processing jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListProcessingJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListProcessingJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProcessingJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProcessingJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListProcessingJobs",
	}
}
