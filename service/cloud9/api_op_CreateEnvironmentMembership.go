// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an environment member to an AWS Cloud9 development environment.
func (c *Client) CreateEnvironmentMembership(ctx context.Context, params *CreateEnvironmentMembershipInput, optFns ...func(*Options)) (*CreateEnvironmentMembershipOutput, error) {
	stack := middleware.NewStack("CreateEnvironmentMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateEnvironmentMembershipMiddlewares(stack)
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
	addOpCreateEnvironmentMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentMembership(options.Region), middleware.Before)
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
			OperationName: "CreateEnvironmentMembership",
			Err:           err,
		}
	}
	out := result.(*CreateEnvironmentMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentMembershipInput struct {

	// The ID of the environment that contains the environment member you want to add.
	//
	// This member is required.
	EnvironmentId *string

	// The type of environment member permissions you want to associate with this
	// environment member. Available values include:
	//
	//     * read-only: Has read-only
	// access to the environment.
	//
	//     * read-write: Has read-write access to the
	// environment.
	//
	// This member is required.
	Permissions types.MemberPermissions

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	//
	// This member is required.
	UserArn *string
}

type CreateEnvironmentMembershipOutput struct {

	// Information about the environment member that was added.
	Membership *types.EnvironmentMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateEnvironmentMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEnvironmentMembership{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEnvironmentMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEnvironmentMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "CreateEnvironmentMembership",
	}
}
