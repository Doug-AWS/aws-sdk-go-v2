// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the field-level encryption profile information.
func (c *Client) GetFieldLevelEncryptionProfile(ctx context.Context, params *GetFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*GetFieldLevelEncryptionProfileOutput, error) {
	stack := middleware.NewStack("GetFieldLevelEncryptionProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetFieldLevelEncryptionProfileMiddlewares(stack)
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
	addOpGetFieldLevelEncryptionProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfile(options.Region), middleware.Before)
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
			OperationName: "GetFieldLevelEncryptionProfile",
			Err:           err,
		}
	}
	out := result.(*GetFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFieldLevelEncryptionProfileInput struct {

	// Get the ID for the field-level encryption profile information.
	//
	// This member is required.
	Id *string
}

type GetFieldLevelEncryptionProfileOutput struct {

	// The current version of the field level encryption profile. For example:
	// E2QWRUHAPOMQZL.
	ETag *string

	// Return the field-level encryption profile information.
	FieldLevelEncryptionProfile *types.FieldLevelEncryptionProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetFieldLevelEncryptionProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetFieldLevelEncryptionProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetFieldLevelEncryptionProfile",
	}
}
