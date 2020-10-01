// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the available capacity providers and the default capacity provider
// strategy for a cluster. You must specify both the available capacity providers
// and a default capacity provider strategy for the cluster. If the specified
// cluster has existing capacity providers associated with it, you must specify all
// existing capacity providers in addition to any new ones you want to add. Any
// existing capacity providers associated with a cluster that are omitted from a
// PutClusterCapacityProviders () API call will be disassociated with the cluster.
// You can only disassociate an existing capacity provider from a cluster if it's
// not being used by any existing tasks. When creating a service or running a task
// on a cluster, if no capacity provider or launch type is specified, then the
// cluster's default capacity provider strategy is used. It is recommended to
// define a default capacity provider strategy for your cluster, however you may
// specify an empty array ([]) to bypass defining a default strategy.
func (c *Client) PutClusterCapacityProviders(ctx context.Context, params *PutClusterCapacityProvidersInput, optFns ...func(*Options)) (*PutClusterCapacityProvidersOutput, error) {
	stack := middleware.NewStack("PutClusterCapacityProviders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutClusterCapacityProvidersMiddlewares(stack)
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
	addOpPutClusterCapacityProvidersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutClusterCapacityProviders(options.Region), middleware.Before)
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
			OperationName: "PutClusterCapacityProviders",
			Err:           err,
		}
	}
	out := result.(*PutClusterCapacityProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutClusterCapacityProvidersInput struct {

	// The name of one or more capacity providers to associate with the cluster. If
	// specifying a capacity provider that uses an Auto Scaling group, the capacity
	// provider must already be created. New capacity providers can be created with the
	// CreateCapacityProvider () API operation. To use a AWS Fargate capacity provider,
	// specify either the FARGATE or FARGATE_SPOT capacity providers. The AWS Fargate
	// capacity providers are available to all accounts and only need to be associated
	// with a cluster to be used.
	//
	// This member is required.
	CapacityProviders []*string

	// The short name or full Amazon Resource Name (ARN) of the cluster to modify the
	// capacity provider settings for. If you do not specify a cluster, the default
	// cluster is assumed.
	//
	// This member is required.
	Cluster *string

	// The capacity provider strategy to use by default for the cluster. When creating
	// a service or running a task on a cluster, if no capacity provider or launch type
	// is specified then the default capacity provider strategy for the cluster is
	// used. A capacity provider strategy consists of one or more capacity providers
	// along with the base and weight to assign to them. A capacity provider must be
	// associated with the cluster to be used in a capacity provider strategy. The
	// PutClusterCapacityProviders () API is used to associate a capacity provider with
	// a cluster. Only capacity providers with an ACTIVE or UPDATING status can be
	// used. If specifying a capacity provider that uses an Auto Scaling group, the
	// capacity provider must already be created. New capacity providers can be created
	// with the CreateCapacityProvider () API operation. To use a AWS Fargate capacity
	// provider, specify either the FARGATE or FARGATE_SPOT capacity providers. The AWS
	// Fargate capacity providers are available to all accounts and only need to be
	// associated with a cluster to be used.
	//
	// This member is required.
	DefaultCapacityProviderStrategy []*types.CapacityProviderStrategyItem
}

type PutClusterCapacityProvidersOutput struct {

	// A regional grouping of one or more container instances on which you can run task
	// requests. Each account receives a default cluster the first time you use the
	// Amazon ECS service, but you may also create other clusters. Clusters may contain
	// more than one instance type simultaneously.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutClusterCapacityProvidersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutClusterCapacityProviders{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutClusterCapacityProviders{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutClusterCapacityProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "PutClusterCapacityProviders",
	}
}
