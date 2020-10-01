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

// This API action puts a bucket policy to an Amazon S3 on Outposts bucket. To put
// a policy on an S3 bucket, see PutBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html) in
// the Amazon Simple Storage Service API. Applies an Amazon S3 bucket policy to an
// Outposts bucket. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
// Amazon Simple Storage Service Developer Guide. If you are using an identity
// other than the root user of the AWS account that owns the Outposts bucket, the
// calling identity must have the PutBucketPolicy permissions on the specified
// Outposts bucket and belong to the bucket owner's account in order to use this
// operation.  <p>If you don't have <code>PutBucketPolicy</code> permissions,
// Amazon S3 returns a <code>403 Access Denied</code> error. If you have the
// correct permissions, but you're not using an identity that belongs to the bucket
// owner's account, Amazon S3 returns a <code>405 Method Not Allowed</code>
// error.</p> <important> <p> As a security precaution, the root user of the AWS
// account that owns a bucket can always use this operation, even if the policy
// explicitly denies the root user the ability to perform this action. </p>
// </important> <p>For more information about bucket policies, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html">Using
// Bucket Policies and User Policies</a>.</p> <p>All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of outpost-id to be
// passed with the request and an S3 on Outposts endpoint hostname prefix instead
// of s3-control. For an example of the request syntax for Amazon S3 on Outposts
// that uses the S3 on Outposts endpoint hostname prefix and the outpost-id derived
// using the access point ARN, see the <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutBucketPolicy.html#API_control_PutBucketPolicy_Examples">
// Example</a> section below.</p> <p>The following actions are related to
// <code>PutBucketPolicy</code>:</p> <ul> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html">GetBucketPolicy</a>
// </p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html">DeleteBucketPolicy</a>
// </p> </li> </ul>
func (c *Client) PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error) {
	stack := middleware.NewStack("PutBucketPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketPolicyMiddlewares(stack)
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
	addOpPutBucketPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketPolicy(options.Region), middleware.Before)
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
			OperationName: "PutBucketPolicy",
			Err:           err,
		}
	}
	out := result.(*PutBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketPolicyInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	// The bucket policy as a JSON document.
	//
	// This member is required.
	Policy *string

	// Set this parameter to true to confirm that you want to remove your permissions
	// to change this bucket policy in the future. This is not supported by Amazon S3
	// on Outposts buckets.
	ConfirmRemoveSelfBucketAccess *bool
}

type PutBucketPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketPolicy",
	}
}
