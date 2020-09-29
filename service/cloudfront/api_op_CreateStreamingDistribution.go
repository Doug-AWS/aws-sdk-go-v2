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

// Creates a new RTMP distribution. An RTMP distribution is similar to a web
// distribution, but an RTMP distribution streams media files using the Adobe
// Real-Time Messaging Protocol (RTMP) instead of serving files using HTTP. To
// create a new distribution, submit a POST request to the CloudFront API
// version/distribution resource. The request body must include a document with a
// StreamingDistributionConfig element. The response echoes the
// StreamingDistributionConfig element and returns other information about the RTMP
// distribution. To get the status of your request, use the GET
// StreamingDistribution API action. When the value of Enabled is true and the
// value of Status is Deployed, your distribution is ready. A distribution usually
// deploys in less than 15 minutes. For more information about web distributions,
// see Working with RTMP Distributions
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-rtmp.html)
// in the Amazon CloudFront Developer Guide. Beginning with the 2012-05-05 version
// of the CloudFront API, we made substantial changes to the format of the XML
// document that you include in the request body when you create or update a web
// distribution or an RTMP distribution, and when you invalidate objects. With
// previous versions of the API, we discovered that it was too easy to accidentally
// delete one or more values for an element that accepts multiple values, for
// example, CNAMEs and trusted signers. Our changes for the 2012-05-05 release are
// intended to prevent these accidental deletions and to notify you when there's a
// mismatch between the number of values you say you're specifying in the Quantity
// element and the number of values specified.
func (c *Client) CreateStreamingDistribution(ctx context.Context, params *CreateStreamingDistributionInput, optFns ...func(*Options)) (*CreateStreamingDistributionOutput, error) {
	stack := middleware.NewStack("CreateStreamingDistribution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateStreamingDistributionMiddlewares(stack)
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
	addOpCreateStreamingDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamingDistribution(options.Region), middleware.Before)
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
			OperationName: "CreateStreamingDistribution",
			Err:           err,
		}
	}
	out := result.(*CreateStreamingDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new streaming distribution.
type CreateStreamingDistributionInput struct {
	// The streaming distribution's configuration information.
	StreamingDistributionConfig *types.StreamingDistributionConfig
}

// The returned result of the corresponding request.
type CreateStreamingDistributionOutput struct {
	// The fully qualified URI of the new streaming distribution resource just created.
	Location *string
	// The streaming distribution's information.
	StreamingDistribution *types.StreamingDistribution
	// The current version of the streaming distribution created.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateStreamingDistributionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateStreamingDistribution{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateStreamingDistribution{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStreamingDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateStreamingDistribution",
	}
}
