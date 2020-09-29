// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a specified pipeline with edits or changes to its structure. Use a JSON
// file with the pipeline structure and UpdatePipeline to provide the full
// structure of the pipeline. Updating the pipeline increases the version number of
// the pipeline by 1.
func (c *Client) UpdatePipeline(ctx context.Context, params *UpdatePipelineInput, optFns ...func(*Options)) (*UpdatePipelineOutput, error) {
	stack := middleware.NewStack("UpdatePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdatePipelineMiddlewares(stack)
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
	addOpUpdatePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePipeline(options.Region), middleware.Before)
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
			OperationName: "UpdatePipeline",
			Err:           err,
		}
	}
	out := result.(*UpdatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdatePipeline action.
type UpdatePipelineInput struct {
	// The name of the pipeline to be updated.
	Pipeline *types.PipelineDeclaration
}

// Represents the output of an UpdatePipeline action.
type UpdatePipelineOutput struct {
	// The structure of the updated pipeline.
	Pipeline *types.PipelineDeclaration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdatePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "UpdatePipeline",
	}
}
