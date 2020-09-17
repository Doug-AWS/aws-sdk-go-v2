// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the request payment configuration of a bucket. To use this version of
// the operation, you must be the bucket owner. For more information, see Requester
// Pays Buckets
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html).
// <p>The following operations are related to
// <code>GetBucketRequestPayment</code>:</p> <ul> <li> <p> <a>ListObjects</a> </p>
// </li> </ul>
func (c *Client) GetBucketRequestPayment(ctx context.Context, params *GetBucketRequestPaymentInput, optFns ...func(*Options)) (*GetBucketRequestPaymentOutput, error) {
	stack := middleware.NewStack("GetBucketRequestPayment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketRequestPaymentMiddlewares(stack)
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
	addOpGetBucketRequestPaymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketRequestPayment(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)

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
			OperationName: "GetBucketRequestPayment",
			Err:           err,
		}
	}
	out := result.(*GetBucketRequestPaymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketRequestPaymentInput struct {
	// The name of the bucket for which to get the payment request configuration
	Bucket *string
}

type GetBucketRequestPaymentOutput struct {
	// Specifies who pays for the download and request fees.
	Payer types.Payer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketRequestPaymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketRequestPayment{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketRequestPayment{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketRequestPayment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketRequestPayment",
	}
}