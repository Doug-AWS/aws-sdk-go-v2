// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an account with an Amazon Macie master account.
func (c *Client) CreateMember(ctx context.Context, params *CreateMemberInput, optFns ...func(*Options)) (*CreateMemberOutput, error) {
	stack := middleware.NewStack("CreateMember", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMemberMiddlewares(stack)
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
	addOpCreateMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMember(options.Region), middleware.Before)
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
			OperationName: "CreateMember",
			Err:           err,
		}
	}
	out := result.(*CreateMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMemberInput struct {

	// The details for the account to associate with the master account.
	//
	// This member is required.
	Account *types.AccountDetail

	// A map of key-value pairs that specifies the tags to associate with the account
	// in Amazon Macie. An account can have a maximum of 50 tags. Each tag consists of
	// a required tag key and an associated tag value. The maximum length of a tag key
	// is 128 characters. The maximum length of a tag value is 256 characters.
	Tags map[string]*string
}

type CreateMemberOutput struct {

	// The Amazon Resource Name (ARN) of the account that was associated with the
	// master account.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMemberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMember{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMember{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateMember",
	}
}
