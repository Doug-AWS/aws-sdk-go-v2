// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes a backup plan. A backup plan can only be deleted after all associated
// selections of resources have been deleted. Deleting a backup plan deletes the
// current version of a backup plan. Previous versions, if any, will still exist.
func (c *Client) DeleteBackupPlan(ctx context.Context, params *DeleteBackupPlanInput, optFns ...func(*Options)) (*DeleteBackupPlanOutput, error) {
	stack := middleware.NewStack("DeleteBackupPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteBackupPlanMiddlewares(stack)
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
	addOpDeleteBackupPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBackupPlan(options.Region), middleware.Before)
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
			OperationName: "DeleteBackupPlan",
			Err:           err,
		}
	}
	out := result.(*DeleteBackupPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBackupPlanInput struct {

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string
}

type DeleteBackupPlanOutput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// The date and time a backup plan is deleted, in Unix format and Coordinated
	// Universal Time (UTC). The value of DeletionDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	DeletionDate *time.Time

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version Ids cannot be edited.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteBackupPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBackupPlan{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBackupPlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBackupPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DeleteBackupPlan",
	}
}
