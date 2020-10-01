// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe an instance's RAID arrays. This call accepts only one
// resource-identifying parameter. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeRaidArrays(ctx context.Context, params *DescribeRaidArraysInput, optFns ...func(*Options)) (*DescribeRaidArraysOutput, error) {
	stack := middleware.NewStack("DescribeRaidArrays", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeRaidArraysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRaidArrays(options.Region), middleware.Before)
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
			OperationName: "DescribeRaidArrays",
			Err:           err,
		}
	}
	out := result.(*DescribeRaidArraysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRaidArraysInput struct {

	// The instance ID. If you use this parameter, DescribeRaidArrays returns
	// descriptions of the RAID arrays associated with the specified instance.
	InstanceId *string

	// An array of RAID array IDs. If you use this parameter, DescribeRaidArrays
	// returns descriptions of the specified arrays. Otherwise, it returns a
	// description of every array.
	RaidArrayIds []*string

	// The stack ID.
	StackId *string
}

// Contains the response to a DescribeRaidArrays request.
type DescribeRaidArraysOutput struct {

	// A RaidArrays object that describes the specified RAID arrays.
	RaidArrays []*types.RaidArray

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeRaidArraysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRaidArrays{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRaidArrays{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRaidArrays(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeRaidArrays",
	}
}
