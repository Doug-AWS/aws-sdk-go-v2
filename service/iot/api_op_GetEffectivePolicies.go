// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the policies that have an effect on the authorization behavior of
// the specified device when it connects to the AWS IoT device gateway.
func (c *Client) GetEffectivePolicies(ctx context.Context, params *GetEffectivePoliciesInput, optFns ...func(*Options)) (*GetEffectivePoliciesOutput, error) {
	stack := middleware.NewStack("GetEffectivePolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetEffectivePoliciesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectivePolicies(options.Region), middleware.Before)
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
			OperationName: "GetEffectivePolicies",
			Err:           err,
		}
	}
	out := result.(*GetEffectivePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectivePoliciesInput struct {
	// The thing name.
	ThingName *string
	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string
	// The principal.
	Principal *string
}

type GetEffectivePoliciesOutput struct {
	// The effective policies.
	EffectivePolicies []*types.EffectivePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetEffectivePoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetEffectivePolicies{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEffectivePolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEffectivePolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetEffectivePolicies",
	}
}
