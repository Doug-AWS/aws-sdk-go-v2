// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified WorkSpaces. You can filter the results by using the
// bundle identifier, directory identifier, or owner, but you can specify only one
// filter at a time.
func (c *Client) DescribeWorkspaces(ctx context.Context, params *DescribeWorkspacesInput, optFns ...func(*Options)) (*DescribeWorkspacesOutput, error) {
	stack := middleware.NewStack("DescribeWorkspaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeWorkspacesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaces(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkspaces",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspacesInput struct {
	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string
	// The name of the directory user. You must specify this parameter with
	// DirectoryId.
	UserName *string
	// The maximum number of items to return.
	Limit *int32
	// The identifier of the bundle. All WorkSpaces that are created from this bundle
	// are retrieved. You cannot combine this parameter with any other filter.
	BundleId *string
	// The identifier of the directory. In addition, you can optionally specify a
	// specific directory user (see UserName). You cannot combine this parameter with
	// any other filter.
	DirectoryId *string
	// The identifiers of the WorkSpaces. You cannot combine this parameter with any
	// other filter. Because the CreateWorkspaces () operation is asynchronous, the
	// identifier it returns is not immediately available. If you immediately call
	// DescribeWorkspaces () with this identifier, no information is returned.
	WorkspaceIds []*string
}

type DescribeWorkspacesOutput struct {
	// Information about the WorkSpaces. Because CreateWorkspaces () is an asynchronous
	// operation, some of the returned information could be incomplete.
	Workspaces []*types.Workspace
	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeWorkspacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkspaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaces",
	}
}
