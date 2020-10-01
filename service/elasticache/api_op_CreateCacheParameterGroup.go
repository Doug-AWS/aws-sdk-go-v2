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

// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the
// CacheParameterGroup. A newly created CacheParameterGroup is an exact duplicate
// of the default parameter group for the CacheParameterGroupFamily. To customize
// the newly created CacheParameterGroup you can change the values of specific
// parameters. For more information, see:
//
//     * ModifyCacheParameterGroup
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html)
// in the ElastiCache API Reference.
//
//     * Parameters and Parameter Groups
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ParameterGroups.html)
// in the ElastiCache User Guide.
func (c *Client) CreateCacheParameterGroup(ctx context.Context, params *CreateCacheParameterGroupInput, optFns ...func(*Options)) (*CreateCacheParameterGroupOutput, error) {
	stack := middleware.NewStack("CreateCacheParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateCacheParameterGroupMiddlewares(stack)
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
	addOpCreateCacheParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheParameterGroup(options.Region), middleware.Before)
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
			OperationName: "CreateCacheParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CreateCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheParameterGroup operation.
type CreateCacheParameterGroupInput struct {

	// The name of the cache parameter group family that the cache parameter group can
	// be used with. Valid values are: memcached1.4 | memcached1.5 | redis2.6 |
	// redis2.8 | redis3.2 | redis4.0 | redis5.0 |
	//
	// This member is required.
	CacheParameterGroupFamily *string

	// A user-specified name for the cache parameter group.
	//
	// This member is required.
	CacheParameterGroupName *string

	// A user-specified description for the cache parameter group.
	//
	// This member is required.
	Description *string
}

type CreateCacheParameterGroupOutput struct {

	// Represents the output of a CreateCacheParameterGroup operation.
	CacheParameterGroup *types.CacheParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateCacheParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCacheParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheParameterGroup",
	}
}
