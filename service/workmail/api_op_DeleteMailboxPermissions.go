// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes permissions granted to a member (user or group).
func (c *Client) DeleteMailboxPermissions(ctx context.Context, params *DeleteMailboxPermissionsInput, optFns ...func(*Options)) (*DeleteMailboxPermissionsOutput, error) {
	stack := middleware.NewStack("DeleteMailboxPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteMailboxPermissionsMiddlewares(stack)
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
	addOpDeleteMailboxPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMailboxPermissions(options.Region), middleware.Before)
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
			OperationName: "DeleteMailboxPermissions",
			Err:           err,
		}
	}
	out := result.(*DeleteMailboxPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMailboxPermissionsInput struct {

	// The identifier of the member (user or group)that owns the mailbox.
	//
	// This member is required.
	EntityId *string

	// The identifier of the member (user or group) for which to delete granted
	// permissions.
	//
	// This member is required.
	GranteeId *string

	// The identifier of the organization under which the member (user or group)
	// exists.
	//
	// This member is required.
	OrganizationId *string
}

type DeleteMailboxPermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteMailboxPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMailboxPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMailboxPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMailboxPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DeleteMailboxPermissions",
	}
}
