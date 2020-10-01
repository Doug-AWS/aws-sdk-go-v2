// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of names for one or more on-premises instances. Unless otherwise
// specified, both registered and deregistered on-premises instance names are
// listed. To list only registered or deregistered on-premises instance names, use
// the registration status parameter.
func (c *Client) ListOnPremisesInstances(ctx context.Context, params *ListOnPremisesInstancesInput, optFns ...func(*Options)) (*ListOnPremisesInstancesOutput, error) {
	stack := middleware.NewStack("ListOnPremisesInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListOnPremisesInstancesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOnPremisesInstances(options.Region), middleware.Before)
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
			OperationName: "ListOnPremisesInstances",
			Err:           err,
		}
	}
	out := result.(*ListOnPremisesInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListOnPremisesInstances operation.
type ListOnPremisesInstancesInput struct {

	// An identifier returned from the previous list on-premises instances call. It can
	// be used to return the next set of on-premises instances in the list.
	NextToken *string

	// The registration status of the on-premises instances:
	//
	//     * Deregistered:
	// Include deregistered on-premises instances in the resulting list.
	//
	//     *
	// Registered: Include registered on-premises instances in the resulting list.
	RegistrationStatus types.RegistrationStatus

	// The on-premises instance tags that are used to restrict the on-premises instance
	// names returned.
	TagFilters []*types.TagFilter
}

// Represents the output of the list on-premises instances operation.
type ListOnPremisesInstancesOutput struct {

	// The list of matching on-premises instance names.
	InstanceNames []*string

	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list on-premises instances call to return the next
	// set of on-premises instances in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListOnPremisesInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListOnPremisesInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOnPremisesInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOnPremisesInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListOnPremisesInstances",
	}
}
