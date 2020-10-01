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

// Creates an origin request policy. After you create an origin request policy, you
// can attach it to one or more cache behaviors. When it’s attached to a cache
// behavior, the origin request policy determines the values that CloudFront
// includes in requests that it sends to the origin. Each request that CloudFront
// sends to the origin includes the following:
//
//     * The request body and the URL
// path (without the domain name) from the viewer request.
//
//     * The headers that
// CloudFront automatically includes in every origin request, including Host,
// User-Agent, and X-Amz-Cf-Id.
//
//     * All HTTP headers, cookies, and URL query
// strings that are specified in the cache policy or the origin request policy.
// These can include items from the viewer request and, in the case of headers,
// additional ones that are added by CloudFront.
//
// CloudFront sends a request when
// it can’t find a valid object in its cache that matches the request. If you want
// to send values to the origin and also include them in the cache key, use
// CreateCachePolicy. For more information about origin request policies, see
// Controlling origin requests
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateOriginRequestPolicy(ctx context.Context, params *CreateOriginRequestPolicyInput, optFns ...func(*Options)) (*CreateOriginRequestPolicyOutput, error) {
	stack := middleware.NewStack("CreateOriginRequestPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateOriginRequestPolicyMiddlewares(stack)
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
	addOpCreateOriginRequestPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOriginRequestPolicy(options.Region), middleware.Before)
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
			OperationName: "CreateOriginRequestPolicy",
			Err:           err,
		}
	}
	out := result.(*CreateOriginRequestPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOriginRequestPolicyInput struct {

	// An origin request policy configuration.
	//
	// This member is required.
	OriginRequestPolicyConfig *types.OriginRequestPolicyConfig
}

type CreateOriginRequestPolicyOutput struct {

	// The current version of the origin request policy.
	ETag *string

	// The fully qualified URI of the origin request policy just created.
	Location *string

	// An origin request policy.
	OriginRequestPolicy *types.OriginRequestPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateOriginRequestPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateOriginRequestPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateOriginRequestPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateOriginRequestPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateOriginRequestPolicy",
	}
}
