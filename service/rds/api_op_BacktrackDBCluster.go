// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Backtracks a DB cluster to a specific time, without creating a new DB cluster.
// For more information on backtracking, see  Backtracking an Aurora DB Cluster
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora MySQL DB
// clusters.
func (c *Client) BacktrackDBCluster(ctx context.Context, params *BacktrackDBClusterInput, optFns ...func(*Options)) (*BacktrackDBClusterOutput, error) {
	stack := middleware.NewStack("BacktrackDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpBacktrackDBClusterMiddlewares(stack)
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
	addOpBacktrackDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBacktrackDBCluster(options.Region), middleware.Before)
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
			OperationName: "BacktrackDBCluster",
			Err:           err,
		}
	}
	out := result.(*BacktrackDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type BacktrackDBClusterInput struct {

	// The timestamp of the time to backtrack the DB cluster to, specified in ISO 8601
	// format. For more information about ISO 8601, see the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) If the specified time isn't a consistent
	// time for the DB cluster, Aurora automatically chooses the nearest possible
	// consistent time for the DB cluster. Constraints:
	//
	//     * Must contain a valid ISO
	// 8601 timestamp.
	//
	//     * Can't contain a timestamp set in the future.
	//
	// Example:
	// 2017-07-08T18:00Z
	//
	// This member is required.
	BacktrackTo *time.Time

	// The DB cluster identifier of the DB cluster to be backtracked. This parameter is
	// stored as a lowercase string. Constraints:
	//
	//     * Must contain from 1 to 63
	// alphanumeric characters or hyphens.
	//
	//     * First character must be a letter.
	//
	//
	// * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example:
	// my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// A value that indicates whether to force the DB cluster to backtrack when binary
	// logging is enabled. Otherwise, an error occurs when binary logging is enabled.
	Force *bool

	// A value that indicates whether to backtrack the DB cluster to the earliest
	// possible backtrack time when BacktrackTo is set to a timestamp earlier than the
	// earliest backtrack time. When this parameter is disabled and BacktrackTo is set
	// to a timestamp earlier than the earliest backtrack time, an error occurs.
	UseEarliestTimeOnPointInTimeUnavailable *bool
}

// This data type is used as a response element in the DescribeDBClusterBacktracks
// action.
type BacktrackDBClusterOutput struct {

	// Contains the backtrack identifier.
	BacktrackIdentifier *string

	// The timestamp of the time at which the backtrack was requested.
	BacktrackRequestCreationTime *time.Time

	// The timestamp of the time to which the DB cluster was backtracked.
	BacktrackTo *time.Time

	// The timestamp of the time from which the DB cluster was backtracked.
	BacktrackedFrom *time.Time

	// Contains a user-supplied DB cluster identifier. This identifier is the unique
	// key that identifies a DB cluster.
	DBClusterIdentifier *string

	// The status of the backtrack. This property returns one of the following
	// values:
	//
	//     * applying - The backtrack is currently being applied to or rolled
	// back from the DB cluster.
	//
	//     * completed - The backtrack has successfully been
	// applied to or rolled back from the DB cluster.
	//
	//     * failed - An error occurred
	// while the backtrack was applied to or rolled back from the DB cluster.
	//
	//     *
	// pending - The backtrack is currently pending application to or rollback from the
	// DB cluster.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpBacktrackDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpBacktrackDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpBacktrackDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opBacktrackDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "BacktrackDBCluster",
	}
}
