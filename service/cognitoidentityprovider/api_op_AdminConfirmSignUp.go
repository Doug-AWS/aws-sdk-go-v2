// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Confirms user registration as an admin without using a confirmation code. Works
// on any user. Calling this action requires developer credentials.
func (c *Client) AdminConfirmSignUp(ctx context.Context, params *AdminConfirmSignUpInput, optFns ...func(*Options)) (*AdminConfirmSignUpOutput, error) {
	stack := middleware.NewStack("AdminConfirmSignUp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminConfirmSignUpMiddlewares(stack)
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
	addOpAdminConfirmSignUpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminConfirmSignUp(options.Region), middleware.Before)
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
			OperationName: "AdminConfirmSignUp",
			Err:           err,
		}
	}
	out := result.(*AdminConfirmSignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to confirm user registration.
type AdminConfirmSignUpInput struct {

	// The user pool ID for which you want to confirm user registration.
	//
	// This member is required.
	UserPoolId *string

	// The user name for which you want to confirm user registration.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. If your user pool configuration includes
	// triggers, the AdminConfirmSignUp API action invokes the AWS Lambda function that
	// is specified for the post confirmation trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. In
	// this payload, the clientMetadata attribute provides the data that you assigned
	// to the ClientMetadata parameter in your AdminConfirmSignUp request. In your
	// function code in AWS Lambda, you can process the ClientMetadata value to enhance
	// your workflow for your specific needs. For more information, see Customizing
	// User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string
}

// Represents the response from the server for the request to confirm registration.
type AdminConfirmSignUpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminConfirmSignUpMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminConfirmSignUp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminConfirmSignUp{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminConfirmSignUp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminConfirmSignUp",
	}
}
