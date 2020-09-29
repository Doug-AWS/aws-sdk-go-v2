// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides identifying details of the resource being migrated so that it can be
// associated in the Application Discovery Service repository. This association
// occurs asynchronously after PutResourceAttributes returns.
//
//     * Keep in mind
// that subsequent calls to PutResourceAttributes will override previously stored
// attributes. For example, if it is first called with a MAC address, but later, it
// is desired to add an IP address, it will then be required to call it with both
// the IP and MAC addresses to prevent overriding the MAC address.
//
//     * Note the
// instructions regarding the special use case of the ResourceAttributeList
// (https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#migrationhub-PutResourceAttributes-request-ResourceAttributeList)
// parameter when specifying any "VM" related value.
//
//     <note> <p>Because this is
// an asynchronous call, it will always return 200, whether an association occurs
// or not. To confirm if an association was found based on the provided details,
// call <code>ListDiscoveredResources</code>.</p> </note>
func (c *Client) PutResourceAttributes(ctx context.Context, params *PutResourceAttributesInput, optFns ...func(*Options)) (*PutResourceAttributesOutput, error) {
	stack := middleware.NewStack("PutResourceAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutResourceAttributesMiddlewares(stack)
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
	addOpPutResourceAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourceAttributes(options.Region), middleware.Before)
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
			OperationName: "PutResourceAttributes",
			Err:           err,
		}
	}
	out := result.(*PutResourceAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourceAttributesInput struct {
	// The name of the ProgressUpdateStream.
	ProgressUpdateStream *string
	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	MigrationTaskName *string
	// Information about the resource that is being migrated. This data will be used to
	// map the task to a resource in the Application Discovery Service repository.
	// Takes the object array of ResourceAttribute where the Type field is reserved for
	// the following values: IPV4_ADDRESS | IPV6_ADDRESS | MAC_ADDRESS | FQDN |
	// VM_MANAGER_ID | VM_MANAGED_OBJECT_REFERENCE | VM_NAME | VM_PATH | BIOS_ID |
	// MOTHERBOARD_SERIAL_NUMBER where the identifying value can be a string up to 256
	// characters.
	//
	//     * <p>If any "VM" related value is set for a
	// <code>ResourceAttribute</code> object, it is required that
	// <code>VM_MANAGER_ID</code>, as a minimum, is always set. If
	// <code>VM_MANAGER_ID</code> is not set, then all "VM" fields will be discarded
	// and "VM" fields will not be used for matching the migration task to a server in
	// Application Discovery Service repository. See the <a
	// href="https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#API_PutResourceAttributes_Examples">Example</a>
	// section below for a use case of specifying "VM" related values.</p> </li> <li>
	// <p> If a server you are trying to match has multiple IP or MAC addresses, you
	// should provide as many as you know in separate type/value pairs passed to the
	// <code>ResourceAttributeList</code> parameter to maximize the chances of
	// matching.</p> </li> </ul> </important>
	ResourceAttributeList []*types.ResourceAttribute
	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type PutResourceAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutResourceAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourceAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourceAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutResourceAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "PutResourceAttributes",
	}
}
