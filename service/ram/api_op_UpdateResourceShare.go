// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified resource share that you own.
func (c *Client) UpdateResourceShare(ctx context.Context, params *UpdateResourceShareInput, optFns ...func(*Options)) (*UpdateResourceShareOutput, error) {
	stack := middleware.NewStack("UpdateResourceShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateResourceShareMiddlewares(stack)
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
	addOpUpdateResourceShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceShare(options.Region), middleware.Before)
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
			OperationName: "UpdateResourceShare",
			Err:           err,
		}
	}
	out := result.(*UpdateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceShareInput struct {

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// This member is required.
	ResourceShareArn *string

	// Indicates whether principals outside your AWS organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The name of the resource share.
	Name *string
}

type UpdateResourceShareOutput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Information about the resource share.
	ResourceShare *types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateResourceShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourceShare{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourceShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateResourceShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "UpdateResourceShare",
	}
}
