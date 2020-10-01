// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of offering promotions. Each offering promotion record contains
// the ID and description of the promotion. The API returns a NotEligible error if
// the caller is not permitted to invoke the operation. Contact
// aws-devicefarm-support@amazon.com (mailto:aws-devicefarm-support@amazon.com) if
// you must be able to invoke this operation.
func (c *Client) ListOfferingPromotions(ctx context.Context, params *ListOfferingPromotionsInput, optFns ...func(*Options)) (*ListOfferingPromotionsOutput, error) {
	stack := middleware.NewStack("ListOfferingPromotions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListOfferingPromotionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferingPromotions(options.Region), middleware.Before)
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
			OperationName: "ListOfferingPromotions",
			Err:           err,
		}
	}
	out := result.(*ListOfferingPromotionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOfferingPromotionsInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

type ListOfferingPromotionsOutput struct {

	// An identifier to be used in the next call to this operation, to return the next
	// set of items in the list.
	NextToken *string

	// Information about the offering promotions.
	OfferingPromotions []*types.OfferingPromotion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListOfferingPromotionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListOfferingPromotions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOfferingPromotions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOfferingPromotions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListOfferingPromotions",
	}
}
