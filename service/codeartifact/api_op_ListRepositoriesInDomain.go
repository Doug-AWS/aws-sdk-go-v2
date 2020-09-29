// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of RepositorySummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html)
// objects. Each RepositorySummary contains information about a repository in the
// specified domain and that matches the input parameters.
func (c *Client) ListRepositoriesInDomain(ctx context.Context, params *ListRepositoriesInDomainInput, optFns ...func(*Options)) (*ListRepositoriesInDomainOutput, error) {
	stack := middleware.NewStack("ListRepositoriesInDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListRepositoriesInDomainMiddlewares(stack)
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
	addOpListRepositoriesInDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositoriesInDomain(options.Region), middleware.Before)
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
			OperationName: "ListRepositoriesInDomain",
			Err:           err,
		}
	}
	out := result.(*ListRepositoriesInDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoriesInDomainInput struct {
	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
	// Filter the list of repositories to only include those that are managed by the
	// AWS account ID.
	AdministratorAccount *string
	// The maximum number of results to return per page.
	MaxResults *int32
	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
	// The name of the domain that contains the returned list of repositories.
	Domain *string
	// A prefix used to filter returned repositories. Only repositories with names that
	// start with repositoryPrefix are returned.
	RepositoryPrefix *string
}

type ListRepositoriesInDomainOutput struct {
	// If there are additional results, this is the token for the next set of results.
	NextToken *string
	// The returned list of repositories.
	Repositories []*types.RepositorySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListRepositoriesInDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListRepositoriesInDomain{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListRepositoriesInDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRepositoriesInDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListRepositoriesInDomain",
	}
}
