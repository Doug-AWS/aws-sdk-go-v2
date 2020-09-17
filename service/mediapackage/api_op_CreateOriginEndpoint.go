// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new OriginEndpoint record.
func (c *Client) CreateOriginEndpoint(ctx context.Context, params *CreateOriginEndpointInput, optFns ...func(*Options)) (*CreateOriginEndpointOutput, error) {
	stack := middleware.NewStack("CreateOriginEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateOriginEndpointMiddlewares(stack)
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
	addOpCreateOriginEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOriginEndpoint(options.Region), middleware.Before)

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
			OperationName: "CreateOriginEndpoint",
			Err:           err,
		}
	}
	out := result.(*CreateOriginEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Configuration parameters used to create a new OriginEndpoint.
type CreateOriginEndpointInput struct {
	// Amount of delay (seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds *int32
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage
	// The ID of the Channel that the OriginEndpoint will be associated with. This
	// cannot be changed after the OriginEndpoint is created.
	ChannelId *string
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage
	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []*string
	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *types.CmafPackageCreateOrUpdateParameters
	// A collection of tags associated with a resource
	Tags map[string]*string
	// A short string that will be used as the filename of the OriginEndpoint URL
	// (defaults to "index").
	ManifestName *string
	// CDN Authorization credentials
	Authorization *types.Authorization
	// Maximum duration (seconds) of content to retain for startover playback. If not
	// specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds *int32
	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination types.Origination
	// The ID of the OriginEndpoint. The ID must be unique within the region and it
	// cannot be changed after the OriginEndpoint is created.
	Id *string
	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *types.MssPackage
	// A short text description of the OriginEndpoint.
	Description *string
}

type CreateOriginEndpointOutput struct {
	// Maximum duration (seconds) of content to retain for startover playback. If not
	// specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds *int32
	// A collection of tags associated with a resource
	Tags map[string]*string
	// The URL of the packaged OriginEndpoint for consumption.
	Url *string
	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination types.Origination
	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage
	// A short string appended to the end of the OriginEndpoint URL.
	ManifestName *string
	// CDN Authorization credentials
	Authorization *types.Authorization
	// The ID of the OriginEndpoint.
	Id *string
	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage
	// The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
	Arn *string
	// A short text description of the OriginEndpoint.
	Description *string
	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *types.CmafPackage
	// Amount of delay (seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds *int32
	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *types.MssPackage
	// The ID of the Channel the OriginEndpoint is associated with.
	ChannelId *string
	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateOriginEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateOriginEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOriginEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateOriginEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "CreateOriginEndpoint",
	}
}