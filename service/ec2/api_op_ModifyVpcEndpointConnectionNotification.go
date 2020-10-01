// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a connection notification for VPC endpoint or VPC endpoint service. You
// can change the SNS topic for the notification, or the events for which to be
// notified.
func (c *Client) ModifyVpcEndpointConnectionNotification(ctx context.Context, params *ModifyVpcEndpointConnectionNotificationInput, optFns ...func(*Options)) (*ModifyVpcEndpointConnectionNotificationOutput, error) {
	stack := middleware.NewStack("ModifyVpcEndpointConnectionNotification", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyVpcEndpointConnectionNotificationMiddlewares(stack)
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
	addOpModifyVpcEndpointConnectionNotificationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpointConnectionNotification(options.Region), middleware.Before)
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
			OperationName: "ModifyVpcEndpointConnectionNotification",
			Err:           err,
		}
	}
	out := result.(*ModifyVpcEndpointConnectionNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointConnectionNotificationInput struct {

	// The ID of the notification.
	//
	// This member is required.
	ConnectionNotificationId *string

	// One or more events for the endpoint. Valid values are Accept, Connect, Delete,
	// and Reject.
	ConnectionEvents []*string

	// The ARN for the SNS topic for the notification.
	ConnectionNotificationArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyVpcEndpointConnectionNotificationOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyVpcEndpointConnectionNotificationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpointConnectionNotification{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpointConnectionNotification{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyVpcEndpointConnectionNotification(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointConnectionNotification",
	}
}
