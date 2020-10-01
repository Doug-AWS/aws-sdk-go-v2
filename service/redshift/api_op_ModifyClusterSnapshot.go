// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the settings for a snapshot. This exanmple modifies the manual
// retention period setting for a cluster snapshot.
func (c *Client) ModifyClusterSnapshot(ctx context.Context, params *ModifyClusterSnapshotInput, optFns ...func(*Options)) (*ModifyClusterSnapshotOutput, error) {
	stack := middleware.NewStack("ModifyClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyClusterSnapshotMiddlewares(stack)
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
	addOpModifyClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSnapshot(options.Region), middleware.Before)
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
			OperationName: "ModifyClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*ModifyClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterSnapshotInput struct {

	// The identifier of the snapshot whose setting you want to modify.
	//
	// This member is required.
	SnapshotIdentifier *string

	// A Boolean option to override an exception if the retention period has already
	// passed.
	Force *bool

	// The number of days that a manual snapshot is retained. If the value is -1, the
	// manual snapshot is retained indefinitely. If the manual snapshot falls outside
	// of the new retention period, you can specify the force option to immediately
	// delete the snapshot. The value must be either -1 or an integer between 1 and
	// 3,653.
	ManualSnapshotRetentionPeriod *int32
}

type ModifyClusterSnapshotOutput struct {

	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSnapshot",
	}
}
