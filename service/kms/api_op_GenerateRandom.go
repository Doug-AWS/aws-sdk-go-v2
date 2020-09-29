// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a random byte string that is cryptographically secure. By default, the
// random byte string is generated in AWS KMS. To generate the byte string in the
// AWS CloudHSM cluster that is associated with a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// specify the custom key store ID. For more information about entropy and random
// number generation, see the AWS Key Management Service Cryptographic Details
// (https://d0.awsstatic.com/whitepapers/KMS-Cryptographic-Details.pdf) whitepaper.
func (c *Client) GenerateRandom(ctx context.Context, params *GenerateRandomInput, optFns ...func(*Options)) (*GenerateRandomOutput, error) {
	stack := middleware.NewStack("GenerateRandom", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGenerateRandomMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateRandom(options.Region), middleware.Before)
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
			OperationName: "GenerateRandom",
			Err:           err,
		}
	}
	out := result.(*GenerateRandomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateRandomInput struct {
	// The length of the byte string.
	NumberOfBytes *int32
	// Generates the random byte string in the AWS CloudHSM cluster that is associated
	// with the specified custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// To find the ID of a custom key store, use the DescribeCustomKeyStores ()
	// operation.
	CustomKeyStoreId *string
}

type GenerateRandomOutput struct {
	// The random byte string. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	Plaintext []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGenerateRandomMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateRandom{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateRandom{}, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateRandom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GenerateRandom",
	}
}
