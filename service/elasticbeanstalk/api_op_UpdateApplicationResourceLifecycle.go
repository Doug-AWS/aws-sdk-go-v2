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

// Modifies lifecycle settings for an application.
func (c *Client) UpdateApplicationResourceLifecycle(ctx context.Context, params *UpdateApplicationResourceLifecycleInput, optFns ...func(*Options)) (*UpdateApplicationResourceLifecycleOutput, error) {
	stack := middleware.NewStack("UpdateApplicationResourceLifecycle", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateApplicationResourceLifecycleMiddlewares(stack)
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
	addOpUpdateApplicationResourceLifecycleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplicationResourceLifecycle(options.Region), middleware.Before)
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
			OperationName: "UpdateApplicationResourceLifecycle",
			Err:           err,
		}
	}
	out := result.(*UpdateApplicationResourceLifecycleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationResourceLifecycleInput struct {
	// The lifecycle configuration.
	ResourceLifecycleConfig *types.ApplicationResourceLifecycleConfig
	// The name of the application.
	ApplicationName *string
}

type UpdateApplicationResourceLifecycleOutput struct {
	// The name of the application.
	ApplicationName *string
	// The lifecycle configuration.
	ResourceLifecycleConfig *types.ApplicationResourceLifecycleConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateApplicationResourceLifecycleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateApplicationResourceLifecycle{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateApplicationResourceLifecycle{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApplicationResourceLifecycle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "UpdateApplicationResourceLifecycle",
	}
}
