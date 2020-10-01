// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends an Amazon Macie membership invitation to one or more accounts.
func (c *Client) CreateInvitations(ctx context.Context, params *CreateInvitationsInput, optFns ...func(*Options)) (*CreateInvitationsOutput, error) {
	stack := middleware.NewStack("CreateInvitations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateInvitationsMiddlewares(stack)
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
	addOpCreateInvitationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInvitations(options.Region), middleware.Before)
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
			OperationName: "CreateInvitations",
			Err:           err,
		}
	}
	out := result.(*CreateInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInvitationsInput struct {

	// An array that lists AWS account IDs, one for each account to send the invitation
	// to.
	//
	// This member is required.
	AccountIds []*string

	// Specifies whether to send an email notification to the root user of each account
	// that the invitation will be sent to. This notification is in addition to an
	// alert that the root user receives in AWS Personal Health Dashboard. To send an
	// email notification to the root user of each account, set this value to true.
	DisableEmailNotification *bool

	// A custom message to include in the invitation. Amazon Macie adds this message to
	// the standard content that it sends for an invitation.
	Message *string
}

type CreateInvitationsOutput struct {

	// An array of objects, one for each account whose invitation hasn't been
	// processed. Each object identifies the account and explains why the invitation
	// hasn't been processed for the account.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateInvitationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateInvitations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInvitations{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateInvitations",
	}
}
