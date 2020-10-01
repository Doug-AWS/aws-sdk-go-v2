// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds additional customer communication to an AWS Support case. Use the caseId
// parameter to identify the case to which to add communication. You can list a set
// of email addresses to copy on the communication by using the ccEmailAddresses
// parameter. The communicationBody value contains the text of the communication.
// <note> <ul> <li> <p>You must have a Business or Enterprise support plan to use
// the AWS Support API. </p> </li> <li> <p>If you call the AWS Support API from an
// account that does not have a Business or Enterprise support plan, the
// <code>SubscriptionRequiredException</code> error message appears. For
// information about changing your support plan, see <a
// href="http://aws.amazon.com/premiumsupport/">AWS Support</a>.</p> </li> </ul>
// </note>
func (c *Client) AddCommunicationToCase(ctx context.Context, params *AddCommunicationToCaseInput, optFns ...func(*Options)) (*AddCommunicationToCaseOutput, error) {
	stack := middleware.NewStack("AddCommunicationToCase", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddCommunicationToCaseMiddlewares(stack)
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
	addOpAddCommunicationToCaseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddCommunicationToCase(options.Region), middleware.Before)
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
			OperationName: "AddCommunicationToCase",
			Err:           err,
		}
	}
	out := result.(*AddCommunicationToCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddCommunicationToCaseInput struct {

	// The body of an email communication to add to the support case.
	//
	// This member is required.
	CommunicationBody *string

	// The ID of a set of one or more attachments for the communication to add to the
	// case. Create the set by calling AddAttachmentsToSet ()
	AttachmentSetId *string

	// The AWS Support case ID requested or returned in the call. The case ID is an
	// alphanumeric string formatted as shown in this example:
	// case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string

	// The email addresses in the CC line of an email to be added to the support case.
	CcEmailAddresses []*string
}

// The result of the AddCommunicationToCase () operation.
type AddCommunicationToCaseOutput struct {

	// True if AddCommunicationToCase () succeeds. Otherwise, returns an error.
	Result *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddCommunicationToCaseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddCommunicationToCase{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddCommunicationToCase{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddCommunicationToCase(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "AddCommunicationToCase",
	}
}
