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

// Gets information about a specified RTMP distribution, including the distribution
// configuration.
func (c *Client) GetStreamingDistribution(ctx context.Context, params *GetStreamingDistributionInput, optFns ...func(*Options)) (*GetStreamingDistributionOutput, error) {
	stack := middleware.NewStack("GetStreamingDistribution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetStreamingDistributionMiddlewares(stack)
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
	addOpGetStreamingDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamingDistribution(options.Region), middleware.Before)
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
			OperationName: "GetStreamingDistribution",
			Err:           err,
		}
	}
	out := result.(*GetStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to get a streaming distribution's information.
type GetStreamingDistributionInput struct {
	// The streaming distribution's ID.
	Id *string
}

// The returned result of the corresponding request.
type GetStreamingDistributionOutput struct {
	// The streaming distribution's information.
	StreamingDistribution *types.StreamingDistribution
	// The current version of the streaming distribution's information. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetStreamingDistributionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetStreamingDistribution{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetStreamingDistribution{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetStreamingDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetStreamingDistribution",
	}
}
