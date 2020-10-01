// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the name and scopes of resource server. All other fields are read-only.
// If you don't provide a value for an attribute, it will be set to the default
// value.
func (c *Client) UpdateResourceServer(ctx context.Context, params *UpdateResourceServerInput, optFns ...func(*Options)) (*UpdateResourceServerOutput, error) {
	stack := middleware.NewStack("UpdateResourceServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateResourceServerMiddlewares(stack)
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
	addOpUpdateResourceServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceServer(options.Region), middleware.Before)
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
			OperationName: "UpdateResourceServer",
			Err:           err,
		}
	}
	out := result.(*UpdateResourceServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceServerInput struct {

	// The identifier for the resource server.
	//
	// This member is required.
	Identifier *string

	// The name of the resource server.
	//
	// This member is required.
	Name *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The scope values to be set for the resource server.
	Scopes []*types.ResourceServerScopeType
}

type UpdateResourceServerOutput struct {

	// The resource server.
	//
	// This member is required.
	ResourceServer *types.ResourceServerType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateResourceServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResourceServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResourceServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateResourceServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateResourceServer",
	}
}
