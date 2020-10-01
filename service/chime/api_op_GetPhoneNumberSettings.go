// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

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

// Retrieves the phone number settings for the administrator's AWS account, such as
// the default outbound calling name.
func (c *Client) GetPhoneNumberSettings(ctx context.Context, params *GetPhoneNumberSettingsInput, optFns ...func(*Options)) (*GetPhoneNumberSettingsOutput, error) {
	stack := middleware.NewStack("GetPhoneNumberSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPhoneNumberSettingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPhoneNumberSettings(options.Region), middleware.Before)
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
			OperationName: "GetPhoneNumberSettings",
			Err:           err,
		}
	}
	out := result.(*GetPhoneNumberSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPhoneNumberSettingsInput struct {
}

type GetPhoneNumberSettingsOutput struct {

	// The default outbound calling name for the account.
	CallingName *string

	// The updated outbound calling name timestamp, in ISO 8601 format.
	CallingNameUpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPhoneNumberSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPhoneNumberSettings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPhoneNumberSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPhoneNumberSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetPhoneNumberSettings",
	}
}
