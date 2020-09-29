// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a subnet group. The specified database subnet group must not be
// associated with any DB instances.
func (c *Client) DeleteDBSubnetGroup(ctx context.Context, params *DeleteDBSubnetGroupInput, optFns ...func(*Options)) (*DeleteDBSubnetGroupOutput, error) {
	stack := middleware.NewStack("DeleteDBSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteDBSubnetGroupMiddlewares(stack)
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
	addOpDeleteDBSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "DeleteDBSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DeleteDBSubnetGroup ().
type DeleteDBSubnetGroupInput struct {
	// The name of the database subnet group to delete. You can't delete the default
	// subnet group. Constraints: Must match the name of an existing DBSubnetGroup.
	// Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string
}

type DeleteDBSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteDBSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDBSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSubnetGroup",
	}
}
