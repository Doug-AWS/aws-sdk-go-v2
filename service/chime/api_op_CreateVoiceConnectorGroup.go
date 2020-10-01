// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Chime Voice Connector group under the administrator's AWS
// account. You can associate Amazon Chime Voice Connectors with the Amazon Chime
// Voice Connector group by including VoiceConnectorItems in the request. You can
// include Amazon Chime Voice Connectors from different AWS Regions in your group.
// This creates a fault tolerant mechanism for fallback in case of availability
// events.
func (c *Client) CreateVoiceConnectorGroup(ctx context.Context, params *CreateVoiceConnectorGroupInput, optFns ...func(*Options)) (*CreateVoiceConnectorGroupOutput, error) {
	stack := middleware.NewStack("CreateVoiceConnectorGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateVoiceConnectorGroupMiddlewares(stack)
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
	addOpCreateVoiceConnectorGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(options.Region), middleware.Before)
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
			OperationName: "CreateVoiceConnectorGroup",
			Err:           err,
		}
	}
	out := result.(*CreateVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceConnectorGroupInput struct {

	// The name of the Amazon Chime Voice Connector group.
	//
	// This member is required.
	Name *string

	// The Amazon Chime Voice Connectors to route inbound calls to.
	VoiceConnectorItems []*types.VoiceConnectorItem
}

type CreateVoiceConnectorGroupOutput struct {

	// The Amazon Chime Voice Connector group details.
	VoiceConnectorGroup *types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateVoiceConnectorGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceConnectorGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceConnectorGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateVoiceConnectorGroup",
	}
}
