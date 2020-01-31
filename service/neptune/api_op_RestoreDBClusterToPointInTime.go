// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreDBClusterToPointInTimeInput struct {
	_ struct{} `type:"structure"`

	// The name of the new DB cluster to be created.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with the new DB cluster.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string `type:"string"`

	// The DB subnet group name to use for the new DB cluster.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// The list of logs that the restored DB cluster is to export to CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The AWS KMS key identifier to use when restoring an encrypted DB cluster
	// from an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are restoring a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// You can restore to a new DB cluster and encrypt the new DB cluster with a
	// KMS key that is different than the KMS key used to encrypt the source DB
	// cluster. The new DB cluster is encrypted with the KMS key identified by the
	// KmsKeyId parameter.
	//
	// If you do not specify a value for the KmsKeyId parameter, then the following
	// will occur:
	//
	//    * If the DB cluster is encrypted, then the restored DB cluster is encrypted
	//    using the KMS key that was used to encrypt the source DB cluster.
	//
	//    * If the DB cluster is not encrypted, then the restored DB cluster is
	//    not encrypted.
	//
	// If DBClusterIdentifier refers to a DB cluster that is not encrypted, then
	// the restore request is rejected.
	KmsKeyId *string `type:"string"`

	// The name of the option group for the new DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the new DB cluster accepts connections.
	//
	// Constraints: Value must be 1150-65535
	//
	// Default: The same port as the original DB cluster.
	Port *int64 `type:"integer"`

	// The date and time to restore the DB cluster to.
	//
	// Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the DB instance
	//
	//    * Must be specified if UseLatestRestorableTime parameter is not provided
	//
	//    * Cannot be specified if UseLatestRestorableTime parameter is true
	//
	//    * Cannot be specified if RestoreType parameter is copy-on-write
	//
	// Example: 2015-03-07T23:45:00Z
	RestoreToTime *time.Time `type:"timestamp"`

	// The type of restore to be performed. You can specify one of the following
	// values:
	//
	//    * full-copy - The new DB cluster is restored as a full copy of the source
	//    DB cluster.
	//
	//    * copy-on-write - The new DB cluster is restored as a clone of the source
	//    DB cluster.
	//
	// If you don't specify a RestoreType value, then the new DB cluster is restored
	// as a full copy of the source DB cluster.
	RestoreType *string `type:"string"`

	// The identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// SourceDBClusterIdentifier is a required field
	SourceDBClusterIdentifier *string `type:"string" required:"true"`

	// The tags to be applied to the restored DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A value that is set to true to restore the DB cluster to the latest restorable
	// backup time, and false otherwise.
	//
	// Default: false
	//
	// Constraints: Cannot be specified if RestoreToTime parameter is provided.
	UseLatestRestorableTime *bool `type:"boolean"`

	// A list of VPC security groups that the new DB cluster belongs to.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBClusterToPointInTimeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBClusterToPointInTimeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBClusterToPointInTimeInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.SourceDBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreDBClusterToPointInTimeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s RestoreDBClusterToPointInTimeOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBClusterToPointInTime = "RestoreDBClusterToPointInTime"

// RestoreDBClusterToPointInTimeRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Restores a DB cluster to an arbitrary point in time. Users can restore to
// any point in time before LatestRestorableTime for up to BackupRetentionPeriod
// days. The target DB cluster is created from the source DB cluster with the
// same configuration as the original DB cluster, except that the new DB cluster
// is created with the default DB security group.
//
// This action only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance action to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier. You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available.
//
//    // Example sending a request using RestoreDBClusterToPointInTimeRequest.
//    req := client.RestoreDBClusterToPointInTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RestoreDBClusterToPointInTime
func (c *Client) RestoreDBClusterToPointInTimeRequest(input *RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterToPointInTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBClusterToPointInTimeInput{}
	}

	req := c.newRequest(op, input, &RestoreDBClusterToPointInTimeOutput{})
	return RestoreDBClusterToPointInTimeRequest{Request: req, Input: input, Copy: c.RestoreDBClusterToPointInTimeRequest}
}

// RestoreDBClusterToPointInTimeRequest is the request type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeRequest struct {
	*aws.Request
	Input *RestoreDBClusterToPointInTimeInput
	Copy  func(*RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest
}

// Send marshals and sends the RestoreDBClusterToPointInTime API request.
func (r RestoreDBClusterToPointInTimeRequest) Send(ctx context.Context) (*RestoreDBClusterToPointInTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterToPointInTimeResponse{
		RestoreDBClusterToPointInTimeOutput: r.Request.Data.(*RestoreDBClusterToPointInTimeOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterToPointInTimeResponse is the response type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeResponse struct {
	*RestoreDBClusterToPointInTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterToPointInTime request.
func (r *RestoreDBClusterToPointInTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
