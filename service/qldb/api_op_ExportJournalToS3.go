// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Exports journal contents within a date and time range from a ledger into a
// specified Amazon Simple Storage Service (Amazon S3) bucket. The data is written
// as files in Amazon Ion format. If the ledger with the given Name doesn't exist,
// then throws ResourceNotFoundException. If the ledger with the given Name is in
// CREATING status, then throws ResourcePreconditionNotMetException. You can
// initiate up to two concurrent journal export requests for each ledger. Beyond
// this limit, journal export requests throw LimitExceededException.
func (c *Client) ExportJournalToS3(ctx context.Context, params *ExportJournalToS3Input, optFns ...func(*Options)) (*ExportJournalToS3Output, error) {
	stack := middleware.NewStack("ExportJournalToS3", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpExportJournalToS3Middlewares(stack)
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
	addOpExportJournalToS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportJournalToS3(options.Region), middleware.Before)
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
			OperationName: "ExportJournalToS3",
			Err:           err,
		}
	}
	out := result.(*ExportJournalToS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportJournalToS3Input struct {
	// The inclusive start date and time for the range of journal contents that you
	// want to export. The InclusiveStartTime must be in ISO 8601 date and time format
	// and in Universal Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z The
	// InclusiveStartTime must be before ExclusiveEndTime. If you provide an
	// InclusiveStartTime that is before the ledger's CreationDateTime, Amazon QLDB
	// defaults it to the ledger's CreationDateTime.
	InclusiveStartTime *time.Time
	// The configuration settings of the Amazon S3 bucket destination for your export
	// request.
	S3ExportConfiguration *types.S3ExportConfiguration
	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal export job to do the following:
	//
	//     * Write objects into your Amazon
	// Simple Storage Service (Amazon S3) bucket.
	//
	//     * (Optional) Use your customer
	// master key (CMK) in AWS Key Management Service (AWS KMS) for server-side
	// encryption of your exported data.
	RoleArn *string
	// The name of the ledger.
	Name *string
	// The exclusive end date and time for the range of journal contents that you want
	// to export. The ExclusiveEndTime must be in ISO 8601 date and time format and in
	// Universal Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z The
	// ExclusiveEndTime must be less than or equal to the current UTC date and time.
	ExclusiveEndTime *time.Time
}

type ExportJournalToS3Output struct {
	// The unique ID that QLDB assigns to each journal export job. To describe your
	// export request and check the status of the job, you can use ExportId to call
	// DescribeJournalS3Export.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpExportJournalToS3Middlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpExportJournalToS3{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpExportJournalToS3{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportJournalToS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ExportJournalToS3",
	}
}
