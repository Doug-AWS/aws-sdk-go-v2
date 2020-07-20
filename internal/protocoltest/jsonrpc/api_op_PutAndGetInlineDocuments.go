// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as part of the payload.
func (c *Client) PutAndGetInlineDocuments(ctx context.Context, params *PutAndGetInlineDocumentsInput, optFns ...func(*Options)) (*PutAndGetInlineDocumentsOutput, error) {
	stack := middleware.NewStack("PutAndGetInlineDocuments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAndGetInlineDocuments(options.Region), middleware.Before)
	addawsAwsjson11_serdeOpPutAndGetInlineDocumentsMiddlewares(stack)

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
			OperationName: "PutAndGetInlineDocuments",
			Err:           err,
		}
	}
	out := result.(*PutAndGetInlineDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAndGetInlineDocumentsInput struct {
	InlineDocument smithy.Document
}

type PutAndGetInlineDocumentsOutput struct {
	InlineDocument smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutAndGetInlineDocumentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutAndGetInlineDocuments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAndGetInlineDocuments{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAndGetInlineDocuments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Json Protocol",
		ServiceID:      "jsonprotocol",
		EndpointPrefix: "jsonprotocol",
		SigningName:    "foo",
		OperationName:  "PutAndGetInlineDocuments",
	}
}
