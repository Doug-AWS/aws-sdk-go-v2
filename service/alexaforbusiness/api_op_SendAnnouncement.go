// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Triggers an asynchronous flow to send text, SSML, or audio announcements to
// rooms that are identified by a search or filter.
func (c *Client) SendAnnouncement(ctx context.Context, params *SendAnnouncementInput, optFns ...func(*Options)) (*SendAnnouncementOutput, error) {
	stack := middleware.NewStack("SendAnnouncement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSendAnnouncementMiddlewares(stack)
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
	addIdempotencyToken_opSendAnnouncementMiddleware(stack, options)
	addOpSendAnnouncementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendAnnouncement(options.Region), middleware.Before)
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
			OperationName: "SendAnnouncement",
			Err:           err,
		}
	}
	out := result.(*SendAnnouncementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendAnnouncementInput struct {
	// The time to live for an announcement. Default is 300. If delivery doesn't occur
	// within this time, the announcement is not delivered.
	TimeToLiveInSeconds *int32
	// The announcement content. This can contain only one of the three possible
	// announcement types (text, SSML or audio).
	Content *types.Content
	// The unique, user-specified identifier for the request that ensures idempotency.
	ClientRequestToken *string
	// The filters to use to send an announcement to a specified list of rooms. The
	// supported filter keys are RoomName, ProfileName, RoomArn, and ProfileArn. To
	// send to all rooms, specify an empty RoomFilters list.
	RoomFilters []*types.Filter
}

type SendAnnouncementOutput struct {
	// The identifier of the announcement.
	AnnouncementArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSendAnnouncementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSendAnnouncement{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendAnnouncement{}, middleware.After)
}

type idempotencyToken_initializeOpSendAnnouncement struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendAnnouncement) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendAnnouncement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendAnnouncementInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendAnnouncementInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendAnnouncementMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpSendAnnouncement{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendAnnouncement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SendAnnouncement",
	}
}
