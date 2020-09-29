// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a reusable delegation set. You can delete a reusable delegation set only
// if it isn't associated with any hosted zones. To verify that the reusable
// delegation set is not associated with any hosted zones, submit a
// GetReusableDelegationSet
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetReusableDelegationSet.html)
// request and specify the ID of the reusable delegation set that you want to
// delete.
func (c *Client) DeleteReusableDelegationSet(ctx context.Context, params *DeleteReusableDelegationSetInput, optFns ...func(*Options)) (*DeleteReusableDelegationSetOutput, error) {
	stack := middleware.NewStack("DeleteReusableDelegationSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteReusableDelegationSetMiddlewares(stack)
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
	addOpDeleteReusableDelegationSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReusableDelegationSet(options.Region), middleware.Before)
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
			OperationName: "DeleteReusableDelegationSet",
			Err:           err,
		}
	}
	out := result.(*DeleteReusableDelegationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a reusable delegation set.
type DeleteReusableDelegationSetInput struct {
	// The ID of the reusable delegation set that you want to delete.
	Id *string
}

// An empty element.
type DeleteReusableDelegationSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteReusableDelegationSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteReusableDelegationSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteReusableDelegationSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteReusableDelegationSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteReusableDelegationSet",
	}
}
