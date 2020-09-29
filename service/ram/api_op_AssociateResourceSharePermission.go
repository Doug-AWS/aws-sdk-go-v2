// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a permission with a resource share.
func (c *Client) AssociateResourceSharePermission(ctx context.Context, params *AssociateResourceSharePermissionInput, optFns ...func(*Options)) (*AssociateResourceSharePermissionOutput, error) {
	stack := middleware.NewStack("AssociateResourceSharePermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateResourceSharePermissionMiddlewares(stack)
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
	addOpAssociateResourceSharePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResourceSharePermission(options.Region), middleware.Before)
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
			OperationName: "AssociateResourceSharePermission",
			Err:           err,
		}
	}
	out := result.(*AssociateResourceSharePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResourceSharePermissionInput struct {
	// The ARN of the AWS RAM permission to associate with the resource share.
	PermissionArn *string
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// Indicates whether the permission should replace the permissions that are
	// currently associated with the resource share. Use true to replace the current
	// permissions. Use false to add the permission to the current permission.
	Replace *bool
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string
}

type AssociateResourceSharePermissionOutput struct {
	// Indicates whether the request succeeded.
	ReturnValue *bool
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateResourceSharePermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateResourceSharePermission{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateResourceSharePermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateResourceSharePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "AssociateResourceSharePermission",
	}
}
