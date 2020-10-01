// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uploads an SSH public key and associates it with the specified IAM user. The SSH
// public key uploaded by this operation can be used only for authenticating the
// associated IAM user to an AWS CodeCommit repository. For more information about
// using SSH keys to authenticate to an AWS CodeCommit repository, see Set up AWS
// CodeCommit for SSH Connections
// (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
func (c *Client) UploadSSHPublicKey(ctx context.Context, params *UploadSSHPublicKeyInput, optFns ...func(*Options)) (*UploadSSHPublicKeyOutput, error) {
	stack := middleware.NewStack("UploadSSHPublicKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUploadSSHPublicKeyMiddlewares(stack)
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
	addOpUploadSSHPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUploadSSHPublicKey(options.Region), middleware.Before)
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
			OperationName: "UploadSSHPublicKey",
			Err:           err,
		}
	}
	out := result.(*UploadSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadSSHPublicKeyInput struct {

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM
	// format. The minimum bit-length of the public key is 2048 bits. For example, you
	// can generate a 2048-bit key, and the resulting PEM file is 1679 bytes long. The
	// regex pattern (http://wikipedia.org/wiki/regex) used to validate this parameter
	// is a string of characters consisting of the following:
	//
	//     * Any printable
	// ASCII character ranging from the space character (\u0020) through the end of the
	// ASCII character range
	//
	//     * The printable characters in the Basic Latin and
	// Latin-1 Supplement character set (through \u00FF)
	//
	//     * The special characters
	// tab (\u0009), line feed (\u000A), and carriage return (\u000D)
	//
	// This member is required.
	SSHPublicKeyBody *string

	// The name of the IAM user to associate the SSH public key with. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

// Contains the response to a successful UploadSSHPublicKey () request.
type UploadSSHPublicKeyOutput struct {

	// Contains information about the SSH public key.
	SSHPublicKey *types.SSHPublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUploadSSHPublicKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUploadSSHPublicKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUploadSSHPublicKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opUploadSSHPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UploadSSHPublicKey",
	}
}
