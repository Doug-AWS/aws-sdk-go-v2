// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified signaling channel. DeleteSignalingChannel is an asynchronous
// operation. If you don't specify the channel's current version, the most recent
// version is deleted.
func (c *Client) DeleteSignalingChannel(ctx context.Context, params *DeleteSignalingChannelInput, optFns ...func(*Options)) (*DeleteSignalingChannelOutput, error) {
	stack := middleware.NewStack("DeleteSignalingChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteSignalingChannelMiddlewares(stack)
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
	addOpDeleteSignalingChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSignalingChannel(options.Region), middleware.Before)
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
			OperationName: "DeleteSignalingChannel",
			Err:           err,
		}
	}
	out := result.(*DeleteSignalingChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSignalingChannelInput struct {
	// The Amazon Resource Name (ARN) of the signaling channel that you want to delete.
	ChannelARN *string
	// The current version of the signaling channel that you want to delete. You can
	// obtain the current version by invoking the DescribeSignalingChannel or
	// ListSignalingChannels API operations.
	CurrentVersion *string
}

type DeleteSignalingChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteSignalingChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSignalingChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSignalingChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSignalingChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "DeleteSignalingChannel",
	}
}
