// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Validates a specified template. AWS CloudFormation first checks if the template
// is valid JSON. If it isn't, AWS CloudFormation checks if the template is valid
// YAML. If both these checks fail, AWS CloudFormation returns a template
// validation error.
func (c *Client) ValidateTemplate(ctx context.Context, params *ValidateTemplateInput, optFns ...func(*Options)) (*ValidateTemplateOutput, error) {
	stack := middleware.NewStack("ValidateTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpValidateTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidateTemplate(options.Region), middleware.Before)
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
			OperationName: "ValidateTemplate",
			Err:           err,
		}
	}
	out := result.(*ValidateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for ValidateTemplate () action.
type ValidateTemplateInput struct {
	// Location of file containing the template body. The URL must point to a template
	// (max size: 460,800 bytes) that is located in an Amazon S3 bucket. For more
	// information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must pass TemplateURL or
	// TemplateBody. If both are passed, only TemplateBody is used.
	TemplateURL *string
	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. For more information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must pass TemplateURL or
	// TemplateBody. If both are passed, only TemplateBody is used.
	TemplateBody *string
}

// The output for ValidateTemplate () action.
type ValidateTemplateOutput struct {
	// A list of TemplateParameter structures.
	Parameters []*types.TemplateParameter
	// The list of resources that generated the values in the Capabilities response
	// element.
	CapabilitiesReason *string
	// A list of the transforms that are declared in the template.
	DeclaredTransforms []*string
	// The capabilities found within the template. If your template contains IAM
	// resources, you must specify the CAPABILITY_IAM or CAPABILITY_NAMED_IAM value for
	// this parameter when you use the CreateStack () or UpdateStack () actions with
	// your template; otherwise, those actions return an InsufficientCapabilities
	// error. For more information, see Acknowledging IAM Resources in AWS
	// CloudFormation Templates
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	Capabilities []types.Capability
	// The description found within the template.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpValidateTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpValidateTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpValidateTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opValidateTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ValidateTemplate",
	}
}
