// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets an origin request policy configuration. To get an origin request policy
// configuration, you must provide the policy’s identifier. If the origin request
// policy is attached to a distribution’s cache behavior, you can get the policy’s
// identifier using ListDistributions or GetDistribution. If the origin request
// policy is not attached to a cache behavior, you can get the identifier using
// ListOriginRequestPolicies.
func (c *Client) GetOriginRequestPolicyConfig(ctx context.Context, params *GetOriginRequestPolicyConfigInput, optFns ...func(*Options)) (*GetOriginRequestPolicyConfigOutput, error) {
	stack := middleware.NewStack("GetOriginRequestPolicyConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetOriginRequestPolicyConfigMiddlewares(stack)
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
	addOpGetOriginRequestPolicyConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOriginRequestPolicyConfig(options.Region), middleware.Before)
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
			OperationName: "GetOriginRequestPolicyConfig",
			Err:           err,
		}
	}
	out := result.(*GetOriginRequestPolicyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOriginRequestPolicyConfigInput struct {

	// The unique identifier for the origin request policy. If the origin request
	// policy is attached to a distribution’s cache behavior, you can get the policy’s
	// identifier using ListDistributions or GetDistribution. If the origin request
	// policy is not attached to a cache behavior, you can get the identifier using
	// ListOriginRequestPolicies.
	//
	// This member is required.
	Id *string
}

type GetOriginRequestPolicyConfigOutput struct {

	// The current version of the origin request policy.
	ETag *string

	// The origin request policy configuration.
	OriginRequestPolicyConfig *types.OriginRequestPolicyConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetOriginRequestPolicyConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetOriginRequestPolicyConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetOriginRequestPolicyConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOriginRequestPolicyConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetOriginRequestPolicyConfig",
	}
}
