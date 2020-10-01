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

// Gets details about a channel
func (c *Client) DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) {
	stack := middleware.NewStack("DescribeChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeChannelMiddlewares(stack)
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
	addOpDescribeChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannel(options.Region), middleware.Before)
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
			OperationName: "DescribeChannel",
			Err:           err,
		}
	}
	out := result.(*DescribeChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeChannelRequest
type DescribeChannelInput struct {

	// channel ID
	//
	// This member is required.
	ChannelId *string
}

// Placeholder documentation for DescribeChannelResponse
type DescribeChannelOutput struct {

	// The unique arn of the channel.
	Arn *string

	// The class for this channel. STANDARD for a channel with two pipelines or
	// SINGLE_PIPELINE for a channel with one pipeline.
	ChannelClass types.ChannelClass

	// A list of destinations of the channel. For UDP outputs, there is one destination
	// per output. For other types (HLS, for example), there is one destination per
	// packager.
	Destinations []*types.OutputDestination

	// The endpoints where outgoing connections initiate from
	EgressEndpoints []*types.ChannelEgressEndpoint

	// Encoder Settings
	EncoderSettings *types.EncoderSettings

	// The unique id of the channel.
	Id *string

	// List of input attachments for channel.
	InputAttachments []*types.InputAttachment

	// Placeholder documentation for InputSpecification
	InputSpecification *types.InputSpecification

	// The log level being written to CloudWatch Logs.
	LogLevel types.LogLevel

	// The name of the channel. (user-mutable)
	Name *string

	// Runtime details for the pipelines of a running channel.
	PipelineDetails []*types.PipelineDetail

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The Amazon Resource Name (ARN) of the role assumed when running the Channel.
	RoleArn *string

	// Placeholder documentation for ChannelState
	State types.ChannelState

	// A collection of key-value pairs.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeChannel",
	}
}
