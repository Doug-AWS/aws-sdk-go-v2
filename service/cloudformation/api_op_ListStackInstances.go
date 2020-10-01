// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about stack instances that are associated with the
// specified stack set. You can filter for stack instances that are associated with
// a specific AWS account name or Region, or that have a specific status.
func (c *Client) ListStackInstances(ctx context.Context, params *ListStackInstancesInput, optFns ...func(*Options)) (*ListStackInstancesOutput, error) {
	stack := middleware.NewStack("ListStackInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListStackInstancesMiddlewares(stack)
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
	addOpListStackInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListStackInstances(options.Region), middleware.Before)
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
			OperationName: "ListStackInstances",
			Err:           err,
		}
	}
	out := result.(*ListStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackInstancesInput struct {

	// The name or unique ID of the stack set that you want to list stack instances
	// for.
	//
	// This member is required.
	StackSetName *string

	// The status that stack instances are filtered by.
	Filters []*types.StackInstanceFilter

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous request didn't return all of the remaining results, the
	// response's NextToken parameter value is set to a token. To retrieve the next set
	// of results, call ListStackInstances again and assign that token to the request
	// object's NextToken parameter. If there are no remaining results, the previous
	// response object's NextToken parameter is set to null.
	NextToken *string

	// The name of the AWS account that you want to list stack instances for.
	StackInstanceAccount *string

	// The name of the Region where you want to list stack instances.
	StackInstanceRegion *string
}

type ListStackInstancesOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListStackInstances again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null.
	NextToken *string

	// A list of StackInstanceSummary structures that contain information about the
	// specified stack instances.
	Summaries []*types.StackInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListStackInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListStackInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opListStackInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStackInstances",
	}
}
