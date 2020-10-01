// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a quantum task.
func (c *Client) CreateQuantumTask(ctx context.Context, params *CreateQuantumTaskInput, optFns ...func(*Options)) (*CreateQuantumTaskOutput, error) {
	stack := middleware.NewStack("CreateQuantumTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateQuantumTaskMiddlewares(stack)
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
	addOpCreateQuantumTaskValidationMiddleware(stack)
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
			OperationName: "CreateQuantumTask",
			Err:           err,
		}
	}
	out := result.(*CreateQuantumTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQuantumTaskInput struct {

	// The action associated with the task.
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Action *string

	// The client token associated with the request.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the device to run the task on.
	//
	// This member is required.
	DeviceArn *string

	// The S3 bucket to store task result files in.
	//
	// This member is required.
	OutputS3Bucket *string

	// The key prefix for the location in the S3 bucket to store task results in.
	//
	// This member is required.
	OutputS3KeyPrefix *string

	// The number of shots to use for the task.
	//
	// This member is required.
	Shots *int64

	// The parameters for the device to run the task on.
	// This value conforms to the media type: application/json
	DeviceParameters *string
}

type CreateQuantumTaskOutput struct {

	// The ARN of the task created by the request.
	//
	// This member is required.
	QuantumTaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateQuantumTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateQuantumTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQuantumTask{}, middleware.After)
}
