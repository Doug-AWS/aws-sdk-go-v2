// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies a user's permissions. For more information, see Security and
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingsecurity.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) SetPermission(ctx context.Context, params *SetPermissionInput, optFns ...func(*Options)) (*SetPermissionOutput, error) {
	stack := middleware.NewStack("SetPermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetPermissionMiddlewares(stack)
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
	addOpSetPermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetPermission(options.Region), middleware.Before)
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
			OperationName: "SetPermission",
			Err:           err,
		}
	}
	out := result.(*SetPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetPermissionInput struct {

	// The user's IAM ARN. This can also be a federated user's ARN.
	//
	// This member is required.
	IamUserArn *string

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The user is allowed to use SSH to communicate with the instance.
	AllowSsh *bool

	// The user is allowed to use sudo to elevate privileges.
	AllowSudo *bool

	// The user's permission level, which must be set to one of the following strings.
	// You cannot set your own permissions level.
	//
	//     * deny
	//
	//     * show
	//
	//     *
	// deploy
	//
	//     * manage
	//
	//     * iam_only
	//
	// For more information about the permissions
	// associated with these levels, see Managing User Permissions
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
	Level *string
}

type SetPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetPermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetPermission{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetPermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetPermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "SetPermission",
	}
}
