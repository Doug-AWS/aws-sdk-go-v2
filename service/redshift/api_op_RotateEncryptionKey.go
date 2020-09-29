// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rotates the encryption keys for a cluster.
func (c *Client) RotateEncryptionKey(ctx context.Context, params *RotateEncryptionKeyInput, optFns ...func(*Options)) (*RotateEncryptionKeyOutput, error) {
	stack := middleware.NewStack("RotateEncryptionKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRotateEncryptionKeyMiddlewares(stack)
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
	addOpRotateEncryptionKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRotateEncryptionKey(options.Region), middleware.Before)
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
			OperationName: "RotateEncryptionKey",
			Err:           err,
		}
	}
	out := result.(*RotateEncryptionKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RotateEncryptionKeyInput struct {
	// The unique identifier of the cluster that you want to rotate the encryption keys
	// for. Constraints: Must be the name of valid cluster that has encryption enabled.
	ClusterIdentifier *string
}

type RotateEncryptionKeyOutput struct {
	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRotateEncryptionKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRotateEncryptionKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRotateEncryptionKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opRotateEncryptionKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RotateEncryptionKey",
	}
}
