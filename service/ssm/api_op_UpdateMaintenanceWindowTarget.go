// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the target of an existing maintenance window. You can change the
// following:  <ul> <li> <p>Name</p> </li> <li> <p>Description</p> </li> <li>
// <p>Owner</p> </li> <li> <p>IDs for an ID target</p> </li> <li> <p>Tags for a Tag
// target</p> </li> <li> <p>From any supported tag type to another. The three
// supported tag types are ID target, Tag target, and resource group. For more
// information, see <a>Target</a>.</p> </li> </ul> <note> <p>If a parameter is
// null, then the corresponding field is not modified.</p> </note>
func (c *Client) UpdateMaintenanceWindowTarget(ctx context.Context, params *UpdateMaintenanceWindowTargetInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowTargetOutput, error) {
	stack := middleware.NewStack("UpdateMaintenanceWindowTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateMaintenanceWindowTargetMiddlewares(stack)
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
	addOpUpdateMaintenanceWindowTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceWindowTarget(options.Region), middleware.Before)
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
			OperationName: "UpdateMaintenanceWindowTarget",
			Err:           err,
		}
	}
	out := result.(*UpdateMaintenanceWindowTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMaintenanceWindowTargetInput struct {

	// The maintenance window ID with which to modify the target.
	//
	// This member is required.
	WindowId *string

	// The target ID to modify.
	//
	// This member is required.
	WindowTargetId *string

	// An optional description for the update.
	Description *string

	// A name for the update.
	Name *string

	// User-provided value that will be included in any CloudWatch events raised while
	// running tasks for these targets in this maintenance window.
	OwnerInformation *string

	// If True, then all fields that are required by the
	// RegisterTargetWithMaintenanceWindow action are also required for this API
	// request. Optional fields that are not specified are set to null.
	Replace *bool

	// The targets to add or replace.
	Targets []*types.Target
}

type UpdateMaintenanceWindowTargetOutput struct {

	// The updated description.
	Description *string

	// The updated name.
	Name *string

	// The updated owner.
	OwnerInformation *string

	// The updated targets.
	Targets []*types.Target

	// The maintenance window ID specified in the update request.
	WindowId *string

	// The target ID specified in the update request.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateMaintenanceWindowTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceWindowTarget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceWindowTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMaintenanceWindowTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateMaintenanceWindowTarget",
	}
}
