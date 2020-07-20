// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML namespace.
func (c *Client) HttpPayloadWithXmlNamespace(ctx context.Context, params *HttpPayloadWithXmlNamespaceInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNamespaceOutput, error) {
	stack := middleware.NewStack("HttpPayloadWithXmlNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespace(options.Region), middleware.Before)
	addawsRestxml_serdeOpHttpPayloadWithXmlNamespaceMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "HttpPayloadWithXmlNamespace",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadWithXmlNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNamespaceInput struct {
	Nested *types.PayloadWithXmlNamespace
}

type HttpPayloadWithXmlNamespaceOutput struct {
	Nested *types.PayloadWithXmlNamespace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpPayloadWithXmlNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlNamespace{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "HttpPayloadWithXmlNamespace",
	}
}
