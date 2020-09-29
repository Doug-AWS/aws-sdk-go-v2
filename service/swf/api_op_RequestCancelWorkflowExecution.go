// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records a WorkflowExecutionCancelRequested event in the currently running
// workflow execution identified by the given domain, workflowId, and runId. This
// logically requests the cancellation of the workflow execution as a whole. It is
// up to the decider to take appropriate actions when it receives an execution
// history with this event.  <note> <p>If the runId isn't specified, the
// <code>WorkflowExecutionCancelRequested</code> event is recorded in the history
// of the current open workflow execution with the specified workflowId in the
// domain.</p> </note> <note> <p>Because this action allows the workflow to
// properly clean up and gracefully close, it should be used instead of
// <a>TerminateWorkflowExecution</a> when possible.</p> </note> <p> <b>Access
// Control</b> </p> <p>You can use IAM policies to control this action's access to
// Amazon SWF resources as follows:</p> <ul> <li> <p>Use a <code>Resource</code>
// element with the domain name to limit the action to only specified domains.</p>
// </li> <li> <p>Use an <code>Action</code> element to allow or deny permission to
// call this action.</p> </li> <li> <p>You cannot use an IAM policy to constrain
// this action's parameters.</p> </li> </ul> <p>If the caller doesn't have
// sufficient permissions to invoke the action, or the parameter values fall
// outside the specified constraints, the action fails. The associated event
// attribute's <code>cause</code> parameter is set to
// <code>OPERATION_NOT_PERMITTED</code>. For details and example IAM policies, see
// <a
// href="https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html">Using
// IAM to Manage Access to Amazon SWF Workflows</a> in the <i>Amazon SWF Developer
// Guide</i>.</p>
func (c *Client) RequestCancelWorkflowExecution(ctx context.Context, params *RequestCancelWorkflowExecutionInput, optFns ...func(*Options)) (*RequestCancelWorkflowExecutionOutput, error) {
	stack := middleware.NewStack("RequestCancelWorkflowExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpRequestCancelWorkflowExecutionMiddlewares(stack)
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
	addOpRequestCancelWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRequestCancelWorkflowExecution(options.Region), middleware.Before)
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
			OperationName: "RequestCancelWorkflowExecution",
			Err:           err,
		}
	}
	out := result.(*RequestCancelWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestCancelWorkflowExecutionInput struct {
	// The name of the domain containing the workflow execution to cancel.
	Domain *string
	// The workflowId of the workflow execution to cancel.
	WorkflowId *string
	// The runId of the workflow execution to cancel.
	RunId *string
}

type RequestCancelWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpRequestCancelWorkflowExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpRequestCancelWorkflowExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpRequestCancelWorkflowExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opRequestCancelWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RequestCancelWorkflowExecution",
	}
}
