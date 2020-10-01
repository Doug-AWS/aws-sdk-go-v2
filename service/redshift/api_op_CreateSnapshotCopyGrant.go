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

// Creates a snapshot copy grant that permits Amazon Redshift to use a customer
// master key (CMK) from AWS Key Management Service (AWS KMS) to encrypt copied
// snapshots in a destination region. For more information about managing snapshot
// copy grants, go to Amazon Redshift Database Encryption
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateSnapshotCopyGrant(ctx context.Context, params *CreateSnapshotCopyGrantInput, optFns ...func(*Options)) (*CreateSnapshotCopyGrantOutput, error) {
	stack := middleware.NewStack("CreateSnapshotCopyGrant", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateSnapshotCopyGrantMiddlewares(stack)
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
	addOpCreateSnapshotCopyGrantValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshotCopyGrant(options.Region), middleware.Before)
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
			OperationName: "CreateSnapshotCopyGrant",
			Err:           err,
		}
	}
	out := result.(*CreateSnapshotCopyGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The result of the CreateSnapshotCopyGrant action.
type CreateSnapshotCopyGrantInput struct {

	// The name of the snapshot copy grant. This name must be unique in the region for
	// the AWS account. Constraints:
	//
	//     * Must contain from 1 to 63 alphanumeric
	// characters or hyphens.
	//
	//     * Alphabetic characters must be lowercase.
	//
	//     *
	// First character must be a letter.
	//
	//     * Cannot end with a hyphen or contain two
	// consecutive hyphens.
	//
	//     * Must be unique for all clusters within an AWS
	// account.
	//
	// This member is required.
	SnapshotCopyGrantName *string

	// The unique identifier of the customer master key (CMK) to which to grant Amazon
	// Redshift permission. If no key is specified, the default key is used.
	KmsKeyId *string

	// A list of tag instances.
	Tags []*types.Tag
}

type CreateSnapshotCopyGrantOutput struct {

	// The snapshot copy grant that grants Amazon Redshift permission to encrypt copied
	// snapshots with the specified customer master key (CMK) from AWS KMS in the
	// destination region. For more information about managing snapshot copy grants, go
	// to Amazon Redshift Database Encryption
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html)
	// in the Amazon Redshift Cluster Management Guide.
	SnapshotCopyGrant *types.SnapshotCopyGrant

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateSnapshotCopyGrantMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateSnapshotCopyGrant{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSnapshotCopyGrant{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSnapshotCopyGrant(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateSnapshotCopyGrant",
	}
}
