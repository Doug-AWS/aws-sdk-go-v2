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

// Gets the UI Customization information for a particular app client's app UI, if
// there is something set. If nothing is set for the particular client, but there
// is an existing pool level customization (app clientId will be ALL), then that is
// returned. If nothing is present, then an empty shape is returned.
func (c *Client) GetUICustomization(ctx context.Context, params *GetUICustomizationInput, optFns ...func(*Options)) (*GetUICustomizationOutput, error) {
	stack := middleware.NewStack("GetUICustomization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetUICustomizationMiddlewares(stack)
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
	addOpGetUICustomizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUICustomization(options.Region), middleware.Before)
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
			OperationName: "GetUICustomization",
			Err:           err,
		}
	}
	out := result.(*GetUICustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUICustomizationInput struct {

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The client ID for the client app.
	ClientId *string
}

type GetUICustomizationOutput struct {

	// The UI customization information.
	//
	// This member is required.
	UICustomization *types.UICustomizationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetUICustomizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetUICustomization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUICustomization{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUICustomization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetUICustomization",
	}
}
