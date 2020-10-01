// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cache policy. You cannot delete a cache policy if it’s attached to a
// cache behavior. First update your distributions to remove the cache policy from
// all cache behaviors, then delete the cache policy. To delete a cache policy, you
// must provide the policy’s identifier and version. To get these values, you can
// use ListCachePolicies or GetCachePolicy.
func (c *Client) DeleteCachePolicy(ctx context.Context, params *DeleteCachePolicyInput, optFns ...func(*Options)) (*DeleteCachePolicyOutput, error) {
	stack := middleware.NewStack("DeleteCachePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteCachePolicyMiddlewares(stack)
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
	addOpDeleteCachePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCachePolicy(options.Region), middleware.Before)
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
			OperationName: "DeleteCachePolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCachePolicyInput struct {

	// The unique identifier for the cache policy that you are deleting. To get the
	// identifier, you can use ListCachePolicies.
	//
	// This member is required.
	Id *string

	// The version of the cache policy that you are deleting. The version is the cache
	// policy’s ETag value, which you can get using ListCachePolicies, GetCachePolicy,
	// or GetCachePolicyConfig.
	IfMatch *string
}

type DeleteCachePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteCachePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteCachePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteCachePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteCachePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteCachePolicy",
	}
}
