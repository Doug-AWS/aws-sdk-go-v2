// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modify a setting for an Amazon Aurora DB cluster. You can change one or more
// database configuration parameters by specifying these parameters and the new
// values in the request. For more information on Amazon Aurora, see  What Is
// Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	stack := middleware.NewStack("ModifyDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBClusterMiddlewares(stack)
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
	addOpModifyDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before)
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
			OperationName: "ModifyDBCluster",
			Err:           err,
		}
	}
	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBClusterInput struct {
	// The new DB cluster identifier for the DB cluster when renaming a DB cluster.
	// This value is stored as a lowercase string. Constraints:
	//
	//     * Must contain
	// from 1 to 63 letters, numbers, or hyphens
	//
	//     * The first character must be a
	// letter
	//
	//     * Can't end with a hyphen or contain two consecutive
	// hyphens
	//
	// Example: my-cluster2
	NewDBClusterIdentifier *string
	// A value that indicates whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB cluster. If this parameter is
	// disabled, changes to the DB cluster are applied during the next maintenance
	// window. The ApplyImmediately parameter only affects the
	// EnableIAMDatabaseAuthentication, MasterUserPassword, and NewDBClusterIdentifier
	// values. If the ApplyImmediately parameter is disabled, then changes to the
	// EnableIAMDatabaseAuthentication, MasterUserPassword, and NewDBClusterIdentifier
	// values are applied during the next maintenance window. All other changes are
	// applied immediately, regardless of the value of the ApplyImmediately parameter.
	// By default, this parameter is disabled.
	ApplyImmediately *bool
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	// When you apply a parameter group using the DBInstanceParameterGroupName
	// parameter, the DB cluster isn't rebooted automatically. Also, parameter changes
	// aren't applied during the next maintenance window but instead are applied
	// immediately. Default: The existing name setting Constraints:
	//
	//     * The DB
	// parameter group must be in the same DB parameter group family as this DB
	// cluster.
	//
	//     * The DBInstanceParameterGroupName parameter is only valid in
	// combination with the AllowMajorVersionUpgrade parameter.
	DBInstanceParameterGroupName *string
	// The version number of the database engine to which you want to upgrade. Changing
	// this parameter results in an outage. The change is applied during the next
	// maintenance window unless ApplyImmediately is enabled. To list all of the
	// available engine versions for aurora (for MySQL 5.6-compatible Aurora), use the
	// following command: aws rds describe-db-engine-versions --engine aurora --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-mysql (for MySQL 5.7-compatible Aurora), use the following command:
	// aws rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-postgresql, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion"
	EngineVersion *string
	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. To see the time blocks available, see  Adjusting the Preferred DB
	// Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Constraints:
	//
	//     * Must be in the format
	// hh24:mi-hh24:mi.
	//
	//     * Must be in Universal Coordinated Time (UTC).
	//
	//     * Must
	// not conflict with the preferred maintenance window.
	//
	//     * Must be at least 30
	// minutes.
	PreferredBackupWindow *string
	// The name of the DB cluster parameter group to use for the DB cluster.
	DBClusterParameterGroupName *string
	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. To see the time blocks available,
	// see  Adjusting the Preferred DB Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string
	// The scaling properties of the DB cluster. You can only modify scaling properties
	// for DB clusters in serverless DB engine mode.
	ScalingConfiguration *types.ScalingConfiguration
	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool
	// The new password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters.
	MasterUserPassword *string
	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//
	//     * Must be a value from 1 to 35
	BackupRetentionPeriod *int32
	// The Active Directory directory ID to move the DB cluster to. Specify none to
	// remove the cluster from its current domain. The domain must be created prior to
	// this operation.
	Domain *string
	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//
	//     * If specified, this value must be set to a number
	// from 0 to 259,200 (72 hours).
	BacktrackWindow *int64
	// A value that indicates whether major version upgrades are allowed. Constraints:
	// You must allow major version upgrades when specifying a value for the
	// EngineVersion parameter that is a different major version than the DB cluster's
	// current version.
	AllowMajorVersionUpgrade *bool
	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string
	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// <p>For more information, see <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html">
	// IAM Database Authentication</a> in the <i>Amazon Aurora User Guide.</i> </p>
	EnableIAMDatabaseAuthentication *bool
	// The port number on which the DB cluster accepts connections. Constraints: Value
	// must be 1150-65535 Default: The same port as the original DB cluster.
	Port *int32
	// A value that indicates whether to enable the HTTP endpoint for an Aurora
	// Serverless DB cluster. By default, the HTTP endpoint is disabled. When enabled,
	// the HTTP endpoint provides a connectionless web service API for running SQL
	// queries on the Aurora Serverless DB cluster. You can also query your database
	// from inside the RDS console with the query editor. For more information, see
	// Using the Data API for Aurora Serverless
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in
	// the Amazon Aurora User Guide.
	EnableHttpEndpoint *bool
	// A value that indicates whether to enable write operations to be forwarded from
	// this cluster to the primary cluster in an Aurora global database. The resulting
	// changes are replicated back to this cluster. This parameter only applies to DB
	// clusters that are secondary clusters in an Aurora global database. By default,
	// Aurora disallows write operations for secondary clusters.
	EnableGlobalWriteForwarding *bool
	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB cluster.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration
	// A value that indicates whether to copy all tags from the DB cluster to snapshots
	// of the DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool
	// A value that indicates that the DB cluster should be associated with the
	// specified option group. Changing this parameter doesn't result in an outage
	// except in the following case, and the change is applied during the next
	// maintenance window unless the ApplyImmediately is enabled for this request. If
	// the parameter change results in an option group that enables OEM, this change
	// can cause a brief (sub-second) period during which new connections are rejected
	// but existing connections are not interrupted. Permanent options can't be removed
	// from an option group. The option group can't be removed from a DB cluster once
	// it is associated with a DB cluster.
	OptionGroupName *string
	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive. Constraints: This identifier must match the identifier of an
	// existing DB cluster.
	DBClusterIdentifier *string
	// A list of VPC security groups that the DB cluster will belong to.
	VpcSecurityGroupIds []*string
}

type ModifyDBClusterOutput struct {
	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBCluster",
	}
}
