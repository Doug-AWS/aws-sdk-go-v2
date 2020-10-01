// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the settings for the specified user, such as phone number settings.
func (c *Client) UpdateUserSettings(ctx context.Context, params *UpdateUserSettingsInput, optFns ...func(*Options)) (*UpdateUserSettingsOutput, error) {
	stack := middleware.NewStack("UpdateUserSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUserSettingsMiddlewares(stack)
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
	addOpUpdateUserSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserSettings(options.Region), middleware.Before)
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
			OperationName: "UpdateUserSettings",
			Err:           err,
		}
	}
	out := result.(*UpdateUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserSettingsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The user ID.
	//
	// This member is required.
	UserId *string

	// The user settings to update.
	//
	// This member is required.
	UserSettings *types.UserSettings
}

type UpdateUserSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUserSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserSettings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateUserSettings",
	}
}
