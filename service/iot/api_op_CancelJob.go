// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a job.
func (c *Client) CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) {
	stack := middleware.NewStack("CancelJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelJobMiddlewares(stack)
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
	addOpCancelJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJob(options.Region), middleware.Before)
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
			OperationName: "CancelJob",
			Err:           err,
		}
	}
	out := result.(*CancelJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJobInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// An optional comment string describing why the job was canceled.
	Comment *string

	// (Optional) If true job executions with status "IN_PROGRESS" and "QUEUED" are
	// canceled, otherwise only job executions with status "QUEUED" are canceled. The
	// default is false. Canceling a job which is "IN_PROGRESS", will cause a device
	// which is executing the job to be unable to update the job execution status. Use
	// caution and ensure that each device executing a job which is canceled is able to
	// recover to a valid state.
	Force *bool

	// (Optional)A reason code string that explains why the job was canceled.
	ReasonCode *string
}

type CancelJobOutput struct {

	// A short text description of the job.
	Description *string

	// The job ARN.
	JobArn *string

	// The unique identifier you assigned to this job when it was created.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CancelJob",
	}
}
