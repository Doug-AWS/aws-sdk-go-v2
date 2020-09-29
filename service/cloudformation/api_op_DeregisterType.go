// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a type or type version from active use in the CloudFormation registry.
// If a type or type version is deregistered, it cannot be used in CloudFormation
// operations. To deregister a type, you must individually deregister all
// registered versions of that type. If a type has only a single registered
// version, deregistering that version results in the type itself being
// deregistered. You cannot deregister the default version of a type, unless it is
// the only registered version of that type, in which case the type itself is
// deregistered as well.
func (c *Client) DeregisterType(ctx context.Context, params *DeregisterTypeInput, optFns ...func(*Options)) (*DeregisterTypeOutput, error) {
	stack := middleware.NewStack("DeregisterType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeregisterTypeMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterType(options.Region), middleware.Before)
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
			OperationName: "DeregisterType",
			Err:           err,
		}
	}
	out := result.(*DeregisterTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTypeInput struct {
	// The Amazon Resource Name (ARN) of the type. Conditional: You must specify either
	// TypeName and Type, or Arn.
	Arn *string
	// The name of the type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	TypeName *string
	// The kind of type. Currently the only valid value is RESOURCE. Conditional: You
	// must specify either TypeName and Type, or Arn.
	Type types.RegistryType
	// The ID of a specific version of the type. The version ID is the value at the end
	// of the Amazon Resource Name (ARN) assigned to the type version when it is
	// registered.
	VersionId *string
}

type DeregisterTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeregisterTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeregisterType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeregisterType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeregisterType",
	}
}
