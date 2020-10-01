// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Performs all the read operations in a batch.
func (c *Client) BatchRead(ctx context.Context, params *BatchReadInput, optFns ...func(*Options)) (*BatchReadOutput, error) {
	stack := middleware.NewStack("BatchRead", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchReadMiddlewares(stack)
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
	addOpBatchReadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchRead(options.Region), middleware.Before)
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
			OperationName: "BatchRead",
			Err:           err,
		}
	}
	out := result.(*BatchReadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchReadInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory (). For
	// more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// A list of operations that are part of the batch.
	//
	// This member is required.
	Operations []*types.BatchReadOperation

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel
}

type BatchReadOutput struct {

	// A list of all the responses for each batch read.
	Responses []*types.BatchReadOperationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchReadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchRead{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchRead{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchRead(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "BatchRead",
	}
}
