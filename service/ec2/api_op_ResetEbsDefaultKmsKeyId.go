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

// Resets the default customer master key (CMK) for EBS encryption for your account
// in this Region to the AWS managed CMK for EBS. After resetting the default CMK
// to the AWS managed CMK, you can continue to encrypt by a customer managed CMK by
// specifying it when you create the volume. For more information, see Amazon EBS
// Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) ResetEbsDefaultKmsKeyId(ctx context.Context, params *ResetEbsDefaultKmsKeyIdInput, optFns ...func(*Options)) (*ResetEbsDefaultKmsKeyIdOutput, error) {
	stack := middleware.NewStack("ResetEbsDefaultKmsKeyId", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpResetEbsDefaultKmsKeyIdMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetEbsDefaultKmsKeyId(options.Region), middleware.Before)
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
			OperationName: "ResetEbsDefaultKmsKeyId",
			Err:           err,
		}
	}
	out := result.(*ResetEbsDefaultKmsKeyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetEbsDefaultKmsKeyIdInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ResetEbsDefaultKmsKeyIdOutput struct {
	// The Amazon Resource Name (ARN) of the default CMK for EBS encryption by default.
	KmsKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpResetEbsDefaultKmsKeyIdMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpResetEbsDefaultKmsKeyId{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpResetEbsDefaultKmsKeyId{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetEbsDefaultKmsKeyId(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ResetEbsDefaultKmsKeyId",
	}
}
