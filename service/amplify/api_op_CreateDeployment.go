// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a deployment for a manually deployed Amplify app. Manually deployed apps
// are not connected to a repository.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	stack := middleware.NewStack("CreateDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDeploymentMiddlewares(stack)
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
	addOpCreateDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before)
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
			OperationName: "CreateDeployment",
			Err:           err,
		}
	}
	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create a new deployment request.
type CreateDeploymentInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the branch, for the job.
	//
	// This member is required.
	BranchName *string

	// An optional file map that contains the file name as the key and the file content
	// md5 hash as the value. If this argument is provided, the service will generate a
	// unique upload URL per file. Otherwise, the service will only generate a single
	// upload URL for the zipped files.
	FileMap map[string]*string
}

// The result structure for the create a new deployment request.
type CreateDeploymentOutput struct {

	// When the fileMap argument is provided in the request, fileUploadUrls will
	// contain a map of file names to upload URLs.
	//
	// This member is required.
	FileUploadUrls map[string]*string

	// When the fileMap argument is not provided in the request, this zipUploadUrl is
	// returned.
	//
	// This member is required.
	ZipUploadUrl *string

	// The job ID for this deployment. will supply to start deployment api.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateDeployment",
	}
}
