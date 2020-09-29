// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified Systems Manager document with the specified instances
// or targets. When you associate a document with one or more instances using
// instance IDs or tags, SSM Agent running on the instance processes the document
// and configures the instance as specified. If you associate a document with an
// instance that already has an associated document, the system returns the
// AssociationAlreadyExists exception.
func (c *Client) CreateAssociationBatch(ctx context.Context, params *CreateAssociationBatchInput, optFns ...func(*Options)) (*CreateAssociationBatchOutput, error) {
	stack := middleware.NewStack("CreateAssociationBatch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAssociationBatchMiddlewares(stack)
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
	addOpCreateAssociationBatchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssociationBatch(options.Region), middleware.Before)
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
			OperationName: "CreateAssociationBatch",
			Err:           err,
		}
	}
	out := result.(*CreateAssociationBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssociationBatchInput struct {
	// One or more associations.
	Entries []*types.CreateAssociationBatchRequestEntry
}

type CreateAssociationBatchOutput struct {
	// Information about the associations that succeeded.
	Successful []*types.AssociationDescription
	// Information about the associations that failed.
	Failed []*types.FailedCreateAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAssociationBatchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssociationBatch{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssociationBatch{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAssociationBatch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateAssociationBatch",
	}
}
