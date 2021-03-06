// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of your Timestream databases. Service quotas apply. For more
// information, see Access Management
// (https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html) in
// the Timestream Developer Guide.
func (c *Client) ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatabases", params, optFns, addOperationListDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatabasesInput struct {

	// The total number of items to return in the output. If the total number of items
	// available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	MaxResults *int32

	// The pagination token. To resume pagination, provide the NextToken value as
	// argument of a subsequent API invocation.
	NextToken *string
}

type ListDatabasesOutput struct {

	// A list of database names.
	Databases []*types.Database

	// The pagination token. This parameter is returned when the response is truncated.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatabases(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "ListDatabases",
	}
}
