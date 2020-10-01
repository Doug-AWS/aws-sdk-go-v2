// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your presets. This will return the
// presets themselves, not just a list of them. To retrieve the next twenty
// presets, use the nextToken string returned with the array.
func (c *Client) ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	stack := middleware.NewStack("ListPresets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPresetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPresets(options.Region), middleware.Before)
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
			OperationName: "ListPresets",
			Err:           err,
		}
	}
	out := result.(*ListPresetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPresetsInput struct {

	// Optionally, specify a preset category to limit responses to only presets from
	// that category.
	Category *string

	// Optional. When you request a list of presets, you can choose to list them
	// alphabetically by NAME or chronologically by CREATION_DATE. If you don't
	// specify, the service will list them by name.
	ListBy types.PresetListBy

	// Optional. Number of presets, up to twenty, that will be returned at one time
	MaxResults *int32

	// Use this string, provided with the response to a previous request, to request
	// the next batch of presets.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order
}

type ListPresetsOutput struct {

	// Use this string to request the next batch of presets.
	NextToken *string

	// List of presets
	Presets []*types.Preset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPresetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPresets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPresets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPresets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "ListPresets",
	}
}
