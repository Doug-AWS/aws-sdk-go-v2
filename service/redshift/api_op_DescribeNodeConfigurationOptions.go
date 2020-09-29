// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns properties of possible node configurations such as node type, number of
// nodes, and disk usage for the specified action type.
func (c *Client) DescribeNodeConfigurationOptions(ctx context.Context, params *DescribeNodeConfigurationOptionsInput, optFns ...func(*Options)) (*DescribeNodeConfigurationOptionsOutput, error) {
	stack := middleware.NewStack("DescribeNodeConfigurationOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeNodeConfigurationOptionsMiddlewares(stack)
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
	addOpDescribeNodeConfigurationOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeConfigurationOptions(options.Region), middleware.Before)
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
			OperationName: "DescribeNodeConfigurationOptions",
			Err:           err,
		}
	}
	out := result.(*DescribeNodeConfigurationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeConfigurationOptionsInput struct {
	// The action type to evaluate for possible node configurations. Specify
	// "restore-cluster" to get configuration combinations based on an existing
	// snapshot. Specify "recommend-node-config" to get configuration recommendations
	// based on an existing cluster or snapshot. Specify "resize-cluster" to get
	// configuration combinations for elastic resize based on an existing cluster.
	ActionType types.ActionType
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 500
	// Constraints: minimum 100, maximum 500.
	MaxRecords *int32
	// A set of name, operator, and value items to filter the results.
	Filters []*types.NodeConfigurationOptionsFilter
	// The identifier of the snapshot to evaluate for possible node configurations.
	SnapshotIdentifier *string
	// The identifier of the cluster to evaluate for possible node configurations.
	ClusterIdentifier *string
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeNodeConfigurationOptions ()
	// request exceed the value specified in MaxRecords, AWS returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string
	// The AWS customer account used to create or copy the snapshot. Required if you
	// are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string
}

type DescribeNodeConfigurationOptionsOutput struct {
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string
	// A list of valid node configurations.
	NodeConfigurationOptionList []*types.NodeConfigurationOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeNodeConfigurationOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeNodeConfigurationOptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeNodeConfigurationOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNodeConfigurationOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeNodeConfigurationOptions",
	}
}
