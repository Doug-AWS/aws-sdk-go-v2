// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a policy to the specified target.
func (c *Client) AttachPolicy(ctx context.Context, params *AttachPolicyInput, optFns ...func(*Options)) (*AttachPolicyOutput, error) {
	stack := middleware.NewStack("AttachPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAttachPolicyMiddlewares(stack)
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
	addOpAttachPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachPolicy(options.Region), middleware.Before)
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
			OperationName: "AttachPolicy",
			Err:           err,
		}
	}
	out := result.(*AttachPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachPolicyInput struct {

	// The name of the policy to attach.
	//
	// This member is required.
	PolicyName *string

	// The identity
	// (https://docs.aws.amazon.com/iot/latest/developerguide/security-iam.html) to
	// which the policy is attached.
	//
	// This member is required.
	Target *string
}

type AttachPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAttachPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAttachPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AttachPolicy",
	}
}
