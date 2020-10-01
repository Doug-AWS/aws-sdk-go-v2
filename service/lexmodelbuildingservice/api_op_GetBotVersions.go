// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about all of the versions of a bot. The GetBotVersions
// operation returns a BotMetadata object for each version of a bot. For example,
// if a bot has three numbered versions, the GetBotVersions operation returns four
// BotMetadata objects in the response, one for each numbered version and one for
// the $LATEST version. The GetBotVersions operation always returns at least one
// version, the $LATEST version. This operation requires permissions for the
// lex:GetBotVersions action.
func (c *Client) GetBotVersions(ctx context.Context, params *GetBotVersionsInput, optFns ...func(*Options)) (*GetBotVersionsOutput, error) {
	stack := middleware.NewStack("GetBotVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBotVersionsMiddlewares(stack)
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
	addOpGetBotVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotVersions(options.Region), middleware.Before)
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
			OperationName: "GetBotVersions",
			Err:           err,
		}
	}
	out := result.(*GetBotVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotVersionsInput struct {

	// The name of the bot for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of bot versions to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string
}

type GetBotVersionsOutput struct {

	// An array of BotMetadata objects, one for each numbered version of the bot plus
	// one for the $LATEST version.
	Bots []*types.BotMetadata

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBotVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBotVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBotVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotVersions",
	}
}
