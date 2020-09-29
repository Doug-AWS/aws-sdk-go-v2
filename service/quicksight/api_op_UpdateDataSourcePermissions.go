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

// Updates the permissions to a data source.
func (c *Client) UpdateDataSourcePermissions(ctx context.Context, params *UpdateDataSourcePermissionsInput, optFns ...func(*Options)) (*UpdateDataSourcePermissionsOutput, error) {
	stack := middleware.NewStack("UpdateDataSourcePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDataSourcePermissionsMiddlewares(stack)
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
	addOpUpdateDataSourcePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSourcePermissions(options.Region), middleware.Before)
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
			OperationName: "UpdateDataSourcePermissions",
			Err:           err,
		}
	}
	out := result.(*UpdateDataSourcePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourcePermissionsInput struct {
	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string
	// A list of resource permissions that you want to grant on the data source.
	GrantPermissions []*types.ResourcePermission
	// The AWS account ID.
	AwsAccountId *string
	// A list of resource permissions that you want to revoke on the data source.
	RevokePermissions []*types.ResourcePermission
}

type UpdateDataSourcePermissionsOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string
	// The Amazon Resource Name (ARN) of the data source.
	DataSourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDataSourcePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSourcePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSourcePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDataSourcePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDataSourcePermissions",
	}
}
