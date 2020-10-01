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

// Describes the available directories that are registered with Amazon WorkSpaces.
func (c *Client) DescribeWorkspaceDirectories(ctx context.Context, params *DescribeWorkspaceDirectoriesInput, optFns ...func(*Options)) (*DescribeWorkspaceDirectoriesOutput, error) {
	stack := middleware.NewStack("DescribeWorkspaceDirectories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeWorkspaceDirectoriesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceDirectories(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkspaceDirectories",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkspaceDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceDirectoriesInput struct {

	// The identifiers of the directories. If the value is null, all directories are
	// retrieved.
	DirectoryIds []*string

	// The maximum number of directories to return.
	Limit *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string
}

type DescribeWorkspaceDirectoriesOutput struct {

	// Information about the directories.
	Directories []*types.WorkspaceDirectory

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeWorkspaceDirectoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceDirectories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceDirectories{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkspaceDirectories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceDirectories",
	}
}
