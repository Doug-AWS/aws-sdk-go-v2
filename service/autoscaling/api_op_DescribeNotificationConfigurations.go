// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the notification actions associated with the specified Auto Scaling
// group.
func (c *Client) DescribeNotificationConfigurations(ctx context.Context, params *DescribeNotificationConfigurationsInput, optFns ...func(*Options)) (*DescribeNotificationConfigurationsOutput, error) {
	stack := middleware.NewStack("DescribeNotificationConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeNotificationConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationConfigurations(options.Region), middleware.Before)
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
			OperationName: "DescribeNotificationConfigurations",
			Err:           err,
		}
	}
	out := result.(*DescribeNotificationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationConfigurationsInput struct {

	// The name of the Auto Scaling group.
	AutoScalingGroupNames []*string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeNotificationConfigurationsOutput struct {

	// The notification configurations.
	//
	// This member is required.
	NotificationConfigurations []*types.NotificationConfiguration

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeNotificationConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeNotificationConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeNotificationConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNotificationConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeNotificationConfigurations",
	}
}
