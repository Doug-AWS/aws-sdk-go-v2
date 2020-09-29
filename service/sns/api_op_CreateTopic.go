// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a topic to which notifications can be published. Users can create at
// most 100,000 standard topics (at most 1,000 FIFO topics). For more information,
// see https://aws.amazon.com/sns (http://aws.amazon.com/sns/). This action is
// idempotent, so if the requester already owns a topic with the specified name,
// that topic's ARN is returned without creating a new topic.
func (c *Client) CreateTopic(ctx context.Context, params *CreateTopicInput, optFns ...func(*Options)) (*CreateTopicOutput, error) {
	stack := middleware.NewStack("CreateTopic", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateTopicMiddlewares(stack)
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
	addOpCreateTopicValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTopic(options.Region), middleware.Before)
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
			OperationName: "CreateTopic",
			Err:           err,
		}
	}
	out := result.(*CreateTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreateTopic action.
type CreateTopicInput struct {
	// The list of tags to add to a new topic. To be able to tag a topic on creation,
	// you must have the sns:CreateTopic and sns:TagResource permissions.
	Tags []*types.Tag
	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// CreateTopic action uses:
	//
	//     * DeliveryPolicy – The policy that defines how
	// Amazon SNS retries failed deliveries to HTTP/S endpoints.
	//
	//     * DisplayName –
	// The display name to use for a topic with SMS subscriptions.
	//
	//     * FifoTopic –
	// Set to true to create a FIFO topic.
	//
	//     * Policy – The policy that defines who
	// can access your topic. By default, only the topic owner can publish or subscribe
	// to the topic.
	//
	//     <p>The following attribute applies only to <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html">server-side-encryption</a>:</p>
	// <ul> <li> <p> <code>KmsMasterKeyId</code> – The ID of an AWS-managed customer
	// master key (CMK) for Amazon SNS or a custom CMK. For more information, see <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms">Key
	// Terms</a>. For more examples, see <a
	// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters">KeyId</a>
	// in the <i>AWS Key Management Service API Reference</i>. </p> </li> </ul> <p>The
	// following attribute applies only to FIFO topics:</p> <ul> <li> <p>
	// <code>ContentBasedDeduplication</code> – Enables content-based deduplication.
	// Amazon SNS uses a SHA-256 hash to generate the
	// <code>MessageDeduplicationId</code> using the body of the message (but not the
	// attributes of the message). </p> </li> <li> <p> When
	// <code>ContentBasedDeduplication</code> is in effect, messages with identical
	// content sent within the deduplication interval are treated as duplicates and
	// only one copy of the message is delivered. </p> </li> <li> <p> If the queue has
	// <code>ContentBasedDeduplication</code> set, your
	// <code>MessageDeduplicationId</code> overrides the generated one. </p> </li>
	// </ul>
	Attributes map[string]*string
	// The name of the topic you want to create. Constraints: Topic names must be made
	// up of only uppercase and lowercase ASCII letters, numbers, underscores, and
	// hyphens, and must be between 1 and 256 characters long. For a FIFO
	// (first-in-first-out) topic, the name must end with the .fifo suffix.
	Name *string
}

// Response from CreateTopic action.
type CreateTopicOutput struct {
	// The Amazon Resource Name (ARN) assigned to the created topic.
	TopicArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateTopicMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateTopic{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateTopic{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTopic(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CreateTopic",
	}
}
