// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new custom verification email template. For more information about
// custom verification email templates, see Using Custom Verification Email
// Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) CreateCustomVerificationEmailTemplate(ctx context.Context, params *CreateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*CreateCustomVerificationEmailTemplateOutput, error) {
	stack := middleware.NewStack("CreateCustomVerificationEmailTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateCustomVerificationEmailTemplateMiddlewares(stack)
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
	addOpCreateCustomVerificationEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomVerificationEmailTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateCustomVerificationEmailTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a custom verification email template.
type CreateCustomVerificationEmailTemplateInput struct {

	// The URL that the recipient of the verification email is sent to if his or her
	// address is not successfully verified.
	//
	// This member is required.
	FailureRedirectionURL *string

	// The email address that the custom verification email is sent from.
	//
	// This member is required.
	FromEmailAddress *string

	// The URL that the recipient of the verification email is sent to if his or her
	// address is successfully verified.
	//
	// This member is required.
	SuccessRedirectionURL *string

	// The content of the custom verification email. The total size of the email must
	// be less than 10 MB. The message body may contain HTML, with some limitations.
	// For more information, see Custom Verification Email Frequently Asked Questions
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html#custom-verification-emails-faq)
	// in the Amazon SES Developer Guide.
	//
	// This member is required.
	TemplateContent *string

	// The name of the custom verification email template.
	//
	// This member is required.
	TemplateName *string

	// The subject line of the custom verification email.
	//
	// This member is required.
	TemplateSubject *string
}

type CreateCustomVerificationEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateCustomVerificationEmailTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCustomVerificationEmailTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCustomVerificationEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateCustomVerificationEmailTemplate",
	}
}
