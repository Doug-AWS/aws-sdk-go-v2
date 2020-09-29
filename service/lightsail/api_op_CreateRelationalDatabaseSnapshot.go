// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a snapshot of your database in Amazon Lightsail. You can use snapshots
// for backups, to make copies of a database, and to save data before deleting a
// database. The create relational database snapshot operation supports tag-based
// access control via request tags. For more information, see the Lightsail Dev
// Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateRelationalDatabaseSnapshot(ctx context.Context, params *CreateRelationalDatabaseSnapshotInput, optFns ...func(*Options)) (*CreateRelationalDatabaseSnapshotOutput, error) {
	stack := middleware.NewStack("CreateRelationalDatabaseSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRelationalDatabaseSnapshotMiddlewares(stack)
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
	addOpCreateRelationalDatabaseSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRelationalDatabaseSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateRelationalDatabaseSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateRelationalDatabaseSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRelationalDatabaseSnapshotInput struct {
	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
	// The name for your new database snapshot. Constraints:
	//
	//     * Must contain from 2
	// to 255 alphanumeric characters, or hyphens.
	//
	//     * The first and last character
	// must be a letter or number.
	RelationalDatabaseSnapshotName *string
	// The name of the database on which to base your new snapshot.
	RelationalDatabaseName *string
}

type CreateRelationalDatabaseSnapshotOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRelationalDatabaseSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRelationalDatabaseSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRelationalDatabaseSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRelationalDatabaseSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateRelationalDatabaseSnapshot",
	}
}
