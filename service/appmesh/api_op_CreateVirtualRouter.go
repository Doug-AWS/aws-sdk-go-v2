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

// Creates a virtual router within a service mesh. Specify a listener for any
// inbound traffic that your virtual router receives. Create a virtual router for
// each protocol and port that you need to route. Virtual routers handle traffic
// for one or more virtual services within your mesh. After you create your virtual
// router, create and associate routes for your virtual router that direct incoming
// requests to different virtual nodes. For more information about virtual routers,
// see Virtual routers
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_routers.html).
func (c *Client) CreateVirtualRouter(ctx context.Context, params *CreateVirtualRouterInput, optFns ...func(*Options)) (*CreateVirtualRouterOutput, error) {
	stack := middleware.NewStack("CreateVirtualRouter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateVirtualRouterMiddlewares(stack)
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
	addOpCreateVirtualRouterValidationMiddleware(stack)
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
			OperationName: "CreateVirtualRouter",
			Err:           err,
		}
	}
	out := result.(*CreateVirtualRouterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateVirtualRouterInput struct {
	// The name to use for the virtual router.
	VirtualRouterName *string
	// The name of the service mesh to create the virtual router in.
	MeshName *string
	// The virtual router specification to apply.
	Spec *types.VirtualRouterSpec
	// Optional metadata that you can apply to the virtual router to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.TagRef
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string
	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type CreateVirtualRouterOutput struct {
	// The full description of your virtual router following the create call.
	VirtualRouter *types.VirtualRouterData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateVirtualRouterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualRouter{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualRouter{}, middleware.After)
}
