// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the approximate count of unique values that match the query.
func (c *Client) GetCardinality(ctx context.Context, params *GetCardinalityInput, optFns ...func(*Options)) (*GetCardinalityOutput, error) {
	stack := middleware.NewStack("GetCardinality", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCardinalityMiddlewares(stack)
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
	addOpGetCardinalityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCardinality(options.Region), middleware.Before)
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
			OperationName: "GetCardinality",
			Err:           err,
		}
	}
	out := result.(*GetCardinalityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCardinalityInput struct {

	// The search query.
	//
	// This member is required.
	QueryString *string

	// The field to aggregate.
	AggregationField *string

	// The name of the index to search.
	IndexName *string

	// The query version.
	QueryVersion *string
}

type GetCardinalityOutput struct {

	// The approximate count of unique values that match the query.
	Cardinality *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCardinalityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCardinality{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCardinality{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCardinality(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetCardinality",
	}
}
