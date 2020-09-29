// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of the specified Amazon Resource Name (ARN) of virtual
// tapes. If a TapeARN is not specified, returns a description of all virtual tapes
// associated with the specified gateway. This operation is only supported in the
// tape gateway type.
func (c *Client) DescribeTapes(ctx context.Context, params *DescribeTapesInput, optFns ...func(*Options)) (*DescribeTapesOutput, error) {
	stack := middleware.NewStack("DescribeTapes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTapesMiddlewares(stack)
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
	addOpDescribeTapesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapes(options.Region), middleware.Before)
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
			OperationName: "DescribeTapes",
			Err:           err,
		}
	}
	out := result.(*DescribeTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapesInput
type DescribeTapesInput struct {
	// A marker value, obtained in a previous call to DescribeTapes. This marker
	// indicates which page of results to retrieve.  <p>If not specified, the first
	// page of results is retrieved.</p>
	Marker *string
	// Specifies that the number of virtual tapes described be limited to the specified
	// number.  <note> <p>Amazon Web Services may impose its own limit, if this field
	// is not set.</p> </note>
	Limit *int32
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// Specifies one or more unique Amazon Resource Names (ARNs) that represent the
	// virtual tapes you want to describe. If this parameter is not specified, Tape
	// gateway returns a description of all virtual tapes associated with the specified
	// gateway.
	TapeARNs []*string
}

// DescribeTapesOutput
type DescribeTapesOutput struct {
	// An array of virtual tape descriptions.
	Tapes []*types.Tape
	// An opaque string which can be used as part of a subsequent DescribeTapes call to
	// retrieve the next page of results.  <p>If a response does not contain a marker,
	// then there are no more results to be retrieved.</p>
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTapesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTapes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeTapes",
	}
}
