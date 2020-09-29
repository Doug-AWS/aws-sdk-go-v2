// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Verifies the specified user attributes in the user pool.
func (c *Client) VerifyUserAttribute(ctx context.Context, params *VerifyUserAttributeInput, optFns ...func(*Options)) (*VerifyUserAttributeOutput, error) {
	stack := middleware.NewStack("VerifyUserAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpVerifyUserAttributeMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpVerifyUserAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyUserAttribute(options.Region), middleware.Before)
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
			OperationName: "VerifyUserAttribute",
			Err:           err,
		}
	}
	out := result.(*VerifyUserAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to verify user attributes.
type VerifyUserAttributeInput struct {
	// Represents the access token of the request to verify user attributes.
	AccessToken *string
	// The verification code in the request to verify user attributes.
	Code *string
	// The attribute name in the request to verify user attributes.
	AttributeName *string
}

// A container representing the response from the server from the request to verify
// user attributes.
type VerifyUserAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpVerifyUserAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpVerifyUserAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifyUserAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opVerifyUserAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyUserAttribute",
	}
}
