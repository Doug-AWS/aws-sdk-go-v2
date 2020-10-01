// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the reservation utilization for your account. Master accounts in an
// organization have access to member accounts. You can filter data by dimensions
// in a time period. You can use GetDimensionValues to determine the possible
// dimension values. Currently, you can group only by SUBSCRIPTION_ID.
func (c *Client) GetReservationUtilization(ctx context.Context, params *GetReservationUtilizationInput, optFns ...func(*Options)) (*GetReservationUtilizationOutput, error) {
	stack := middleware.NewStack("GetReservationUtilization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetReservationUtilizationMiddlewares(stack)
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
	addOpGetReservationUtilizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationUtilization(options.Region), middleware.Before)
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
			OperationName: "GetReservationUtilization",
			Err:           err,
		}
	}
	out := result.(*GetReservationUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservationUtilizationInput struct {

	// Sets the start and end dates for retrieving RI utilization. The start date is
	// inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters utilization data by dimensions. You can filter by the following
	// dimensions:
	//
	//     * AZ
	//
	//     * CACHE_ENGINE
	//
	//     * DEPLOYMENT_OPTION
	//
	//     *
	// INSTANCE_TYPE
	//
	//     * LINKED_ACCOUNT
	//
	//     * OPERATING_SYSTEM
	//
	//     * PLATFORM
	//
	//
	// * REGION
	//
	//     * SERVICE
	//
	//     * SCOPE
	//
	//     * TENANCY
	//
	// GetReservationUtilization
	// uses the same Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension,
	// and nesting is supported up to only one level deep. If there are multiple values
	// for a dimension, they are OR'd together.
	Filter *types.Expression

	// If GroupBy is set, Granularity can't be set. If Granularity isn't set, the
	// response object doesn't include Granularity, either MONTHLY or DAILY. If both
	// GroupBy and Granularity aren't set, GetReservationUtilization defaults to DAILY.
	// The GetReservationUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity

	// Groups only by SUBSCRIPTION_ID. Metadata is included.
	GroupBy []*types.GroupDefinition

	// The token to retrieve the next set of results. AWS provides the token when the
	// response from a previous call has more results than the maximum page size.
	NextPageToken *string
}

type GetReservationUtilizationOutput struct {

	// The amount of time that you used your RIs.
	//
	// This member is required.
	UtilizationsByTime []*types.UtilizationByTime

	// The token for the next set of retrievable results. AWS provides the token when
	// the response from a previous call has more results than the maximum page size.
	NextPageToken *string

	// The total amount of time that you used your RIs.
	Total *types.ReservationAggregates

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetReservationUtilizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationUtilization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationUtilization{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetReservationUtilization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetReservationUtilization",
	}
}
