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

// Returns a list of registration tokens for the specified type(s).
func (c *Client) ListTypeRegistrations(ctx context.Context, params *ListTypeRegistrationsInput, optFns ...func(*Options)) (*ListTypeRegistrationsOutput, error) {
	stack := middleware.NewStack("ListTypeRegistrations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListTypeRegistrationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypeRegistrations(options.Region), middleware.Before)
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
			OperationName: "ListTypeRegistrations",
			Err:           err,
		}
	}
	out := result.(*ListTypeRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypeRegistrationsInput struct {

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// The current status of the type registration request. The default is IN_PROGRESS.
	RegistrationStatusFilter types.RegistrationStatus

	// The kind of type. Currently the only valid value is RESOURCE. Conditional: You
	// must specify either TypeName and Type, or Arn.
	Type types.RegistryType

	// The Amazon Resource Name (ARN) of the type. Conditional: You must specify either
	// TypeName and Type, or Arn.
	TypeArn *string

	// The name of the type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	TypeName *string
}

type ListTypeRegistrationsOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null.
	NextToken *string

	// A list of type registration tokens. Use DescribeTypeRegistration () to return
	// detailed information about a type registration request.
	RegistrationTokenList []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListTypeRegistrationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListTypeRegistrations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypeRegistrations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTypeRegistrations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypeRegistrations",
	}
}
