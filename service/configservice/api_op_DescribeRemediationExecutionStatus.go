// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a detailed view of a Remediation Execution for a set of resources
// including state, timestamps for when steps for the remediation execution occur,
// and any error messages for steps that have failed. When you specify the limit
// and the next token, you receive a paginated response.
func (c *Client) DescribeRemediationExecutionStatus(ctx context.Context, params *DescribeRemediationExecutionStatusInput, optFns ...func(*Options)) (*DescribeRemediationExecutionStatusOutput, error) {
	stack := middleware.NewStack("DescribeRemediationExecutionStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeRemediationExecutionStatusMiddlewares(stack)
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
	addOpDescribeRemediationExecutionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRemediationExecutionStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeRemediationExecutionStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeRemediationExecutionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRemediationExecutionStatusInput struct {
	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	ResourceKeys []*types.ResourceKey
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
	// The maximum number of RemediationExecutionStatuses returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit *int32
	// A list of AWS Config rule names.
	ConfigRuleName *string
}

type DescribeRemediationExecutionStatusOutput struct {
	// Returns a list of remediation execution statuses objects.
	RemediationExecutionStatuses []*types.RemediationExecutionStatus
	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeRemediationExecutionStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRemediationExecutionStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRemediationExecutionStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRemediationExecutionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeRemediationExecutionStatus",
	}
}
