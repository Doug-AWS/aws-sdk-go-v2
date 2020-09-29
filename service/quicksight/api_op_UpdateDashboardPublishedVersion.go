// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the published version of a dashboard.
func (c *Client) UpdateDashboardPublishedVersion(ctx context.Context, params *UpdateDashboardPublishedVersionInput, optFns ...func(*Options)) (*UpdateDashboardPublishedVersionOutput, error) {
	stack := middleware.NewStack("UpdateDashboardPublishedVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDashboardPublishedVersionMiddlewares(stack)
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
	addOpUpdateDashboardPublishedVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDashboardPublishedVersion(options.Region), middleware.Before)
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
			OperationName: "UpdateDashboardPublishedVersion",
			Err:           err,
		}
	}
	out := result.(*UpdateDashboardPublishedVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDashboardPublishedVersionInput struct {
	// The version number of the dashboard.
	VersionNumber *int64
	// The ID of the AWS account that contains the dashboard that you're updating.
	AwsAccountId *string
	// The ID for the dashboard.
	DashboardId *string
}

type UpdateDashboardPublishedVersionOutput struct {
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string
	// The ID for the dashboard.
	DashboardId *string
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDashboardPublishedVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDashboardPublishedVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDashboardPublishedVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDashboardPublishedVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDashboardPublishedVersion",
	}
}
