// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

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

// Stops logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func (c *Client) DisableLogging(ctx context.Context, params *DisableLoggingInput, optFns ...func(*Options)) (*DisableLoggingOutput, error) {
	stack := middleware.NewStack("DisableLogging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDisableLoggingMiddlewares(stack)
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
	addOpDisableLoggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableLogging(options.Region), middleware.Before)
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
			OperationName: "DisableLogging",
			Err:           err,
		}
	}
	out := result.(*DisableLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DisableLoggingInput struct {
	// The identifier of the cluster on which logging is to be stopped. Example:
	// examplecluster
	ClusterIdentifier *string
}

// Describes the status of logging for a cluster.
type DisableLoggingOutput struct {
	// The prefix applied to the log file names.
	S3KeyPrefix *string
	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string
	// true if logging is on, false if logging is off.
	LoggingEnabled *bool
	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time
	// The name of the S3 bucket where the log files are stored.
	BucketName *string
	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDisableLoggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDisableLogging{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableLogging{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableLogging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DisableLogging",
	}
}
