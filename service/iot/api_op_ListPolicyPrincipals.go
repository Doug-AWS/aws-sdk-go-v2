// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the principals associated with the specified policy. Note: This API is
// deprecated. Please use ListTargetsForPolicy () instead.
func (c *Client) ListPolicyPrincipals(ctx context.Context, params *ListPolicyPrincipalsInput, optFns ...func(*Options)) (*ListPolicyPrincipalsOutput, error) {
	stack := middleware.NewStack("ListPolicyPrincipals", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPolicyPrincipalsMiddlewares(stack)
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
	addOpListPolicyPrincipalsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicyPrincipals(options.Region), middleware.Before)

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
			OperationName: "ListPolicyPrincipals",
			Err:           err,
		}
	}
	out := result.(*ListPolicyPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPolicyPrincipals operation.
type ListPolicyPrincipalsInput struct {
	// The result page size.
	PageSize *int32
	// The policy name.
	PolicyName *string
	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool
	// The marker for the next set of results.
	Marker *string
}

// The output from the ListPolicyPrincipals operation.
type ListPolicyPrincipalsOutput struct {
	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string
	// The descriptions of the principals.
	Principals []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPolicyPrincipalsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPolicyPrincipals{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPolicyPrincipals{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPolicyPrincipals(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListPolicyPrincipals",
	}
}