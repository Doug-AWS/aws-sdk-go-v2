// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples serializes a blob shape in the payload. In this example, no XML
// document is synthesized because the payload is not a structure or a union type.
func (c *Client) HttpPayloadTraits(ctx context.Context, params *HttpPayloadTraitsInput, optFns ...func(*Options)) (*HttpPayloadTraitsOutput, error) {
	stack := middleware.NewStack("HttpPayloadTraits", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpHttpPayloadTraitsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadTraits(options.Region), middleware.Before)
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
			OperationName: "HttpPayloadTraits",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadTraitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadTraitsInput struct {
	Foo  *string
	Blob []byte
}

type HttpPayloadTraitsOutput struct {
	Foo  *string
	Blob []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpPayloadTraitsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadTraits{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadTraits{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadTraits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadTraits",
	}
}
