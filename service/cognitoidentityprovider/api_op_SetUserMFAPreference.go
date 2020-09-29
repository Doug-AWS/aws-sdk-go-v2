// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are enabled and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are enabled. If multiple options are enabled and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign in.
func (c *Client) SetUserMFAPreference(ctx context.Context, params *SetUserMFAPreferenceInput, optFns ...func(*Options)) (*SetUserMFAPreferenceOutput, error) {
	stack := middleware.NewStack("SetUserMFAPreference", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetUserMFAPreferenceMiddlewares(stack)
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
	addOpSetUserMFAPreferenceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserMFAPreference(options.Region), middleware.Before)
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
			OperationName: "SetUserMFAPreference",
			Err:           err,
		}
	}
	out := result.(*SetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserMFAPreferenceInput struct {
	// The access token for the user.
	AccessToken *string
	// The time-based one-time password software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType
	// The SMS text message multi-factor authentication (MFA) settings.
	SMSMfaSettings *types.SMSMfaSettingsType
}

type SetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetUserMFAPreferenceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserMFAPreference{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserMFAPreference{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetUserMFAPreference(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUserMFAPreference",
	}
}
