// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of dataset groups created using the CreateDatasetGroup ()
// operation. For each dataset group, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). You can retrieve the
// complete set of properties by using the dataset group ARN with the
// DescribeDatasetGroup () operation.
func (c *Client) ListDatasetGroups(ctx context.Context, params *ListDatasetGroupsInput, optFns ...func(*Options)) (*ListDatasetGroupsOutput, error) {
	stack := middleware.NewStack("ListDatasetGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDatasetGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetGroups(options.Region), middleware.Before)
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
			OperationName: "ListDatasetGroups",
			Err:           err,
		}
	}
	out := result.(*ListDatasetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetGroupsInput struct {

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
}

type ListDatasetGroupsOutput struct {

	// An array of objects that summarize each dataset group's properties.
	DatasetGroups []*types.DatasetGroupSummary

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDatasetGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDatasetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListDatasetGroups",
	}
}
