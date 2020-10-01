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

// Indicates whether the specified AWS resources are compliant. If a resource is
// noncompliant, this action returns the number of AWS Config rules that the
// resource does not comply with. A resource is compliant if it complies with all
// the AWS Config rules that evaluate it. It is noncompliant if it does not comply
// with one or more of these rules. If AWS Config has no current evaluation results
// for the resource, it returns INSUFFICIENT_DATA. This result might indicate one
// of the following conditions about the rules that evaluate the resource:
//
//     *
// AWS Config has never invoked an evaluation for the rule. To check whether it
// has, use the DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime.
//
//     * The rule's AWS
// Lambda function is failing to send evaluation results to AWS Config. Verify that
// the role that you assigned to your configuration recorder includes the
// config:PutEvaluations permission. If the rule is a custom rule, verify that the
// AWS Lambda execution role includes the config:PutEvaluations permission.
//
//     *
// The rule's AWS Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
func (c *Client) DescribeComplianceByResource(ctx context.Context, params *DescribeComplianceByResourceInput, optFns ...func(*Options)) (*DescribeComplianceByResourceOutput, error) {
	stack := middleware.NewStack("DescribeComplianceByResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeComplianceByResourceMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComplianceByResource(options.Region), middleware.Before)
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
			OperationName: "DescribeComplianceByResource",
			Err:           err,
		}
	}
	out := result.(*DescribeComplianceByResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeComplianceByResourceInput struct {

	// Filters the results by compliance. The allowed values are COMPLIANT,
	// NON_COMPLIANT, and INSUFFICIENT_DATA.
	ComplianceTypes []types.ComplianceType

	// The maximum number of evaluation results returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The ID of the AWS resource for which you want compliance information. You can
	// specify only one resource ID. If you specify a resource ID, you must also
	// specify a type for ResourceType.
	ResourceId *string

	// The types of AWS resources for which you want compliance information (for
	// example, AWS::EC2::Instance). For this action, you can specify that the resource
	// type is an AWS account by specifying AWS::::Account.
	ResourceType *string
}

//
type DescribeComplianceByResourceOutput struct {

	// Indicates whether the specified AWS resource complies with all of the AWS Config
	// rules that evaluate it.
	ComplianceByResources []*types.ComplianceByResource

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeComplianceByResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeComplianceByResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeComplianceByResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeComplianceByResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeComplianceByResource",
	}
}
