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

// Returns a list of the available cache engines and their versions.
func (c *Client) DescribeCacheEngineVersions(ctx context.Context, params *DescribeCacheEngineVersionsInput, optFns ...func(*Options)) (*DescribeCacheEngineVersionsOutput, error) {
	stack := middleware.NewStack("DescribeCacheEngineVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeCacheEngineVersionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheEngineVersions(options.Region), middleware.Before)
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
			OperationName: "DescribeCacheEngineVersions",
			Err:           err,
		}
	}
	out := result.(*DescribeCacheEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsInput struct {

	// The name of a specific cache parameter group family to return details for. Valid
	// values are: memcached1.4 | memcached1.5 | redis2.6 | redis2.8 | redis3.2 |
	// redis4.0 | redis5.0 | Constraints:
	//
	//     * Must be 1 to 255 alphanumeric
	// characters
	//
	//     * First character must be a letter
	//
	//     * Cannot end with a
	// hyphen or contain two consecutive hyphens
	CacheParameterGroupFamily *string

	// If true, specifies that only the default version of the specified engine or
	// engine and major version combination is to be returned.
	DefaultOnly *bool

	// The cache engine to return. Valid values: memcached | redis
	Engine *string

	// The cache engine version to return. Example: 1.4.14
	EngineVersion *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.  <p>Default: 100</p>
	// <p>Constraints: minimum 20; maximum 100.</p>
	MaxRecords *int32
}

// Represents the output of a DescribeCacheEngineVersions () operation.
type DescribeCacheEngineVersionsOutput struct {

	// A list of cache engine version details. Each element in the list contains
	// detailed information about one cache engine version.
	CacheEngineVersions []*types.CacheEngineVersion

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeCacheEngineVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheEngineVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheEngineVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCacheEngineVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheEngineVersions",
	}
}
