// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the ThreatIntelSets of the GuardDuty service specified by the detector ID.
// If you use this operation from a member account, the ThreatIntelSets associated
// with the master account are returned.
func (c *Client) ListThreatIntelSets(ctx context.Context, params *ListThreatIntelSetsInput, optFns ...func(*Options)) (*ListThreatIntelSetsOutput, error) {
	stack := middleware.NewStack("ListThreatIntelSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThreatIntelSetsMiddlewares(stack)
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
	addOpListThreatIntelSetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThreatIntelSets(options.Region), middleware.Before)
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
			OperationName: "ListThreatIntelSets",
			Err:           err,
		}
	}
	out := result.(*ListThreatIntelSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThreatIntelSetsInput struct {

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// This member is required.
	DetectorId *string

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int32

	// You can use this parameter to paginate results in the response. Set the value of
	// this parameter to null on your first call to the list action. For subsequent
	// calls to the action, fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string
}

type ListThreatIntelSetsOutput struct {

	// The IDs of the ThreatIntelSet resources.
	//
	// This member is required.
	ThreatIntelSetIds []*string

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThreatIntelSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThreatIntelSets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThreatIntelSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThreatIntelSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListThreatIntelSets",
	}
}
