// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns the specified hierarchy group to the specified user.
func (c *Client) UpdateUserHierarchy(ctx context.Context, params *UpdateUserHierarchyInput, optFns ...func(*Options)) (*UpdateUserHierarchyOutput, error) {
	stack := middleware.NewStack("UpdateUserHierarchy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUserHierarchyMiddlewares(stack)
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
	addOpUpdateUserHierarchyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserHierarchy(options.Region), middleware.Before)
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
			OperationName: "UpdateUserHierarchy",
			Err:           err,
		}
	}
	out := result.(*UpdateUserHierarchyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserHierarchyInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier of the user account.
	//
	// This member is required.
	UserId *string

	// The identifier of the hierarchy group.
	HierarchyGroupId *string
}

type UpdateUserHierarchyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUserHierarchyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserHierarchy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserHierarchy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserHierarchy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateUserHierarchy",
	}
}
