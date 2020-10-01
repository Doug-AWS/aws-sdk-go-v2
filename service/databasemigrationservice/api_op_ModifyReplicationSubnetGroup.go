// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the settings for the specified replication subnet group.
func (c *Client) ModifyReplicationSubnetGroup(ctx context.Context, params *ModifyReplicationSubnetGroupInput, optFns ...func(*Options)) (*ModifyReplicationSubnetGroupOutput, error) {
	stack := middleware.NewStack("ModifyReplicationSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpModifyReplicationSubnetGroupMiddlewares(stack)
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
	addOpModifyReplicationSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyReplicationSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyReplicationSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyReplicationSubnetGroupInput struct {

	// The name of the replication instance subnet group.
	//
	// This member is required.
	ReplicationSubnetGroupIdentifier *string

	// A list of subnet IDs.
	//
	// This member is required.
	SubnetIds []*string

	// A description for the replication instance subnet group.
	ReplicationSubnetGroupDescription *string
}

//
type ModifyReplicationSubnetGroupOutput struct {

	// The modified replication subnet group.
	ReplicationSubnetGroup *types.ReplicationSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpModifyReplicationSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReplicationSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReplicationSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyReplicationSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyReplicationSubnetGroup",
	}
}
