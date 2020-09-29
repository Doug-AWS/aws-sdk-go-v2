// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an asynchronous batch translation job. Batch translation jobs can be used
// to translate large volumes of text across multiple documents at once. For more
// information, see async ().  <p>Batch translation jobs can be described with the
// <a>DescribeTextTranslationJob</a> operation, listed with the
// <a>ListTextTranslationJobs</a> operation, and stopped with the
// <a>StopTextTranslationJob</a> operation.</p> <note> <p>Amazon Translate does not
// support batch translation of multiple source languages at once.</p> </note>
func (c *Client) StartTextTranslationJob(ctx context.Context, params *StartTextTranslationJobInput, optFns ...func(*Options)) (*StartTextTranslationJobOutput, error) {
	stack := middleware.NewStack("StartTextTranslationJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartTextTranslationJobMiddlewares(stack)
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
	addIdempotencyToken_opStartTextTranslationJobMiddleware(stack, options)
	addOpStartTextTranslationJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartTextTranslationJob(options.Region), middleware.Before)
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
			OperationName: "StartTextTranslationJob",
			Err:           err,
		}
	}
	out := result.(*StartTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTextTranslationJobInput struct {
	// The client token of the EC2 instance calling the request. This token is
	// auto-generated when using the Amazon Translate SDK. Otherwise, use the
	// DescribeInstances () EC2 operation to retreive an instance's client token. For
	// more information, see Client Tokens () in the EC2 User Guide.
	ClientToken *string
	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that grants Amazon Translate read access to your input data. For more
	// nformation, see identity-and-access-management ().
	DataAccessRoleArn *string
	// Specifies the format and S3 location of the input documents for the translation
	// job.
	InputDataConfig *types.InputDataConfig
	// Specifies the S3 folder to which your job output will be saved.
	OutputDataConfig *types.OutputDataConfig
	// The name of the batch translation job to be performed.
	JobName *string
	// The language code of the input language. For a list of language codes, see
	// what-is-languages (). Amazon Translate does not automatically detect a source
	// language during batch translation jobs.
	SourceLanguageCode *string
	// The name of the terminology to use in the batch translation job. For a list of
	// available terminologies, use the ListTerminologies () operation.
	TerminologyNames []*string
	// The language code of the output language.
	TargetLanguageCodes []*string
}

type StartTextTranslationJobOutput struct {
	// The identifier generated for the job. To get the status of a job, use this ID
	// with the DescribeTextTranslationJob () operation.
	JobId *string
	// The status of the job. Possible values include:
	//
	//     * SUBMITTED - The job has
	// been received and is queued for processing.
	//
	//     * IN_PROGRESS - Amazon
	// Translate is processing the job.
	//
	//     * COMPLETED - The job was successfully
	// completed and the output is available.
	//
	//     * COMPLETED_WITH_ERRORS - The job
	// was completed with errors. The errors can be analyzed in the job's output.
	//
	//
	// * FAILED - The job did not complete. To get details, use the
	// DescribeTextTranslationJob () operation.
	//
	//     * STOP_REQUESTED - The user who
	// started the job has requested that it be stopped.
	//
	//     * STOPPED - The job has
	// been stopped.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartTextTranslationJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartTextTranslationJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTextTranslationJob{}, middleware.After)
}

type idempotencyToken_initializeOpStartTextTranslationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartTextTranslationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartTextTranslationJobInput ")
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
func addIdempotencyToken_opStartTextTranslationJobMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartTextTranslationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartTextTranslationJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "StartTextTranslationJob",
	}
}
