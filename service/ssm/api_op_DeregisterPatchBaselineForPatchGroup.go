// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a patch group from a patch baseline.
func (c *Client) DeregisterPatchBaselineForPatchGroup(ctx context.Context, params *DeregisterPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*DeregisterPatchBaselineForPatchGroupOutput, error) {
	stack := middleware.NewStack("DeregisterPatchBaselineForPatchGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterPatchBaselineForPatchGroupMiddlewares(stack)
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
	addOpDeregisterPatchBaselineForPatchGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterPatchBaselineForPatchGroup(options.Region), middleware.Before)
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
			OperationName: "DeregisterPatchBaselineForPatchGroup",
			Err:           err,
		}
	}
	out := result.(*DeregisterPatchBaselineForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterPatchBaselineForPatchGroupInput struct {

	// The ID of the patch baseline to deregister the patch group from.
	//
	// This member is required.
	BaselineId *string

	// The name of the patch group that should be deregistered from the patch baseline.
	//
	// This member is required.
	PatchGroup *string
}

type DeregisterPatchBaselineForPatchGroupOutput struct {

	// The ID of the patch baseline the patch group was deregistered from.
	BaselineId *string

	// The name of the patch group deregistered from the patch baseline.
	PatchGroup *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterPatchBaselineForPatchGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterPatchBaselineForPatchGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterPatchBaselineForPatchGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterPatchBaselineForPatchGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeregisterPatchBaselineForPatchGroup",
	}
}
