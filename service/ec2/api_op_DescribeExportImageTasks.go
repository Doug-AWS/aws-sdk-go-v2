// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified export image tasks or all of your export image tasks.
func (c *Client) DescribeExportImageTasks(ctx context.Context, params *DescribeExportImageTasksInput, optFns ...func(*Options)) (*DescribeExportImageTasksOutput, error) {
	stack := middleware.NewStack("DescribeExportImageTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeExportImageTasksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportImageTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeExportImageTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeExportImageTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportImageTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IDs of the export image tasks.
	ExportImageTaskIds []*string

	// Filter tasks using the task-state filter and one of the following values:
	// active, completed, deleting, or deleted.
	Filters []*types.Filter

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// A token that indicates the next page of results.
	NextToken *string
}

type DescribeExportImageTasksOutput struct {

	// Information about the export image tasks.
	ExportImageTasks []*types.ExportImageTask

	// The token to use to get the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeExportImageTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeExportImageTasks{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeExportImageTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeExportImageTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeExportImageTasks",
	}
}
