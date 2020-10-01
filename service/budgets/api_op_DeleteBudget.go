// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a budget. You can delete your budget at any time. Deleting a budget also
// deletes the notifications and subscribers that are associated with that budget.
func (c *Client) DeleteBudget(ctx context.Context, params *DeleteBudgetInput, optFns ...func(*Options)) (*DeleteBudgetOutput, error) {
	stack := middleware.NewStack("DeleteBudget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteBudgetMiddlewares(stack)
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
	addOpDeleteBudgetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBudget(options.Region), middleware.Before)
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
			OperationName: "DeleteBudget",
			Err:           err,
		}
	}
	out := result.(*DeleteBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DeleteBudget
type DeleteBudgetInput struct {

	// The accountId that is associated with the budget that you want to delete.
	//
	// This member is required.
	AccountId *string

	// The name of the budget that you want to delete.
	//
	// This member is required.
	BudgetName *string
}

// Response of DeleteBudget
type DeleteBudgetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteBudgetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBudget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBudget{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBudget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DeleteBudget",
	}
}
