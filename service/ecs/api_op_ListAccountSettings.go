// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the account settings for a specified principal.
func (c *Client) ListAccountSettings(ctx context.Context, params *ListAccountSettingsInput, optFns ...func(*Options)) (*ListAccountSettingsOutput, error) {
	stack := middleware.NewStack("ListAccountSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListAccountSettingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountSettings(options.Region), middleware.Before)
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
			OperationName: "ListAccountSettings",
			Err:           err,
		}
	}
	out := result.(*ListAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountSettingsInput struct {

	// Specifies whether to return the effective settings. If true, the account
	// settings for the root user or the default setting for the principalArn are
	// returned. If false, the account settings for the principalArn are returned if
	// they are set. Otherwise, no account settings are returned.
	EffectiveSettings *bool

	// The maximum number of account setting results returned by ListAccountSettings in
	// paginated output. When this parameter is used, ListAccountSettings only returns
	// maxResults results in a single page along with a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListAccountSettings request with the returned nextToken value. This value can be
	// between 1 and 10. If this parameter is not used, then ListAccountSettings
	// returns up to 10 results and a nextToken value if applicable.
	MaxResults *int32

	// The name of the account setting you want to list the settings for.
	Name types.SettingName

	// The nextToken value returned from a ListAccountSettings request indicating that
	// more results are available to fulfill the request and further calls will be
	// needed. If maxResults was provided, it is possible the number of results to be
	// fewer than maxResults. This token should be treated as an opaque identifier that
	// is only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// The ARN of the principal, which can be an IAM user, IAM role, or the root user.
	// If this field is omitted, the account settings are listed only for the
	// authenticated user.
	PrincipalArn *string

	// The value of the account settings with which to filter results. You must also
	// specify an account setting name to use this parameter.
	Value *string
}

type ListAccountSettingsOutput struct {

	// The nextToken value to include in a future ListAccountSettings request. When the
	// results of a ListAccountSettings request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// The account settings for the resource.
	Settings []*types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListAccountSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAccountSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListAccountSettings",
	}
}
