// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides organization conformance pack deployment status for an organization.
// Only a master account and a delegated administrator account can call this API.
// When calling this API with a delegated administrator, you must ensure AWS
// Organizations ListDelegatedAdministrator permissions are added. The status is
// not considered successful until organization conformance pack is successfully
// deployed in all the member accounts with an exception of excluded accounts. When
// you specify the limit and the next token, you receive a paginated response.
// Limit and next token are not applicable if you specify organization conformance
// pack names. They are only applicable, when you request all the organization
// conformance packs.
func (c *Client) DescribeOrganizationConformancePackStatuses(ctx context.Context, params *DescribeOrganizationConformancePackStatusesInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePackStatusesOutput, error) {
	stack := middleware.NewStack("DescribeOrganizationConformancePackStatuses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeOrganizationConformancePackStatusesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(options.Region), middleware.Before)
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
			OperationName: "DescribeOrganizationConformancePackStatuses",
			Err:           err,
		}
	}
	out := result.(*DescribeOrganizationConformancePackStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConformancePackStatusesInput struct {

	// The maximum number of OrganizationConformancePackStatuses returned on each page.
	// If you do no specify a number, AWS Config uses the default. The default is 100.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization conformance packs for which you want status details.
	// If you do not specify any names, AWS Config returns details for all your
	// organization conformance packs.
	OrganizationConformancePackNames []*string
}

type DescribeOrganizationConformancePackStatusesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConformancePackStatus objects.
	OrganizationConformancePackStatuses []*types.OrganizationConformancePackStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeOrganizationConformancePackStatusesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConformancePackStatuses{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrganizationConformancePackStatuses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeOrganizationConformancePackStatuses",
	}
}
