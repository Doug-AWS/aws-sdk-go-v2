// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists logging levels.
func (c *Client) ListV2LoggingLevels(ctx context.Context, params *ListV2LoggingLevelsInput, optFns ...func(*Options)) (*ListV2LoggingLevelsOutput, error) {
	stack := middleware.NewStack("ListV2LoggingLevels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListV2LoggingLevelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListV2LoggingLevels(options.Region), middleware.Before)
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
			OperationName: "ListV2LoggingLevels",
			Err:           err,
		}
	}
	out := result.(*ListV2LoggingLevelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListV2LoggingLevelsInput struct {
	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string
	// The maximum number of results to return at one time.
	MaxResults *int32
	// The type of resource for which you are configuring logging. Must be THING_Group.
	TargetType types.LogTargetType
}

type ListV2LoggingLevelsOutput struct {
	// The logging configuration for a target.
	LogTargetConfigurations []*types.LogTargetConfiguration
	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListV2LoggingLevelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListV2LoggingLevels{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListV2LoggingLevels{}, middleware.After)
}

func newServiceMetadataMiddleware_opListV2LoggingLevels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListV2LoggingLevels",
	}
}
