// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the phone numbers for the specified Amazon Chime account, Amazon Chime
// user, Amazon Chime Voice Connector, or Amazon Chime Voice Connector group.
func (c *Client) ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	stack := middleware.NewStack("ListPhoneNumbers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPhoneNumbersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbers(options.Region), middleware.Before)
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
			OperationName: "ListPhoneNumbers",
			Err:           err,
		}
	}
	out := result.(*ListPhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumbersInput struct {
	// The value to use for the filter.
	FilterValue *string
	// The filter to use to limit the number of results.
	FilterName types.PhoneNumberAssociationName
	// The phone number product type.
	ProductType types.PhoneNumberProductType
	// The phone number status.
	Status types.PhoneNumberStatus
	// The maximum number of results to return in a single call.
	MaxResults *int32
	// The token to use to retrieve the next page of results.
	NextToken *string
}

type ListPhoneNumbersOutput struct {
	// The phone number details.
	PhoneNumbers []*types.PhoneNumber
	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPhoneNumbersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumbers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumbers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPhoneNumbers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListPhoneNumbers",
	}
}
