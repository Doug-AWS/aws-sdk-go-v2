// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

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

// Generates a temporary authentication token for accessing repositories in the
// domain. This API requires the codeartifact:GetAuthorizationToken and
// sts:GetServiceBearerToken permissions. CodeArtifact authorization tokens are
// valid for a period of 12 hours when created with the login command. You can call
// login periodically to refresh the token. When you create an authorization token
// with the GetAuthorizationToken API, you can set a custom authorization period,
// up to a maximum of 12 hours, with the durationSeconds parameter. The
// authorization period begins after login or GetAuthorizationToken is called. If
// login or GetAuthorizationToken is called while assuming a role, the token
// lifetime is independent of the maximum session duration of the role. For
// example, if you call sts assume-role and specify a session duration of 15
// minutes, then generate a CodeArtifact authorization token, the token will be
// valid for the full authorization period even though this is longer than the
// 15-minute session duration. See Using IAM Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) for more
// information on controlling session duration.
func (c *Client) GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) {
	stack := middleware.NewStack("GetAuthorizationToken", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAuthorizationTokenMiddlewares(stack)
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
	addOpGetAuthorizationTokenValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthorizationToken(options.Region), middleware.Before)
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
			OperationName: "GetAuthorizationToken",
			Err:           err,
		}
	}
	out := result.(*GetAuthorizationTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAuthorizationTokenInput struct {

	// The name of the domain that is in scope for the generated authorization token.
	//
	// This member is required.
	Domain *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The time, in seconds, that the generated authorization token is valid.
	DurationSeconds *int64
}

type GetAuthorizationTokenOutput struct {

	// The returned authentication token.
	AuthorizationToken *string

	// A timestamp that specifies the date and time the authorization token expires.
	Expiration *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAuthorizationTokenMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAuthorizationToken{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAuthorizationToken{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAuthorizationToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetAuthorizationToken",
	}
}
