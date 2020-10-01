// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets revisions of the specified workflow. Only the last 100 revisions are
// stored. If the workflow has been deprecated, this action will return revisions
// that occurred before the deprecation. This action won't work for workflows that
// have been deleted.
func (c *Client) GetFlowTemplateRevisions(ctx context.Context, params *GetFlowTemplateRevisionsInput, optFns ...func(*Options)) (*GetFlowTemplateRevisionsOutput, error) {
	stack := middleware.NewStack("GetFlowTemplateRevisions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFlowTemplateRevisionsMiddlewares(stack)
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
	addOpGetFlowTemplateRevisionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFlowTemplateRevisions(options.Region), middleware.Before)
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
			OperationName: "GetFlowTemplateRevisions",
			Err:           err,
		}
	}
	out := result.(*GetFlowTemplateRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFlowTemplateRevisionsInput struct {

	// The ID of the workflow. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:workflow:WORKFLOWNAME
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type GetFlowTemplateRevisionsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that provide summary data about each revision.
	Summaries []*types.FlowTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFlowTemplateRevisionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFlowTemplateRevisions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFlowTemplateRevisions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFlowTemplateRevisions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetFlowTemplateRevisions",
	}
}
