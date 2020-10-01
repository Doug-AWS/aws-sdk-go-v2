// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the association between an Amazon Lex bot and a messaging platform. This
// operation requires permission for the lex:DeleteBotChannelAssociation action.
func (c *Client) DeleteBotChannelAssociation(ctx context.Context, params *DeleteBotChannelAssociationInput, optFns ...func(*Options)) (*DeleteBotChannelAssociationOutput, error) {
	stack := middleware.NewStack("DeleteBotChannelAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteBotChannelAssociationMiddlewares(stack)
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
	addOpDeleteBotChannelAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBotChannelAssociation(options.Region), middleware.Before)
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
			OperationName: "DeleteBotChannelAssociation",
			Err:           err,
		}
	}
	out := result.(*DeleteBotChannelAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotChannelAssociationInput struct {

	// An alias that points to the specific version of the Amazon Lex bot to which this
	// association is being made.
	//
	// This member is required.
	BotAlias *string

	// The name of the Amazon Lex bot.
	//
	// This member is required.
	BotName *string

	// The name of the association. The name is case sensitive.
	//
	// This member is required.
	Name *string
}

type DeleteBotChannelAssociationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteBotChannelAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBotChannelAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBotChannelAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBotChannelAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBotChannelAssociation",
	}
}
