// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of aliases in the caller's AWS account and region. You cannot list
// aliases in other accounts. For more information about aliases, see CreateAlias.
// By default, the ListAliases command returns all aliases in the account and
// region. To get only the aliases that point to a particular customer master key
// (CMK), use the KeyId parameter. The ListAliases response can include aliases
// that you created and associated with your customer managed CMKs, and aliases
// that AWS created and associated with AWS managed CMKs in your account. You can
// recognize AWS aliases because their names have the format aws/, such as
// aws/dynamodb. The response might also include aliases that have no TargetKeyId
// field. These are predefined aliases that AWS has created but has not yet
// associated with a CMK. Aliases that AWS creates in your account, including
// predefined aliases, do not count against your AWS KMS aliases quota
// (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#aliases-limit).
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, addOperationListAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {

	// Lists only aliases that refer to the specified CMK. The value of this parameter
	// can be the ID or Amazon Resource Name (ARN) of a CMK in the caller's account and
	// region. You cannot use an alias name or alias ARN in this value. This parameter
	// is optional. If you omit it, ListAliases returns all aliases in the account and
	// region.
	KeyId *string

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

type ListAliasesOutput struct {

	// A list of aliases.
	Aliases []*types.AliasListEntry

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

func addOperationListAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListAliases",
	}
}
