// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the approval states for a specified pull request.
// Approval states only apply to pull requests that have one or more approval rules
// applied to them.
func (c *Client) GetPullRequestApprovalStates(ctx context.Context, params *GetPullRequestApprovalStatesInput, optFns ...func(*Options)) (*GetPullRequestApprovalStatesOutput, error) {
	stack := middleware.NewStack("GetPullRequestApprovalStates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPullRequestApprovalStatesMiddlewares(stack)
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
	addOpGetPullRequestApprovalStatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPullRequestApprovalStates(options.Region), middleware.Before)
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
			OperationName: "GetPullRequestApprovalStates",
			Err:           err,
		}
	}
	out := result.(*GetPullRequestApprovalStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPullRequestApprovalStatesInput struct {
	// The system-generated ID for the pull request revision.
	RevisionId *string
	// The system-generated ID for the pull request.
	PullRequestId *string
}

type GetPullRequestApprovalStatesOutput struct {
	// Information about users who have approved the pull request.
	Approvals []*types.Approval

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPullRequestApprovalStatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPullRequestApprovalStates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPullRequestApprovalStates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPullRequestApprovalStates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetPullRequestApprovalStates",
	}
}
