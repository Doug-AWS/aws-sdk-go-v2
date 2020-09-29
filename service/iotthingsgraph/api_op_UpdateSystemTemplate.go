// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified system. You don't need to run this action after updating a
// workflow. Any deployment that uses the system will see the changes in the system
// when it is redeployed.
func (c *Client) UpdateSystemTemplate(ctx context.Context, params *UpdateSystemTemplateInput, optFns ...func(*Options)) (*UpdateSystemTemplateOutput, error) {
	stack := middleware.NewStack("UpdateSystemTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSystemTemplateMiddlewares(stack)
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
	addOpUpdateSystemTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSystemTemplate(options.Region), middleware.Before)
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
			OperationName: "UpdateSystemTemplate",
			Err:           err,
		}
	}
	out := result.(*UpdateSystemTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSystemTemplateInput struct {
	// The DefinitionDocument that contains the updated system definition.
	Definition *types.DefinitionDocument
	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace. If no value is specified, the latest version is used by
	// default.
	CompatibleNamespaceVersion *int64
	// The ID of the system to be updated. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	Id *string
}

type UpdateSystemTemplateOutput struct {
	// An object containing summary information about the updated system.
	Summary *types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSystemTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSystemTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSystemTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSystemTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "UpdateSystemTemplate",
	}
}
