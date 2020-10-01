// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about the current status and configuration settings for an
// Amazon Macie account.
func (c *Client) GetMacieSession(ctx context.Context, params *GetMacieSessionInput, optFns ...func(*Options)) (*GetMacieSessionOutput, error) {
	stack := middleware.NewStack("GetMacieSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetMacieSessionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMacieSession(options.Region), middleware.Before)
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
			OperationName: "GetMacieSession",
			Err:           err,
		}
	}
	out := result.(*GetMacieSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMacieSessionInput struct {
}

type GetMacieSessionOutput struct {

	// The date and time, in UTC and extended ISO 8601 format, when the Amazon Macie
	// account was created.
	CreatedAt *time.Time

	// The frequency with which Amazon Macie publishes updates to policy findings for
	// the account. This includes publishing updates to AWS Security Hub and Amazon
	// EventBridge (formerly called Amazon CloudWatch Events).
	FindingPublishingFrequency types.FindingPublishingFrequency

	// The Amazon Resource Name (ARN) of the service-linked role that allows Amazon
	// Macie to monitor and analyze data in AWS resources for the account.
	ServiceRole *string

	// The current status of the Amazon Macie account. Possible values are: PAUSED, the
	// account is enabled but all Amazon Macie activities are suspended (paused) for
	// the account; and, ENABLED, the account is enabled and all Amazon Macie
	// activities are enabled for the account.
	Status types.MacieStatus

	// The date and time, in UTC and extended ISO 8601 format, of the most recent
	// change to the status of the Amazon Macie account.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetMacieSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetMacieSession{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMacieSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMacieSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetMacieSession",
	}
}
