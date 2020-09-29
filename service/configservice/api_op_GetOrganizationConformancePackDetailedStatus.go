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

// Returns detailed status for each member account within an organization for a
// given organization conformance pack. Only a master account and a delegated
// administrator account can call this API. When calling this API with a delegated
// administrator, you must ensure AWS Organizations ListDelegatedAdministrator
// permissions are added.
func (c *Client) GetOrganizationConformancePackDetailedStatus(ctx context.Context, params *GetOrganizationConformancePackDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error) {
	stack := middleware.NewStack("GetOrganizationConformancePackDetailedStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetOrganizationConformancePackDetailedStatusMiddlewares(stack)
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
	addOpGetOrganizationConformancePackDetailedStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(options.Region), middleware.Before)
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
			OperationName: "GetOrganizationConformancePackDetailedStatus",
			Err:           err,
		}
	}
	out := result.(*GetOrganizationConformancePackDetailedStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationConformancePackDetailedStatusInput struct {
	// An OrganizationResourceDetailedStatusFilters object.
	Filters *types.OrganizationResourceDetailedStatusFilters
	// The name of organization conformance pack for which you want status details for
	// member accounts.
	OrganizationConformancePackName *string
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of OrganizationConformancePackDetailedStatuses returned on
	// each page. If you do not specify a number, AWS Config uses the default. The
	// default is 100.
	Limit *int32
}

type GetOrganizationConformancePackDetailedStatusOutput struct {
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// A list of OrganizationConformancePackDetailedStatus objects.
	OrganizationConformancePackDetailedStatuses []*types.OrganizationConformancePackDetailedStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetOrganizationConformancePackDetailedStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetOrganizationConformancePackDetailedStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOrganizationConformancePackDetailedStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetOrganizationConformancePackDetailedStatus",
	}
}
