// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the settings for an existing DB proxy.
func (c *Client) ModifyDBProxy(ctx context.Context, params *ModifyDBProxyInput, optFns ...func(*Options)) (*ModifyDBProxyOutput, error) {
	stack := middleware.NewStack("ModifyDBProxy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBProxyMiddlewares(stack)
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
	addOpModifyDBProxyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBProxy(options.Region), middleware.Before)
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
			OperationName: "ModifyDBProxy",
			Err:           err,
		}
	}
	out := result.(*ModifyDBProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBProxyInput struct {
	// Whether Transport Layer Security (TLS) encryption is required for connections to
	// the proxy. By enabling this setting, you can enforce encrypted TLS connections
	// to the proxy, even if the associated database doesn't use TLS.
	RequireTLS *bool
	// The new list of security groups for the DBProxy.
	SecurityGroups []*string
	// The new identifier for the DBProxy. An identifier must begin with a letter and
	// must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen
	// or contain two consecutive hyphens.
	NewDBProxyName *string
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access
	// secrets in AWS Secrets Manager.
	RoleArn *string
	// Whether the proxy includes detailed information about SQL statements in its
	// logs. This information helps you to debug issues involving SQL behavior or the
	// performance and scalability of the proxy connections. The debug information
	// includes the text of SQL statements that you submit through the proxy. Thus,
	// only enable this setting when needed for debugging, and only when you have
	// security measures in place to safeguard any sensitive information that appears
	// in the logs.
	DebugLogging *bool
	// The number of seconds that a connection to the proxy can be inactive before the
	// proxy disconnects it. You can set this value higher or lower than the connection
	// timeout limit for the associated database.
	IdleClientTimeout *int32
	// The identifier for the DBProxy to modify.
	DBProxyName *string
	// The new authentication settings for the DBProxy.
	Auth []*types.UserAuthConfig
}

type ModifyDBProxyOutput struct {
	// The DBProxy object representing the new settings for the proxy.
	DBProxy *types.DBProxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBProxyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBProxy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBProxy{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBProxy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBProxy",
	}
}
