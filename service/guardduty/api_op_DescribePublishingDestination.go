// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the publishing destination specified by the provided
// destinationId.
func (c *Client) DescribePublishingDestination(ctx context.Context, params *DescribePublishingDestinationInput, optFns ...func(*Options)) (*DescribePublishingDestinationOutput, error) {
	stack := middleware.NewStack("DescribePublishingDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePublishingDestinationMiddlewares(stack)
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
	addOpDescribePublishingDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePublishingDestination(options.Region), middleware.Before)
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
			OperationName: "DescribePublishingDestination",
			Err:           err,
		}
	}
	out := result.(*DescribePublishingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePublishingDestinationInput struct {

	// The ID of the publishing destination to retrieve.
	//
	// This member is required.
	DestinationId *string

	// The unique ID of the detector associated with the publishing destination to
	// retrieve.
	//
	// This member is required.
	DetectorId *string
}

type DescribePublishingDestinationOutput struct {

	// The ID of the publishing destination.
	//
	// This member is required.
	DestinationId *string

	// A DestinationProperties object that includes the DestinationArn and KmsKeyArn of
	// the publishing destination.
	//
	// This member is required.
	DestinationProperties *types.DestinationProperties

	// The type of publishing destination. Currently, only Amazon S3 buckets are
	// supported.
	//
	// This member is required.
	DestinationType types.DestinationType

	// The time, in epoch millisecond format, at which GuardDuty was first unable to
	// publish findings to the destination.
	//
	// This member is required.
	PublishingFailureStartTimestamp *int64

	// The status of the publishing destination.
	//
	// This member is required.
	Status types.PublishingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePublishingDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePublishingDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePublishingDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePublishingDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "DescribePublishingDestination",
	}
}
