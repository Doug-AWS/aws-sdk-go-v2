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
)

// Lists all global tables that have a replica in the specified Region. This
// operation only applies to Version 2017.11.29
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.
func (c *Client) ListGlobalTables(ctx context.Context, params *ListGlobalTablesInput, optFns ...func(*Options)) (*ListGlobalTablesOutput, error) {
	stack := middleware.NewStack("ListGlobalTables", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListGlobalTablesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGlobalTables(options.Region), middleware.Before)
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
			OperationName: "ListGlobalTables",
			Err:           err,
		}
	}
	out := result.(*ListGlobalTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGlobalTablesInput struct {
	// The first global table name that this operation will evaluate.
	ExclusiveStartGlobalTableName *string
	// The maximum number of table names to return, if the parameter is not specified
	// DynamoDB defaults to 100. If the number of global tables DynamoDB finds reaches
	// this limit, it stops the operation and returns the table names collected up to
	// that point, with a table name in the LastEvaluatedGlobalTableName to apply in a
	// subsequent operation to the ExclusiveStartGlobalTableName parameter.
	Limit *int32
	// Lists the global tables in a specific Region.
	RegionName *string
}

type ListGlobalTablesOutput struct {
	// List of global table names.
	GlobalTables []*types.GlobalTable
	// Last evaluated global table name.
	LastEvaluatedGlobalTableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListGlobalTablesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListGlobalTables{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListGlobalTables{}, middleware.After)
}

func newServiceMetadataMiddleware_opListGlobalTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListGlobalTables",
	}
}
