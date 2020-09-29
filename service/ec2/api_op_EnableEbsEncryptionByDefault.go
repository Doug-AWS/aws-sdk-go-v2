// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables EBS encryption by default for your account in the current Region. After
// you enable encryption by default, the EBS volumes that you create are are always
// encrypted, either using the default CMK or the CMK that you specified when you
// created each volume. For more information, see Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. You can specify the default CMK for
// encryption by default using ModifyEbsDefaultKmsKeyId () or
// ResetEbsDefaultKmsKeyId (). Enabling encryption by default has no effect on the
// encryption status of your existing volumes. After you enable encryption by
// default, you can no longer launch instances using instance types that do not
// support encryption. For more information, see Supported Instance Types
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
func (c *Client) EnableEbsEncryptionByDefault(ctx context.Context, params *EnableEbsEncryptionByDefaultInput, optFns ...func(*Options)) (*EnableEbsEncryptionByDefaultOutput, error) {
	stack := middleware.NewStack("EnableEbsEncryptionByDefault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpEnableEbsEncryptionByDefaultMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(options.Region), middleware.Before)
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
			OperationName: "EnableEbsEncryptionByDefault",
			Err:           err,
		}
	}
	out := result.(*EnableEbsEncryptionByDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableEbsEncryptionByDefaultInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type EnableEbsEncryptionByDefaultOutput struct {
	// The updated status of encryption by default.
	EbsEncryptionByDefault *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpEnableEbsEncryptionByDefaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpEnableEbsEncryptionByDefault{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpEnableEbsEncryptionByDefault{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableEbsEncryptionByDefault",
	}
}
