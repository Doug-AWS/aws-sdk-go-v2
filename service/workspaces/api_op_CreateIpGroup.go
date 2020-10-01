// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an IP access control group. An IP access control group provides you with
// the ability to control the IP addresses from which users are allowed to access
// their WorkSpaces. To specify the CIDR address ranges, add rules to your IP
// access control group and then associate the group with your directory. You can
// add rules when you create the group or at any time using AuthorizeIpRules ().
// <p>There is a default IP access control group associated with your directory. If
// you don't associate an IP access control group with your directory, the default
// group is used. The default group includes a default rule that allows users to
// access their WorkSpaces from anywhere. You cannot modify the default IP access
// control group for your directory.</p>
func (c *Client) CreateIpGroup(ctx context.Context, params *CreateIpGroupInput, optFns ...func(*Options)) (*CreateIpGroupOutput, error) {
	stack := middleware.NewStack("CreateIpGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateIpGroupMiddlewares(stack)
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
	addOpCreateIpGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIpGroup(options.Region), middleware.Before)
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
			OperationName: "CreateIpGroup",
			Err:           err,
		}
	}
	out := result.(*CreateIpGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIpGroupInput struct {

	// The name of the group.
	//
	// This member is required.
	GroupName *string

	// The description of the group.
	GroupDesc *string

	// The tags. Each WorkSpaces resource can have a maximum of 50 tags.
	Tags []*types.Tag

	// The rules to add to the group.
	UserRules []*types.IpRuleItem
}

type CreateIpGroupOutput struct {

	// The identifier of the group.
	GroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateIpGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIpGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIpGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIpGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateIpGroup",
	}
}
