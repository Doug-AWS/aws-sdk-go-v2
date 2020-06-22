// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests are serialized when there's no input payload but
// there are HTTP labels.
func (c *Client) HttpRequestWithLabels(ctx context.Context, params *HttpRequestWithLabelsInput, optFns ...func(*Options)) (*HttpRequestWithLabelsOutput, error) {
	stack := middleware.NewStack("HttpRequestWithLabels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Serialize.Add(&awsRestjson1_serializeOpHttpRequestWithLabels{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpRequestWithLabels{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "HttpRequestWithLabels",
			Err:           err,
		}
	}
	out := result.(*HttpRequestWithLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsInput struct {
	// Serialized in the path as true or false.
	Boolean *bool
	Double  *float64
	Float   *float32
	Integer *int32
	Long    *int64
	Short   *int16
	String_ *string
	// Note that this member has no format, so it's serialized as an RFC 3399
	// date-time.
	Timestamp *time.Time
}

type HttpRequestWithLabelsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}