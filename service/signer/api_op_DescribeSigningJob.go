// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a specific code signing job. You specify the job by
// using the jobId value that is returned by the StartSigningJob () operation.
func (c *Client) DescribeSigningJob(ctx context.Context, params *DescribeSigningJobInput, optFns ...func(*Options)) (*DescribeSigningJobOutput, error) {
	stack := middleware.NewStack("DescribeSigningJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeSigningJobMiddlewares(stack)
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
	addOpDescribeSigningJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSigningJob(options.Region), middleware.Before)
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
			OperationName: "DescribeSigningJob",
			Err:           err,
		}
	}
	out := result.(*DescribeSigningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSigningJobInput struct {
	// The ID of the signing job on input.
	JobId *string
}

type DescribeSigningJobOutput struct {
	// The object that contains the name of your S3 bucket or your raw code.
	Source *types.Source
	// The ID of the signing job on output.
	JobId *string
	// The Amazon Resource Name (ARN) of your code signing certificate.
	SigningMaterial *types.SigningMaterial
	// The microcontroller platform to which your signed code image will be
	// distributed.
	PlatformId *string
	// Status of the signing job.
	Status types.SigningStatus
	// Name of the S3 bucket where the signed code image is saved by code signing.
	SignedObject *types.SignedObject
	// String value that contains the status reason.
	StatusReason *string
	// The IAM principal that requested the signing job.
	RequestedBy *string
	// Date and time that the signing job was completed.
	CompletedAt *time.Time
	// A list of any overrides that were applied to the signing operation.
	Overrides *types.SigningPlatformOverrides
	// Map of user-assigned key-value pairs used during signing. These values contain
	// any information that you specified for use in your signing job.
	SigningParameters map[string]*string
	// Date and time that the signing job was created.
	CreatedAt *time.Time
	// The name of the profile that initiated the signing operation.
	ProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeSigningJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSigningJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSigningJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSigningJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "DescribeSigningJob",
	}
}
