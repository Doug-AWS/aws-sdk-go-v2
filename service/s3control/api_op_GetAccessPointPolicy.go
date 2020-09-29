// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the access point policy associated with the specified access point.
func (c *Client) GetAccessPointPolicy(ctx context.Context, params *GetAccessPointPolicyInput, optFns ...func(*Options)) (*GetAccessPointPolicyOutput, error) {
	stack := middleware.NewStack("GetAccessPointPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetAccessPointPolicyMiddlewares(stack)
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
	addOpGetAccessPointPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicy(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "GetAccessPointPolicy",
			Err:           err,
		}
	}
	out := result.(*GetAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyInput struct {
	// The name of the access point whose policy you want to retrieve.
	Name *string
	// The account ID for the account that owns the specified access point.
	AccountId *string
}

type GetAccessPointPolicyOutput struct {
	// The access point policy associated with the specified access point.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetAccessPointPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPointPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointPolicy",
	}
}
