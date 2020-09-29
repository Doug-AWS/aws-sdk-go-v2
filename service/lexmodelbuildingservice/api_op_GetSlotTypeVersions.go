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

// Gets information about all versions of a slot type. The GetSlotTypeVersions
// operation returns a SlotTypeMetadata object for each version of a slot type. For
// example, if a slot type has three numbered versions, the GetSlotTypeVersions
// operation returns four SlotTypeMetadata objects in the response, one for each
// numbered version and one for the $LATEST version. The GetSlotTypeVersions
// operation always returns at least one version, the $LATEST version. This
// operation requires permissions for the lex:GetSlotTypeVersions action.
func (c *Client) GetSlotTypeVersions(ctx context.Context, params *GetSlotTypeVersionsInput, optFns ...func(*Options)) (*GetSlotTypeVersionsOutput, error) {
	stack := middleware.NewStack("GetSlotTypeVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSlotTypeVersionsMiddlewares(stack)
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
	addOpGetSlotTypeVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSlotTypeVersions(options.Region), middleware.Before)
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
			OperationName: "GetSlotTypeVersions",
			Err:           err,
		}
	}
	out := result.(*GetSlotTypeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSlotTypeVersionsInput struct {
	// The maximum number of slot type versions to return in the response. The default
	// is 10.
	MaxResults *int32
	// The name of the slot type for which versions should be returned.
	Name *string
	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string
}

type GetSlotTypeVersionsOutput struct {
	// An array of SlotTypeMetadata objects, one for each numbered version of the slot
	// type plus one for the $LATEST version.
	SlotTypes []*types.SlotTypeMetadata
	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSlotTypeVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSlotTypeVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSlotTypeVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSlotTypeVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSlotTypeVersions",
	}
}
