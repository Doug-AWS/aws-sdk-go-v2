// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a variable.
func (c *Client) UpdateVariable(ctx context.Context, params *UpdateVariableInput, optFns ...func(*Options)) (*UpdateVariableOutput, error) {
	stack := middleware.NewStack("UpdateVariable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateVariableMiddlewares(stack)
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
	addOpUpdateVariableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVariable(options.Region), middleware.Before)
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
			OperationName: "UpdateVariable",
			Err:           err,
		}
	}
	out := result.(*UpdateVariableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVariableInput struct {

	// The name of the variable.
	//
	// This member is required.
	Name *string

	// The new default value of the variable.
	DefaultValue *string

	// The new description.
	Description *string

	// The variable type. For more information see Variable types
	// (https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types).
	VariableType *string
}

type UpdateVariableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateVariableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVariable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVariable{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateVariable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateVariable",
	}
}
