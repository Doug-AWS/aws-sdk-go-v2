// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Validates the specified pipeline definition to ensure that it is well formed and
// can be run without error.
func (c *Client) ValidatePipelineDefinition(ctx context.Context, params *ValidatePipelineDefinitionInput, optFns ...func(*Options)) (*ValidatePipelineDefinitionOutput, error) {
	stack := middleware.NewStack("ValidatePipelineDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpValidatePipelineDefinitionMiddlewares(stack)
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
	addOpValidatePipelineDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePipelineDefinition(options.Region), middleware.Before)
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
			OperationName: "ValidatePipelineDefinition",
			Err:           err,
		}
	}
	out := result.(*ValidatePipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ValidatePipelineDefinition.
type ValidatePipelineDefinitionInput struct {
	// The objects that define the pipeline changes to validate against the pipeline.
	PipelineObjects []*types.PipelineObject
	// The parameter objects used with the pipeline.
	ParameterObjects []*types.ParameterObject
	// The ID of the pipeline.
	PipelineId *string
	// The parameter values used with the pipeline.
	ParameterValues []*types.ParameterValue
}

// Contains the output of ValidatePipelineDefinition.
type ValidatePipelineDefinitionOutput struct {
	// Any validation errors that were found.
	ValidationErrors []*types.ValidationError
	// Indicates whether there were validation errors.
	Errored *bool
	// Any validation warnings that were found.
	ValidationWarnings []*types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpValidatePipelineDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpValidatePipelineDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidatePipelineDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opValidatePipelineDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "ValidatePipelineDefinition",
	}
}
