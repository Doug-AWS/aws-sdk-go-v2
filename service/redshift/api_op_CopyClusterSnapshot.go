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

// Copies the specified automated cluster snapshot to a new manual cluster
// snapshot. The source must be an automated snapshot and it must be in the
// available state. When you delete a cluster, Amazon Redshift deletes any
// automated snapshots of the cluster. Also, when the retention period of the
// snapshot expires, Amazon Redshift automatically deletes it. If you want to keep
// an automated snapshot for a longer period, you can make a manual copy of the
// snapshot. Manual snapshots are retained until you delete them. For more
// information about working with snapshots, go to Amazon Redshift Snapshots
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CopyClusterSnapshot(ctx context.Context, params *CopyClusterSnapshotInput, optFns ...func(*Options)) (*CopyClusterSnapshotOutput, error) {
	stack := middleware.NewStack("CopyClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCopyClusterSnapshotMiddlewares(stack)
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
	addOpCopyClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyClusterSnapshot(options.Region), middleware.Before)
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
			OperationName: "CopyClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*CopyClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CopyClusterSnapshotInput struct {

	// The identifier for the source snapshot. Constraints:
	//
	//     * Must be the
	// identifier for a valid automated snapshot whose state is available.
	//
	// This member is required.
	SourceSnapshotIdentifier *string

	// The identifier given to the new manual snapshot. Constraints:
	//
	//     * Cannot be
	// null, empty, or blank.
	//
	//     * Must contain from 1 to 255 alphanumeric characters
	// or hyphens.
	//
	//     * First character must be a letter.
	//
	//     * Cannot end with a
	// hyphen or contain two consecutive hyphens.
	//
	//     * Must be unique for the AWS
	// account that is making the request.
	//
	// This member is required.
	TargetSnapshotIdentifier *string

	// The number of days that a manual snapshot is retained. If the value is -1, the
	// manual snapshot is retained indefinitely. The value must be either -1 or an
	// integer between 1 and 3,653. The default value is -1.
	ManualSnapshotRetentionPeriod *int32

	// The identifier of the cluster the source snapshot was created from. This
	// parameter is required if your IAM user has a policy containing a snapshot
	// resource element that specifies anything other than * for the cluster name.
	// Constraints:
	//
	//     * Must be the identifier for a valid cluster.
	SourceSnapshotClusterIdentifier *string
}

type CopyClusterSnapshotOutput struct {

	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCopyClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCopyClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CopyClusterSnapshot",
	}
}
