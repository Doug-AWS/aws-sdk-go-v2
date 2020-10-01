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

// Updates a multiplex.
func (c *Client) UpdateMultiplex(ctx context.Context, params *UpdateMultiplexInput, optFns ...func(*Options)) (*UpdateMultiplexOutput, error) {
	stack := middleware.NewStack("UpdateMultiplex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateMultiplexMiddlewares(stack)
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
	addOpUpdateMultiplexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMultiplex(options.Region), middleware.Before)
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
			OperationName: "UpdateMultiplex",
			Err:           err,
		}
	}
	out := result.(*UpdateMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a multiplex.
type UpdateMultiplexInput struct {

	// ID of the multiplex to update.
	//
	// This member is required.
	MultiplexId *string

	// The new settings for a multiplex.
	MultiplexSettings *types.MultiplexSettings

	// Name of the multiplex.
	Name *string
}

// Placeholder documentation for UpdateMultiplexResponse
type UpdateMultiplexOutput struct {

	// The updated multiplex.
	Multiplex *types.Multiplex

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateMultiplexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMultiplex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMultiplex{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMultiplex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateMultiplex",
	}
}
