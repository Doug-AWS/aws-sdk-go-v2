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

// Gets information about the specified thing.
func (c *Client) DescribeThing(ctx context.Context, params *DescribeThingInput, optFns ...func(*Options)) (*DescribeThingOutput, error) {
	stack := middleware.NewStack("DescribeThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeThingMiddlewares(stack)
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
	addOpDescribeThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThing(options.Region), middleware.Before)
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
			OperationName: "DescribeThing",
			Err:           err,
		}
	}
	out := result.(*DescribeThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeThing operation.
type DescribeThingInput struct {
	// The name of the thing.
	ThingName *string
}

// The output from the DescribeThing operation.
type DescribeThingOutput struct {
	// The ARN of the thing to describe.
	ThingArn *string
	// The thing type name.
	ThingTypeName *string
	// The name of the billing group the thing belongs to.
	BillingGroupName *string
	// The name of the thing.
	ThingName *string
	// The default MQTT client ID. For a typical device, the thing name is also used as
	// the default MQTT client ID. Although we don’t require a mapping between a
	// thing's registry name and its use of MQTT client IDs, certificates, or shadow
	// state, we recommend that you choose a thing name and use it as the MQTT client
	// ID for the registry and the Device Shadow service. This lets you better organize
	// your AWS IoT fleet without removing the flexibility of the underlying device
	// certificate model or shadows.
	DefaultClientId *string
	// The thing attributes.
	Attributes map[string]*string
	// The current version of the thing record in the registry. To avoid unintentional
	// changes to the information in the registry, you can pass the version information
	// in the expectedVersion parameter of the UpdateThing and DeleteThing calls.
	Version *int64
	// The ID of the thing to describe.
	ThingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeThing",
	}
}
