// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns AWS resources for this environment.
func (c *Client) DescribeEnvironmentResources(ctx context.Context, params *DescribeEnvironmentResourcesInput, optFns ...func(*Options)) (*DescribeEnvironmentResourcesOutput, error) {
	stack := middleware.NewStack("DescribeEnvironmentResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeEnvironmentResourcesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentResources(options.Region), middleware.Before)

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
			OperationName: "DescribeEnvironmentResources",
			Err:           err,
		}
	}
	out := result.(*DescribeEnvironmentResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe the resources in an environment.
type DescribeEnvironmentResourcesInput struct {
	// The name of the environment to retrieve AWS resource usage data. Condition: You
	// must specify either this or an EnvironmentId, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string
	// The ID of the environment to retrieve AWS resource usage data. Condition: You
	// must specify either this or an EnvironmentName, or both. If you do not specify
	// either, AWS Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentId *string
}

// Result message containing a list of environment resource descriptions.
type DescribeEnvironmentResourcesOutput struct {
	// A list of EnvironmentResourceDescription ().
	EnvironmentResources *types.EnvironmentResourceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeEnvironmentResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEnvironmentResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentResources",
	}
}