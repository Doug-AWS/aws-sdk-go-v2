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

// Change the settings for the account-level suppression list.
func (c *Client) PutAccountSuppressionAttributes(ctx context.Context, params *PutAccountSuppressionAttributesInput, optFns ...func(*Options)) (*PutAccountSuppressionAttributesOutput, error) {
	stack := middleware.NewStack("PutAccountSuppressionAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutAccountSuppressionAttributesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSuppressionAttributes(options.Region), middleware.Before)
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
			OperationName: "PutAccountSuppressionAttributes",
			Err:           err,
		}
	}
	out := result.(*PutAccountSuppressionAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change your account's suppression preferences.
type PutAccountSuppressionAttributesInput struct {
	// A list that contains the reasons that email addresses will be automatically
	// added to the suppression list for your account. This list can contain any or all
	// of the following:
	//
	//     * COMPLAINT – Amazon SES adds an email address to the
	// suppression list for your account when a message sent to that address results in
	// a complaint.
	//
	//     * BOUNCE – Amazon SES adds an email address to the suppression
	// list for your account when a message sent to that address results in a hard
	// bounce.
	SuppressedReasons []types.SuppressionListReason
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutAccountSuppressionAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutAccountSuppressionAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountSuppressionAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountSuppressionAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAccountSuppressionAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutAccountSuppressionAttributes",
	}
}
