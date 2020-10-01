// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the image scanning configuration for the specified repository.
func (c *Client) PutImageScanningConfiguration(ctx context.Context, params *PutImageScanningConfigurationInput, optFns ...func(*Options)) (*PutImageScanningConfigurationOutput, error) {
	stack := middleware.NewStack("PutImageScanningConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutImageScanningConfigurationMiddlewares(stack)
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
	addOpPutImageScanningConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutImageScanningConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutImageScanningConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutImageScanningConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImageScanningConfigurationInput struct {

	// The image scanning configuration for the repository. This setting determines
	// whether images are scanned for known vulnerabilities after being pushed to the
	// repository.
	//
	// This member is required.
	ImageScanningConfiguration *types.ImageScanningConfiguration

	// The name of the repository in which to update the image scanning configuration
	// setting.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the repository in
	// which to update the image scanning configuration setting. If you do not specify
	// a registry, the default registry is assumed.
	RegistryId *string
}

type PutImageScanningConfigurationOutput struct {

	// The image scanning configuration setting for the repository.
	ImageScanningConfiguration *types.ImageScanningConfiguration

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutImageScanningConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutImageScanningConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutImageScanningConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutImageScanningConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutImageScanningConfiguration",
	}
}
