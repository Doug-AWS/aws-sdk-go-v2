// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets code to perform a specified mapping.
func (c *Client) GetPlan(ctx context.Context, params *GetPlanInput, optFns ...func(*Options)) (*GetPlanOutput, error) {
	stack := middleware.NewStack("GetPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPlanMiddlewares(stack)
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
	addOpGetPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlan(options.Region), middleware.Before)
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
			OperationName: "GetPlan",
			Err:           err,
		}
	}
	out := result.(*GetPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlanInput struct {

	// The list of mappings from a source table to target tables.
	//
	// This member is required.
	Mapping []*types.MappingEntry

	// The source table.
	//
	// This member is required.
	Source *types.CatalogEntry

	// The programming language of the code to perform the mapping.
	Language types.Language

	// The parameters for the mapping.
	Location *types.Location

	// The target tables.
	Sinks []*types.CatalogEntry
}

type GetPlanOutput struct {

	// A Python script to perform the mapping.
	PythonScript *string

	// The Scala code to perform the mapping.
	ScalaCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPlan{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetPlan",
	}
}
