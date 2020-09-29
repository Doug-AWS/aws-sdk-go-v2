// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an activity. This operation is eventually consistent. The results are
// best effort and may not reflect very recent updates and changes.
func (c *Client) DescribeActivity(ctx context.Context, params *DescribeActivityInput, optFns ...func(*Options)) (*DescribeActivityOutput, error) {
	stack := middleware.NewStack("DescribeActivity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeActivityMiddlewares(stack)
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
	addOpDescribeActivityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivity(options.Region), middleware.Before)
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
			OperationName: "DescribeActivity",
			Err:           err,
		}
	}
	out := result.(*DescribeActivityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivityInput struct {
	// The Amazon Resource Name (ARN) of the activity to describe.
	ActivityArn *string
}

type DescribeActivityOutput struct {
	// The date the activity is created.
	CreationDate *time.Time
	// The name of the activity. A name must not contain:
	//
	//     * white space
	//
	//     *
	// brackets < > { } [ ]
	//
	//     * wildcard characters ? *
	//
	//     * special characters "
	// # % \ ^ | ~ ` $ & , ; : /
	//
	//     * control characters (U+0000-001F,
	// U+007F-009F)
	//
	// To enable logging with CloudWatch Logs, the name should only
	// contain 0-9, A-Z, a-z, - and _.
	Name *string
	// The Amazon Resource Name (ARN) that identifies the activity.
	ActivityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeActivityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeActivity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeActivity{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeActivity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeActivity",
	}
}
