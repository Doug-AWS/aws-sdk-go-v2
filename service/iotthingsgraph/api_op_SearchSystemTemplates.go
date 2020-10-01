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

// Searches for summary information about systems in the user's account. You can
// filter by the ID of a workflow to return only systems that use the specified
// workflow.
func (c *Client) SearchSystemTemplates(ctx context.Context, params *SearchSystemTemplatesInput, optFns ...func(*Options)) (*SearchSystemTemplatesOutput, error) {
	stack := middleware.NewStack("SearchSystemTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchSystemTemplatesMiddlewares(stack)
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
	addOpSearchSystemTemplatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSystemTemplates(options.Region), middleware.Before)
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
			OperationName: "SearchSystemTemplates",
			Err:           err,
		}
	}
	out := result.(*SearchSystemTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSystemTemplatesInput struct {

	// An array of filters that limit the result set. The only valid filter is
	// FLOW_TEMPLATE_ID.
	Filters []*types.SystemTemplateFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type SearchSystemTemplatesOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary information about each system
	// deployment in the result set.
	Summaries []*types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchSystemTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSystemTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSystemTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchSystemTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchSystemTemplates",
	}
}
