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

// Returns the template body for a specified stack. You can get the template for
// running or deleted stacks. For deleted stacks, GetTemplate returns the template
// for up to 90 days after the stack has been deleted. If the template does not
// exist, a ValidationError is returned.
func (c *Client) GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) {
	stack := middleware.NewStack("GetTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplate(options.Region), middleware.Before)
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
			OperationName: "GetTemplate",
			Err:           err,
		}
	}
	out := result.(*GetTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for a GetTemplate () action.
type GetTemplateInput struct {

	// The name or Amazon Resource Name (ARN) of a change set for which AWS
	// CloudFormation returns the associated template. If you specify a name, you must
	// also specify the StackName.
	ChangeSetName *string

	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	//     * Running stacks: You can specify either the
	// stack's name or its unique stack ID.
	//
	//     * Deleted stacks: You must specify the
	// unique stack ID.
	//
	// Default: There is no default value.
	StackName *string

	// For templates that include transforms, the stage of the template that AWS
	// CloudFormation returns. To get the user-submitted template, specify Original. To
	// get the template after AWS CloudFormation has processed all transforms, specify
	// Processed. If the template doesn't include transforms, Original and Processed
	// return the same template. By default, AWS CloudFormation specifies Original.
	TemplateStage types.TemplateStage
}

// The output for GetTemplate () action.
type GetTemplateOutput struct {

	// The stage of the template that you can retrieve. For stacks, the Original and
	// Processed templates are always available. For change sets, the Original template
	// is always available. After AWS CloudFormation finishes creating the change set,
	// the Processed template becomes available.
	StagesAvailable []types.TemplateStage

	// Structure containing the template body. (For more information, go to Template
	// Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.) AWS CloudFormation returns the same
	// template that was used when the stack was created.
	TemplateBody *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "GetTemplate",
	}
}
