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

// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API call),
// domain controllers will be added or removed. It may take up to 45 minutes for
// any new domain controllers to become fully active once the requested number of
// domain controllers is updated. During this time, you cannot make another update
// request.
func (c *Client) UpdateNumberOfDomainControllers(ctx context.Context, params *UpdateNumberOfDomainControllersInput, optFns ...func(*Options)) (*UpdateNumberOfDomainControllersOutput, error) {
	stack := middleware.NewStack("UpdateNumberOfDomainControllers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateNumberOfDomainControllersMiddlewares(stack)
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
	addOpUpdateNumberOfDomainControllersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(options.Region), middleware.Before)
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
			OperationName: "UpdateNumberOfDomainControllers",
			Err:           err,
		}
	}
	out := result.(*UpdateNumberOfDomainControllersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNumberOfDomainControllersInput struct {

	// The number of domain controllers desired in the directory.
	//
	// This member is required.
	DesiredNumber *int32

	// Identifier of the directory to which the domain controllers will be added or
	// removed.
	//
	// This member is required.
	DirectoryId *string
}

type UpdateNumberOfDomainControllersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateNumberOfDomainControllersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNumberOfDomainControllers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNumberOfDomainControllers{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "UpdateNumberOfDomainControllers",
	}
}
