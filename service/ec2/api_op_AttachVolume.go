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
	"time"
)

// Attaches an EBS volume to a running or stopped instance and exposes it to the
// instance with the specified device name. Encrypted EBS volumes must be attached
// to instances that support Amazon EBS encryption. For more information, see
// Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. After you attach an EBS volume, you
// must make it available. For more information, see Making an EBS Volume Available
// For Use
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-using-volumes.html). If
// a volume has an AWS Marketplace product code:
//
//     * The volume can be attached
// only to a stopped instance.
//
//     * AWS Marketplace product codes are copied from
// the volume to the instance.
//
//     * You must be subscribed to the product.
//
//     *
// The instance type and operating system of the instance must support the product.
// For example, you can't detach a volume from a Windows instance and attach it to
// a Linux instance.
//
// For more information, see Attaching Amazon EBS Volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-attaching-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) AttachVolume(ctx context.Context, params *AttachVolumeInput, optFns ...func(*Options)) (*AttachVolumeOutput, error) {
	stack := middleware.NewStack("AttachVolume", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAttachVolumeMiddlewares(stack)
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
	addOpAttachVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachVolume(options.Region), middleware.Before)

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
			OperationName: "AttachVolume",
			Err:           err,
		}
	}
	out := result.(*AttachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachVolumeInput struct {
	// The ID of the EBS volume. The volume and instance must be within the same
	// Availability Zone.
	VolumeId *string
	// The device name (for example, /dev/sdh or xvdh).
	Device *string
	// The ID of the instance.
	InstanceId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Describes volume attachment details.
type AttachVolumeOutput struct {
	// The ID of the volume.
	VolumeId *string
	// The ID of the instance.
	InstanceId *string
	// The time stamp when the attachment initiated.
	AttachTime *time.Time
	// The device name.
	Device *string
	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool
	// The attachment state of the volume.
	State types.VolumeAttachmentState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAttachVolumeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAttachVolume{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAttachVolume{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AttachVolume",
	}
}