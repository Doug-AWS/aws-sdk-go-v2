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

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure TOTP software token MFA. To configure either type
// of MFA, use the AdminSetUserMFAPreference () action instead.
func (c *Client) AdminSetUserSettings(ctx context.Context, params *AdminSetUserSettingsInput, optFns ...func(*Options)) (*AdminSetUserSettingsOutput, error) {
	stack := middleware.NewStack("AdminSetUserSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminSetUserSettingsMiddlewares(stack)
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
	addOpAdminSetUserSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserSettings(options.Region), middleware.Before)
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
			OperationName: "AdminSetUserSettings",
			Err:           err,
		}
	}
	out := result.(*AdminSetUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// You can use this parameter to set an MFA configuration that uses the SMS
// delivery medium.
type AdminSetUserSettingsInput struct {

	// You can use this parameter only to set an SMS configuration that uses SMS for
	// delivery.
	//
	// This member is required.
	MFAOptions []*types.MFAOptionType

	// The ID of the user pool that contains the user that you are setting options for.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user that you are setting options for.
	//
	// This member is required.
	Username *string
}

// Represents the response from the server to set user settings as an
// administrator.
type AdminSetUserSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminSetUserSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminSetUserSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminSetUserSettings",
	}
}
