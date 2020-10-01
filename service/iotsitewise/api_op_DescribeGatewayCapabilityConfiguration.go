// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a gateway capability configuration. Each gateway
// capability defines data sources for a gateway. A capability configuration can
// contain multiple data source configurations. If you define OPC-UA sources for a
// gateway in the AWS IoT SiteWise console, all of your OPC-UA sources are stored
// in one capability configuration. To list all capability configurations for a
// gateway, use DescribeGateway
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGateway.html).
func (c *Client) DescribeGatewayCapabilityConfiguration(ctx context.Context, params *DescribeGatewayCapabilityConfigurationInput, optFns ...func(*Options)) (*DescribeGatewayCapabilityConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeGatewayCapabilityConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeGatewayCapabilityConfigurationMiddlewares(stack)
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
	addOpDescribeGatewayCapabilityConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGatewayCapabilityConfiguration(options.Region), middleware.Before)
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
			OperationName: "DescribeGatewayCapabilityConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeGatewayCapabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGatewayCapabilityConfigurationInput struct {

	// The namespace of the capability configuration. For example, if you configure
	// OPC-UA sources from the AWS IoT SiteWise console, your OPC-UA capability
	// configuration has the namespace iotsitewise:opcuacollector:version, where
	// version is a number such as 1.
	//
	// This member is required.
	CapabilityNamespace *string

	// The ID of the gateway that defines the capability configuration.
	//
	// This member is required.
	GatewayId *string
}

type DescribeGatewayCapabilityConfigurationOutput struct {

	// The JSON document that defines the gateway capability's configuration. For more
	// information, see Configuring data sources (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/configure-sources.html#configure-source-cli)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	CapabilityConfiguration *string

	// The namespace of the gateway capability.
	//
	// This member is required.
	CapabilityNamespace *string

	// The synchronization status of the capability configuration. The sync status can
	// be one of the following:
	//
	//     * IN_SYNC – The gateway is running the capability
	// configuration.
	//
	//     * OUT_OF_SYNC – The gateway hasn't received the capability
	// configuration.
	//
	//     * SYNC_FAILED – The gateway rejected the capability
	// configuration.
	//
	// This member is required.
	CapabilitySyncStatus types.CapabilitySyncStatus

	// The ID of the gateway that defines the capability configuration.
	//
	// This member is required.
	GatewayId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeGatewayCapabilityConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGatewayCapabilityConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGatewayCapabilityConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGatewayCapabilityConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeGatewayCapabilityConfiguration",
	}
}
