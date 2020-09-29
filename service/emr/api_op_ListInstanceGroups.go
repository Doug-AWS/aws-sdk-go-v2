// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides all available details about the instance groups in a cluster.
func (c *Client) ListInstanceGroups(ctx context.Context, params *ListInstanceGroupsInput, optFns ...func(*Options)) (*ListInstanceGroupsOutput, error) {
	stack := middleware.NewStack("ListInstanceGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListInstanceGroupsMiddlewares(stack)
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
	addOpListInstanceGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceGroups(options.Region), middleware.Before)
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
			OperationName: "ListInstanceGroups",
			Err:           err,
		}
	}
	out := result.(*ListInstanceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which instance groups to retrieve.
type ListInstanceGroupsInput struct {
	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
	// The identifier of the cluster for which to list the instance groups.
	ClusterId *string
}

// This input determines which instance groups to retrieve.
type ListInstanceGroupsOutput struct {
	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
	// The list of instance groups for the cluster and given filters.
	InstanceGroups []*types.InstanceGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListInstanceGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListInstanceGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstanceGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListInstanceGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstanceGroups",
	}
}
