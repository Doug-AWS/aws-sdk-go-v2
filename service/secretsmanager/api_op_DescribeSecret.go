// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the details of a secret. It does not include the encrypted fields.
// Secrets Manager only returns fields populated with a value in the response.
// Minimum permissions To run this command, you must have the following
// permissions:
//
//     * secretsmanager:DescribeSecret
//
// Related operations
//
//     * To
// create a secret, use CreateSecret ().
//
//     * To modify a secret, use
// UpdateSecret ().
//
//     * To retrieve the encrypted secret information in a
// version of the secret, use GetSecretValue ().
//
//     * To list all of the secrets
// in the AWS account, use ListSecrets ().
func (c *Client) DescribeSecret(ctx context.Context, params *DescribeSecretInput, optFns ...func(*Options)) (*DescribeSecretOutput, error) {
	stack := middleware.NewStack("DescribeSecret", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeSecretMiddlewares(stack)
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
	addOpDescribeSecretValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecret(options.Region), middleware.Before)
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
			OperationName: "DescribeSecret",
			Err:           err,
		}
	}
	out := result.(*DescribeSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecretInput struct {

	// The identifier of the secret whose details you want to retrieve. You can specify
	// either the Amazon Resource Name (ARN) or the friendly name of the secret. If you
	// specify an ARN, we generally recommend that you specify a complete ARN. You can
	// specify a partial ARN too—for example, if you don’t include the final hyphen and
	// six random characters that Secrets Manager adds at the end of the ARN when you
	// created the secret. A partial ARN match can work as long as it uniquely matches
	// only one secret. However, if your secret has a name that ends in a hyphen
	// followed by six characters (before Secrets Manager adds the hyphen and six
	// characters to the ARN) and you try to use that as a partial ARN, then those
	// characters cause Secrets Manager to assume that you’re specifying a complete
	// ARN. This confusion can cause unexpected results. To avoid this situation, we
	// recommend that you don’t create secret names ending with a hyphen followed by
	// six characters. If you specify an incomplete ARN without the random suffix, and
	// instead provide the 'friendly name', you must not include the random suffix. If
	// you do include the random suffix added by Secrets Manager, you receive either a
	// ResourceNotFoundException or an AccessDeniedException error, depending on your
	// permissions.
	//
	// This member is required.
	SecretId *string
}

type DescribeSecretOutput struct {

	// The ARN of the secret.
	ARN *string

	// The date that the secret was created.
	CreatedDate *time.Time

	// This value exists if the secret is scheduled for deletion. Some time after the
	// specified date and time, Secrets Manager deletes the secret and all of its
	// versions. If a secret is scheduled for deletion, then its details, including the
	// encrypted secret information, is not accessible. To cancel a scheduled deletion
	// and restore access, use RestoreSecret ().
	DeletedDate *time.Time

	// The user-provided description of the secret.
	Description *string

	// The ARN or alias of the AWS KMS customer master key (CMK) that's used to encrypt
	// the SecretString or SecretBinary fields in each version of the secret. If you
	// don't provide a key, then Secrets Manager defaults to encrypting the secret
	// fields with the default AWS KMS CMK (the one named awssecretsmanager) for this
	// account.
	KmsKeyId *string

	// The last date that this secret was accessed. This value is truncated to midnight
	// of the date and therefore shows only the date, not the time.
	LastAccessedDate *time.Time

	// The last date and time that this secret was modified in any way.
	LastChangedDate *time.Time

	// The most recent date and time that the Secrets Manager rotation process was
	// successfully completed. This value is null if the secret has never rotated.
	LastRotatedDate *time.Time

	// The user-provided friendly name of the secret.
	Name *string

	// Returns the name of the service that created this secret.
	OwningService *string

	// Specifies whether automatic rotation is enabled for this secret. To enable
	// rotation, use RotateSecret () with AutomaticallyRotateAfterDays set to a value
	// greater than 0. To disable rotation, use CancelRotateSecret ().
	RotationEnabled *bool

	// The ARN of a Lambda function that's invoked by Secrets Manager to rotate the
	// secret either automatically per the schedule or manually by a call to
	// RotateSecret.
	RotationLambdaARN *string

	// A structure that contains the rotation configuration for this secret.
	RotationRules *types.RotationRulesType

	// The list of user-defined tags that are associated with the secret. To add tags
	// to a secret, use TagResource (). To remove tags, use UntagResource ().
	Tags []*types.Tag

	// A list of all of the currently assigned VersionStage staging labels and the
	// VersionId that each is attached to. Staging labels are used to keep track of the
	// different versions during the rotation process. A version that does not have any
	// staging labels attached is considered deprecated and subject to deletion. Such
	// versions are not included in this list.
	VersionIdsToStages map[string][]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeSecretMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSecret{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSecret{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSecret(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "DescribeSecret",
	}
}
