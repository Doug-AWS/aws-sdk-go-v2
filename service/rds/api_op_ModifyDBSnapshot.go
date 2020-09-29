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

// Updates a manual DB snapshot with a new engine version. The snapshot can be
// encrypted or unencrypted, but not shared or public.  </p> <p>Amazon RDS supports
// upgrading DB snapshots for MySQL, Oracle, and PostgreSQL. </p>
func (c *Client) ModifyDBSnapshot(ctx context.Context, params *ModifyDBSnapshotInput, optFns ...func(*Options)) (*ModifyDBSnapshotOutput, error) {
	stack := middleware.NewStack("ModifyDBSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBSnapshotMiddlewares(stack)
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
	addOpModifyDBSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBSnapshot(options.Region), middleware.Before)
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
			OperationName: "ModifyDBSnapshot",
			Err:           err,
		}
	}
	out := result.(*ModifyDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBSnapshotInput struct {
	// The identifier of the DB snapshot to modify.
	DBSnapshotIdentifier *string
	// The engine version to upgrade the DB snapshot to.  <p>The following are the
	// database engines and engine versions that are available when you upgrade a DB
	// snapshot. </p> <p> <b>MySQL</b> </p> <ul> <li> <p> <code>5.5.46</code>
	// (supported for 5.1 DB snapshots)</p> </li> </ul> <p> <b>Oracle</b> </p> <ul>
	// <li> <p> <code>12.1.0.2.v8</code> (supported for 12.1.0.1 DB snapshots)</p>
	// </li> <li> <p> <code>11.2.0.4.v12</code> (supported for 11.2.0.2 DB
	// snapshots)</p> </li> <li> <p> <code>11.2.0.4.v11</code> (supported for 11.2.0.3
	// DB snapshots)</p> </li> </ul> <p> <b>PostgreSQL</b> </p> <p>For the list of
	// engine versions that are available for upgrading a DB snapshot, see <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.html#USER_UpgradeDBInstance.PostgreSQL.MajorVersion">
	// Upgrading the PostgreSQL DB Engine for Amazon RDS</a>. </p>
	EngineVersion *string
	// The option group to identify with the upgraded DB snapshot.  <p>You can specify
	// this parameter when you upgrade an Oracle DB snapshot. The same option group
	// considerations apply when upgrading a DB snapshot as when upgrading a DB
	// instance. For more information, see <a
	// href="http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Oracle.html#USER_UpgradeDBInstance.Oracle.OGPG.OG">Option
	// Group Considerations</a> in the <i>Amazon RDS User Guide.</i> </p>
	OptionGroupName *string
}

type ModifyDBSnapshotOutput struct {
	// Contains the details of an Amazon RDS DB snapshot. This data type is used as a
	// response element in the DescribeDBSnapshots action.
	DBSnapshot *types.DBSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBSnapshot",
	}
}
