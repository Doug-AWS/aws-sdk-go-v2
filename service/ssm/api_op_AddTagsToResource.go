// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or overwrites one or more tags for the specified resource. Tags are
// metadata that you can assign to your documents, managed instances, maintenance
// windows, Parameter Store parameters, and patch baselines. Tags enable you to
// categorize your resources in different ways, for example, by purpose, owner, or
// environment. Each tag consists of a key and an optional value, both of which you
// define. For example, you could define a set of tags for your account's managed
// instances that helps you track each instance's owner and stack level. For
// example: Key=Owner and Value=DbAdmin, SysAdmin, or Dev. Or Key=Stack and
// Value=Production, Pre-Production, or Test.  <p>Each resource can have a maximum
// of 50 tags. </p> <p>We recommend that you devise a set of tag keys that meets
// your needs for each resource type.  Using a consistent set of tag keys makes it
// easier for you to manage your resources. You can search and filter the resources
// based on the tags you add. Tags don't have any semantic meaning to and are
// interpreted strictly as a string of characters. For more information about using
// tags with EC2 instances, see Tagging your Amazon EC2 resources
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
// Amazon EC2 User Guide.
func (c *Client) AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) {
	stack := middleware.NewStack("AddTagsToResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddTagsToResourceMiddlewares(stack)
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
	addOpAddTagsToResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToResource(options.Region), middleware.Before)
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
			OperationName: "AddTagsToResource",
			Err:           err,
		}
	}
	out := result.(*AddTagsToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddTagsToResourceInput struct {

	// The resource ID you want to tag. Use the ID of the resource. Here are some
	// examples: ManagedInstance: mi-012345abcde MaintenanceWindow: mw-012345abcde
	// PatchBaseline: pb-012345abcde For the Document and Parameter values, use the
	// name of the resource. The ManagedInstance type for this API action is only for
	// on-premises managed instances. You must specify the name of the managed instance
	// in the following format: mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// This member is required.
	ResourceId *string

	// Specifies the type of resource you are tagging. The ManagedInstance type for
	// this API action is for on-premises managed instances. You must specify the name
	// of the managed instance in the following format: mi-ID_number. For example,
	// mi-1a2b3c4d5e6f.
	//
	// This member is required.
	ResourceType types.ResourceTypeForTagging

	// One or more tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value to an
	// empty string. Do not enter personally identifiable information in this field.
	//
	// This member is required.
	Tags []*types.Tag
}

type AddTagsToResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddTagsToResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTagsToResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "AddTagsToResource",
	}
}
