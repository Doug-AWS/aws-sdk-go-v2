// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables single-sign on for a directory.
func (c *Client) DisableSso(ctx context.Context, params *DisableSsoInput, optFns ...func(*Options)) (*DisableSsoOutput, error) {
	stack := middleware.NewStack("DisableSso", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableSsoMiddlewares(stack)
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
	addOpDisableSsoValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableSso(options.Region), middleware.Before)
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
			OperationName: "DisableSso",
			Err:           err,
		}
	}
	out := result.(*DisableSsoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DisableSso () operation.
type DisableSsoInput struct {

	// The identifier of the directory for which to disable single-sign on.
	//
	// This member is required.
	DirectoryId *string

	// The password of an alternate account to use to disable single-sign on. This is
	// only used for AD Connector directories. For more information, see the UserName
	// parameter.
	Password *string

	// The username of an alternate account to use to disable single-sign on. This is
	// only used for AD Connector directories. This account must have privileges to
	// remove a service principal name. If the AD Connector service account does not
	// have privileges to remove a service principal name, you can specify an alternate
	// account with the UserName and Password parameters. These credentials are only
	// used to disable single sign-on and are not stored by the service. The AD
	// Connector service account is not changed.
	UserName *string
}

// Contains the results of the DisableSso () operation.
type DisableSsoOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableSsoMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableSso{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableSso{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableSso(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DisableSso",
	}
}
