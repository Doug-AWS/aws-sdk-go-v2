// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a broker. Note: This API is asynchronous.
func (c *Client) CreateBroker(ctx context.Context, params *CreateBrokerInput, optFns ...func(*Options)) (*CreateBrokerOutput, error) {
	stack := middleware.NewStack("CreateBroker", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateBrokerMiddlewares(stack)
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
	addIdempotencyToken_opCreateBrokerMiddleware(stack, options)
	addOpCreateBrokerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBroker(options.Region), middleware.Before)
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
			OperationName: "CreateBroker",
			Err:           err,
		}
	}
	out := result.(*CreateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a broker using the specified properties.
type CreateBrokerInput struct {

	// The authentication strategy used to secure the broker.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. Enables automatic upgrades to new minor versions for brokers, as
	// Apache releases the versions. The automatic upgrades occur during the
	// maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade *bool

	// Required. The name of the broker. This value must be unique in your AWS account,
	// 1-50 characters long, must contain only letters, numbers, dashes, and
	// underscores, and must not contain whitespaces, brackets, wildcard characters, or
	// special characters.
	BrokerName *string

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// The unique ID that the requester receives for the created broker. Amazon MQ
	// passes your ID with the API action. Note: We recommend using a Universally
	// Unique Identifier (UUID) for the creatorRequestId. You may omit the
	// creatorRequestId if your application doesn't require idempotency.
	CreatorRequestId *string

	// Required. The deployment mode of the broker.
	DeploymentMode types.DeploymentMode

	// Encryption options for the broker.
	EncryptionOptions *types.EncryptionOptions

	// Required. The type of broker engine. Note: Currently, Amazon MQ supports only
	// ACTIVEMQ.
	EngineType types.EngineType

	// Required. The version of the broker engine. For a list of supported engine
	// versions, see
	// https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// Required. The broker's instance type.
	HostInstanceType *string

	// The metadata of the LDAP server used to authenticate and authorize connections
	// to the broker.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// Required. Enables connections from applications outside of the VPC that hosts
	// the broker's subnets.
	PubliclyAccessible *bool

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []*string

	// The broker's storage type.
	StorageType types.BrokerStorageType

	// The list of groups (2 maximum) that define which subnets and IP ranges the
	// broker can use from different Availability Zones. A SINGLE_INSTANCE deployment
	// requires one subnet (for example, the default subnet). An
	// ACTIVE_STANDBY_MULTI_AZ deployment requires two subnets.
	SubnetIds []*string

	// Create tags when creating the broker.
	Tags map[string]*string

	// Required. The list of ActiveMQ users (persons or applications) who can access
	// queues and topics. This value can contain only alphanumeric characters, dashes,
	// periods, underscores, and tildes (- . _ ~). This value must be 2-100 characters
	// long.
	Users []*types.User
}

type CreateBrokerOutput struct {

	// The Amazon Resource Name (ARN) of the broker.
	BrokerArn *string

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateBrokerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateBroker{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBroker{}, middleware.After)
}

type idempotencyToken_initializeOpCreateBroker struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBroker) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBrokerInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBrokerMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateBroker{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateBroker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "CreateBroker",
	}
}
