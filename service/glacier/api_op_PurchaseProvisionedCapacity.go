// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation purchases a provisioned capacity unit for an AWS account.
func (c *Client) PurchaseProvisionedCapacity(ctx context.Context, params *PurchaseProvisionedCapacityInput, optFns ...func(*Options)) (*PurchaseProvisionedCapacityOutput, error) {
	stack := middleware.NewStack("PurchaseProvisionedCapacity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPurchaseProvisionedCapacityMiddlewares(stack)
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
	addOpPurchaseProvisionedCapacityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseProvisionedCapacity(options.Region), middleware.Before)

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
			OperationName: "PurchaseProvisionedCapacity",
			Err:           err,
		}
	}
	out := result.(*PurchaseProvisionedCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PurchaseProvisionedCapacityInput struct {
	// The AWS account ID of the account that owns the vault. You can either specify an
	// AWS account ID or optionally a single '-' (hyphen), in which case Amazon S3
	// Glacier uses the AWS account ID associated with the credentials used to sign the
	// request. If you use an account ID, don't include any hyphens ('-') in the ID.
	AccountId *string
}

type PurchaseProvisionedCapacityOutput struct {
	// The ID that identifies the provisioned capacity unit.
	CapacityId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPurchaseProvisionedCapacityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseProvisionedCapacity{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseProvisionedCapacity{}, middleware.After)
}

func newServiceMetadataMiddleware_opPurchaseProvisionedCapacity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "PurchaseProvisionedCapacity",
	}
}