// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// List backups associated with an AWS account. To list backups for a given table,
// specify TableName. ListBackups returns a paginated list of results with at most
// 1 MB worth of items in a page. You can also specify a limit for the maximum
// number of entries to be returned in a page. In the request, start time is
// inclusive, but end time is exclusive. Note that these limits are for the time at
// which the original backup was requested. You can call ListBackups a maximum of
// five times per second.
func (c *Client) ListBackups(ctx context.Context, params *ListBackupsInput, optFns ...func(*Options)) (*ListBackupsOutput, error) {
	stack := middleware.NewStack("ListBackups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListBackupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBackups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "ListBackups",
			Err:           err,
		}
	}
	out := result.(*ListBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupsInput struct {

	// The backups from the table specified by BackupType are listed. Where BackupType
	// can be:
	//
	//     * USER - On-demand backup created by you.
	//
	//     * SYSTEM - On-demand
	// backup automatically created by DynamoDB.
	//
	//     * ALL - All types of on-demand
	// backups (USER and SYSTEM).
	BackupType types.BackupTypeFilter

	// LastEvaluatedBackupArn is the Amazon Resource Name (ARN) of the backup last
	// evaluated when the current page of results was returned, inclusive of the
	// current page of results. This value may be specified as the
	// ExclusiveStartBackupArn of a new ListBackups operation in order to fetch the
	// next page of results.
	ExclusiveStartBackupArn *string

	// Maximum number of backups to return at once.
	Limit *int32

	// The backups from the table specified by TableName are listed.
	TableName *string

	// Only backups created after this time are listed. TimeRangeLowerBound is
	// inclusive.
	TimeRangeLowerBound *time.Time

	// Only backups created before this time are listed. TimeRangeUpperBound is
	// exclusive.
	TimeRangeUpperBound *time.Time
}

type ListBackupsOutput struct {

	// List of BackupSummary objects.
	BackupSummaries []*types.BackupSummary

	// The ARN of the backup last evaluated when the current page of results was
	// returned, inclusive of the current page of results. This value may be specified
	// as the ExclusiveStartBackupArn of a new ListBackups operation in order to fetch
	// the next page of results. If LastEvaluatedBackupArn is empty, then the last page
	// of results has been processed and there are no more results to be retrieved. If
	// LastEvaluatedBackupArn is not empty, this may or may not indicate that there is
	// more data to be returned. All results are guaranteed to have been returned if
	// and only if no value for LastEvaluatedBackupArn is returned.
	LastEvaluatedBackupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListBackupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListBackups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListBackups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBackups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListBackups",
	}
}
