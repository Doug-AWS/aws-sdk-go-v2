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

// Creates a new table from a table in an Amazon Redshift cluster snapshot. You
// must create the new table within the Amazon Redshift cluster that the snapshot
// was taken from. You cannot use RestoreTableFromClusterSnapshot to restore a
// table with the same name as an existing table in an Amazon Redshift cluster.
// That is, you cannot overwrite an existing table in a cluster with a restored
// table. If you want to replace your original table with a new, restored table,
// then rename or drop your original table before you call
// RestoreTableFromClusterSnapshot. When you have renamed your original table, then
// you can pass the original name of the table as the NewTableName parameter value
// in the call to RestoreTableFromClusterSnapshot. This way, you can replace the
// original table with the table created from the snapshot.
func (c *Client) RestoreTableFromClusterSnapshot(ctx context.Context, params *RestoreTableFromClusterSnapshotInput, optFns ...func(*Options)) (*RestoreTableFromClusterSnapshotOutput, error) {
	stack := middleware.NewStack("RestoreTableFromClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRestoreTableFromClusterSnapshotMiddlewares(stack)
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
	addOpRestoreTableFromClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableFromClusterSnapshot(options.Region), middleware.Before)
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
			OperationName: "RestoreTableFromClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*RestoreTableFromClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreTableFromClusterSnapshotInput struct {

	// The identifier of the Amazon Redshift cluster to restore the table to.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the table to create as a result of the current request.
	//
	// This member is required.
	NewTableName *string

	// The identifier of the snapshot to restore the table from. This snapshot must
	// have been created from the Amazon Redshift cluster specified by the
	// ClusterIdentifier parameter.
	//
	// This member is required.
	SnapshotIdentifier *string

	// The name of the source database that contains the table to restore from.
	//
	// This member is required.
	SourceDatabaseName *string

	// The name of the source table to restore from.
	//
	// This member is required.
	SourceTableName *string

	// The name of the source schema that contains the table to restore from. If you do
	// not specify a SourceSchemaName value, the default is public.
	SourceSchemaName *string

	// The name of the database to restore the table to.
	TargetDatabaseName *string

	// The name of the schema to restore the table to.
	TargetSchemaName *string
}

type RestoreTableFromClusterSnapshotOutput struct {

	// Describes the status of a RestoreTableFromClusterSnapshot () operation.
	TableRestoreStatus *types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRestoreTableFromClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRestoreTableFromClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreTableFromClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreTableFromClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RestoreTableFromClusterSnapshot",
	}
}
