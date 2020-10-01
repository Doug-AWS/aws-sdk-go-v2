// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the specified metrics. You can use the returned metrics with GetMetricData
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html)
// to obtain statistical data. Up to 500 results are returned for any one call. To
// retrieve additional results, use the returned token with subsequent calls. After
// you create a metric, allow up to 15 minutes before the metric appears. You can
// see statistics about the metric sooner by using GetMetricData
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html).
// <p> <code>ListMetrics</code> doesn't return information about metrics if those
// metrics haven't reported data in the past two weeks. To retrieve those metrics,
// use <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html">GetMetricData</a>
// or <a
// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html">GetMetricStatistics</a>.</p>
func (c *Client) ListMetrics(ctx context.Context, params *ListMetricsInput, optFns ...func(*Options)) (*ListMetricsOutput, error) {
	stack := middleware.NewStack("ListMetrics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListMetricsMiddlewares(stack)
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
	addOpListMetricsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMetrics(options.Region), middleware.Before)
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
			OperationName: "ListMetrics",
			Err:           err,
		}
	}
	out := result.(*ListMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetricsInput struct {

	// The dimensions to filter against.
	Dimensions []*types.DimensionFilter

	// The name of the metric to filter against.
	MetricName *string

	// The namespace to filter against.
	Namespace *string

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string

	// To filter the results to show only metrics that have had data points published
	// in the past three hours, specify this parameter with a value of PT3H. This is
	// the only valid value for this parameter. The results that are returned are an
	// approximation of the value you specify. There is a low probability that the
	// returned results include metrics with last published data as much as 40 minutes
	// more than the specified time interval.
	RecentlyActive types.RecentlyActive
}

type ListMetricsOutput struct {

	// The metrics that match your request.
	Metrics []*types.Metric

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListMetricsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListMetrics{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListMetrics{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMetrics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "ListMetrics",
	}
}
