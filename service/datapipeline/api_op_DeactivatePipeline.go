// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deactivates the specified running pipeline. The pipeline is set to the
// DEACTIVATING state until the deactivation process completes. To resume a
// deactivated pipeline, use ActivatePipeline (). By default, the pipeline resumes
// from the last completed execution. Optionally, you can specify the date and time
// to resume the pipeline.
func (c *Client) DeactivatePipeline(ctx context.Context, params *DeactivatePipelineInput, optFns ...func(*Options)) (*DeactivatePipelineOutput, error) {
	stack := middleware.NewStack("DeactivatePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeactivatePipelineMiddlewares(stack)
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
	addOpDeactivatePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivatePipeline(options.Region), middleware.Before)
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
			OperationName: "DeactivatePipeline",
			Err:           err,
		}
	}
	out := result.(*DeactivatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeactivatePipeline.
type DeactivatePipelineInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether to cancel any running objects. The default is true, which sets
	// the state of any running objects to CANCELED. If this value is false, the
	// pipeline is deactivated after all running objects finish.
	CancelActive *bool
}

// Contains the output of DeactivatePipeline.
type DeactivatePipelineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeactivatePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeactivatePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeactivatePipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeactivatePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DeactivatePipeline",
	}
}
