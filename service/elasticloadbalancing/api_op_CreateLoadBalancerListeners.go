// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more listeners for the specified load balancer. If a listener
// with the specified port does not already exist, it is created; otherwise, the
// properties of the new listener must match the properties of the existing
// listener. For more information, see Listeners for Your Classic Load Balancer
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
// in the Classic Load Balancers Guide.
func (c *Client) CreateLoadBalancerListeners(ctx context.Context, params *CreateLoadBalancerListenersInput, optFns ...func(*Options)) (*CreateLoadBalancerListenersOutput, error) {
	stack := middleware.NewStack("CreateLoadBalancerListeners", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateLoadBalancerListenersMiddlewares(stack)
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
	addOpCreateLoadBalancerListenersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancerListeners(options.Region), middleware.Before)

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
			OperationName: "CreateLoadBalancerListeners",
			Err:           err,
		}
	}
	out := result.(*CreateLoadBalancerListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancerListeners.
type CreateLoadBalancerListenersInput struct {
	// The name of the load balancer.
	LoadBalancerName *string
	// The listeners.
	Listeners []*types.Listener
}

// Contains the parameters for CreateLoadBalancerListener.
type CreateLoadBalancerListenersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateLoadBalancerListenersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancerListeners{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancerListeners{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLoadBalancerListeners(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancerListeners",
	}
}