// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the upload buffer of a gateway. This operation is
// supported for the stored volume, cached volume, and tape gateway types.  <p>The
// response includes disk IDs that are configured as upload buffer space, and it
// includes the amount of upload buffer space allocated and used.</p>
func (c *Client) DescribeUploadBuffer(ctx context.Context, params *DescribeUploadBufferInput, optFns ...func(*Options)) (*DescribeUploadBufferOutput, error) {
	stack := middleware.NewStack("DescribeUploadBuffer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeUploadBufferMiddlewares(stack)
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
	addOpDescribeUploadBufferValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUploadBuffer(options.Region), middleware.Before)
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
			OperationName: "DescribeUploadBuffer",
			Err:           err,
		}
	}
	out := result.(*DescribeUploadBufferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUploadBufferInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

type DescribeUploadBufferOutput struct {

	// An array of the gateway's local disk IDs that are configured as working storage.
	// Each local disk ID is specified as a string (minimum length of 1 and maximum
	// length of 300). If no local disks are configured as working storage, then the
	// DiskIds array is empty.
	DiskIds []*string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// The total number of bytes allocated in the gateway's as upload buffer.
	UploadBufferAllocatedInBytes *int64

	// The total number of bytes being used in the gateway's upload buffer.
	UploadBufferUsedInBytes *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeUploadBufferMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUploadBuffer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUploadBuffer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUploadBuffer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeUploadBuffer",
	}
}
