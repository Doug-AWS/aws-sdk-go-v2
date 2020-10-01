// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a member (user or group) to the resource's set of delegates.
func (c *Client) AssociateDelegateToResource(ctx context.Context, params *AssociateDelegateToResourceInput, optFns ...func(*Options)) (*AssociateDelegateToResourceOutput, error) {
	stack := middleware.NewStack("AssociateDelegateToResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateDelegateToResourceMiddlewares(stack)
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
	addOpAssociateDelegateToResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDelegateToResource(options.Region), middleware.Before)
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
			OperationName: "AssociateDelegateToResource",
			Err:           err,
		}
	}
	out := result.(*AssociateDelegateToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDelegateToResourceInput struct {

	// The member (user or group) to associate to the resource.
	//
	// This member is required.
	EntityId *string

	// The organization under which the resource exists.
	//
	// This member is required.
	OrganizationId *string

	// The resource for which members (users or groups) are associated.
	//
	// This member is required.
	ResourceId *string
}

type AssociateDelegateToResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateDelegateToResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDelegateToResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDelegateToResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateDelegateToResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "AssociateDelegateToResource",
	}
}
