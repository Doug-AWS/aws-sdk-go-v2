// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the configuration information about a distribution.
func (c *Client) GetDistributionConfig(ctx context.Context, params *GetDistributionConfigInput, optFns ...func(*Options)) (*GetDistributionConfigOutput, error) {
	stack := middleware.NewStack("GetDistributionConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetDistributionConfigMiddlewares(stack)
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
	addOpGetDistributionConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributionConfig(options.Region), middleware.Before)
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
			OperationName: "GetDistributionConfig",
			Err:           err,
		}
	}
	out := result.(*GetDistributionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to get a distribution configuration.
type GetDistributionConfigInput struct {

	// The distribution's ID. If the ID is empty, an empty distribution configuration
	// is returned.
	//
	// This member is required.
	Id *string
}

// The returned result of the corresponding request.
type GetDistributionConfigOutput struct {

	// The distribution's configuration information.
	DistributionConfig *types.DistributionConfig

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetDistributionConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetDistributionConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetDistributionConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDistributionConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetDistributionConfig",
	}
}
