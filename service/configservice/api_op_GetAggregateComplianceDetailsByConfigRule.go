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

// Returns the evaluation results for the specified AWS Config rule for a specific
// resource in a rule. The results indicate which AWS resources were evaluated by
// the rule, when each resource was last evaluated, and whether each resource
// complies with the rule. The results can return an empty result page. But if you
// have a nextToken, the results are displayed on the next page.
func (c *Client) GetAggregateComplianceDetailsByConfigRule(ctx context.Context, params *GetAggregateComplianceDetailsByConfigRuleInput, optFns ...func(*Options)) (*GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	stack := middleware.NewStack("GetAggregateComplianceDetailsByConfigRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAggregateComplianceDetailsByConfigRuleMiddlewares(stack)
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
	addOpGetAggregateComplianceDetailsByConfigRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAggregateComplianceDetailsByConfigRule(options.Region), middleware.Before)
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
			OperationName: "GetAggregateComplianceDetailsByConfigRule",
			Err:           err,
		}
	}
	out := result.(*GetAggregateComplianceDetailsByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAggregateComplianceDetailsByConfigRuleInput struct {

	// The 12-digit account ID of the source account.
	//
	// This member is required.
	AccountId *string

	// The source region from where the data is aggregated.
	//
	// This member is required.
	AwsRegion *string

	// The name of the AWS Config rule for which you want compliance information.
	//
	// This member is required.
	ConfigRuleName *string

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The resource compliance status. For the
	// GetAggregateComplianceDetailsByConfigRuleRequest data type, AWS Config supports
	// only the COMPLIANT and NON_COMPLIANT. AWS Config does not support the
	// NOT_APPLICABLE and INSUFFICIENT_DATA values.
	ComplianceType types.ComplianceType

	// The maximum number of evaluation results returned on each page. The default is
	// 50. You cannot specify a number greater than 100. If you specify 0, AWS Config
	// uses the default.
	Limit *int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

type GetAggregateComplianceDetailsByConfigRuleOutput struct {

	// Returns an AggregateEvaluationResults object.
	AggregateEvaluationResults []*types.AggregateEvaluationResult

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAggregateComplianceDetailsByConfigRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAggregateComplianceDetailsByConfigRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAggregateComplianceDetailsByConfigRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAggregateComplianceDetailsByConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetAggregateComplianceDetailsByConfigRule",
	}
}
