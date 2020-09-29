// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enable or disable the Deliverability dashboard for your Amazon Pinpoint account.
// When you enable the Deliverability dashboard, you gain access to reputation,
// deliverability, and other metrics for the domains that you use to send email
// using Amazon Pinpoint. You also gain the ability to perform predictive inbox
// placement tests. When you use the Deliverability dashboard, you pay a monthly
// subscription charge, in addition to any other fees that you accrue by using
// Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing
// (http://aws.amazon.com/pinpoint/pricing/).
func (c *Client) PutDeliverabilityDashboardOption(ctx context.Context, params *PutDeliverabilityDashboardOptionInput, optFns ...func(*Options)) (*PutDeliverabilityDashboardOptionOutput, error) {
	stack := middleware.NewStack("PutDeliverabilityDashboardOption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutDeliverabilityDashboardOptionMiddlewares(stack)
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
	addOpPutDeliverabilityDashboardOptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(options.Region), middleware.Before)
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
			OperationName: "PutDeliverabilityDashboardOption",
			Err:           err,
		}
	}
	out := result.(*PutDeliverabilityDashboardOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Enable or disable the Deliverability dashboard for your Amazon Pinpoint account.
// When you enable the Deliverability dashboard, you gain access to reputation,
// deliverability, and other metrics for the domains that you use to send email
// using Amazon Pinpoint. You also gain the ability to perform predictive inbox
// placement tests. When you use the Deliverability dashboard, you pay a monthly
// subscription charge, in addition to any other fees that you accrue by using
// Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing
// (http://aws.amazon.com/pinpoint/pricing/).
type PutDeliverabilityDashboardOptionInput struct {
	// An array of objects, one for each verified domain that you use to send email and
	// enabled the Deliverability dashboard for.
	SubscribedDomains []*types.DomainDeliverabilityTrackingOption
	// Specifies whether to enable the Deliverability dashboard for your Amazon
	// Pinpoint account. To enable the dashboard, set this value to true.
	DashboardEnabled *bool
}

// A response that indicates whether the Deliverability dashboard is enabled for
// your Amazon Pinpoint account.
type PutDeliverabilityDashboardOptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutDeliverabilityDashboardOptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutDeliverabilityDashboardOption{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDeliverabilityDashboardOption{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDeliverabilityDashboardOption",
	}
}
