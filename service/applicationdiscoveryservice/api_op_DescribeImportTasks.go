// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of import tasks for your account, including status information,
// times, IDs, the Amazon S3 Object URL for the import file, and more.
func (c *Client) DescribeImportTasks(ctx context.Context, params *DescribeImportTasksInput, optFns ...func(*Options)) (*DescribeImportTasksOutput, error) {
	stack := middleware.NewStack("DescribeImportTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeImportTasksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImportTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeImportTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeImportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportTasksInput struct {

	// An array of name-value pairs that you provide to filter the results for the
	// DescribeImportTask request to a specific subset of results. Currently, wildcard
	// values aren't supported for filters.
	Filters []*types.ImportTaskFilter

	// The maximum number of results that you want this request to return, up to 100.
	MaxResults *int32

	// The token to request a specific page of results.
	NextToken *string
}

type DescribeImportTasksOutput struct {

	// The token to request the next page of results.
	NextToken *string

	// A returned array of import tasks that match any applied filters, up to the
	// specified number of maximum results.
	Tasks []*types.ImportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeImportTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImportTasks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImportTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImportTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeImportTasks",
	}
}
