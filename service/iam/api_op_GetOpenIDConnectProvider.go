// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about the specified OpenID Connect (OIDC) provider resource
// object in IAM.
func (c *Client) GetOpenIDConnectProvider(ctx context.Context, params *GetOpenIDConnectProviderInput, optFns ...func(*Options)) (*GetOpenIDConnectProviderOutput, error) {
	stack := middleware.NewStack("GetOpenIDConnectProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetOpenIDConnectProviderMiddlewares(stack)
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
	addOpGetOpenIDConnectProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenIDConnectProvider(options.Region), middleware.Before)
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
			OperationName: "GetOpenIDConnectProvider",
			Err:           err,
		}
	}
	out := result.(*GetOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpenIDConnectProviderInput struct {

	// The Amazon Resource Name (ARN) of the OIDC provider resource object in IAM to
	// get information for. You can get a list of OIDC provider resource ARNs by using
	// the ListOpenIDConnectProviders () operation. For more information about ARNs,
	// see Amazon Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	OpenIDConnectProviderArn *string
}

// Contains the response to a successful GetOpenIDConnectProvider () request.
type GetOpenIDConnectProviderOutput struct {

	// A list of client IDs (also known as audiences) that are associated with the
	// specified IAM OIDC provider resource object. For more information, see
	// CreateOpenIDConnectProvider ().
	ClientIDList []*string

	// The date and time when the IAM OIDC provider resource object was created in the
	// AWS account.
	CreateDate *time.Time

	// A list of certificate thumbprints that are associated with the specified IAM
	// OIDC provider resource object. For more information, see
	// CreateOpenIDConnectProvider ().
	ThumbprintList []*string

	// The URL that the IAM OIDC provider resource object is associated with. For more
	// information, see CreateOpenIDConnectProvider ().
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetOpenIDConnectProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetOpenIDConnectProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetOpenIDConnectProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOpenIDConnectProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetOpenIDConnectProvider",
	}
}
