// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the position of the specified receipt rule in the receipt rule set. For
// information about managing receipt rules, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rules.html).
// You can execute this operation no more than once per second.
func (c *Client) SetReceiptRulePosition(ctx context.Context, params *SetReceiptRulePositionInput, optFns ...func(*Options)) (*SetReceiptRulePositionOutput, error) {
	stack := middleware.NewStack("SetReceiptRulePosition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetReceiptRulePositionMiddlewares(stack)
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
	addOpSetReceiptRulePositionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetReceiptRulePosition(options.Region), middleware.Before)
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
			OperationName: "SetReceiptRulePosition",
			Err:           err,
		}
	}
	out := result.(*SetReceiptRulePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to set the position of a receipt rule in a receipt rule
// set. You use receipt rule sets to receive email with Amazon SES. For more
// information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type SetReceiptRulePositionInput struct {

	// The name of the receipt rule to reposition.
	//
	// This member is required.
	RuleName *string

	// The name of the receipt rule set that contains the receipt rule to reposition.
	//
	// This member is required.
	RuleSetName *string

	// The name of the receipt rule after which to place the specified receipt rule.
	After *string
}

// An empty element returned on a successful request.
type SetReceiptRulePositionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetReceiptRulePositionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetReceiptRulePosition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetReceiptRulePosition{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetReceiptRulePosition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetReceiptRulePosition",
	}
}
