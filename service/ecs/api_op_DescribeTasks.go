// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a specified task or tasks.
func (c *Client) DescribeTasks(ctx context.Context, params *DescribeTasksInput, optFns ...func(*Options)) (*DescribeTasksOutput, error) {
	stack := middleware.NewStack("DescribeTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTasksMiddlewares(stack)
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
	addOpDescribeTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTasksInput struct {

	// A list of up to 100 task IDs or full ARN entries.
	//
	// This member is required.
	Tasks []*string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// task or tasks to describe. If you do not specify a cluster, the default cluster
	// is assumed. This parameter is required if the task or tasks you are describing
	// were launched in any cluster other than the default cluster.
	Cluster *string

	// Specifies whether you want to see the resource tags for the task. If TAGS is
	// specified, the tags are included in the response. If this field is omitted, tags
	// are not included in the response.
	Include []types.TaskField
}

type DescribeTasksOutput struct {

	// Any failures associated with the call.
	Failures []*types.Failure

	// The list of tasks.
	Tasks []*types.Task

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeTasks",
	}
}
