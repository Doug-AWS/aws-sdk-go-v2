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

// Adds tasks, schedules, and preconditions to the specified pipeline. You can use
// PutPipelineDefinition to populate a new pipeline. PutPipelineDefinition also
// validates the configuration as it adds it to the pipeline. Changes to the
// pipeline are saved unless one of the following validation errors exist in the
// pipeline.
//
//     * An object is missing a name or identifier field.
//
//     * A
// string or reference field is empty.
//
//     * The number of objects in the pipeline
// exceeds the allowed maximum number of objects.
//
//     * The pipeline is in a
// FINISHED state.
//
// Pipeline object definitions are passed to the
// PutPipelineDefinition action and returned by the GetPipelineDefinition ()
// action.
func (c *Client) PutPipelineDefinition(ctx context.Context, params *PutPipelineDefinitionInput, optFns ...func(*Options)) (*PutPipelineDefinitionOutput, error) {
	stack := middleware.NewStack("PutPipelineDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutPipelineDefinitionMiddlewares(stack)
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
	addOpPutPipelineDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPipelineDefinition(options.Region), middleware.Before)
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
			OperationName: "PutPipelineDefinition",
			Err:           err,
		}
	}
	out := result.(*PutPipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PutPipelineDefinition.
type PutPipelineDefinitionInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The objects that define the pipeline. These objects overwrite the existing
	// pipeline definition.
	//
	// This member is required.
	PipelineObjects []*types.PipelineObject

	// The parameter objects used with the pipeline.
	ParameterObjects []*types.ParameterObject

	// The parameter values used with the pipeline.
	ParameterValues []*types.ParameterValue
}

// Contains the output of PutPipelineDefinition.
type PutPipelineDefinitionOutput struct {

	// Indicates whether there were validation errors, and the pipeline definition is
	// stored but cannot be activated until you correct the pipeline and call
	// PutPipelineDefinition to commit the corrected pipeline.
	//
	// This member is required.
	Errored *bool

	// The validation errors that are associated with the objects defined in
	// pipelineObjects.
	ValidationErrors []*types.ValidationError

	// The validation warnings that are associated with the objects defined in
	// pipelineObjects.
	ValidationWarnings []*types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutPipelineDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutPipelineDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPipelineDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutPipelineDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "PutPipelineDefinition",
	}
}
