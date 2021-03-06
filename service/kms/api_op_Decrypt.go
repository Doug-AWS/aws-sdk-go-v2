// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Decrypts ciphertext that was encrypted by a AWS KMS customer master key (CMK)
// using any of the following operations:
//
// * Encrypt
//
// * GenerateDataKey
//
// *
// GenerateDataKeyPair
//
// * GenerateDataKeyWithoutPlaintext
//
// *
// GenerateDataKeyPairWithoutPlaintext
//
// You can use this operation to decrypt
// ciphertext that was encrypted under a symmetric or asymmetric CMK. When the CMK
// is asymmetric, you must specify the CMK and the encryption algorithm that was
// used to encrypt the ciphertext. For information about symmetric and asymmetric
// CMKs, see Using Symmetric and Asymmetric CMKs
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the AWS Key Management Service Developer Guide. The Decrypt operation also
// decrypts ciphertext that was encrypted outside of AWS KMS by the public key in
// an AWS KMS asymmetric CMK. However, it cannot decrypt ciphertext produced by
// other libraries, such as the AWS Encryption SDK
// (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/) or Amazon
// S3 client-side encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html).
// These libraries return a ciphertext format that is incompatible with AWS KMS. If
// the ciphertext was encrypted under a symmetric CMK, you do not need to specify
// the CMK or the encryption algorithm. AWS KMS can get this information from
// metadata that it adds to the symmetric ciphertext blob. However, if you prefer,
// you can specify the KeyId to ensure that a particular CMK is used to decrypt the
// ciphertext. If you specify a different CMK than the one used to encrypt the
// ciphertext, the Decrypt operation fails. Whenever possible, use key policies to
// give users permission to call the Decrypt operation on a particular CMK, instead
// of using IAM policies. Otherwise, you might create an IAM user policy that gives
// the user Decrypt permission on all CMKs. This user could decrypt ciphertext that
// was encrypted by CMKs in other accounts if the key policy for the cross-account
// CMK permits it. If you must use an IAM policy for Decrypt permissions, limit the
// user to particular CMKs or particular trusted accounts. The CMK that you use for
// this operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) Decrypt(ctx context.Context, params *DecryptInput, optFns ...func(*Options)) (*DecryptOutput, error) {
	if params == nil {
		params = &DecryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Decrypt", params, optFns, addOperationDecryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecryptInput struct {

	// Ciphertext to be decrypted. The blob includes metadata.
	//
	// This member is required.
	CiphertextBlob []byte

	// Specifies the encryption algorithm that will be used to decrypt the ciphertext.
	// Specify the same algorithm that was used to encrypt the data. If you specify a
	// different algorithm, the Decrypt operation fails. This parameter is required
	// only when the ciphertext was encrypted under an asymmetric CMK. The default
	// value, SYMMETRIC_DEFAULT, represents the only supported algorithm that is valid
	// for symmetric CMKs.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context to use when decrypting the data. An encryption
	// context is valid only for cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// with a symmetric CMK. The standard asymmetric encryption algorithms that AWS KMS
	// uses do not support an encryption context. An encryption context is a collection
	// of non-secret key-value pairs that represents additional authenticated data.
	// When you use an encryption context to encrypt data, you must specify the same
	// (an exact case-sensitive match) encryption context to decrypt the data. An
	// encryption context is optional when encrypting with a symmetric CMK, but it is
	// highly recommended. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]*string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string

	// Specifies the customer master key (CMK) that AWS KMS will use to decrypt the
	// ciphertext. Enter a key ID of the CMK that was used to encrypt the ciphertext.
	// If you specify a KeyId value, the Decrypt operation succeeds only if the
	// specified CMK was used to encrypt the ciphertext. This parameter is required
	// only when the ciphertext was encrypted under an asymmetric CMK. Otherwise, AWS
	// KMS uses the metadata that it adds to the ciphertext blob to determine which CMK
	// was used to encrypt the ciphertext. However, you can use this parameter to
	// ensure that a particular CMK (of any kind) is used to decrypt the ciphertext. To
	// specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias
	// ARN. When using an alias name, prefix it with "alias/". For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys or DescribeKey. To get the alias name and alias ARN,
	// use ListAliases.
	KeyId *string
}

type DecryptOutput struct {

	// The encryption algorithm that was used to decrypt the ciphertext.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the CMK that was used to decrypt the ciphertext.
	KeyId *string

	// Decrypted plaintext data. When you use the HTTP API or the AWS CLI, the value is
	// Base64-encoded. Otherwise, it is not Base64-encoded.
	Plaintext []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDecryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDecrypt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecrypt{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDecryptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecrypt(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDecrypt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Decrypt",
	}
}
