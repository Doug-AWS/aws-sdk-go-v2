// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the details of an existing subscription. Only enter values for
// parameters you want to change. Empty parameters are not updated.
func (c *Client) UpdateSubscription(ctx context.Context, params *UpdateSubscriptionInput, optFns ...func(*Options)) (*UpdateSubscriptionOutput, error) {
	stack := middleware.NewStack("UpdateSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSubscriptionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubscription(options.Region), middleware.Before)

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
			OperationName: "UpdateSubscription",
			Err:           err,
		}
	}
	out := result.(*UpdateSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSubscriptionInput struct {
	// When you initally create a subscription, AutoRenew is set to ENABLED. If
	// ENABLED, the subscription will be automatically renewed at the end of the
	// existing subscription period. You can change this by submitting an
	// UpdateSubscription request. If the UpdateSubscription request does not included
	// a value for AutoRenew, the existing value for AutoRenew remains unchanged.
	AutoRenew types.AutoRenew
}

type UpdateSubscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "UpdateSubscription",
	}
}