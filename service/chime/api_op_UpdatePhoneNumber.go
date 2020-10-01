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

// Updates phone number details, such as product type or calling name, for the
// specified phone number ID. You can update one phone number detail at a time. For
// example, you can update either the product type or the calling name in one
// action. For toll-free numbers, you must use the Amazon Chime Voice Connector
// product type. Updates to outbound calling names can take up to 72 hours to
// complete. Pending updates to outbound calling names must be complete before you
// can request another update.
func (c *Client) UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) {
	stack := middleware.NewStack("UpdatePhoneNumber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdatePhoneNumberMiddlewares(stack)
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
	addOpUpdatePhoneNumberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePhoneNumber(options.Region), middleware.Before)
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
			OperationName: "UpdatePhoneNumber",
			Err:           err,
		}
	}
	out := result.(*UpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePhoneNumberInput struct {

	// The phone number ID.
	//
	// This member is required.
	PhoneNumberId *string

	// The outbound calling name associated with the phone number.
	CallingName *string

	// The product type.
	ProductType types.PhoneNumberProductType
}

type UpdatePhoneNumberOutput struct {

	// The updated phone number details.
	PhoneNumber *types.PhoneNumber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdatePhoneNumberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePhoneNumber{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePhoneNumber{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePhoneNumber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdatePhoneNumber",
	}
}
