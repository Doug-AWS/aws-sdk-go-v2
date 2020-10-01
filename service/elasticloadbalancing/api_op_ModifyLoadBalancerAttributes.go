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

// Modifies the attributes of the specified load balancer. You can modify the load
// balancer attributes, such as AccessLogs, ConnectionDraining, and
// CrossZoneLoadBalancing by either enabling or disabling them. Or, you can modify
// the load balancer attribute ConnectionSettings by specifying an idle connection
// timeout value for your load balancer. For more information, see the following in
// the Classic Load Balancers Guide:
//
//     * Cross-Zone Load Balancing
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
//
//
// * Connection Draining
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
//
//
// * Access Logs
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html)
//
//
// * Idle Connection Timeout
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
func (c *Client) ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error) {
	stack := middleware.NewStack("ModifyLoadBalancerAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyLoadBalancerAttributesMiddlewares(stack)
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
	addOpModifyLoadBalancerAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(options.Region), middleware.Before)
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
			OperationName: "ModifyLoadBalancerAttributes",
			Err:           err,
		}
	}
	out := result.(*ModifyLoadBalancerAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesInput struct {

	// The attributes for the load balancer.
	//
	// This member is required.
	LoadBalancerAttributes *types.LoadBalancerAttributes

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the output of ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesOutput struct {

	// Information about the load balancer attributes.
	LoadBalancerAttributes *types.LoadBalancerAttributes

	// The name of the load balancer.
	LoadBalancerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyLoadBalancerAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyLoadBalancerAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyLoadBalancerAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyLoadBalancerAttributes",
	}
}
