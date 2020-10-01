// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns table metadata for the specified catalog, database, and table.
func (c *Client) GetTableMetadata(ctx context.Context, params *GetTableMetadataInput, optFns ...func(*Options)) (*GetTableMetadataOutput, error) {
	stack := middleware.NewStack("GetTableMetadata", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTableMetadataMiddlewares(stack)
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
	addOpGetTableMetadataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTableMetadata(options.Region), middleware.Before)
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
			OperationName: "GetTableMetadata",
			Err:           err,
		}
	}
	out := result.(*GetTableMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableMetadataInput struct {

	// The name of the data catalog that contains the database and table metadata to
	// return.
	//
	// This member is required.
	CatalogName *string

	// The name of the database that contains the table metadata to return.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table for which metadata is returned.
	//
	// This member is required.
	TableName *string
}

type GetTableMetadataOutput struct {

	// An object that contains table metadata.
	TableMetadata *types.TableMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTableMetadataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTableMetadata{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTableMetadata{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTableMetadata(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetTableMetadata",
	}
}
