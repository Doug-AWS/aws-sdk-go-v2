// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of resource metadata for a given list of development endpoint
// names. After calling the ListDevEndpoints operation, you can call this operation
// to access the data to which you have been granted permissions. This operation
// supports all IAM permissions, including permission conditions that uses tags.
func (c *Client) BatchGetDevEndpoints(ctx context.Context, params *BatchGetDevEndpointsInput, optFns ...func(*Options)) (*BatchGetDevEndpointsOutput, error) {
	stack := middleware.NewStack("BatchGetDevEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetDevEndpointsMiddlewares(stack)
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
	addOpBatchGetDevEndpointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDevEndpoints(options.Region), middleware.Before)
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
			OperationName: "BatchGetDevEndpoints",
			Err:           err,
		}
	}
	out := result.(*BatchGetDevEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetDevEndpointsInput struct {
	// The list of DevEndpoint names, which might be the names returned from the
	// ListDevEndpoint operation.
	DevEndpointNames []*string
}

type BatchGetDevEndpointsOutput struct {
	// A list of DevEndpoint definitions.
	DevEndpoints []*types.DevEndpoint
	// A list of DevEndpoints not found.
	DevEndpointsNotFound []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetDevEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDevEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDevEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetDevEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchGetDevEndpoints",
	}
}
