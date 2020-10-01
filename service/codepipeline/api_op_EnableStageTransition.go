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

// Enables artifacts in a pipeline to transition to a stage in a pipeline.
func (c *Client) EnableStageTransition(ctx context.Context, params *EnableStageTransitionInput, optFns ...func(*Options)) (*EnableStageTransitionOutput, error) {
	stack := middleware.NewStack("EnableStageTransition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableStageTransitionMiddlewares(stack)
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
	addOpEnableStageTransitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableStageTransition(options.Region), middleware.Before)
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
			OperationName: "EnableStageTransition",
			Err:           err,
		}
	}
	out := result.(*EnableStageTransitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an EnableStageTransition action.
type EnableStageTransitionInput struct {

	// The name of the pipeline in which you want to enable the flow of artifacts from
	// one stage to another.
	//
	// This member is required.
	PipelineName *string

	// The name of the stage where you want to enable the transition of artifacts,
	// either into the stage (inbound) or from that stage to the next stage (outbound).
	//
	// This member is required.
	StageName *string

	// Specifies whether artifacts are allowed to enter the stage and be processed by
	// the actions in that stage (inbound) or whether already processed artifacts are
	// allowed to transition to the next stage (outbound).
	//
	// This member is required.
	TransitionType types.StageTransitionType
}

type EnableStageTransitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableStageTransitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableStageTransition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableStageTransition{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableStageTransition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "EnableStageTransition",
	}
}
