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
	"time"
)

// Returns information about the most recent High Availability monitoring test that
// was performed on the host in a cluster. If a test isn't performed, the status
// and start time in the response would be null.
func (c *Client) DescribeAvailabilityMonitorTest(ctx context.Context, params *DescribeAvailabilityMonitorTestInput, optFns ...func(*Options)) (*DescribeAvailabilityMonitorTestOutput, error) {
	stack := middleware.NewStack("DescribeAvailabilityMonitorTest", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAvailabilityMonitorTestMiddlewares(stack)
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
	addOpDescribeAvailabilityMonitorTestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailabilityMonitorTest(options.Region), middleware.Before)
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
			OperationName: "DescribeAvailabilityMonitorTest",
			Err:           err,
		}
	}
	out := result.(*DescribeAvailabilityMonitorTestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAvailabilityMonitorTestInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type DescribeAvailabilityMonitorTestOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// The status of the High Availability monitoring test. If a test hasn't been
	// performed, the value of this field is null.
	Status types.AvailabilityMonitorTestStatus
	// The time the High Availability monitoring test was started. If a test hasn't
	// been performed, the value of this field is null.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAvailabilityMonitorTestMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAvailabilityMonitorTest{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAvailabilityMonitorTest{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAvailabilityMonitorTest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeAvailabilityMonitorTest",
	}
}
