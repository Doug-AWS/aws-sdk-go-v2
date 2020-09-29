// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests authorization to create or delete a peer connection between the VPC for
// your Amazon GameLift fleet and a virtual private cloud (VPC) in your AWS
// account. VPC peering enables the game servers on your fleet to communicate
// directly with other AWS resources. Once you've received authorization, call
// CreateVpcPeeringConnection () to establish the peering connection. For more
// information, see VPC Peering with Amazon GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
// You can peer with VPCs that are owned by any AWS account you have access to,
// including the account that you use to manage your Amazon GameLift fleets. You
// cannot peer with VPCs that are in different Regions. To request authorization to
// create a connection, call this operation from the AWS account with the VPC that
// you want to peer to your Amazon GameLift fleet. For example, to enable your game
// servers to retrieve data from a DynamoDB table, use the account that manages
// that DynamoDB resource. Identify the following values: (1) The ID of the VPC
// that you want to peer with, and (2) the ID of the AWS account that you use to
// manage Amazon GameLift. If successful, VPC peering is authorized for the
// specified VPC. To request authorization to delete a connection, call this
// operation from the AWS account with the VPC that is peered with your Amazon
// GameLift fleet. Identify the following values: (1) VPC ID that you want to
// delete the peering connection for, and (2) ID of the AWS account that you use to
// manage Amazon GameLift. The authorization remains valid for 24 hours unless it
// is canceled by a call to DeleteVpcPeeringAuthorization (). You must create or
// delete the peering connection while the authorization is valid.
//
//     *
// CreateVpcPeeringAuthorization ()
//
//     * DescribeVpcPeeringAuthorizations ()
//
//
// * DeleteVpcPeeringAuthorization ()
//
//     * CreateVpcPeeringConnection ()
//
//     *
// DescribeVpcPeeringConnections ()
//
//     * DeleteVpcPeeringConnection ()
func (c *Client) CreateVpcPeeringAuthorization(ctx context.Context, params *CreateVpcPeeringAuthorizationInput, optFns ...func(*Options)) (*CreateVpcPeeringAuthorizationOutput, error) {
	stack := middleware.NewStack("CreateVpcPeeringAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateVpcPeeringAuthorizationMiddlewares(stack)
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
	addOpCreateVpcPeeringAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringAuthorization(options.Region), middleware.Before)
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
			OperationName: "CreateVpcPeeringAuthorization",
			Err:           err,
		}
	}
	out := result.(*CreateVpcPeeringAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateVpcPeeringAuthorizationInput struct {
	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering with
	// Amazon GameLift Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string
	// A unique identifier for the AWS account that you use to manage your Amazon
	// GameLift fleet. You can find your Account ID in the AWS Management Console under
	// account settings.
	GameLiftAwsAccountId *string
}

// Represents the returned data in response to a request action.
type CreateVpcPeeringAuthorizationOutput struct {
	// Details on the requested VPC peering authorization, including expiration.
	VpcPeeringAuthorization *types.VpcPeeringAuthorization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateVpcPeeringAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVpcPeeringAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVpcPeeringAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVpcPeeringAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateVpcPeeringAuthorization",
	}
}
