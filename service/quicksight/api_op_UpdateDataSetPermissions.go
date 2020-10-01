// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the permissions on a dataset. The permissions resource is
// arn:aws:quicksight:region:aws-account-id:dataset/data-set-id.
func (c *Client) UpdateDataSetPermissions(ctx context.Context, params *UpdateDataSetPermissionsInput, optFns ...func(*Options)) (*UpdateDataSetPermissionsOutput, error) {
	stack := middleware.NewStack("UpdateDataSetPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDataSetPermissionsMiddlewares(stack)
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
	addOpUpdateDataSetPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSetPermissions(options.Region), middleware.Before)
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
			OperationName: "UpdateDataSetPermissions",
			Err:           err,
		}
	}
	out := result.(*UpdateDataSetPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSetPermissionsInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	//
	// This member is required.
	DataSetId *string

	// The resource permissions that you want to grant to the dataset.
	GrantPermissions []*types.ResourcePermission

	// The resource permissions that you want to revoke from the dataset.
	RevokePermissions []*types.ResourcePermission
}

type UpdateDataSetPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	DataSetArn *string

	// The ID for the dataset whose permissions you want to update. This ID is unique
	// per AWS Region for each AWS account.
	DataSetId *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDataSetPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSetPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSetPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDataSetPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDataSetPermissions",
	}
}
