// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a reference data source to an existing SQL-based Amazon Kinesis Data
// Analytics application. Kinesis Data Analytics reads reference data (that is, an
// Amazon S3 object) and creates an in-application table within your application.
// In the request, you provide the source (S3 bucket name and object key name),
// name of the in-application table to create, and the necessary mapping
// information that describes how data in an Amazon S3 object maps to columns in
// the resulting in-application table.
func (c *Client) AddApplicationReferenceDataSource(ctx context.Context, params *AddApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*AddApplicationReferenceDataSourceOutput, error) {
	stack := middleware.NewStack("AddApplicationReferenceDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddApplicationReferenceDataSourceMiddlewares(stack)
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
	addOpAddApplicationReferenceDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationReferenceDataSource(options.Region), middleware.Before)
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
			OperationName: "AddApplicationReferenceDataSource",
			Err:           err,
		}
	}
	out := result.(*AddApplicationReferenceDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationReferenceDataSourceInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The version of the application for which you are adding the reference data
	// source. You can use the DescribeApplication () operation to get the current
	// application version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The reference data source can be an object in your Amazon S3 bucket. Kinesis
	// Data Analytics reads the object and copies the data into the in-application
	// table that is created. You provide an S3 bucket, object key name, and the
	// resulting in-application table that is created.
	//
	// This member is required.
	ReferenceDataSource *types.ReferenceDataSource
}

type AddApplicationReferenceDataSourceOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// The updated application version ID. Amazon Kinesis Data Analytics increments
	// this ID when the application is updated.
	ApplicationVersionId *int64

	// Describes reference data sources configured for the application. </p>
	ReferenceDataSourceDescriptions []*types.ReferenceDataSourceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddApplicationReferenceDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationReferenceDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationReferenceDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddApplicationReferenceDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationReferenceDataSource",
	}
}
