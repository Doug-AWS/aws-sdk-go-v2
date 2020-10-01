// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the configured values for the user enrollment invitation email
// template.
func (c *Client) GetInvitationConfiguration(ctx context.Context, params *GetInvitationConfigurationInput, optFns ...func(*Options)) (*GetInvitationConfigurationOutput, error) {
	stack := middleware.NewStack("GetInvitationConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInvitationConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInvitationConfiguration(options.Region), middleware.Before)
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
			OperationName: "GetInvitationConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetInvitationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInvitationConfigurationInput struct {
}

type GetInvitationConfigurationOutput struct {

	// The email ID of the organization or individual contact that the enrolled user
	// can use.
	ContactEmail *string

	// The name of the organization sending the enrollment invite to a user.
	OrganizationName *string

	// The list of private skill IDs that you want to recommend to the user to enable
	// in the invitation.
	PrivateSkillIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInvitationConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInvitationConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInvitationConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInvitationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "GetInvitationConfiguration",
	}
}
