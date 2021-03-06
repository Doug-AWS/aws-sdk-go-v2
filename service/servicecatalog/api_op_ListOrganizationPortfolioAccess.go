// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the organization nodes that have access to the specified portfolio. This
// API can only be called by the management account in the organization or by a
// delegated admin. If a delegated admin is de-registered, they can no longer
// perform this operation.
func (c *Client) ListOrganizationPortfolioAccess(ctx context.Context, params *ListOrganizationPortfolioAccessInput, optFns ...func(*Options)) (*ListOrganizationPortfolioAccessOutput, error) {
	if params == nil {
		params = &ListOrganizationPortfolioAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationPortfolioAccess", params, optFns, addOperationListOrganizationPortfolioAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationPortfolioAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationPortfolioAccessInput struct {

	// The organization node type that will be returned in the output.
	//
	// * ORGANIZATION
	// - Organization that has access to the portfolio.
	//
	// * ORGANIZATIONAL_UNIT -
	// Organizational unit that has access to the portfolio within your
	// organization.
	//
	// * ACCOUNT - Account that has access to the portfolio within your
	// organization.
	//
	// This member is required.
	OrganizationNodeType types.OrganizationNodeType

	// The portfolio identifier. For example, port-2abcdext3y5fk.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
}

type ListOrganizationPortfolioAccessOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Displays information about the organization nodes.
	OrganizationNodes []*types.OrganizationNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOrganizationPortfolioAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOrganizationPortfolioAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOrganizationPortfolioAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListOrganizationPortfolioAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationPortfolioAccess(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListOrganizationPortfolioAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListOrganizationPortfolioAccess",
	}
}
