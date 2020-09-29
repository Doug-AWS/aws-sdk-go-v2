// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified budget with the specified resource.
func (c *Client) AssociateBudgetWithResource(ctx context.Context, params *AssociateBudgetWithResourceInput, optFns ...func(*Options)) (*AssociateBudgetWithResourceOutput, error) {
	stack := middleware.NewStack("AssociateBudgetWithResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateBudgetWithResourceMiddlewares(stack)
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
	addOpAssociateBudgetWithResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateBudgetWithResource(options.Region), middleware.Before)
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
			OperationName: "AssociateBudgetWithResource",
			Err:           err,
		}
	}
	out := result.(*AssociateBudgetWithResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateBudgetWithResourceInput struct {
	// The name of the budget you want to associate.
	BudgetName *string
	// The resource identifier. Either a portfolio-id or a product-id.
	ResourceId *string
}

type AssociateBudgetWithResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateBudgetWithResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateBudgetWithResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateBudgetWithResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateBudgetWithResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "AssociateBudgetWithResource",
	}
}
