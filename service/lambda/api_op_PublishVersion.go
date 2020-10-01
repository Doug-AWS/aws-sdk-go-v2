// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a version
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the
// current code and configuration of a function. Use versions to create a snapshot
// of your function code and configuration that doesn't change.  <p>AWS Lambda
// doesn't publish a version if the function's configuration and code haven't
// changed since the last version. Use <a>UpdateFunctionCode</a> or
// <a>UpdateFunctionConfiguration</a> to update the function before publishing a
// version.</p> <p>Clients can invoke versions directly or with an alias. To create
// an alias, use <a>CreateAlias</a>.</p>
func (c *Client) PublishVersion(ctx context.Context, params *PublishVersionInput, optFns ...func(*Options)) (*PublishVersionOutput, error) {
	stack := middleware.NewStack("PublishVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPublishVersionMiddlewares(stack)
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
	addOpPublishVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPublishVersion(options.Region), middleware.Before)
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
			OperationName: "PublishVersion",
			Err:           err,
		}
	}
	out := result.(*PublishVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishVersionInput struct {

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// MyFunction.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//     * Partial ARN -
	// 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	//
	// This member is required.
	FunctionName *string

	// Only publish a version if the hash value matches the value that's specified. Use
	// this option to avoid publishing a version if the function code has changed since
	// you last updated it. You can get the hash for the version that you uploaded from
	// the output of UpdateFunctionCode ().
	CodeSha256 *string

	// A description for the version to override the description in the function
	// configuration.
	Description *string

	// Only update the function if the revision ID matches the ID that's specified. Use
	// this option to avoid publishing a version if the function configuration has
	// changed since you last updated it.
	RevisionId *string
}

// Details about a function's configuration.
type PublishVersionOutput struct {

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize *int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables.
	Environment *types.EnvironmentResponse

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []*types.Layer

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string

	// The memory that's allocated to the function.
	MemorySize *int32

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime types.Runtime

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPublishVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPublishVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opPublishVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "PublishVersion",
	}
}
