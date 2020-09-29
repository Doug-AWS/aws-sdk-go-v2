// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a tag that is assigned to a GameLift resource. Resource tags are used to
// organize AWS resources for a range of purposes. This action handles the
// permissions necessary to manage tags for the following GameLift resource
// types:
//
//     * Build
//
//     * Script
//
//     * Fleet
//
//     * Alias
//
//     *
// GameSessionQueue
//
//     * MatchmakingConfiguration
//
//     * MatchmakingRuleSet
//
// To
// remove a tag from a resource, specify the unique ARN value for the resource and
// provide a string list containing one or more tags to be removed. This action
// succeeds even if the list includes tags that are not currently assigned to the
// specified resource. Learn more Tagging AWS Resources
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
// General Reference  AWS Tagging Strategies
// (http://aws.amazon.com/answers/account-management/aws-tagging-strategies/)
// Related operations
//
//     * TagResource ()
//
//     * UntagResource ()
//
//     *
// ListTagsForResource ()
func (c *Client) UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) {
	stack := middleware.NewStack("UntagResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUntagResourceMiddlewares(stack)
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
	addOpUntagResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagResource(options.Region), middleware.Before)
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
			OperationName: "UntagResource",
			Err:           err,
		}
	}
	out := result.(*UntagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagResourceInput struct {
	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to and uniquely identifies the GameLift resource that you want to
	// remove tags from. GameLift resource ARNs are included in the data object for the
	// resource, which can be retrieved by calling a List or Describe action for the
	// resource type.
	ResourceARN *string
	// A list of one or more tag keys to remove from the specified GameLift resource.
	// An AWS resource can have only one tag with a specific tag key, so specifying the
	// tag key identifies which tag to remove.
	TagKeys []*string
}

type UntagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUntagResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUntagResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUntagResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opUntagResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UntagResource",
	}
}
