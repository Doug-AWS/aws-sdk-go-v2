// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new MediaPackage VOD PackagingConfiguration resource.
func (c *Client) CreatePackagingConfiguration(ctx context.Context, params *CreatePackagingConfigurationInput, optFns ...func(*Options)) (*CreatePackagingConfigurationOutput, error) {
	stack := middleware.NewStack("CreatePackagingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePackagingConfigurationMiddlewares(stack)
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
	addOpCreatePackagingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackagingConfiguration(options.Region), middleware.Before)
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
			OperationName: "CreatePackagingConfiguration",
			Err:           err,
		}
	}
	out := result.(*CreatePackagingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A new MediaPackage VOD PackagingConfiguration resource configuration.
type CreatePackagingConfigurationInput struct {
	// A collection of tags associated with a resource
	Tags map[string]*string
	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage
	// The ID of the PackagingConfiguration.
	Id *string
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage
	// The ID of a PackagingGroup.
	PackagingGroupId *string
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage
}

type CreatePackagingConfigurationOutput struct {
	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage
	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage
	// The ID of a PackagingGroup.
	PackagingGroupId *string
	// The ID of the PackagingConfiguration.
	Id *string
	// A collection of tags associated with a resource
	Tags map[string]*string
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage
	// The ARN of the PackagingConfiguration.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePackagingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackagingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackagingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePackagingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "CreatePackagingConfiguration",
	}
}
