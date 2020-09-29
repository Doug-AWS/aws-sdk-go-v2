// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of configuration items as specified by the value passed to the
// required parameter configurationType. Optional filtering may be applied to
// refine search results.
func (c *Client) ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
	stack := middleware.NewStack("ListConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListConfigurationsMiddlewares(stack)
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
	addOpListConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurations(options.Region), middleware.Before)
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
			OperationName: "ListConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationsInput struct {
	// The total number of items to return. The maximum value is 100.
	MaxResults *int32
	// Certain filter criteria return output that can be sorted in ascending or
	// descending order. For a list of output characteristics for each filter, see
	// Using the ListConfigurations Action
	// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#ListConfigurations)
	// in the AWS Application Discovery Service User Guide.
	OrderBy []*types.OrderByElement
	// Token to retrieve the next set of results. For example, if a previous call to
	// ListConfigurations returned 100 items, but you set
	// ListConfigurationsRequest$maxResults to 10, you received a set of 10 results
	// along with a token. Use that token in this query to get the next set of 10.
	NextToken *string
	// A valid configuration identified by Application Discovery Service.
	ConfigurationType types.ConfigurationItemType
	// You can filter the request using various logical operators and a key-value
	// format. For example: {"key": "serverType", "value": "webServer"} For a complete
	// list of filter options and guidance about using them with this action, see Using
	// the ListConfigurations Action
	// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#ListConfigurations)
	// in the AWS Application Discovery Service User Guide.
	Filters []*types.Filter
}

type ListConfigurationsOutput struct {
	// Returns configuration details, including the configuration ID, attribute names,
	// and attribute values.
	Configurations []map[string]*string
	// Token to retrieve the next set of results. For example, if your call to
	// ListConfigurations returned 100 items, but you set
	// ListConfigurationsRequest$maxResults to 10, you received a set of 10 results
	// along with this token. Use this token in the next query to retrieve the next set
	// of 10.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "ListConfigurations",
	}
}
