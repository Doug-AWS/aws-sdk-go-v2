// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Produces details about an input
func (c *Client) DescribeInput(ctx context.Context, params *DescribeInputInput, optFns ...func(*Options)) (*DescribeInputOutput, error) {
	stack := middleware.NewStack("DescribeInput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeInputMiddlewares(stack)
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
	addOpDescribeInputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInput(options.Region), middleware.Before)
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
			OperationName: "DescribeInput",
			Err:           err,
		}
	}
	out := result.(*DescribeInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputRequest
type DescribeInputInput struct {

	// Unique ID of the input
	//
	// This member is required.
	InputId *string
}

// Placeholder documentation for DescribeInputResponse
type DescribeInputOutput struct {

	// The Unique ARN of the input (generated, immutable).
	Arn *string

	// A list of channel IDs that that input is attached to (currently an input can
	// only be attached to one channel).
	AttachedChannels []*string

	// A list of the destinations of the input (PUSH-type).
	Destinations []*types.InputDestination

	// The generated ID of the input (unique for user account, immutable).
	Id *string

	// STANDARD - MediaLive expects two sources to be connected to this input. If the
	// channel is also STANDARD, both sources will be ingested. If the channel is
	// SINGLE_PIPELINE, only the first source will be ingested; the second source will
	// always be ignored, even if the first source fails. SINGLE_PIPELINE - You can
	// connect only one source to this input. If the ChannelClass is also
	// SINGLE_PIPELINE, this value is valid. If the ChannelClass is STANDARD, this
	// value is not valid because the channel requires two sources in the input.
	InputClass types.InputClass

	// Settings for the input devices.
	InputDevices []*types.InputDeviceSettings

	// Certain pull input sources can be dynamic, meaning that they can have their
	// URL's dynamically changes during input switch actions. Presently, this
	// functionality only works with MP4_FILE inputs.
	InputSourceType types.InputSourceType

	// A list of MediaConnect Flows for this input.
	MediaConnectFlows []*types.MediaConnectFlow

	// The user-assigned name (This is a mutable value).
	Name *string

	// The Amazon Resource Name (ARN) of the role this input assumes during and after
	// creation.
	RoleArn *string

	// A list of IDs for all the Input Security Groups attached to the input.
	SecurityGroups []*string

	// A list of the sources of the input (PULL-type).
	Sources []*types.InputSource

	// Placeholder documentation for InputState
	State types.InputState

	// A collection of key-value pairs.
	Tags map[string]*string

	// Placeholder documentation for InputType
	Type types.InputType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeInputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInput{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInput{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeInput",
	}
}
