// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes Region-specific settings for a global table. This method only applies
// to Version 2017.11.29
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
// of global tables.
func (c *Client) DescribeGlobalTableSettings(ctx context.Context, params *DescribeGlobalTableSettingsInput, optFns ...func(*Options)) (*DescribeGlobalTableSettingsOutput, error) {
	stack := middleware.NewStack("DescribeGlobalTableSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeGlobalTableSettingsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addOpDescribeGlobalTableSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalTableSettings(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "DescribeGlobalTableSettings",
			Err:           err,
		}
	}
	out := result.(*DescribeGlobalTableSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalTableSettingsInput struct {
	// The name of the global table to describe.
	GlobalTableName *string
}

type DescribeGlobalTableSettingsOutput struct {
	// The name of the global table.
	GlobalTableName *string
	// The Region-specific settings for the global table.
	ReplicaSettings []*types.ReplicaSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeGlobalTableSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeGlobalTableSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeGlobalTableSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGlobalTableSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "DescribeGlobalTableSettings",
	}
}
