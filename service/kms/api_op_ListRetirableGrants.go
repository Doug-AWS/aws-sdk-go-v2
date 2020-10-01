// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all grants for which the grant's RetiringPrincipal matches the
// one specified. A typical use is to list all grants that you are able to retire.
// To retire a grant, use RetireGrant ().
func (c *Client) ListRetirableGrants(ctx context.Context, params *ListRetirableGrantsInput, optFns ...func(*Options)) (*ListRetirableGrantsOutput, error) {
	stack := middleware.NewStack("ListRetirableGrants", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRetirableGrantsMiddlewares(stack)
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
	addOpListRetirableGrantsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRetirableGrants(options.Region), middleware.Before)
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
			OperationName: "ListRetirableGrants",
			Err:           err,
		}
	}
	out := result.(*ListRetirableGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetirableGrantsInput struct {

	// The retiring principal for which to list grants. To specify the retiring
	// principal, use the Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an AWS principal. Valid AWS principals include AWS accounts (root), IAM users,
	// federated users, and assumed role users. For examples of the ARN syntax for
	// specifying a principal, see AWS Identity and Access Management (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the Amazon Web Services General Reference.
	//
	// This member is required.
	RetiringPrincipal *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, AWS KMS does not return more than the specified number of
	// items, but it might return fewer. This value is optional. If you include a
	// value, it must be between 1 and 100, inclusive. If you do not include a value,
	// it defaults to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string
}

type ListRetirableGrantsOutput struct {

	// A list of grants.
	Grants []*types.GrantListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRetirableGrantsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRetirableGrants{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRetirableGrants{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRetirableGrants(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListRetirableGrants",
	}
}
