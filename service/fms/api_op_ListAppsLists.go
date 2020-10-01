// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of AppsListDataSummary objects.
func (c *Client) ListAppsLists(ctx context.Context, params *ListAppsListsInput, optFns ...func(*Options)) (*ListAppsListsOutput, error) {
	stack := middleware.NewStack("ListAppsLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAppsListsMiddlewares(stack)
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
	addOpListAppsListsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAppsLists(options.Region), middleware.Before)
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
			OperationName: "ListAppsLists",
			Err:           err,
		}
	}
	out := result.(*ListAppsListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppsListsInput struct {

	// The maximum number of objects that you want AWS Firewall Manager to return for
	// this request. If more objects are available, in the response, AWS Firewall
	// Manager provides a NextToken value that you can use in a subsequent call to get
	// the next batch of objects. If you don't specify this, AWS Firewall Manager
	// returns all available objects.
	//
	// This member is required.
	MaxResults *int32

	// Specifies whether the lists to retrieve are default lists owned by AWS Firewall
	// Manager.
	DefaultLists *bool

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, AWS Firewall Manager returns this token in the
	// response. For all but the first request, you provide the token returned by the
	// prior request in the request parameters, to retrieve the next batch of objects.
	NextToken *string
}

type ListAppsListsOutput struct {

	// An array of AppsListDataSummary objects.
	AppsLists []*types.AppsListDataSummary

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, AWS Firewall Manager returns this token in the
	// response. You can use this token in subsequent requests to retrieve the next
	// batch of objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAppsListsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAppsLists{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAppsLists{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAppsLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListAppsLists",
	}
}
