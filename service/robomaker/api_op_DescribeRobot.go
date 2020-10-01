// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a robot.
func (c *Client) DescribeRobot(ctx context.Context, params *DescribeRobotInput, optFns ...func(*Options)) (*DescribeRobotOutput, error) {
	stack := middleware.NewStack("DescribeRobot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRobotMiddlewares(stack)
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
	addOpDescribeRobotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRobot(options.Region), middleware.Before)
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
			OperationName: "DescribeRobot",
			Err:           err,
		}
	}
	out := result.(*DescribeRobotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRobotInput struct {

	// The Amazon Resource Name (ARN) of the robot to be described.
	//
	// This member is required.
	Robot *string
}

type DescribeRobotOutput struct {

	// The target architecture of the robot application.
	Architecture types.Architecture

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string

	// The Greengrass group id.
	GreengrassGroupId *string

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The time of the last deployment job.
	LastDeploymentTime *time.Time

	// The name of the robot.
	Name *string

	// The status of the fleet.
	Status types.RobotStatus

	// The list of all tags added to the specified robot.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRobotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRobot{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRobot{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRobot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeRobot",
	}
}
