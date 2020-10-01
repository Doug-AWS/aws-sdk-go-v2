// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a batch inference job. The operation can handle up to 50 million records
// and the input file must be in JSON format. For more information, see
// recommendations-batch ().
func (c *Client) CreateBatchInferenceJob(ctx context.Context, params *CreateBatchInferenceJobInput, optFns ...func(*Options)) (*CreateBatchInferenceJobOutput, error) {
	stack := middleware.NewStack("CreateBatchInferenceJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateBatchInferenceJobMiddlewares(stack)
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
	addOpCreateBatchInferenceJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchInferenceJob(options.Region), middleware.Before)
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
			OperationName: "CreateBatchInferenceJob",
			Err:           err,
		}
	}
	out := result.(*CreateBatchInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchInferenceJobInput struct {

	// The Amazon S3 path that leads to the input file to base your recommendations on.
	// The input material must be in JSON format.
	//
	// This member is required.
	JobInput *types.BatchInferenceJobInput

	// The name of the batch inference job to create.
	//
	// This member is required.
	JobName *string

	// The path to the Amazon S3 bucket where the job's output will be stored.
	//
	// This member is required.
	JobOutput *types.BatchInferenceJobOutput

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and out Amazon S3 buckets respectively.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version that will be used to
	// generate the batch inference recommendations.
	//
	// This member is required.
	SolutionVersionArn *string

	// The ARN of the filter to apply to the batch inference job. For more information
	// on using filters, see Using Filters with Amazon Personalize.
	FilterArn *string

	// The number of recommendations to retreive.
	NumResults *int32
}

type CreateBatchInferenceJobOutput struct {

	// The ARN of the batch inference job.
	BatchInferenceJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateBatchInferenceJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchInferenceJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchInferenceJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBatchInferenceJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateBatchInferenceJob",
	}
}
