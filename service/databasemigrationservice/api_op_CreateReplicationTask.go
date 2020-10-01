// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a replication task using the specified parameters.
func (c *Client) CreateReplicationTask(ctx context.Context, params *CreateReplicationTaskInput, optFns ...func(*Options)) (*CreateReplicationTaskOutput, error) {
	stack := middleware.NewStack("CreateReplicationTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateReplicationTaskMiddlewares(stack)
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
	addOpCreateReplicationTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationTask(options.Region), middleware.Before)
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
			OperationName: "CreateReplicationTask",
			Err:           err,
		}
	}
	out := result.(*CreateReplicationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateReplicationTaskInput struct {

	// The migration type. Valid values: full-load | cdc | full-load-and-cdc
	//
	// This member is required.
	MigrationType types.MigrationTypeValue

	// The Amazon Resource Name (ARN) of a replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// An identifier for the replication task. Constraints:
	//
	//     * Must contain 1-255
	// alphanumeric characters or hyphens.
	//
	//     * First character must be a letter.
	//
	//
	// * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	ReplicationTaskIdentifier *string

	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	//
	// This member is required.
	SourceEndpointArn *string

	// The table mappings for the task, in JSON format. For more information, see Using
	// Table Mapping to Specify Task Settings
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	// in the AWS Database Migration Service User Guide.
	//
	// This member is required.
	TableMappings *string

	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	//
	// This member is required.
	TargetEndpointArn *string

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error. The value can be in date,
	// checkpoint, or LSN/SCN format. Date Example: --cdc-start-position
	// “2018-03-08T12:12:12” Checkpoint Example: --cdc-start-position
	// "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373” When you use
	// this task setting with a source PostgreSQL database, a logical replication slot
	// should already be created and associated with the source endpoint. You can
	// verify this by setting the slotName extra connection attribute to the name of
	// this logical replication slot. For more information, see Extra Connection
	// Attributes When Using PostgreSQL as a Source for AWS DMS
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib).
	CdcStartPosition *string

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation to
	// start. Specifying both values results in an error. Timestamp Example:
	// --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time

	// Indicates when you want a change data capture (CDC) operation to stop. The value
	// can be either server time or commit time. Server time example:
	// --cdc-stop-position “server_time:3018-02-09T12:12:12” Commit time example:
	// --cdc-stop-position “commit_time: 3018-02-09T12:12:12 “
	CdcStopPosition *string

	// Overall settings for the task, in JSON format. For more information, see
	// Specifying Task Settings for AWS Database Migration Service Tasks
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html)
	// in the AWS Database Migration User Guide.
	ReplicationTaskSettings *string

	// One or more tags to be assigned to the replication task.
	Tags []*types.Tag

	// Supplemental information that the task requires to migrate the data for certain
	// source and target endpoints. For more information, see Specifying Supplemental
	// Data for Task Settings
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html) in
	// the AWS Database Migration Service User Guide.
	TaskData *string
}

//
type CreateReplicationTaskOutput struct {

	// The replication task that was created.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateReplicationTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateReplicationTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateReplicationTask",
	}
}
