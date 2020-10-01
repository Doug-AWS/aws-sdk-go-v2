// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all schema extensions applied to a Microsoft AD Directory.
func (c *Client) ListSchemaExtensions(ctx context.Context, params *ListSchemaExtensionsInput, optFns ...func(*Options)) (*ListSchemaExtensionsOutput, error) {
	stack := middleware.NewStack("ListSchemaExtensions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListSchemaExtensionsMiddlewares(stack)
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
	addOpListSchemaExtensionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaExtensions(options.Region), middleware.Before)
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
			OperationName: "ListSchemaExtensions",
			Err:           err,
		}
	}
	out := result.(*ListSchemaExtensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaExtensionsInput struct {

	// The identifier of the directory from which to retrieve the schema extension
	// information.
	//
	// This member is required.
	DirectoryId *string

	// The maximum number of items to return.
	Limit *int32

	// The ListSchemaExtensions.NextToken value from a previous call to
	// ListSchemaExtensions. Pass null if this is the first call.
	NextToken *string
}

type ListSchemaExtensionsOutput struct {

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListSchemaExtensions to retrieve the next set
	// of items.
	NextToken *string

	// Information about the schema extensions applied to the directory.
	SchemaExtensionsInfo []*types.SchemaExtensionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListSchemaExtensionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemaExtensions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemaExtensions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSchemaExtensions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListSchemaExtensions",
	}
}
