// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the roles for an identity pool. These roles are used when making calls to
// GetCredentialsForIdentity () action. You must use AWS Developer credentials to
// call this API.
func (c *Client) SetIdentityPoolRoles(ctx context.Context, params *SetIdentityPoolRolesInput, optFns ...func(*Options)) (*SetIdentityPoolRolesOutput, error) {
	stack := middleware.NewStack("SetIdentityPoolRoles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetIdentityPoolRolesMiddlewares(stack)
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
	addOpSetIdentityPoolRolesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityPoolRoles(options.Region), middleware.Before)
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
			OperationName: "SetIdentityPoolRoles",
			Err:           err,
		}
	}
	out := result.(*SetIdentityPoolRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the SetIdentityPoolRoles action.
type SetIdentityPoolRolesInput struct {
	// The map of roles associated with this pool. For a given role, the key will be
	// either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]*string
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string
	// How users for a specific identity provider are to mapped to roles. This is a
	// string to RoleMapping () object map. The string identifies the identity
	// provider, for example, "graph.facebook.com" or
	// "cognito-idp-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Up to 25
	// rules can be specified per identity provider.
	RoleMappings map[string]*types.RoleMapping
}

type SetIdentityPoolRolesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetIdentityPoolRolesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetIdentityPoolRoles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetIdentityPoolRoles{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetIdentityPoolRoles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "SetIdentityPoolRoles",
	}
}
