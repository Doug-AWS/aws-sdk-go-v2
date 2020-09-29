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

// Returns a list of organization conformance packs. Only a master account and a
// delegated administrator account can call this API. When calling this API with a
// delegated administrator, you must ensure AWS Organizations
// ListDelegatedAdministrator permissions are added. When you specify the limit and
// the next token, you receive a paginated response. Limit and next token are not
// applicable if you specify organization conformance packs names. They are only
// applicable, when you request all the organization conformance packs.
func (c *Client) DescribeOrganizationConformancePacks(ctx context.Context, params *DescribeOrganizationConformancePacksInput, optFns ...func(*Options)) (*DescribeOrganizationConformancePacksOutput, error) {
	stack := middleware.NewStack("DescribeOrganizationConformancePacks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeOrganizationConformancePacksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConformancePacks(options.Region), middleware.Before)
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
			OperationName: "DescribeOrganizationConformancePacks",
			Err:           err,
		}
	}
	out := result.(*DescribeOrganizationConformancePacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConformancePacksInput struct {
	// The name that you assign to an organization conformance pack.
	OrganizationConformancePackNames []*string
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of organization config packs returned on each page. If you do
	// no specify a number, AWS Config uses the default. The default is 100.
	Limit *int32
}

type DescribeOrganizationConformancePacksOutput struct {
	// Returns a list of OrganizationConformancePacks objects.
	OrganizationConformancePacks []*types.OrganizationConformancePack
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeOrganizationConformancePacksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConformancePacks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConformancePacks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrganizationConformancePacks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeOrganizationConformancePacks",
	}
}
