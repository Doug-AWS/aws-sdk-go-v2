// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes your managed prefix lists and any AWS-managed prefix lists. To view
// the entries for your prefix list, use GetManagedPrefixListEntries ().
func (c *Client) DescribeManagedPrefixLists(ctx context.Context, params *DescribeManagedPrefixListsInput, optFns ...func(*Options)) (*DescribeManagedPrefixListsOutput, error) {
	stack := middleware.NewStack("DescribeManagedPrefixLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeManagedPrefixListsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeManagedPrefixLists(options.Region), middleware.Before)
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
			OperationName: "DescribeManagedPrefixLists",
			Err:           err,
		}
	}
	out := result.(*DescribeManagedPrefixListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeManagedPrefixListsInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The token for the next page of results.
	NextToken *string
	// One or more filters.
	//
	//     * owner-id - The ID of the prefix list owner.
	//
	//     *
	// prefix-list-id - The ID of the prefix list.
	//
	//     * prefix-list-name - The name
	// of the prefix list.
	Filters []*types.Filter
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// One or more prefix list IDs.
	PrefixListIds []*string
}

type DescribeManagedPrefixListsOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// Information about the prefix lists.
	PrefixLists []*types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeManagedPrefixListsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeManagedPrefixLists{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeManagedPrefixLists{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeManagedPrefixLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeManagedPrefixLists",
	}
}
