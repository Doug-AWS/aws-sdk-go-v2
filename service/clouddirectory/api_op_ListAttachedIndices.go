// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists indices attached to the specified object.
func (c *Client) ListAttachedIndices(ctx context.Context, params *ListAttachedIndicesInput, optFns ...func(*Options)) (*ListAttachedIndicesOutput, error) {
	stack := middleware.NewStack("ListAttachedIndices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAttachedIndicesMiddlewares(stack)
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
	addOpListAttachedIndicesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedIndices(options.Region), middleware.Before)
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
			OperationName: "ListAttachedIndices",
			Err:           err,
		}
	}
	out := result.(*ListAttachedIndicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedIndicesInput struct {
	// The pagination token.
	NextToken *string
	// A reference to the object that has indices attached.
	TargetReference *types.ObjectReference
	// The consistency level to use for this operation.
	ConsistencyLevel types.ConsistencyLevel
	// The maximum number of results to retrieve.
	MaxResults *int32
	// The ARN of the directory.
	DirectoryArn *string
}

type ListAttachedIndicesOutput struct {
	// The pagination token.
	NextToken *string
	// The indices attached to the specified object.
	IndexAttachments []*types.IndexAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAttachedIndicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAttachedIndices{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttachedIndices{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAttachedIndices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListAttachedIndices",
	}
}
