// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the delegates associated with a resource. Users and groups can be resource
// delegates and answer requests on behalf of the resource.
func (c *Client) ListResourceDelegates(ctx context.Context, params *ListResourceDelegatesInput, optFns ...func(*Options)) (*ListResourceDelegatesOutput, error) {
	stack := middleware.NewStack("ListResourceDelegates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListResourceDelegatesMiddlewares(stack)
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
	addOpListResourceDelegatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceDelegates(options.Region), middleware.Before)
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
			OperationName: "ListResourceDelegates",
			Err:           err,
		}
	}
	out := result.(*ListResourceDelegatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceDelegatesInput struct {

	// The identifier for the organization that contains the resource for which
	// delegates are listed.
	//
	// This member is required.
	OrganizationId *string

	// The identifier for the resource whose delegates are listed.
	//
	// This member is required.
	ResourceId *string

	// The number of maximum results in a page.
	MaxResults *int32

	// The token used to paginate through the delegates associated with a resource.
	NextToken *string
}

type ListResourceDelegatesOutput struct {

	// One page of the resource's delegates.
	Delegates []*types.Delegate

	// The token used to paginate through the delegates associated with a resource.
	// While results are still available, it has an associated value. When the last
	// page is reached, the token is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListResourceDelegatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceDelegates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceDelegates{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResourceDelegates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListResourceDelegates",
	}
}
