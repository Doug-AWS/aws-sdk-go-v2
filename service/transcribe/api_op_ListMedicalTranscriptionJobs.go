// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists medical transcription jobs with a specified status or substring that
// matches their names.
func (c *Client) ListMedicalTranscriptionJobs(ctx context.Context, params *ListMedicalTranscriptionJobsInput, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) {
	stack := middleware.NewStack("ListMedicalTranscriptionJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMedicalTranscriptionJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(options.Region), middleware.Before)
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
			OperationName: "ListMedicalTranscriptionJobs",
			Err:           err,
		}
	}
	out := result.(*ListMedicalTranscriptionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMedicalTranscriptionJobsInput struct {

	// When specified, the jobs returned in the list are limited to jobs whose name
	// contains the specified string.
	JobNameContains *string

	// The maximum number of medical transcription jobs to return in the response. IF
	// there are fewer results in the list, this response contains only the actual
	// results.
	MaxResults *int32

	// If you a receive a truncated result in the previous request of
	// ListMedicalTranscriptionJobs, include NextToken to fetch the next set of jobs.
	NextToken *string

	// When specified, returns only medical transcription jobs with the specified
	// status. Jobs are ordered by creation date, with the newest jobs returned first.
	// If you don't specify a status, Amazon Transcribe Medical returns all
	// transcription jobs ordered by creation date.
	Status types.TranscriptionJobStatus
}

type ListMedicalTranscriptionJobsOutput struct {

	// A list of objects containing summary information for a transcription job.
	MedicalTranscriptionJobSummaries []*types.MedicalTranscriptionJobSummary

	// The ListMedicalTranscriptionJobs operation returns a page of jobs at a time. The
	// maximum size of the page is set by the MaxResults parameter. If the number of
	// jobs exceeds what can fit on a page, Amazon Transcribe Medical returns the
	// NextPage token. Include the token in the next request to the
	// ListMedicalTranscriptionJobs operation to return in the next page of jobs.
	NextToken *string

	// The requested status of the medical transcription jobs returned.
	Status types.TranscriptionJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMedicalTranscriptionJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMedicalTranscriptionJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMedicalTranscriptionJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListMedicalTranscriptionJobs",
	}
}
