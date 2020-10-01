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

// Adds a member (user or group) to the group's set.
func (c *Client) AssociateMemberToGroup(ctx context.Context, params *AssociateMemberToGroupInput, optFns ...func(*Options)) (*AssociateMemberToGroupOutput, error) {
	stack := middleware.NewStack("AssociateMemberToGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateMemberToGroupMiddlewares(stack)
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
	addOpAssociateMemberToGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMemberToGroup(options.Region), middleware.Before)
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
			OperationName: "AssociateMemberToGroup",
			Err:           err,
		}
	}
	out := result.(*AssociateMemberToGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMemberToGroupInput struct {

	// The group to which the member (user or group) is associated.
	//
	// This member is required.
	GroupId *string

	// The member (user or group) to associate to the group.
	//
	// This member is required.
	MemberId *string

	// The organization under which the group exists.
	//
	// This member is required.
	OrganizationId *string
}

type AssociateMemberToGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateMemberToGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateMemberToGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateMemberToGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateMemberToGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "AssociateMemberToGroup",
	}
}
