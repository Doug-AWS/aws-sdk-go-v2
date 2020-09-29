// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the current user's special folders; the RootFolder and the RecycleBin.
// RootFolder is the root of user's files and folders and RecycleBin is the root of
// recycled items. This is not a valid action for SigV4 (administrative API)
// clients. This action requires an authentication token. To get an authentication
// token, register an application with Amazon WorkDocs. For more information, see
// Authentication and Access Control for User Applications
// (https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html)
// in the Amazon WorkDocs Developer Guide.
func (c *Client) DescribeRootFolders(ctx context.Context, params *DescribeRootFoldersInput, optFns ...func(*Options)) (*DescribeRootFoldersOutput, error) {
	stack := middleware.NewStack("DescribeRootFolders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRootFoldersMiddlewares(stack)
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
	addOpDescribeRootFoldersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRootFolders(options.Region), middleware.Before)
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
			OperationName: "DescribeRootFolders",
			Err:           err,
		}
	}
	out := result.(*DescribeRootFoldersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRootFoldersInput struct {
	// The maximum number of items to return.
	Limit *int32
	// Amazon WorkDocs authentication token.
	AuthenticationToken *string
	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string
}

type DescribeRootFoldersOutput struct {
	// The marker for the next set of results.
	Marker *string
	// The user's special folders.
	Folders []*types.FolderMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRootFoldersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRootFolders{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRootFolders{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRootFolders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeRootFolders",
	}
}
