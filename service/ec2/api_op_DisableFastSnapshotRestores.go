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

// Disables fast snapshot restores for the specified snapshots in the specified
// Availability Zones.
func (c *Client) DisableFastSnapshotRestores(ctx context.Context, params *DisableFastSnapshotRestoresInput, optFns ...func(*Options)) (*DisableFastSnapshotRestoresOutput, error) {
	stack := middleware.NewStack("DisableFastSnapshotRestores", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisableFastSnapshotRestoresMiddlewares(stack)
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
	addOpDisableFastSnapshotRestoresValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableFastSnapshotRestores(options.Region), middleware.Before)
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
			OperationName: "DisableFastSnapshotRestores",
			Err:           err,
		}
	}
	out := result.(*DisableFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableFastSnapshotRestoresInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// One or more Availability Zones. For example, us-east-2a.
	AvailabilityZones []*string
	// The IDs of one or more snapshots. For example, snap-1234567890abcdef0.
	SourceSnapshotIds []*string
}

type DisableFastSnapshotRestoresOutput struct {
	// Information about the snapshots for which fast snapshot restores could not be
	// disabled.
	Unsuccessful []*types.DisableFastSnapshotRestoreErrorItem
	// Information about the snapshots for which fast snapshot restores were
	// successfully disabled.
	Successful []*types.DisableFastSnapshotRestoreSuccessItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisableFastSnapshotRestoresMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisableFastSnapshotRestores{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisableFastSnapshotRestores{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableFastSnapshotRestores(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableFastSnapshotRestores",
	}
}
