// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a platform application object for one of the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging), to which devices and
// mobile apps may register. You must specify PlatformPrincipal and
// PlatformCredential attributes when using the CreatePlatformApplication action.
// PlatformPrincipal and PlatformCredential are received from the notification
// service.
//
//     * For ADM, PlatformPrincipal is client id and PlatformCredential
// is client secret.
//
//     * For Baidu, PlatformPrincipal is API key and
// PlatformCredential is secret key.
//
//     * For APNS and APNS_SANDBOX,
// PlatformPrincipal is SSL certificate and PlatformCredential is private key.
//
//
// * For GCM (Firebase Cloud Messaging), there is no PlatformPrincipal and the
// PlatformCredential is API key.
//
//     * For MPNS, PlatformPrincipal is TLS
// certificate and PlatformCredential is private key.
//
//     * For WNS,
// PlatformPrincipal is Package Security Identifier and PlatformCredential is
// secret key.
//
// You can use the returned PlatformApplicationArn as an attribute for
// the CreatePlatformEndpoint action.
func (c *Client) CreatePlatformApplication(ctx context.Context, params *CreatePlatformApplicationInput, optFns ...func(*Options)) (*CreatePlatformApplicationOutput, error) {
	stack := middleware.NewStack("CreatePlatformApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreatePlatformApplicationMiddlewares(stack)
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
	addOpCreatePlatformApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlatformApplication(options.Region), middleware.Before)
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
			OperationName: "CreatePlatformApplication",
			Err:           err,
		}
	}
	out := result.(*CreatePlatformApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreatePlatformApplication action.
type CreatePlatformApplicationInput struct {
	// For a list of attributes, see SetPlatformApplicationAttributes
	// (https://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html)
	Attributes map[string]*string
	// The following platforms are supported: ADM (Amazon Device Messaging), APNS
	// (Apple Push Notification Service), APNS_SANDBOX, and GCM (Firebase Cloud
	// Messaging).
	Platform *string
	// Application names must be made up of only uppercase and lowercase ASCII letters,
	// numbers, underscores, hyphens, and periods, and must be between 1 and 256
	// characters long.
	Name *string
}

// Response from CreatePlatformApplication action.
type CreatePlatformApplicationOutput struct {
	// PlatformApplicationArn is returned.
	PlatformApplicationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreatePlatformApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreatePlatformApplication{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePlatformApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePlatformApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CreatePlatformApplication",
	}
}
