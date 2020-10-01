// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a subscription definition version.
func (c *Client) GetSubscriptionDefinitionVersion(ctx context.Context, params *GetSubscriptionDefinitionVersionInput, optFns ...func(*Options)) (*GetSubscriptionDefinitionVersionOutput, error) {
	stack := middleware.NewStack("GetSubscriptionDefinitionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSubscriptionDefinitionVersionMiddlewares(stack)
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
	addOpGetSubscriptionDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscriptionDefinitionVersion(options.Region), middleware.Before)
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
			OperationName: "GetSubscriptionDefinitionVersion",
			Err:           err,
		}
	}
	out := result.(*GetSubscriptionDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionDefinitionVersionInput struct {

	// The ID of the subscription definition.
	//
	// This member is required.
	SubscriptionDefinitionId *string

	// The ID of the subscription definition version. This value maps to the
	// ''Version'' property of the corresponding ''VersionInformation'' object, which
	// is returned by ''ListSubscriptionDefinitionVersions'' requests. If the version
	// is the last one that was associated with a subscription definition, the value
	// also maps to the ''LatestVersion'' property of the corresponding
	// ''DefinitionInformation'' object.
	//
	// This member is required.
	SubscriptionDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type GetSubscriptionDefinitionVersionOutput struct {

	// The ARN of the subscription definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the subscription definition
	// version was created.
	CreationTimestamp *string

	// Information about the subscription definition version.
	Definition *types.SubscriptionDefinitionVersion

	// The ID of the subscription definition version.
	Id *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The version of the subscription definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSubscriptionDefinitionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscriptionDefinitionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscriptionDefinitionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSubscriptionDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetSubscriptionDefinitionVersion",
	}
}
