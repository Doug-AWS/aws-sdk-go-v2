// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searchs for dashboards that belong to a user.
func (c *Client) SearchDashboards(ctx context.Context, params *SearchDashboardsInput, optFns ...func(*Options)) (*SearchDashboardsOutput, error) {
	stack := middleware.NewStack("SearchDashboards", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSearchDashboardsMiddlewares(stack)
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
	addOpSearchDashboardsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDashboards(options.Region), middleware.Before)
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
			OperationName: "SearchDashboards",
			Err:           err,
		}
	}
	out := result.(*SearchDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDashboardsInput struct {

	// The ID of the AWS account that contains the user whose dashboards you're
	// searching for.
	//
	// This member is required.
	AwsAccountId *string

	// The filters to apply to the search. Currently, you can search only by user name,
	// for example, "Filters": [ { "Name": "QUICKSIGHT_USER", "Operator":
	// "StringEquals", "Value": "arn:aws:quicksight:us-east-1:1:user/default/UserName1"
	// } ]
	//
	// This member is required.
	Filters []*types.DashboardSearchFilter

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type SearchDashboardsOutput struct {

	// The list of dashboards owned by the user specified in Filters in your request.
	DashboardSummaryList []*types.DashboardSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSearchDashboardsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSearchDashboards{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchDashboards{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchDashboards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchDashboards",
	}
}
