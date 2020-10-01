// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the description of an endpoint configuration created using the
// CreateEndpointConfig API.
func (c *Client) DescribeEndpointConfig(ctx context.Context, params *DescribeEndpointConfigInput, optFns ...func(*Options)) (*DescribeEndpointConfigOutput, error) {
	stack := middleware.NewStack("DescribeEndpointConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEndpointConfigMiddlewares(stack)
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
	addOpDescribeEndpointConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointConfig(options.Region), middleware.Before)
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
			OperationName: "DescribeEndpointConfig",
			Err:           err,
		}
	}
	out := result.(*DescribeEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointConfigInput struct {

	// The name of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string
}

type DescribeEndpointConfigOutput struct {

	// A timestamp that shows when the endpoint configuration was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigArn *string

	// Name of the Amazon SageMaker endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint.
	//
	// This member is required.
	ProductionVariants []*types.ProductionVariant

	//
	DataCaptureConfig *types.DataCaptureConfig

	// AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML
	// storage volume attached to the instance.
	KmsKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEndpointConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpointConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpointConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEndpointConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeEndpointConfig",
	}
}
