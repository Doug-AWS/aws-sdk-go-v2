// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the event buses in your account, including the default event bus,
// custom event buses, and partner event buses.
func (c *Client) ListEventBuses(ctx context.Context, params *ListEventBusesInput, optFns ...func(*Options)) (*ListEventBusesOutput, error) {
	stack := middleware.NewStack("ListEventBuses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEventBusesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEventBuses(options.Region), middleware.Before)
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
			OperationName: "ListEventBuses",
			Err:           err,
		}
	}
	out := result.(*ListEventBusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventBusesInput struct {

	// Specifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// Specifying this limits the results to only those event buses with names that
	// start with the specified prefix.
	NamePrefix *string

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string
}

type ListEventBusesOutput struct {

	// This list of event buses.
	EventBuses []*types.EventBus

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEventBusesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEventBuses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventBuses{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEventBuses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListEventBuses",
	}
}
