// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Writes a block of data to a block in the snapshot. If the specified block
// contains data, the existing data is overwritten. The target snapshot must be in
// the pending state. Data written to a snapshot must be aligned with 512-byte
// sectors.
func (c *Client) PutSnapshotBlock(ctx context.Context, params *PutSnapshotBlockInput, optFns ...func(*Options)) (*PutSnapshotBlockOutput, error) {
	stack := middleware.NewStack("PutSnapshotBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutSnapshotBlockMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddUnsignedPayloadMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutSnapshotBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutSnapshotBlock(options.Region), middleware.Before)
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
			OperationName: "PutSnapshotBlock",
			Err:           err,
		}
	}
	out := result.(*PutSnapshotBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSnapshotBlockInput struct {

	// The data to write to the block. The block data is not signed as part of the
	// Signature Version 4 signing process. As a result, you must generate and provide
	// a Base64-encoded SHA256 checksum for the block data using the x-amz-Checksum
	// header. Also, you must specify the checksum algorithm using the
	// x-amz-Checksum-Algorithm header. The checksum that you provide is part of the
	// Signature Version 4 signing process. It is validated against a checksum
	// generated by Amazon EBS to ensure the validity and authenticity of the data. If
	// the checksums do not correspond, the request fails. For more information, see
	// Using checksums with the EBS direct APIs
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapis-using-checksums)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// This member is required.
	BlockData io.Reader

	// The block index of the block in which to write the data. A block index is the
	// offset position of a block within a snapshot, and it is used to identify the
	// block. To identify the logical offset of the data in the logical volume,
	// multiply the block index with the block size (Block index * 512 bytes).
	//
	// This member is required.
	BlockIndex *int32

	// A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are
	// supported.
	//
	// This member is required.
	Checksum *string

	// The algorithm used to generate the checksum. Currently, the only supported
	// algorithm is SHA256.
	//
	// This member is required.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The size of the data to write to the block, in bytes. Currently, the only
	// supported size is 524288. Valid values: 524288
	//
	// This member is required.
	DataLength *int32

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// The progress of the write process, as a percentage.
	Progress *int32
}

type PutSnapshotBlockOutput struct {

	// The SHA256 checksum generated for the block data by Amazon EBS.
	Checksum *string

	// The algorithm used by Amazon EBS to generate the checksum.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutSnapshotBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutSnapshotBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSnapshotBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutSnapshotBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "PutSnapshotBlock",
	}
}
