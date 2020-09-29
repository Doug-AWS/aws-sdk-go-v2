// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the network profile.
func (c *Client) UpdateNetworkProfile(ctx context.Context, params *UpdateNetworkProfileInput, optFns ...func(*Options)) (*UpdateNetworkProfileOutput, error) {
	stack := middleware.NewStack("UpdateNetworkProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateNetworkProfileMiddlewares(stack)
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
	addOpUpdateNetworkProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNetworkProfile(options.Region), middleware.Before)
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
			OperationName: "UpdateNetworkProfile",
			Err:           err,
		}
	}
	out := result.(*UpdateNetworkProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNetworkProfileInput struct {
	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int32
	// The Amazon Resource Name (ARN) of the project for which you want to update
	// network profile settings.
	Arn *string
	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	UplinkDelayMs *int64
	// The description of the network profile about which you are returning
	// information.
	Description *string
	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64
	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64
	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64
	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int32
	// The type of network profile to return information about. Valid values are listed
	// here.
	Type types.NetworkProfileType
	// Delay time for all packets to destination in milliseconds as an integer from 0
	// to 2000.
	DownlinkDelayMs *int64
	// The name of the network profile about which you are returning information.
	Name *string
	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64
}

type UpdateNetworkProfileOutput struct {
	// A list of the available network profiles.
	NetworkProfile *types.NetworkProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateNetworkProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNetworkProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNetworkProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNetworkProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateNetworkProfile",
	}
}
