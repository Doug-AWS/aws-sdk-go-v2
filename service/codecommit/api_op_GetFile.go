// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the base-64 encoded contents of a specified file and its metadata.
func (c *Client) GetFile(ctx context.Context, params *GetFileInput, optFns ...func(*Options)) (*GetFileOutput, error) {
	stack := middleware.NewStack("GetFile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFileMiddlewares(stack)
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
	addOpGetFileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFile(options.Region), middleware.Before)
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
			OperationName: "GetFile",
			Err:           err,
		}
	}
	out := result.(*GetFileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFileInput struct {

	// The fully qualified path to the file, including the full name and extension of
	// the file. For example, /examples/file.md is the fully qualified path to a file
	// named file.md in a folder named examples.
	//
	// This member is required.
	FilePath *string

	// The name of the repository that contains the file.
	//
	// This member is required.
	RepositoryName *string

	// The fully quaified reference that identifies the commit that contains the file.
	// For example, you can specify a full commit ID, a tag, a branch name, or a
	// reference such as refs/heads/master. If none is provided, the head commit is
	// used.
	CommitSpecifier *string
}

type GetFileOutput struct {

	// The blob ID of the object that represents the file content.
	//
	// This member is required.
	BlobId *string

	// The full commit ID of the commit that contains the content returned by GetFile.
	//
	// This member is required.
	CommitId *string

	// The base-64 encoded binary data object that represents the content of the file.
	//
	// This member is required.
	FileContent []byte

	// The extrapolated file mode permissions of the blob. Valid values include strings
	// such as EXECUTABLE and not numeric values. The file mode permissions returned by
	// this API are not the standard file mode permission values, such as 100644, but
	// rather extrapolated values. See the supported return values.
	//
	// This member is required.
	FileMode types.FileModeTypeEnum

	// The fully qualified path to the specified file. Returns the name and extension
	// of the file.
	//
	// This member is required.
	FilePath *string

	// The size of the contents of the file, in bytes.
	//
	// This member is required.
	FileSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFile{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetFile",
	}
}
