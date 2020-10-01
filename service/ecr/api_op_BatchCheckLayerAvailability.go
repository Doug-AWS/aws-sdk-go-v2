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

// Checks the availability of one or more image layers in a repository. When an
// image is pushed to a repository, each image layer is checked to verify if it has
// been uploaded before. If it has been uploaded, then the image layer is skipped.
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func (c *Client) BatchCheckLayerAvailability(ctx context.Context, params *BatchCheckLayerAvailabilityInput, optFns ...func(*Options)) (*BatchCheckLayerAvailabilityOutput, error) {
	stack := middleware.NewStack("BatchCheckLayerAvailability", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchCheckLayerAvailabilityMiddlewares(stack)
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
	addOpBatchCheckLayerAvailabilityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCheckLayerAvailability(options.Region), middleware.Before)
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
			OperationName: "BatchCheckLayerAvailability",
			Err:           err,
		}
	}
	out := result.(*BatchCheckLayerAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCheckLayerAvailabilityInput struct {

	// The digests of the image layers to check.
	//
	// This member is required.
	LayerDigests []*string

	// The name of the repository that is associated with the image layers to check.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the image layers
	// to check. If you do not specify a registry, the default registry is assumed.
	RegistryId *string
}

type BatchCheckLayerAvailabilityOutput struct {

	// Any failures associated with the call.
	Failures []*types.LayerFailure

	// A list of image layer objects corresponding to the image layer references in the
	// request.
	Layers []*types.Layer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchCheckLayerAvailabilityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchCheckLayerAvailability{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchCheckLayerAvailability{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchCheckLayerAvailability(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "BatchCheckLayerAvailability",
	}
}
