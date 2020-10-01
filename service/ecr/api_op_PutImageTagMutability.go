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

// Updates the image tag mutability settings for the specified repository. For more
// information, see Image Tag Mutability
// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-tag-mutability.html)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) PutImageTagMutability(ctx context.Context, params *PutImageTagMutabilityInput, optFns ...func(*Options)) (*PutImageTagMutabilityOutput, error) {
	stack := middleware.NewStack("PutImageTagMutability", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutImageTagMutabilityMiddlewares(stack)
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
	addOpPutImageTagMutabilityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutImageTagMutability(options.Region), middleware.Before)
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
			OperationName: "PutImageTagMutability",
			Err:           err,
		}
	}
	out := result.(*PutImageTagMutabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutImageTagMutabilityInput struct {

	// The tag mutability setting for the repository. If MUTABLE is specified, image
	// tags can be overwritten. If IMMUTABLE is specified, all image tags within the
	// repository will be immutable which will prevent them from being overwritten.
	//
	// This member is required.
	ImageTagMutability types.ImageTagMutability

	// The name of the repository in which to update the image tag mutability settings.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the repository in
	// which to update the image tag mutability settings. If you do not specify a
	// registry, the default registry is assumed.
	RegistryId *string
}

type PutImageTagMutabilityOutput struct {

	// The image tag mutability setting for the repository.
	ImageTagMutability types.ImageTagMutability

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutImageTagMutabilityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutImageTagMutability{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutImageTagMutability{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutImageTagMutability(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutImageTagMutability",
	}
}
