// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by an AWS Lambda function to deliver evaluation results to AWS Config. This
// action is required in every AWS Lambda function that is invoked by an AWS Config
// rule.
func (c *Client) PutEvaluations(ctx context.Context, params *PutEvaluationsInput, optFns ...func(*Options)) (*PutEvaluationsOutput, error) {
	stack := middleware.NewStack("PutEvaluations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutEvaluationsMiddlewares(stack)
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
	addOpPutEvaluationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEvaluations(options.Region), middleware.Before)
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
			OperationName: "PutEvaluations",
			Err:           err,
		}
	}
	out := result.(*PutEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PutEvaluationsInput struct {

	// An encrypted token that associates an evaluation with an AWS Config rule.
	// Identifies the rule and the event that triggered the evaluation.
	//
	// This member is required.
	ResultToken *string

	// The assessments that the AWS Lambda function performs. Each evaluation
	// identifies an AWS resource and indicates whether it complies with the AWS Config
	// rule that invokes the AWS Lambda function.
	Evaluations []*types.Evaluation

	// Use this parameter to specify a test run for PutEvaluations. You can verify
	// whether your AWS Lambda function will deliver evaluation results to AWS Config.
	// No updates occur to your existing evaluations, and evaluation results are not
	// sent to AWS Config.  <note> <p>When <code>TestMode</code> is <code>true</code>,
	// <code>PutEvaluations</code> doesn't require a valid value for the
	// <code>ResultToken</code> parameter, but the value cannot be null.</p> </note>
	TestMode *bool
}

//
type PutEvaluationsOutput struct {

	// Requests that failed because of a client or server error.
	FailedEvaluations []*types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutEvaluationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutEvaluations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEvaluations{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutEvaluations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutEvaluations",
	}
}
