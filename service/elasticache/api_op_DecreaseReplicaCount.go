// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Dynamically decreases the number of replicas in a Redis (cluster mode disabled)
// replication group or the number of replica nodes in one or more node groups
// (shards) of a Redis (cluster mode enabled) replication group. This operation is
// performed with no cluster down time.
func (c *Client) DecreaseReplicaCount(ctx context.Context, params *DecreaseReplicaCountInput, optFns ...func(*Options)) (*DecreaseReplicaCountOutput, error) {
	stack := middleware.NewStack("DecreaseReplicaCount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDecreaseReplicaCountMiddlewares(stack)
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
	addOpDecreaseReplicaCountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseReplicaCount(options.Region), middleware.Before)
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
			OperationName: "DecreaseReplicaCount",
			Err:           err,
		}
	}
	out := result.(*DecreaseReplicaCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseReplicaCountInput struct {

	// If True, the number of replica nodes is decreased immediately.
	// <code>ApplyImmediately=False</code> is not currently supported.</p>
	//
	// This member is required.
	ApplyImmediately *bool

	// The id of the replication group from which you want to remove replica nodes.
	//
	// This member is required.
	ReplicationGroupId *string

	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number of
	// replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the
	// replication group's node groups. The minimum number of replicas in a shard or
	// replication group is:
	//
	//     * Redis (cluster mode disabled)
	//
	//         * If
	// Multi-AZ is enabled: 1
	//
	//         * If Multi-AZ is not enabled: 0
	//
	//     * Redis
	// (cluster mode enabled): 0 (though you will not be able to failover to a replica
	// if your primary node fails)
	NewReplicaCount *int32

	// A list of ConfigureShard objects that can be used to configure each shard in a
	// Redis (cluster mode enabled) replication group. The ConfigureShard has three
	// members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []*types.ConfigureShard

	// A list of the node ids to remove from the replication group or node group
	// (shard).
	ReplicasToRemove []*string
}

type DecreaseReplicaCountOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDecreaseReplicaCountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDecreaseReplicaCount{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDecreaseReplicaCount{}, middleware.After)
}

func newServiceMetadataMiddleware_opDecreaseReplicaCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DecreaseReplicaCount",
	}
}
