// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a theme alias for a theme.
func (c *Client) CreateThemeAlias(ctx context.Context, params *CreateThemeAliasInput, optFns ...func(*Options)) (*CreateThemeAliasOutput, error) {
	stack := middleware.NewStack("CreateThemeAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateThemeAliasMiddlewares(stack)
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
	addOpCreateThemeAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThemeAlias(options.Region), middleware.Before)
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
			OperationName: "CreateThemeAlias",
			Err:           err,
		}
	}
	out := result.(*CreateThemeAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThemeAliasInput struct {

	// The name that you want to give to the theme alias that you are creating. The
	// alias name can't begin with a $. Alias names that start with $ are reserved by
	// Amazon QuickSight.
	//
	// This member is required.
	AliasName *string

	// The ID of the AWS account that contains the theme for the new theme alias.
	//
	// This member is required.
	AwsAccountId *string

	// An ID for the theme alias.
	//
	// This member is required.
	ThemeId *string

	// The version number of the theme.
	//
	// This member is required.
	ThemeVersionNumber *int64
}

type CreateThemeAliasOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Information about the theme alias.
	ThemeAlias *types.ThemeAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateThemeAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateThemeAlias{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThemeAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateThemeAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateThemeAlias",
	}
}
