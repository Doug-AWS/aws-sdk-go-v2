// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a package with an Amazon ES domain.
func (c *Client) AssociatePackage(ctx context.Context, params *AssociatePackageInput, optFns ...func(*Options)) (*AssociatePackageOutput, error) {
	stack := middleware.NewStack("AssociatePackage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociatePackageMiddlewares(stack)
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
	addOpAssociatePackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePackage(options.Region), middleware.Before)
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
			OperationName: "AssociatePackage",
			Err:           err,
		}
	}
	out := result.(*AssociatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to AssociatePackage () operation.
type AssociatePackageInput struct {
	// Name of the domain that you want to associate the package with.
	DomainName *string
	// Internal ID of the package that you want to associate with a domain. Use
	// DescribePackages to find this value.
	PackageID *string
}

// Container for response returned by AssociatePackage () operation.
type AssociatePackageOutput struct {
	// DomainPackageDetails
	DomainPackageDetails *types.DomainPackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociatePackageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociatePackage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociatePackage{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociatePackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "AssociatePackage",
	}
}
