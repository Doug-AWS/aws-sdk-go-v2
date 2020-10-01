// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the specified managed prefix list. Adding or removing entries in a
// prefix list creates a new version of the prefix list. Changing the name of the
// prefix list does not affect the version. If you specify a current version number
// that does not match the true current version number, the request fails.
func (c *Client) ModifyManagedPrefixList(ctx context.Context, params *ModifyManagedPrefixListInput, optFns ...func(*Options)) (*ModifyManagedPrefixListOutput, error) {
	stack := middleware.NewStack("ModifyManagedPrefixList", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyManagedPrefixListMiddlewares(stack)
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
	addOpModifyManagedPrefixListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyManagedPrefixList(options.Region), middleware.Before)
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
			OperationName: "ModifyManagedPrefixList",
			Err:           err,
		}
	}
	out := result.(*ModifyManagedPrefixListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyManagedPrefixListInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// One or more entries to add to the prefix list.
	AddEntries []*types.AddPrefixListEntry

	// The current version of the prefix list.
	CurrentVersion *int64

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// A name for the prefix list.
	PrefixListName *string

	// One or more entries to remove from the prefix list.
	RemoveEntries []*types.RemovePrefixListEntry
}

type ModifyManagedPrefixListOutput struct {

	// Information about the prefix list.
	PrefixList *types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyManagedPrefixListMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyManagedPrefixList{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyManagedPrefixList{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyManagedPrefixList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyManagedPrefixList",
	}
}
