// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing Cost Category. Changes made to the Cost Category rules will
// be used to categorize the current month’s expenses and future expenses. This
// won’t change categorization for the previous months.
func (c *Client) UpdateCostCategoryDefinition(ctx context.Context, params *UpdateCostCategoryDefinitionInput, optFns ...func(*Options)) (*UpdateCostCategoryDefinitionOutput, error) {
	stack := middleware.NewStack("UpdateCostCategoryDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateCostCategoryDefinitionMiddlewares(stack)
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
	addOpUpdateCostCategoryDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(options.Region), middleware.Before)
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
			OperationName: "UpdateCostCategoryDefinition",
			Err:           err,
		}
	}
	out := result.(*UpdateCostCategoryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCostCategoryDefinitionInput struct {
	// The Expression object used to categorize costs. For more information, see
	// CostCategoryRule
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html).
	Rules []*types.CostCategoryRule
	// The unique identifier for your Cost Category.
	CostCategoryArn *string
	// The rule schema version in this particular Cost Category.
	RuleVersion types.CostCategoryRuleVersion
}

type UpdateCostCategoryDefinitionOutput struct {
	// The Cost Category's effective start date.
	EffectiveStart *string
	// The unique identifier for your Cost Category.
	CostCategoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateCostCategoryDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCostCategoryDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCostCategoryDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "UpdateCostCategoryDefinition",
	}
}
