// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a volume attribute. By default, all I/O operations for the volume are
// suspended when the data on the volume is determined to be potentially
// inconsistent, to prevent undetectable, latent data corruption. The I/O access to
// the volume can be resumed by first enabling I/O access and then checking the
// data consistency on your volume. You can change the default behavior to resume
// I/O operations. We recommend that you change this only for boot volumes or for
// volumes that are stateless or disposable.
func (c *Client) ModifyVolumeAttribute(ctx context.Context, params *ModifyVolumeAttributeInput, optFns ...func(*Options)) (*ModifyVolumeAttributeOutput, error) {
	stack := middleware.NewStack("ModifyVolumeAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyVolumeAttributeMiddlewares(stack)
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
	addOpModifyVolumeAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVolumeAttribute(options.Region), middleware.Before)
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
			OperationName: "ModifyVolumeAttribute",
			Err:           err,
		}
	}
	out := result.(*ModifyVolumeAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVolumeAttributeInput struct {

	// The ID of the volume.
	//
	// This member is required.
	VolumeId *string

	// Indicates whether the volume should be auto-enabled for I/O operations.
	AutoEnableIO *types.AttributeBooleanValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyVolumeAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyVolumeAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyVolumeAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVolumeAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyVolumeAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVolumeAttribute",
	}
}
