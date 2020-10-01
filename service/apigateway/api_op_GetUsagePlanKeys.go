// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all the usage plan keys representing the API keys added to a specified
// usage plan.
func (c *Client) GetUsagePlanKeys(ctx context.Context, params *GetUsagePlanKeysInput, optFns ...func(*Options)) (*GetUsagePlanKeysOutput, error) {
	stack := middleware.NewStack("GetUsagePlanKeys", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUsagePlanKeysMiddlewares(stack)
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
	addOpGetUsagePlanKeysValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlanKeys(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetUsagePlanKeys",
			Err:           err,
		}
	}
	out := result.(*GetUsagePlanKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get all the usage plan keys representing the API keys added
// to a specified usage plan.
type GetUsagePlanKeysInput struct {

	// [Required] The Id of the UsagePlan () resource representing the usage plan
	// containing the to-be-retrieved UsagePlanKey () resource representing a plan
	// customer.
	//
	// This member is required.
	UsagePlanId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	Name *string

	// A query parameter specifying the name of the to-be-returned usage plan keys.
	NameQuery *string

	// The current pagination position in the paged result set.
	Position *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents the collection of usage plan keys added to usage plans for the
// associated API keys and, possibly, other types of keys. Create and Use Usage
// Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanKeysOutput struct {

	// The current page of elements from this collection.
	Items []*types.UsagePlanKey

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUsagePlanKeysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlanKeys{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlanKeys{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUsagePlanKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlanKeys",
	}
}
