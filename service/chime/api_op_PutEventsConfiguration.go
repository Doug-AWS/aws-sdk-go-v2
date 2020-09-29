// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an events configuration that allows a bot to receive outgoing events
// sent by Amazon Chime. Choose either an HTTPS endpoint or a Lambda function ARN.
// For more information, see Bot ().
func (c *Client) PutEventsConfiguration(ctx context.Context, params *PutEventsConfigurationInput, optFns ...func(*Options)) (*PutEventsConfigurationOutput, error) {
	stack := middleware.NewStack("PutEventsConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutEventsConfigurationMiddlewares(stack)
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
	addOpPutEventsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventsConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutEventsConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutEventsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventsConfigurationInput struct {
	// Lambda function ARN that allows the bot to receive outgoing events.
	LambdaFunctionArn *string
	// HTTPS endpoint that allows the bot to receive outgoing events.
	OutboundEventsHTTPSEndpoint *string
	// The Amazon Chime account ID.
	AccountId *string
	// The bot ID.
	BotId *string
}

type PutEventsConfigurationOutput struct {
	// The configuration that allows a bot to receive outgoing events. Can be either an
	// HTTPS endpoint or a Lambda function ARN.
	EventsConfiguration *types.EventsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutEventsConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutEventsConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEventsConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutEventsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutEventsConfiguration",
	}
}
