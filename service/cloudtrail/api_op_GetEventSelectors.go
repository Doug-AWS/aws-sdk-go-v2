// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the settings for the event selectors that you configured for your
// trail. The information returned for your event selectors includes the
// following:
//
//     * If your event selector includes read-only events, write-only
// events, or all events. This applies to both management events and data events.
//
//
// * If your event selector includes management events.
//
//     * If your event
// selector includes data events, the Amazon S3 objects or AWS Lambda functions
// that you are logging for data events.
//
// For more information, see Logging Data
// and Management Events for Trails
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html)
// in the AWS CloudTrail User Guide.
func (c *Client) GetEventSelectors(ctx context.Context, params *GetEventSelectorsInput, optFns ...func(*Options)) (*GetEventSelectorsOutput, error) {
	stack := middleware.NewStack("GetEventSelectors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetEventSelectorsMiddlewares(stack)
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
	addOpGetEventSelectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventSelectors(options.Region), middleware.Before)

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
			OperationName: "GetEventSelectors",
			Err:           err,
		}
	}
	out := result.(*GetEventSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventSelectorsInput struct {
	// Specifies the name of the trail or trail ARN. If you specify a trail name, the
	// string must meet the following requirements:
	//
	//     * Contain only ASCII letters
	// (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	//
	//     *
	// Start with a letter or number, and end with a letter or number
	//
	//     * Be between
	// 3 and 128 characters
	//
	//     * Have no adjacent periods, underscores or dashes.
	// Names like my-_namespace and my--namespace are not valid.
	//
	//     * Not be in IP
	// address format (for example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must
	// be in the format: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailName *string
}

type GetEventSelectorsOutput struct {
	// The specified trail ARN that has the event selectors.
	TrailARN *string
	// The event selectors that are configured for the trail.
	EventSelectors []*types.EventSelector

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetEventSelectorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetEventSelectors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEventSelectors{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEventSelectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetEventSelectors",
	}
}