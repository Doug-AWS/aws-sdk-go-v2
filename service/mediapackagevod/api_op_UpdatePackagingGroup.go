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

// Updates a specific packaging group. You can't change the id attribute or any
// other system-generated attributes.
func (c *Client) UpdatePackagingGroup(ctx context.Context, params *UpdatePackagingGroupInput, optFns ...func(*Options)) (*UpdatePackagingGroupOutput, error) {
	stack := middleware.NewStack("UpdatePackagingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdatePackagingGroupMiddlewares(stack)
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
	addOpUpdatePackagingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePackagingGroup(options.Region), middleware.Before)
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
			OperationName: "UpdatePackagingGroup",
			Err:           err,
		}
	}
	out := result.(*UpdatePackagingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A MediaPackage VOD PackagingGroup resource configuration.
type UpdatePackagingGroupInput struct {

	// The ID of a MediaPackage VOD PackagingGroup resource.
	//
	// This member is required.
	Id *string

	// CDN Authorization credentials
	Authorization *types.Authorization
}

type UpdatePackagingGroupOutput struct {

	// The ARN of the PackagingGroup.
	Arn *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string

	// The ID of the PackagingGroup.
	Id *string

	// A collection of tags associated with a resource
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdatePackagingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePackagingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePackagingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePackagingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "UpdatePackagingGroup",
	}
}
