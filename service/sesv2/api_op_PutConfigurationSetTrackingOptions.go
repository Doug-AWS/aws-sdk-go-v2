// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specify a custom domain to use for open and click tracking elements in email
// that you send.
func (c *Client) PutConfigurationSetTrackingOptions(ctx context.Context, params *PutConfigurationSetTrackingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetTrackingOptionsOutput, error) {
	stack := middleware.NewStack("PutConfigurationSetTrackingOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutConfigurationSetTrackingOptionsMiddlewares(stack)
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
	addOpPutConfigurationSetTrackingOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetTrackingOptions(options.Region), middleware.Before)
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
			OperationName: "PutConfigurationSetTrackingOptions",
			Err:           err,
		}
	}
	out := result.(*PutConfigurationSetTrackingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add a custom domain for tracking open and click events to a
// configuration set.
type PutConfigurationSetTrackingOptionsInput struct {
	// The domain that you want to use to track open and click events.
	CustomRedirectDomain *string
	// The name of the configuration set that you want to add a custom tracking domain
	// to.
	ConfigurationSetName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetTrackingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutConfigurationSetTrackingOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetTrackingOptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetTrackingOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutConfigurationSetTrackingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetTrackingOptions",
	}
}
