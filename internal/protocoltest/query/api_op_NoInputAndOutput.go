// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request payload or response members. While this should be rare, code generators
// must support this.
func (c *Client) NoInputAndOutput(ctx context.Context, params *NoInputAndOutputInput, optFns ...func(*Options)) (*NoInputAndOutputOutput, error) {
	stack := middleware.NewStack("NoInputAndOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNoInputAndOutput(options.Region), middleware.Before)

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
			OperationName: "NoInputAndOutput",
			Err:           err,
		}
	}
	out := result.(*NoInputAndOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NoInputAndOutputInput struct {
}

type NoInputAndOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opNoInputAndOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "NoInputAndOutput",
	}
}
