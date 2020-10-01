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

// Lists the action executions that have occurred in a pipeline.
func (c *Client) ListActionExecutions(ctx context.Context, params *ListActionExecutionsInput, optFns ...func(*Options)) (*ListActionExecutionsOutput, error) {
	stack := middleware.NewStack("ListActionExecutions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListActionExecutionsMiddlewares(stack)
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
	addOpListActionExecutionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListActionExecutions(options.Region), middleware.Before)
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
			OperationName: "ListActionExecutions",
			Err:           err,
		}
	}
	out := result.(*ListActionExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActionExecutionsInput struct {

	// The name of the pipeline for which you want to list action execution history.
	//
	// This member is required.
	PipelineName *string

	// Input information used to filter action execution history.
	Filter *types.ActionExecutionFilter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. Action
	// execution history is retained for up to 12 months, based on action execution
	// start times. Default value is 100. Detailed execution history is available for
	// executions run on or after February 21, 2019.
	MaxResults *int32

	// The token that was returned from the previous ListActionExecutions call, which
	// can be used to return the next set of action executions in the list.
	NextToken *string
}

type ListActionExecutionsOutput struct {

	// The details for a list of recent executions, such as action execution ID.
	ActionExecutionDetails []*types.ActionExecutionDetail

	// If the amount of returned information is significantly large, an identifier is
	// also returned and can be used in a subsequent ListActionExecutions call to
	// return the next set of action executions in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListActionExecutionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListActionExecutions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActionExecutions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListActionExecutions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListActionExecutions",
	}
}
