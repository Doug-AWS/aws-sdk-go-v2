// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a sample mapping template that can be used to transform a payload into
// the structure of a model.
func (c *Client) GetModelTemplate(ctx context.Context, params *GetModelTemplateInput, optFns ...func(*Options)) (*GetModelTemplateOutput, error) {
	stack := middleware.NewStack("GetModelTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetModelTemplateMiddlewares(stack)
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
	addOpGetModelTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetModelTemplate",
			Err:           err,
		}
	}
	out := result.(*GetModelTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to generate a sample mapping template used to transform the payload.
type GetModelTemplateInput struct {

	// [Required] The name of the model for which to generate a template.
	//
	// This member is required.
	ModelName *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a mapping template used to transform a payload. Mapping Templates
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html#models-mappings-mappings)
type GetModelTemplateOutput struct {

	// The Apache Velocity Template Language (VTL)
	// (https://velocity.apache.org/engine/devel/vtl-reference.html) template content
	// used for the template resource.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetModelTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetModelTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModelTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetModelTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetModelTemplate",
	}
}
