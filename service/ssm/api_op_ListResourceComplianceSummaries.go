// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a resource-level summary count. The summary includes information about
// compliant and non-compliant statuses and detailed compliance-item severity
// counts, according to the filter criteria you specify.
func (c *Client) ListResourceComplianceSummaries(ctx context.Context, params *ListResourceComplianceSummariesInput, optFns ...func(*Options)) (*ListResourceComplianceSummariesOutput, error) {
	stack := middleware.NewStack("ListResourceComplianceSummaries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListResourceComplianceSummariesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceComplianceSummaries(options.Region), middleware.Before)
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
			OperationName: "ListResourceComplianceSummaries",
			Err:           err,
		}
	}
	out := result.(*ListResourceComplianceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceComplianceSummariesInput struct {

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []*types.ComplianceStringFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type ListResourceComplianceSummariesOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A summary count for specified or targeted managed instances. Summary count
	// includes information about compliant and non-compliant State Manager
	// associations, patch status, or custom items according to the filter criteria
	// that you specify.
	ResourceComplianceSummaryItems []*types.ResourceComplianceSummaryItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListResourceComplianceSummariesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceComplianceSummaries{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceComplianceSummaries{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResourceComplianceSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListResourceComplianceSummaries",
	}
}
