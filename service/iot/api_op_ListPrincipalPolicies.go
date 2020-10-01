// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the policies attached to the specified principal. If you use an Cognito
// identity, the ID must be in AmazonCognito Identity format
// (https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetCredentialsForIdentity.html#API_GetCredentialsForIdentity_RequestSyntax).
// Note: This API is deprecated. Please use ListAttachedPolicies () instead.
func (c *Client) ListPrincipalPolicies(ctx context.Context, params *ListPrincipalPoliciesInput, optFns ...func(*Options)) (*ListPrincipalPoliciesOutput, error) {
	stack := middleware.NewStack("ListPrincipalPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPrincipalPoliciesMiddlewares(stack)
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
	addOpListPrincipalPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalPolicies(options.Region), middleware.Before)
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
			OperationName: "ListPrincipalPolicies",
			Err:           err,
		}
	}
	out := result.(*ListPrincipalPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPrincipalPolicies operation.
type ListPrincipalPoliciesInput struct {

	// The principal.
	//
	// This member is required.
	Principal *string

	// Specifies the order for results. If true, results are returned in ascending
	// creation order.
	AscendingOrder *bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32
}

// The output from the ListPrincipalPolicies operation.
type ListPrincipalPoliciesOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The policies.
	Policies []*types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPrincipalPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipalPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipalPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPrincipalPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPrincipalPolicies",
	}
}
