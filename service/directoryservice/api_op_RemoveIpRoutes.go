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

// Removes IP address blocks from a directory.
func (c *Client) RemoveIpRoutes(ctx context.Context, params *RemoveIpRoutesInput, optFns ...func(*Options)) (*RemoveIpRoutesOutput, error) {
	stack := middleware.NewStack("RemoveIpRoutes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveIpRoutesMiddlewares(stack)
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
	addOpRemoveIpRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveIpRoutes(options.Region), middleware.Before)
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
			OperationName: "RemoveIpRoutes",
			Err:           err,
		}
	}
	out := result.(*RemoveIpRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveIpRoutesInput struct {
	// Identifier (ID) of the directory from which you want to remove the IP addresses.
	DirectoryId *string
	// IP address blocks that you want to remove.
	CidrIps []*string
}

type RemoveIpRoutesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveIpRoutesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveIpRoutes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveIpRoutes{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveIpRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "RemoveIpRoutes",
	}
}
