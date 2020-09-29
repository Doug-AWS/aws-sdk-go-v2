// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the custom action targets in Security Hub in your account.
func (c *Client) DescribeActionTargets(ctx context.Context, params *DescribeActionTargetsInput, optFns ...func(*Options)) (*DescribeActionTargetsOutput, error) {
	stack := middleware.NewStack("DescribeActionTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeActionTargetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActionTargets(options.Region), middleware.Before)
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
			OperationName: "DescribeActionTargets",
			Err:           err,
		}
	}
	out := result.(*DescribeActionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActionTargetsInput struct {
	// The maximum number of results to return.
	MaxResults *int32
	// The token that is required for pagination. On your first call to the
	// DescribeActionTargets operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string
	// A list of custom action target ARNs for the custom action targets to retrieve.
	ActionTargetArns []*string
}

type DescribeActionTargetsOutput struct {
	// The pagination token to use to request the next page of results.
	NextToken *string
	// A list of ActionTarget objects. Each object includes the ActionTargetArn,
	// Description, and Name of a custom action target available in Security Hub.
	ActionTargets []*types.ActionTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeActionTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeActionTargets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeActionTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeActionTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeActionTargets",
	}
}
