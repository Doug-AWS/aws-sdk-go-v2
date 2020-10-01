// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a list of steps for the cluster in reverse order unless you specify
// stepIds with the request of filter by StepStates. You can specify a maximum of
// ten stepIDs.
func (c *Client) ListSteps(ctx context.Context, params *ListStepsInput, optFns ...func(*Options)) (*ListStepsOutput, error) {
	stack := middleware.NewStack("ListSteps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListStepsMiddlewares(stack)
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
	addOpListStepsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSteps(options.Region), middleware.Before)
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
			OperationName: "ListSteps",
			Err:           err,
		}
	}
	out := result.(*ListStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which steps to list.
type ListStepsInput struct {

	// The identifier of the cluster for which to list the steps.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// The filter to limit the step list based on the identifier of the steps. You can
	// specify a maximum of ten Step IDs. The character constraint applies to the
	// overall length of the array.
	StepIds []*string

	// The filter to limit the step list based on certain states.
	StepStates []types.StepState
}

// This output contains the list of steps returned in reverse order. This means
// that the last step is the first element in the list.
type ListStepsOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// The filtered list of steps for the cluster.
	Steps []*types.StepSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListStepsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSteps{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSteps{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSteps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListSteps",
	}
}
