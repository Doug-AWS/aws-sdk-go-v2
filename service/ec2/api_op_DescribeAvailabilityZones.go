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

// Describes the Availability Zones and Local Zones that are available to you. If
// there is an event impacting an Availability Zone or Local Zone, you can use this
// request to view the state and any provided messages for that Availability Zone
// or Local Zone. For more information about Availability Zones and Local Zones,
// see Regions and Availability Zones
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeAvailabilityZones(ctx context.Context, params *DescribeAvailabilityZonesInput, optFns ...func(*Options)) (*DescribeAvailabilityZonesOutput, error) {
	stack := middleware.NewStack("DescribeAvailabilityZones", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeAvailabilityZonesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailabilityZones(options.Region), middleware.Before)
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
			OperationName: "DescribeAvailabilityZones",
			Err:           err,
		}
	}
	out := result.(*DescribeAvailabilityZonesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAvailabilityZonesInput struct {

	// Include all Availability Zones and Local Zones regardless of your opt in status.
	// If you do not use this parameter, the results include only the zones for the
	// Regions where you have chosen the option to opt in.
	AllAvailabilityZones *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	//     * group-name - For Availability Zones, use the Region name.
	// For Local Zones, use the name of the group associated with the Local Zone (for
	// example, us-west-2-lax-1).
	//
	//     * message - The Zone message.
	//
	//     *
	// opt-in-status - The opt in status (opted-in, and not-opted-in |
	// opt-in-not-required).
	//
	//     * The ID of the zone that handles some of the Local
	// Zone control plane operations, such as API calls.
	//
	//     * region-name - The name
	// of the Region for the Zone (for example, us-east-1).
	//
	//     * state - The state of
	// the Availability Zone or Local Zone (available | information | impaired |
	// unavailable).
	//
	//     * zone-id - The ID of the Availability Zone (for example,
	// use1-az1) or the Local Zone (for example, use usw2-lax1-az1).
	//
	//     * zone-type -
	// The type of zone, for example, local-zone.
	//
	//     * zone-name - The name of the
	// Availability Zone (for example, us-east-1a) or the Local Zone (for example, use
	// us-west-2-lax-1a).
	//
	//     * zone-type - The type of zone, for example, local-zone.
	Filters []*types.Filter

	// The IDs of the Zones.
	ZoneIds []*string

	// The names of the Zones.
	ZoneNames []*string
}

type DescribeAvailabilityZonesOutput struct {

	// Information about the Zones.
	AvailabilityZones []*types.AvailabilityZone

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeAvailabilityZonesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeAvailabilityZones{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAvailabilityZones{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAvailabilityZones(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAvailabilityZones",
	}
}
