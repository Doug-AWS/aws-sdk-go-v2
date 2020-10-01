// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the definitions of some or all of the tables in a given Database.
func (c *Client) GetTables(ctx context.Context, params *GetTablesInput, optFns ...func(*Options)) (*GetTablesOutput, error) {
	stack := middleware.NewStack("GetTables", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTablesMiddlewares(stack)
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
	addOpGetTablesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTables(options.Region), middleware.Before)
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
			OperationName: "GetTables",
			Err:           err,
		}
	}
	out := result.(*GetTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTablesInput struct {

	// The database in the catalog whose tables to list. For Hive compatibility, this
	// name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The ID of the Data Catalog where the tables reside. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string

	// A regular expression pattern. If present, only those tables whose names match
	// the pattern are returned.
	Expression *string

	// The maximum number of tables to return in a single response.
	MaxResults *int32

	// A continuation token, included if this is a continuation call.
	NextToken *string
}

type GetTablesOutput struct {

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// A list of the requested Table objects.
	TableList []*types.Table

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTablesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTables{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTables{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetTables",
	}
}
