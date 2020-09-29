// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the configuration for delivering audit streams to the customer
// account.
func (c *Client) DescribeAuditStreamConfiguration(ctx context.Context, params *DescribeAuditStreamConfigurationInput, optFns ...func(*Options)) (*DescribeAuditStreamConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeAuditStreamConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAuditStreamConfigurationMiddlewares(stack)
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
	addOpDescribeAuditStreamConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAuditStreamConfiguration(options.Region), middleware.Before)
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
			OperationName: "DescribeAuditStreamConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeAuditStreamConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAuditStreamConfigurationInput struct {
	// The ARN of the fleet.
	FleetArn *string
}

type DescribeAuditStreamConfigurationOutput struct {
	// The ARN of the Amazon Kinesis data stream that will receive the audit events.
	AuditStreamArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAuditStreamConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAuditStreamConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAuditStreamConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAuditStreamConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeAuditStreamConfiguration",
	}
}
