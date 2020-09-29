// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all provisioned DAX clusters if no cluster identifier
// is specified, or about a specific DAX cluster if a cluster identifier is
// supplied. If the cluster is in the CREATING state, only cluster level
// information will be displayed until all of the nodes are successfully
// provisioned. If the cluster is in the DELETING state, only cluster level
// information will be displayed. If nodes are currently being added to the DAX
// cluster, node endpoint information and creation time for the additional nodes
// will not be displayed until they are completely provisioned. When the DAX
// cluster state is available, the cluster is ready for use. If nodes are currently
// being removed from the DAX cluster, no endpoint information for the removed
// nodes is displayed.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	stack := middleware.NewStack("DescribeClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {
	// The names of the DAX clusters being described.
	ClusterNames []*string
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved. The value for MaxResults must be between
	// 20 and 100.
	MaxResults *int32
}

type DescribeClustersOutput struct {
	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string
	// The descriptions of your DAX clusters, in response to a DescribeClusters
	// request.
	Clusters []*types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "DescribeClusters",
	}
}
