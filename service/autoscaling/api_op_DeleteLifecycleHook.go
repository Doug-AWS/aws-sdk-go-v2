// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified lifecycle hook. If there are any outstanding lifecycle
// actions, they are completed first (ABANDON for launching instances, CONTINUE for
// terminating instances).
func (c *Client) DeleteLifecycleHook(ctx context.Context, params *DeleteLifecycleHookInput, optFns ...func(*Options)) (*DeleteLifecycleHookOutput, error) {
	stack := middleware.NewStack("DeleteLifecycleHook", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteLifecycleHookMiddlewares(stack)
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
	addOpDeleteLifecycleHookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLifecycleHook(options.Region), middleware.Before)
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
			OperationName: "DeleteLifecycleHook",
			Err:           err,
		}
	}
	out := result.(*DeleteLifecycleHookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLifecycleHookInput struct {
	// The name of the lifecycle hook.
	LifecycleHookName *string
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
}

type DeleteLifecycleHookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteLifecycleHookMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLifecycleHook{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLifecycleHook{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLifecycleHook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DeleteLifecycleHook",
	}
}
