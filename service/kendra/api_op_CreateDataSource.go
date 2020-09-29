// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a data source that you use to with an Amazon Kendra index. You specify a
// name, connector type and description for your data source. You can choose
// between an S3 connector, a SharePoint Online connector, and a database
// connector. You also specify configuration information such as document metadata
// (author, source URI, and so on) and user context information. CreateDataSource
// is a synchronous operation. The operation returns 200 if the data source was
// successfully created. Otherwise, an exception is raised.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	stack := middleware.NewStack("CreateDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDataSourceMiddlewares(stack)
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
	addOpCreateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before)
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
			OperationName: "CreateDataSource",
			Err:           err,
		}
	}
	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {
	// A description for the data source.
	Description *string
	// The type of repository that contains the data source.
	Type types.DataSourceType
	// A list of key-value pairs that identify the data source. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []*types.Tag
	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source. For more information, see IAM Roles for Amazon Kendra
	// (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn *string
	// A unique name for the data source. A data source name can't be changed without
	// deleting and recreating the data source.
	Name *string
	// Sets the frequency that Amazon Kendra will check the documents in your
	// repository and update the index. If you don't set a schedule Amazon Kendra will
	// not periodically update the index. You can call the StartDataSourceSyncJob
	// operation to update the index.
	Schedule *string
	// The identifier of the index that should be associated with this data source.
	IndexId *string
	// The connector configuration information that is required to access the
	// repository.
	Configuration *types.DataSourceConfiguration
}

type CreateDataSourceOutput struct {
	// A unique identifier for the data source.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateDataSource",
	}
}
