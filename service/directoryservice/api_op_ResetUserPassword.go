// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets the password for any user in your AWS Managed Microsoft AD or Simple AD
// directory. You can reset the password for any user in your directory with the
// following exceptions:
//
//     * For Simple AD, you cannot reset the password for
// any user that is a member of either the Domain Admins or Enterprise Admins group
// except for the administrator user.
//
//     * For AWS Managed Microsoft AD, you can
// only reset the password for a user that is in an OU based off of the NetBIOS
// name that you typed when you created your directory. For example, you cannot
// reset the password for a user in the AWS Reserved OU. For more information about
// the OU structure for an AWS Managed Microsoft AD directory, see What Gets
// Created
// (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started_what_gets_created.html)
// in the AWS Directory Service Administration Guide.
func (c *Client) ResetUserPassword(ctx context.Context, params *ResetUserPasswordInput, optFns ...func(*Options)) (*ResetUserPasswordOutput, error) {
	stack := middleware.NewStack("ResetUserPassword", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResetUserPasswordMiddlewares(stack)
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
	addOpResetUserPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetUserPassword(options.Region), middleware.Before)
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
			OperationName: "ResetUserPassword",
			Err:           err,
		}
	}
	out := result.(*ResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetUserPasswordInput struct {

	// Identifier of the AWS Managed Microsoft AD or Simple AD directory in which the
	// user resides.
	//
	// This member is required.
	DirectoryId *string

	// The new password that will be reset.
	//
	// This member is required.
	NewPassword *string

	// The user name of the user whose password will be reset.
	//
	// This member is required.
	UserName *string
}

type ResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResetUserPasswordMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResetUserPassword{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetUserPassword{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetUserPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ResetUserPassword",
	}
}
