// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unlinks a federated identity from an existing account. Unlinked logins will be
// considered new identities next time they are seen. Removing the last linked
// login will make this identity inaccessible. This is a public API. You do not
// need any credentials to call this API.
func (c *Client) UnlinkIdentity(ctx context.Context, params *UnlinkIdentityInput, optFns ...func(*Options)) (*UnlinkIdentityOutput, error) {
	stack := middleware.NewStack("UnlinkIdentity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUnlinkIdentityMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUnlinkIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnlinkIdentity(options.Region), middleware.Before)
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
			OperationName: "UnlinkIdentity",
			Err:           err,
		}
	}
	out := result.(*UnlinkIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the UnlinkIdentity action.
type UnlinkIdentityInput struct {

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	//
	// This member is required.
	Logins map[string]*string

	// Provider names to unlink from this identity.
	//
	// This member is required.
	LoginsToRemove []*string
}

type UnlinkIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUnlinkIdentityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUnlinkIdentity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnlinkIdentity{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnlinkIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "UnlinkIdentity",
	}
}
