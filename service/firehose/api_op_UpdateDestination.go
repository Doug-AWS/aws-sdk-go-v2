// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified destination of the specified delivery stream.  <p>Use this
// operation to change the destination type (for example, to replace the Amazon S3
// destination with Amazon Redshift) or change the parameters associated with a
// destination (for example, to change the bucket name of the Amazon S3
// destination). The update might not occur immediately. The target delivery stream
// remains active while the configurations are updated, so data writes to the
// delivery stream can continue during this process. The updated configurations are
// usually effective within a few minutes.</p> <p>Switching between Amazon ES and
// other services is not supported. For an Amazon ES destination, you can only
// update to another Amazon ES destination.</p> <p>If the destination type is the
// same, Kinesis Data Firehose merges the configuration parameters specified with
// the destination configuration that already exists on the delivery stream. If any
// of the parameters are not specified in the call, the existing values are
// retained. For example, in the Amazon S3 destination, if
// <a>EncryptionConfiguration</a> is not specified, then the existing
// <code>EncryptionConfiguration</code> is maintained on the destination.</p> <p>If
// the destination type is not the same, for example, changing the destination from
// Amazon S3 to Amazon Redshift, Kinesis Data Firehose does not merge any
// parameters. In this case, all parameters must be specified.</p> <p>Kinesis Data
// Firehose uses <code>CurrentDeliveryStreamVersionId</code> to avoid race
// conditions and conflicting merges. This is a required field, and the service
// updates the configuration only if the existing configuration has a version ID
// that matches. After the update is applied successfully, the version ID is
// updated, and can be retrieved using <a>DescribeDeliveryStream</a>. Use the new
// version ID to set <code>CurrentDeliveryStreamVersionId</code> in the next
// call.</p>
func (c *Client) UpdateDestination(ctx context.Context, params *UpdateDestinationInput, optFns ...func(*Options)) (*UpdateDestinationOutput, error) {
	stack := middleware.NewStack("UpdateDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDestinationMiddlewares(stack)
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
	addOpUpdateDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDestination(options.Region), middleware.Before)
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
			OperationName: "UpdateDestination",
			Err:           err,
		}
	}
	out := result.(*UpdateDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDestinationInput struct {
	// Describes an update to the specified HTTP endpoint destination.
	HttpEndpointDestinationUpdate *types.HttpEndpointDestinationUpdate
	// The name of the delivery stream.
	DeliveryStreamName *string
	// Describes an update for a destination in Amazon ES.
	ElasticsearchDestinationUpdate *types.ElasticsearchDestinationUpdate
	// Describes an update for a destination in Amazon S3.
	ExtendedS3DestinationUpdate *types.ExtendedS3DestinationUpdate
	// [Deprecated] Describes an update for a destination in Amazon S3.
	S3DestinationUpdate *types.S3DestinationUpdate
	// Describes an update for a destination in Amazon Redshift.
	RedshiftDestinationUpdate *types.RedshiftDestinationUpdate
	// Obtain this value from the VersionId result of DeliveryStreamDescription ().
	// This value is required, and helps the service perform conditional operations.
	// For example, if there is an interleaving update and this value is null, then the
	// update destination fails. After the update is successful, the VersionId value is
	// updated. The service then performs a merge of the old configuration with the new
	// configuration.
	CurrentDeliveryStreamVersionId *string
	// Describes an update for a destination in Splunk.
	SplunkDestinationUpdate *types.SplunkDestinationUpdate
	// The ID of the destination.
	DestinationId *string
}

type UpdateDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDestination{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "UpdateDestination",
	}
}
