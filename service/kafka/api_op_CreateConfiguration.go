// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new MSK configuration.
func (c *Client) CreateConfiguration(ctx context.Context, params *CreateConfigurationInput, optFns ...func(*Options)) (*CreateConfigurationOutput, error) {
	stack := middleware.NewStack("CreateConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateConfigurationMiddlewares(stack)
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
	addOpCreateConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguration(options.Region), middleware.Before)
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
			OperationName: "CreateConfiguration",
			Err:           err,
		}
	}
	out := result.(*CreateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfigurationInput struct {
	// Contents of the server.properties file. When using the API, you must ensure that
	// the contents of the file are base64 encoded. When using the AWS Management
	// Console, the SDK, or the AWS CLI, the contents of server.properties can be in
	// plaintext.
	ServerProperties []byte
	// The versions of Apache Kafka with which you can use this MSK configuration.
	KafkaVersions []*string
	// The description of the configuration.
	Description *string
	// The name of the configuration.
	Name *string
}

type CreateConfigurationOutput struct {
	// The name of the configuration.
	Name *string
	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string
	// The time when the configuration was created.
	CreationTime *time.Time
	// Latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "CreateConfiguration",
	}
}
