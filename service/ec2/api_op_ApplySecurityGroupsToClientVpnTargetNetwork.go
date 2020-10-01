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

// Applies a security group to the association between the target network and the
// Client VPN endpoint. This action replaces the existing security groups with the
// specified security groups.
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetwork(ctx context.Context, params *ApplySecurityGroupsToClientVpnTargetNetworkInput, optFns ...func(*Options)) (*ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	stack := middleware.NewStack("ApplySecurityGroupsToClientVpnTargetNetwork", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpApplySecurityGroupsToClientVpnTargetNetworkMiddlewares(stack)
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
	addOpApplySecurityGroupsToClientVpnTargetNetworkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(options.Region), middleware.Before)
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
			OperationName: "ApplySecurityGroupsToClientVpnTargetNetwork",
			Err:           err,
		}
	}
	out := result.(*ApplySecurityGroupsToClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplySecurityGroupsToClientVpnTargetNetworkInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IDs of the security groups to apply to the associated target network. Up to
	// 5 security groups can be applied to an associated target network.
	//
	// This member is required.
	SecurityGroupIds []*string

	// The ID of the VPC in which the associated target network is located.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ApplySecurityGroupsToClientVpnTargetNetworkOutput struct {

	// The IDs of the applied security groups.
	SecurityGroupIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpApplySecurityGroupsToClientVpnTargetNetworkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
}

func newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ApplySecurityGroupsToClientVpnTargetNetwork",
	}
}
