// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new DB cluster from a DB snapshot or DB cluster snapshot. If a DB
// snapshot is specified, the target DB cluster is created from the source DB
// snapshot with a default configuration and default security group. If a DB
// cluster snapshot is specified, the target DB cluster is created from the source
// DB cluster restore point with the same configuration as the original source DB
// cluster, except that the new DB cluster is created with the default security
// group.
func (c *Client) RestoreDBClusterFromSnapshot(ctx context.Context, params *RestoreDBClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreDBClusterFromSnapshotOutput, error) {
	stack := middleware.NewStack("RestoreDBClusterFromSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRestoreDBClusterFromSnapshotMiddlewares(stack)
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
	addOpRestoreDBClusterFromSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(options.Region), middleware.Before)
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
			OperationName: "RestoreDBClusterFromSnapshot",
			Err:           err,
		}
	}
	out := result.(*RestoreDBClusterFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterFromSnapshotInput struct {
	// The tags to be assigned to the restored DB cluster.
	Tags []*types.Tag
	// The name of the DB cluster parameter group to associate with the new DB cluster.
	// Constraints:
	//
	//     * If supplied, must match the name of an existing
	// DBClusterParameterGroup.
	DBClusterParameterGroupName *string
	// A list of VPC security groups that the new DB cluster will belong to.
	VpcSecurityGroupIds []*string
	// (Not supported by Neptune)
	OptionGroupName *string
	// Not supported.
	DatabaseName *string
	// The name of the DB subnet group to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetgroup
	DBSubnetGroupName *string
	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool
	// The name of the DB cluster to create from the DB snapshot or DB cluster
	// snapshot. This parameter isn't case-sensitive. Constraints:
	//
	//     * Must contain
	// from 1 to 63 letters, numbers, or hyphens
	//
	//     * First character must be a
	// letter
	//
	//     * Cannot end with a hyphen or contain two consecutive
	// hyphens
	//
	// Example: my-snapshot-id
	DBClusterIdentifier *string
	// The version of the database engine to use for the new DB cluster.
	EngineVersion *string
	// The identifier for the DB snapshot or DB cluster snapshot to restore from. You
	// can use either the name or the Amazon Resource Name (ARN) to specify a DB
	// cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	// Constraints:
	//
	//     * Must match the identifier of an existing Snapshot.
	SnapshotIdentifier *string
	// The port number on which the new DB cluster accepts connections. Constraints:
	// Value must be 1150-65535 Default: The same port as the original DB cluster.
	Port *int32
	// The AWS KMS key identifier to use when restoring an encrypted DB cluster from a
	// DB snapshot or DB cluster snapshot. The KMS key identifier is the Amazon
	// Resource Name (ARN) for the KMS encryption key. If you are restoring a DB
	// cluster with the same AWS account that owns the KMS encryption key used to
	// encrypt the new DB cluster, then you can use the KMS key alias instead of the
	// ARN for the KMS encryption key. If you do not specify a value for the KmsKeyId
	// parameter, then the following will occur:
	//
	//     * If the DB snapshot or DB
	// cluster snapshot in SnapshotIdentifier is encrypted, then the restored DB
	// cluster is encrypted using the KMS key that was used to encrypt the DB snapshot
	// or DB cluster snapshot.
	//
	//     * If the DB snapshot or DB cluster snapshot in
	// SnapshotIdentifier is not encrypted, then the restored DB cluster is not
	// encrypted.
	KmsKeyId *string
	// Provides the list of EC2 Availability Zones that instances in the restored DB
	// cluster can be created in.
	AvailabilityZones []*string
	// True to enable mapping of AWS Identity and Access Management (IAM) accounts to
	// database accounts, and otherwise false. Default: false
	EnableIAMDatabaseAuthentication *bool
	// The database engine to use for the new DB cluster. Default: The same as source
	// Constraint: Must be compatible with the engine of the source
	Engine *string
	// The list of logs that the restored DB cluster is to export to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []*string
}

type RestoreDBClusterFromSnapshotOutput struct {
	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters () action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRestoreDBClusterFromSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterFromSnapshot",
	}
}
