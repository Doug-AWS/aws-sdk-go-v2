// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a traffic policy, which you use to create multiple DNS resource record
// sets for one domain name (such as example.com) or one subdomain name (such as
// www.example.com).
func (c *Client) CreateTrafficPolicy(ctx context.Context, params *CreateTrafficPolicyInput, optFns ...func(*Options)) (*CreateTrafficPolicyOutput, error) {
	stack := middleware.NewStack("CreateTrafficPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateTrafficPolicyMiddlewares(stack)
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
	addOpCreateTrafficPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficPolicy(options.Region), middleware.Before)
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
			OperationName: "CreateTrafficPolicy",
			Err:           err,
		}
	}
	out := result.(*CreateTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the traffic policy that you want
// to create.
type CreateTrafficPolicyInput struct {
	// The name of the traffic policy.
	Name *string
	// The definition of this traffic policy in JSON format. For more information, see
	// Traffic Policy Document Format
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html).
	Document *string
	// (Optional) Any comments that you want to include about the traffic policy.
	Comment *string
}

// A complex type that contains the response information for the
// CreateTrafficPolicy request.
type CreateTrafficPolicyOutput struct {
	// A unique URL that represents a new traffic policy.
	Location *string
	// A complex type that contains settings for the new traffic policy.
	TrafficPolicy *types.TrafficPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateTrafficPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateTrafficPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateTrafficPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrafficPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateTrafficPolicy",
	}
}
