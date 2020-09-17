// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a robot with a fleet.
func (c *Client) RegisterRobot(ctx context.Context, params *RegisterRobotInput, optFns ...func(*Options)) (*RegisterRobotOutput, error) {
	stack := middleware.NewStack("RegisterRobot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegisterRobotMiddlewares(stack)
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
	addOpRegisterRobotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterRobot(options.Region), middleware.Before)

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
			OperationName: "RegisterRobot",
			Err:           err,
		}
	}
	out := result.(*RegisterRobotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterRobotInput struct {
	// The Amazon Resource Name (ARN) of the robot.
	Robot *string
	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string
}

type RegisterRobotOutput struct {
	// The Amazon Resource Name (ARN) of the fleet that the robot will join.
	Fleet *string
	// Information about the robot registration.
	Robot *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegisterRobotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegisterRobot{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterRobot{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterRobot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "RegisterRobot",
	}
}