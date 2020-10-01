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

// Retrieves details for the specified user ID, such as primary email address,
// license type, and personal meeting PIN. To retrieve user details with an email
// address instead of a user ID, use the ListUsers () action, and then filter by
// email address.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	stack := middleware.NewStack("GetUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUserMiddlewares(stack)
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
	addOpGetUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before)
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
			OperationName: "GetUser",
			Err:           err,
		}
	}
	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The user ID.
	//
	// This member is required.
	UserId *string
}

type GetUserOutput struct {

	// The user details.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetUser",
	}
}
