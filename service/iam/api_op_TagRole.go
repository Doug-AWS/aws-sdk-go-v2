// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more tags to an IAM role. The role can be a regular role or a
// service-linked role. If a tag with the same key name already exists, then that
// tag is overwritten with the new value. A tag consists of a key name and an
// associated value. By assigning tags to your resources, you can do the
// following:
//
//     * Administrative grouping and discovery - Attach tags to
// resources to aid in organization and search. For example, you could search for
// all resources with the key name Project and the value MyImportantProject. Or
// search for all resources with the key name Cost Center and the value 41200.
//
//
// * Access control - Reference tags in IAM user-based and resource-based policies.
// You can use tags to restrict access to only an IAM user or role that has a
// specified tag attached. You can also restrict access to only those resources
// that have a certain tag attached. For examples of policies that show how to use
// tags to control access, see Control Access Using IAM Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html) in the IAM
// User Guide.
//
//     * Cost allocation - Use tags to help track which individuals
// and teams are using which AWS resources.
//
//     * Make sure that you have no
// invalid tags and that you do not exceed the allowed number of tags per role. In
// either case, the entire request fails and no tags are added to the role.
//
//     *
// AWS always interprets the tag Value as a single string. If you need to store an
// array, you can store comma-separated values in the string. However, you must
// interpret the value in your code.
//
// For more information about tagging, see
// Tagging IAM Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
// Guide.
func (c *Client) TagRole(ctx context.Context, params *TagRoleInput, optFns ...func(*Options)) (*TagRoleOutput, error) {
	stack := middleware.NewStack("TagRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpTagRoleMiddlewares(stack)
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
	addOpTagRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagRole(options.Region), middleware.Before)
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
			OperationName: "TagRole",
			Err:           err,
		}
	}
	out := result.(*TagRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagRoleInput struct {
	// The list of tags that you want to attach to the role. Each tag consists of a key
	// name and an associated value. You can specify this with a JSON string.
	Tags []*types.Tag
	// The name of the role that you want to add tags to. This parameter accepts
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that consist of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	RoleName *string
}

type TagRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpTagRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpTagRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpTagRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "TagRole",
	}
}
