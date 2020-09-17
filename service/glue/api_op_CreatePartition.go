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

// Creates a new partition.
func (c *Client) CreatePartition(ctx context.Context, params *CreatePartitionInput, optFns ...func(*Options)) (*CreatePartitionOutput, error) {
	stack := middleware.NewStack("CreatePartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePartitionMiddlewares(stack)
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
	addOpCreatePartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePartition(options.Region), middleware.Before)

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
			OperationName: "CreatePartition",
			Err:           err,
		}
	}
	out := result.(*CreatePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePartitionInput struct {
	// The name of the metadata table in which the partition is to be created.
	TableName *string
	// The name of the metadata database in which the partition is to be created.
	DatabaseName *string
	// The AWS account ID of the catalog in which the partition is to be created.
	CatalogId *string
	// A PartitionInput structure defining the partition to be created.
	PartitionInput *types.PartitionInput
}

type CreatePartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreatePartition",
	}
}