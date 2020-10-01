// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the monitoring settings for the cluster. You can use this operation to
// specify which Apache Kafka metrics you want Amazon MSK to send to Amazon
// CloudWatch. You can also specify settings for open monitoring with Prometheus.
func (c *Client) UpdateMonitoring(ctx context.Context, params *UpdateMonitoringInput, optFns ...func(*Options)) (*UpdateMonitoringOutput, error) {
	stack := middleware.NewStack("UpdateMonitoring", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateMonitoringMiddlewares(stack)
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
	addOpUpdateMonitoringValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMonitoring(options.Region), middleware.Before)
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
			OperationName: "UpdateMonitoring",
			Err:           err,
		}
	}
	out := result.(*UpdateMonitoringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request body for UpdateMonitoring.
type UpdateMonitoringInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The version of the MSK cluster to update. Cluster versions aren't simple
	// numbers. You can describe an MSK cluster to find its version. When this update
	// operation is successful, it generates a new cluster version.
	//
	// This member is required.
	CurrentVersion *string

	// Specifies which Apache Kafka metrics Amazon MSK gathers and sends to Amazon
	// CloudWatch for this cluster.
	EnhancedMonitoring types.EnhancedMonitoring

	LoggingInfo *types.LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *types.OpenMonitoringInfo
}

type UpdateMonitoringOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateMonitoringMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMonitoring{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMonitoring{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMonitoring(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateMonitoring",
	}
}
