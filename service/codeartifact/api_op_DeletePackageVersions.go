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

// Deletes one or more versions of a package. A deleted package version cannot be
// restored in your repository. If you want to remove a package version from your
// repository and be able to restore it later, set its status to Archived. Archived
// packages cannot be downloaded from a repository and don't show up with list
// package APIs (for example, ListackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)),
// but you can restore them using UpdatePackageVersionsStatus
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html).
func (c *Client) DeletePackageVersions(ctx context.Context, params *DeletePackageVersionsInput, optFns ...func(*Options)) (*DeletePackageVersionsOutput, error) {
	stack := middleware.NewStack("DeletePackageVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeletePackageVersionsMiddlewares(stack)
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
	addOpDeletePackageVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePackageVersions(options.Region), middleware.Before)
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
			OperationName: "DeletePackageVersions",
			Err:           err,
		}
	}
	out := result.(*DeletePackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePackageVersionsInput struct {

	// The name of the domain that contains the package to delete.
	//
	// This member is required.
	Domain *string

	// The format of the package versions to delete. The valid values are:
	//
	//     * npm
	//
	//
	// * pypi
	//
	//     * maven
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package with the versions to delete.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package versions to delete.
	//
	// This member is required.
	Repository *string

	// An array of strings that specify the versions of the package to delete.
	//
	// This member is required.
	Versions []*string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The expected status of the package version to delete. Valid values are:
	//
	//     *
	// Published
	//
	//     * Unfinished
	//
	//     * Unlisted
	//
	//     * Archived
	//
	//     * Disposed
	ExpectedStatus types.PackageVersionStatus

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string
}

type DeletePackageVersionsOutput struct {

	// A PackageVersionError object that contains a map of errors codes for the deleted
	// package that failed. The possible error codes are:
	//
	//     * ALREADY_EXISTS
	//
	//     *
	// MISMATCHED_REVISION
	//
	//     * MISMATCHED_STATUS
	//
	//     * NOT_ALLOWED
	//
	//     *
	// NOT_FOUND
	//
	//     * SKIPPED
	FailedVersions map[string]*types.PackageVersionError

	// A list of the package versions that were successfully deleted.
	SuccessfulVersions map[string]*types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeletePackageVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeletePackageVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePackageVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePackageVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeletePackageVersions",
	}
}
