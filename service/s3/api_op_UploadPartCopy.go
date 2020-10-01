// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Uploads a part by copying data from an existing object as data source. You
// specify the data source by adding the request header x-amz-copy-source in your
// request and a byte range by adding the request header x-amz-copy-source-range in
// your request. The minimum allowable part size for a multipart upload is 5 MB.
// For more information about multipart upload limits, go to Quick Facts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/qfacts.html) in the Amazon
// Simple Storage Service Developer Guide. Instead of using an existing object as
// part data, you might use the UploadPart () operation and provide data in your
// request.  <p>You must initiate a multipart upload before you can upload any
// part. In response to your initiate request. Amazon S3 returns a unique
// identifier, the upload ID, that you must include in your upload part
// request.</p> <p>For more information about using the <code>UploadPartCopy</code>
// operation, see the following:</p> <ul> <li> <p>For conceptual information about
// multipart uploads, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html">Uploading
// Objects Using Multipart Upload</a> in the <i>Amazon Simple Storage Service
// Developer Guide</i>.</p> </li> <li> <p>For information about permissions
// required to use the multipart upload API, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html">Multipart
// Upload API and Permissions</a> in the <i>Amazon Simple Storage Service Developer
// Guide</i>.</p> </li> <li> <p>For information about copying objects using a
// single atomic operation vs. the multipart upload, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html">Operations
// on Objects</a> in the <i>Amazon Simple Storage Service Developer Guide</i>.</p>
// </li> <li> <p>For information about using server-side encryption with
// customer-provided encryption keys with the UploadPartCopy operation, see
// <a>CopyObject</a> and <a>UploadPart</a>.</p> </li> </ul> <p>Note the following
// additional considerations about the request headers
// <code>x-amz-copy-source-if-match</code>,
// <code>x-amz-copy-source-if-none-match</code>,
// <code>x-amz-copy-source-if-unmodified-since</code>, and
// <code>x-amz-copy-source-if-modified-since</code>:</p> <p> </p> <ul> <li> <p>
// <b>Consideration 1</b> - If both of the <code>x-amz-copy-source-if-match</code>
// and <code>x-amz-copy-source-if-unmodified-since</code> headers are present in
// the request as follows:</p> <p> <code>x-amz-copy-source-if-match</code>
// condition evaluates to <code>true</code>, and;</p> <p>
// <code>x-amz-copy-source-if-unmodified-since</code> condition evaluates to
// <code>false</code>;</p> <p>Amazon S3 returns <code>200 OK</code> and copies the
// data. </p> </li> <li> <p> <b>Consideration 2</b> - If both of the
// <code>x-amz-copy-source-if-none-match</code> and
// <code>x-amz-copy-source-if-modified-since</code> headers are present in the
// request as follows:</p> <p> <code>x-amz-copy-source-if-none-match</code>
// condition evaluates to <code>false</code>, and;</p> <p>
// <code>x-amz-copy-source-if-modified-since</code> condition evaluates to
// <code>true</code>;</p> <p>Amazon S3 returns <code>412 Precondition Failed</code>
// response code. </p> </li> </ul> <p> <b>Versioning</b> </p> <p>If your bucket has
// versioning enabled, you could have multiple versions of the same object. By
// default, <code>x-amz-copy-source</code> identifies the current version of the
// object to copy. If the current version is a delete marker and you don't specify
// a versionId in the <code>x-amz-copy-source</code>, Amazon S3 returns a 404
// error, because the object does not exist. If you specify versionId in the
// <code>x-amz-copy-source</code> and the versionId is a delete marker, Amazon S3
// returns an HTTP 400 error, because you are not allowed to specify a delete
// marker as a version for the <code>x-amz-copy-source</code>. </p> <p>You can
// optionally specify a specific version of the source object to copy by adding the
// <code>versionId</code> subresource as shown in the following example:</p> <p>
// <code>x-amz-copy-source: /bucket/object?versionId=version id</code> </p> <p
// class="title"> <b>Special Errors</b> </p> <ul> <li> <p class="title"> <b></b>
// </p> <ul> <li> <p> <i>Code: NoSuchUpload</i> </p> </li> <li> <p> <i>Cause: The
// specified multipart upload does not exist. The upload ID might be invalid, or
// the multipart upload might have been aborted or completed.</i> </p> </li> <li>
// <p> <i>HTTP Status Code: 404 Not Found</i> </p> </li> </ul> </li> <li> <p
// class="title"> <b></b> </p> <ul> <li> <p> <i>Code: InvalidRequest</i> </p> </li>
// <li> <p> <i>Cause: The specified copy source is not supported as a byte-range
// copy source.</i> </p> </li> <li> <p> <i>HTTP Status Code: 400 Bad Request</i>
// </p> </li> </ul> </li> </ul> <p class="title"> <b>Related Resources</b> </p>
// <ul> <li> <p> <a>CreateMultipartUpload</a> </p> </li> <li> <p> <a>UploadPart</a>
// </p> </li> <li> <p> <a>CompleteMultipartUpload</a> </p> </li> <li> <p>
// <a>AbortMultipartUpload</a> </p> </li> <li> <p> <a>ListParts</a> </p> </li> <li>
// <p> <a>ListMultipartUploads</a> </p> </li> </ul>
func (c *Client) UploadPartCopy(ctx context.Context, params *UploadPartCopyInput, optFns ...func(*Options)) (*UploadPartCopyOutput, error) {
	stack := middleware.NewStack("UploadPartCopy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUploadPartCopyMiddlewares(stack)
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
	addOpUploadPartCopyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUploadPartCopy(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "UploadPartCopy",
			Err:           err,
		}
	}
	out := result.(*UploadPartCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadPartCopyInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// The name of the source bucket and key name of the source object, separated by a
	// slash (/). Must be URL-encoded.
	//
	// This member is required.
	CopySource *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Part number of part being copied. This is a positive integer between 1 and
	// 10,000.
	//
	// This member is required.
	PartNumber *int32

	// Upload ID identifying the multipart upload whose part is being copied.
	//
	// This member is required.
	UploadId *string

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time

	// The range of bytes to copy from the source object. The range value must use the
	// form bytes=first-last, where the first and last are the zero-based byte offsets
	// to copy. For example, bytes=0-9 indicates that you want to copy the first 10
	// bytes of the source. You can copy a range only if the source object is greater
	// than 5 MB.
	CopySourceRange *string

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256).
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one that
	// was used when the source object was created.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header. This must be the same
	// encryption key specified in the initiate multipart upload request.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string
}

type UploadPartCopyOutput struct {

	// Container for all response elements.
	CopyPartResult *types.CopyPartResult

	// The version of the source object that was copied, if you have enabled versioning
	// on the source bucket.
	CopySourceVersionId *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUploadPartCopyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUploadPartCopy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUploadPartCopy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUploadPartCopy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UploadPartCopy",
	}
}
