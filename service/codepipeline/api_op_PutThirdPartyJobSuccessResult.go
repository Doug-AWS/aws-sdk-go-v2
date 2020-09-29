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

// Represents the success of a third party job as returned to the pipeline by a job
// worker. Used for partner actions only.
func (c *Client) PutThirdPartyJobSuccessResult(ctx context.Context, params *PutThirdPartyJobSuccessResultInput, optFns ...func(*Options)) (*PutThirdPartyJobSuccessResultOutput, error) {
	stack := middleware.NewStack("PutThirdPartyJobSuccessResult", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutThirdPartyJobSuccessResultMiddlewares(stack)
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
	addOpPutThirdPartyJobSuccessResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutThirdPartyJobSuccessResult(options.Region), middleware.Before)
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
			OperationName: "PutThirdPartyJobSuccessResult",
			Err:           err,
		}
	}
	out := result.(*PutThirdPartyJobSuccessResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutThirdPartyJobSuccessResult action.
type PutThirdPartyJobSuccessResultInput struct {
	// Represents information about a current revision.
	CurrentRevision *types.CurrentRevision
	// The ID of the job that successfully completed. This is the same ID returned from
	// PollForThirdPartyJobs.
	JobId *string
	// The details of the actions taken and results produced on an artifact as it
	// passes through stages in the pipeline.
	ExecutionDetails *types.ExecutionDetails
	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	ClientToken *string
	// A token generated by a job worker, such as an AWS CodeDeploy deployment ID, that
	// a successful job provides to identify a partner action in progress. Future jobs
	// use this token to identify the running instance of the action. It can be reused
	// to return more information about the progress of the partner action. When the
	// action is complete, no continuation token should be supplied.
	ContinuationToken *string
}

type PutThirdPartyJobSuccessResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutThirdPartyJobSuccessResultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutThirdPartyJobSuccessResult{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutThirdPartyJobSuccessResult{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutThirdPartyJobSuccessResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutThirdPartyJobSuccessResult",
	}
}
