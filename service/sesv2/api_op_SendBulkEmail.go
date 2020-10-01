// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Composes an email message to multiple destinations.
func (c *Client) SendBulkEmail(ctx context.Context, params *SendBulkEmailInput, optFns ...func(*Options)) (*SendBulkEmailOutput, error) {
	stack := middleware.NewStack("SendBulkEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendBulkEmailMiddlewares(stack)
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
	addOpSendBulkEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendBulkEmail(options.Region), middleware.Before)
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
			OperationName: "SendBulkEmail",
			Err:           err,
		}
	}
	out := result.(*SendBulkEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send email messages to multiple destinations using
// Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type SendBulkEmailInput struct {

	// The list of bulk email entry objects.
	//
	// This member is required.
	BulkEmailEntries []*types.BulkEmailEntry

	// An object that contains the body of the message. You can specify a template
	// message.
	//
	// This member is required.
	DefaultContent *types.BulkEmailContent

	// The name of the configuration set that you want to use when sending the email.
	ConfigurationSetName *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using the SendEmail operation. Tags correspond to characteristics of the
	// email that you define, so that you can publish email sending events.
	DefaultEmailTags []*types.MessageTag

	// The address that you want bounce and complaint notifications to be sent to.
	FeedbackForwardingEmailAddress *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the FeedbackForwardingEmailAddress
	// parameter. For example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use feedback@example.com, then you would specify the
	// FeedbackForwardingEmailAddressIdentityArn to be
	// arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// FeedbackForwardingEmailAddress to be feedback@example.com. For more information
	// about sending authorization, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	FeedbackForwardingEmailAddressIdentityArn *string

	// The email address that you want to use as the "From" address for the email. The
	// address that you specify has to be verified.
	FromEmailAddress *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the FromEmailAddress parameter. For
	// example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com) attaches a policy to it
	// that authorizes you to use sender@example.com, then you would specify the
	// FromEmailAddressIdentityArn to be
	// arn:aws:ses:us-east-1:123456789012:identity/example.com, and the
	// FromEmailAddress to be sender@example.com. For more information about sending
	// authorization, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
	FromEmailAddressIdentityArn *string

	// The "Reply-to" email addresses for the message. When the recipient replies to
	// the message, each Reply-to address receives the reply.
	ReplyToAddresses []*string
}

// The following data is returned in JSON format by the service.
type SendBulkEmailOutput struct {

	// A list of BulkMailEntry objects.
	//
	// This member is required.
	BulkEmailEntryResults []*types.BulkEmailEntryResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendBulkEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendBulkEmail{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendBulkEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendBulkEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendBulkEmail",
	}
}
