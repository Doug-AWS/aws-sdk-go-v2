// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a conformance pack. A conformance pack is a collection of AWS
// Config rules that can be easily deployed in an account and a region and across
// AWS Organization. This API creates a service linked role
// AWSServiceRoleForConfigConforms in your account. The service linked role is
// created only when the role does not exist in your account. You must specify
// either the TemplateS3Uri or the TemplateBody parameter, but not both. If you
// provide both AWS Config uses the TemplateS3Uri parameter and ignores the
// TemplateBody parameter.
func (c *Client) PutConformancePack(ctx context.Context, params *PutConformancePackInput, optFns ...func(*Options)) (*PutConformancePackOutput, error) {
	stack := middleware.NewStack("PutConformancePack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutConformancePackMiddlewares(stack)
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
	addOpPutConformancePackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConformancePack(options.Region), middleware.Before)
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
			OperationName: "PutConformancePack",
			Err:           err,
		}
	}
	out := result.(*PutConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConformancePackInput struct {

	// Name of the conformance pack you want to create.
	//
	// This member is required.
	ConformancePackName *string

	// AWS Config stores intermediate files while processing conformance pack template.
	//
	// This member is required.
	DeliveryS3Bucket *string

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []*types.ConformancePackInputParameter

	// The prefix for the Amazon S3 bucket.
	DeliveryS3KeyPrefix *string

	// A string containing full conformance pack template body. Structure containing
	// the template body with a minimum length of 1 byte and a maximum length of 51,200
	// bytes. You can only use a YAML template with one resource type, that is, config
	// rule and a remediation action.
	TemplateBody *string

	// Location of file containing the template body (s3://bucketname/prefix). The uri
	// must point to the conformance pack template (max size: 300 KB) that is located
	// in an Amazon S3 bucket in the same region as the conformance pack. You must have
	// access to read Amazon S3 bucket.
	TemplateS3Uri *string
}

type PutConformancePackOutput struct {

	// ARN of the conformance pack.
	ConformancePackArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutConformancePackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutConformancePack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConformancePack{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutConformancePack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConformancePack",
	}
}
