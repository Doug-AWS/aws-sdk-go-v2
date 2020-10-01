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

// The DeleteDBInstance action deletes a previously provisioned DB instance. When
// you delete a DB instance, all automated backups for that instance are deleted
// and can't be recovered. Manual DB snapshots of the DB instance to be deleted by
// DeleteDBInstance are not deleted. If you request a final DB snapshot the status
// of the Amazon Neptune DB instance is deleting until the DB snapshot is created.
// The API action DescribeDBInstance is used to monitor the status of this
// operation. The action can't be canceled or reverted once submitted. Note that
// when a DB instance is in a failure state and has a status of failed,
// incompatible-restore, or incompatible-network, you can only delete it when the
// SkipFinalSnapshot parameter is set to true. You can't delete a DB instance if it
// is the only instance in the DB cluster, or if it has deletion protection
// enabled.
func (c *Client) DeleteDBInstance(ctx context.Context, params *DeleteDBInstanceInput, optFns ...func(*Options)) (*DeleteDBInstanceOutput, error) {
	stack := middleware.NewStack("DeleteDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteDBInstanceMiddlewares(stack)
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
	addOpDeleteDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBInstance(options.Region), middleware.Before)
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
			OperationName: "DeleteDBInstance",
			Err:           err,
		}
	}
	out := result.(*DeleteDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBInstanceInput struct {

	// The DB instance identifier for the DB instance to be deleted. This parameter
	// isn't case-sensitive. Constraints:
	//
	//     * Must match the name of an existing DB
	// instance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The DBSnapshotIdentifier of the new DBSnapshot created when SkipFinalSnapshot is
	// set to false. Specifying this parameter and also setting the SkipFinalShapshot
	// parameter to true results in an error. Constraints:
	//
	//     * Must be 1 to 255
	// letters or numbers.
	//
	//     * First character must be a letter
	//
	//     * Cannot end
	// with a hyphen or contain two consecutive hyphens
	//
	//     * Cannot be specified when
	// deleting a Read Replica.
	FinalDBSnapshotIdentifier *string

	// Determines whether a final DB snapshot is created before the DB instance is
	// deleted. If true is specified, no DBSnapshot is created. If false is specified,
	// a DB snapshot is created before the DB instance is deleted. Note that when a DB
	// instance is in a failure state and has a status of 'failed',
	// 'incompatible-restore', or 'incompatible-network', it can only be deleted when
	// the SkipFinalSnapshot parameter is set to "true". Specify true when deleting a
	// Read Replica. The FinalDBSnapshotIdentifier parameter must be specified if
	// SkipFinalSnapshot is false. Default: false
	SkipFinalSnapshot *bool
}

type DeleteDBInstanceOutput struct {

	// Contains the details of an Amazon Neptune DB instance. This data type is used as
	// a response element in the DescribeDBInstances () action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBInstance",
	}
}
