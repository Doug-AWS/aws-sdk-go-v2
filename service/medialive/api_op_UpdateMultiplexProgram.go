// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update a program in a multiplex.
func (c *Client) UpdateMultiplexProgram(ctx context.Context, params *UpdateMultiplexProgramInput, optFns ...func(*Options)) (*UpdateMultiplexProgramOutput, error) {
	stack := middleware.NewStack("UpdateMultiplexProgram", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateMultiplexProgramMiddlewares(stack)
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
	addOpUpdateMultiplexProgramValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMultiplexProgram(options.Region), middleware.Before)
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
			OperationName: "UpdateMultiplexProgram",
			Err:           err,
		}
	}
	out := result.(*UpdateMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a program in a multiplex.
type UpdateMultiplexProgramInput struct {

	// The ID of the multiplex of the program to update.
	//
	// This member is required.
	MultiplexId *string

	// The name of the program to update.
	//
	// This member is required.
	ProgramName *string

	// The new settings for a multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings
}

// Placeholder documentation for UpdateMultiplexProgramResponse
type UpdateMultiplexProgramOutput struct {

	// The updated multiplex program.
	MultiplexProgram *types.MultiplexProgram

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateMultiplexProgramMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMultiplexProgram{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMultiplexProgram{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMultiplexProgram(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateMultiplexProgram",
	}
}
