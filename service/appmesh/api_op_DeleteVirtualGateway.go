// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing virtual gateway. You cannot delete a virtual gateway if any
// gateway routes are associated to it.
func (c *Client) DeleteVirtualGateway(ctx context.Context, params *DeleteVirtualGatewayInput, optFns ...func(*Options)) (*DeleteVirtualGatewayOutput, error) {
	stack := middleware.NewStack("DeleteVirtualGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteVirtualGatewayMiddlewares(stack)
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
	addOpDeleteVirtualGatewayValidationMiddleware(stack)
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
			OperationName: "DeleteVirtualGateway",
			Err:           err,
		}
	}
	out := result.(*DeleteVirtualGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVirtualGatewayInput struct {

	// The name of the service mesh to delete the virtual gateway from.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual gateway to delete.
	//
	// This member is required.
	VirtualGatewayName *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

type DeleteVirtualGatewayOutput struct {

	// The virtual gateway that was deleted.
	//
	// This member is required.
	VirtualGateway *types.VirtualGatewayData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteVirtualGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVirtualGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVirtualGateway{}, middleware.After)
}
