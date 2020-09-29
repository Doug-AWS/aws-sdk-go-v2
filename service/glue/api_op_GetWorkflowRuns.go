// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves metadata for all runs of a given workflow.
func (c *Client) GetWorkflowRuns(ctx context.Context, params *GetWorkflowRunsInput, optFns ...func(*Options)) (*GetWorkflowRunsOutput, error) {
	stack := middleware.NewStack("GetWorkflowRuns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetWorkflowRunsMiddlewares(stack)
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
	addOpGetWorkflowRunsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRuns(options.Region), middleware.Before)
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
			OperationName: "GetWorkflowRuns",
			Err:           err,
		}
	}
	out := result.(*GetWorkflowRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunsInput struct {
	// Name of the workflow whose metadata of runs should be returned.
	Name *string
	// The maximum size of the response.
	NextToken *string
	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool
	// The maximum number of workflow runs to be included in the response.
	MaxResults *int32
}

type GetWorkflowRunsOutput struct {
	// A continuation token, if not all requested workflow runs have been returned.
	NextToken *string
	// A list of workflow run metadata objects.
	Runs []*types.WorkflowRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetWorkflowRunsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkflowRuns{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkflowRuns{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetWorkflowRuns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetWorkflowRuns",
	}
}
