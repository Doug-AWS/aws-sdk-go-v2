// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a block storage disk to a running or stopped Lightsail instance and
// exposes it to the instance with the specified disk name. The attach disk
// operation supports tag-based access control via resource tags applied to the
// resource identified by disk name. For more information, see the Lightsail Dev
// Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) AttachDisk(ctx context.Context, params *AttachDiskInput, optFns ...func(*Options)) (*AttachDiskOutput, error) {
	stack := middleware.NewStack("AttachDisk", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAttachDiskMiddlewares(stack)
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
	addOpAttachDiskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachDisk(options.Region), middleware.Before)
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
			OperationName: "AttachDisk",
			Err:           err,
		}
	}
	out := result.(*AttachDiskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachDiskInput struct {
	// The disk path to expose to the instance (e.g., /dev/xvdf).
	DiskPath *string
	// The name of the Lightsail instance where you want to utilize the storage disk.
	InstanceName *string
	// The unique Lightsail disk name (e.g., my-disk).
	DiskName *string
}

type AttachDiskOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAttachDiskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAttachDisk{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachDisk{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachDisk(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachDisk",
	}
}
