// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a specific email address that's on the suppression
// list for your account.
func (c *Client) GetSuppressedDestination(ctx context.Context, params *GetSuppressedDestinationInput, optFns ...func(*Options)) (*GetSuppressedDestinationOutput, error) {
	stack := middleware.NewStack("GetSuppressedDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSuppressedDestinationMiddlewares(stack)
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
	addOpGetSuppressedDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSuppressedDestination(options.Region), middleware.Before)
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
			OperationName: "GetSuppressedDestination",
			Err:           err,
		}
	}
	out := result.(*GetSuppressedDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve information about an email address that's on the
// suppression list for your account.
type GetSuppressedDestinationInput struct {
	// The email address that's on the account suppression list.
	EmailAddress *string
}

// Information about the suppressed email address.
type GetSuppressedDestinationOutput struct {
	// An object containing information about the suppressed email address.
	SuppressedDestination *types.SuppressedDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSuppressedDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSuppressedDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSuppressedDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSuppressedDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetSuppressedDestination",
	}
}
