// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Calling this operation enables AWS Health to work with AWS Organizations. This
// applies a Service Linked Role (SLR) to the master account in the organization.
// To learn more about the steps in this process, visit enabling service access for
// AWS Health in AWS Organizations. To call this operation, you must sign in as an
// IAM user, assume an IAM role, or sign in as the root user (not recommended) in
// the organization's master account.
func (c *Client) EnableHealthServiceAccessForOrganization(ctx context.Context, params *EnableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*EnableHealthServiceAccessForOrganizationOutput, error) {
	stack := middleware.NewStack("EnableHealthServiceAccessForOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableHealthServiceAccessForOrganizationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(options.Region), middleware.Before)

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
			OperationName: "EnableHealthServiceAccessForOrganization",
			Err:           err,
		}
	}
	out := result.(*EnableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableHealthServiceAccessForOrganizationInput struct {
}

type EnableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "EnableHealthServiceAccessForOrganization",
	}
}