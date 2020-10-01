// Code generated by smithy-go-codegen DO NOT EDIT.

package savingsplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Savings Plans offering rates.
func (c *Client) DescribeSavingsPlansOfferingRates(ctx context.Context, params *DescribeSavingsPlansOfferingRatesInput, optFns ...func(*Options)) (*DescribeSavingsPlansOfferingRatesOutput, error) {
	stack := middleware.NewStack("DescribeSavingsPlansOfferingRates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeSavingsPlansOfferingRatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSavingsPlansOfferingRates(options.Region), middleware.Before)
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
			OperationName: "DescribeSavingsPlansOfferingRates",
			Err:           err,
		}
	}
	out := result.(*DescribeSavingsPlansOfferingRatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSavingsPlansOfferingRatesInput struct {

	// The filters.
	Filters []*types.SavingsPlanOfferingRateFilterElement

	// The maximum number of results to return with a single call. To retrieve
	// additional results, make another call with the returned token value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The specific AWS operation for the line item in the billing report.
	Operations []*string

	// The AWS products.
	Products []types.SavingsPlanProductType

	// The IDs of the offerings.
	SavingsPlanOfferingIds []*string

	// The payment options.
	SavingsPlanPaymentOptions []types.SavingsPlanPaymentOption

	// The plan types.
	SavingsPlanTypes []types.SavingsPlanType

	// The services.
	ServiceCodes []types.SavingsPlanRateServiceCode

	// The usage details of the line item in the billing report.
	UsageTypes []*string
}

type DescribeSavingsPlansOfferingRatesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Savings Plans offering rates.
	SearchResults []*types.SavingsPlanOfferingRate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeSavingsPlansOfferingRatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSavingsPlansOfferingRates{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSavingsPlansOfferingRates{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSavingsPlansOfferingRates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "savingsplans",
		OperationName: "DescribeSavingsPlansOfferingRates",
	}
}
