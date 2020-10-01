// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the resource query of a group.
func (c *Client) UpdateGroupQuery(ctx context.Context, params *UpdateGroupQueryInput, optFns ...func(*Options)) (*UpdateGroupQueryOutput, error) {
	stack := middleware.NewStack("UpdateGroupQuery", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateGroupQueryMiddlewares(stack)
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
	addOpUpdateGroupQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGroupQuery(options.Region), middleware.Before)
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
			OperationName: "UpdateGroupQuery",
			Err:           err,
		}
	}
	out := result.(*UpdateGroupQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGroupQueryInput struct {

	// The resource query to determine which AWS resources are members of this resource
	// group.
	//
	// This member is required.
	ResourceQuery *types.ResourceQuery

	// The name or the ARN of the resource group to query.
	Group *string

	// Don't use this parameter. Use Group instead.
	GroupName *string
}

type UpdateGroupQueryOutput struct {

	// The updated resource query associated with the resource group after the update.
	GroupQuery *types.GroupQuery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateGroupQueryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGroupQuery{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGroupQuery{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGroupQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "UpdateGroupQuery",
	}
}
