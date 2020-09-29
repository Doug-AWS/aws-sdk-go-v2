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

// Describes a simulation application.
func (c *Client) DescribeSimulationApplication(ctx context.Context, params *DescribeSimulationApplicationInput, optFns ...func(*Options)) (*DescribeSimulationApplicationOutput, error) {
	stack := middleware.NewStack("DescribeSimulationApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeSimulationApplicationMiddlewares(stack)
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
	addOpDescribeSimulationApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSimulationApplication(options.Region), middleware.Before)
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
			OperationName: "DescribeSimulationApplication",
			Err:           err,
		}
	}
	out := result.(*DescribeSimulationApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSimulationApplicationInput struct {
	// The application information for the simulation application.
	Application *string
	// The version of the simulation application to describe.
	ApplicationVersion *string
}

type DescribeSimulationApplicationOutput struct {
	// The sources of the simulation application.
	Sources []*types.Source
	// The revision id of the simulation application.
	RevisionId *string
	// The version of the simulation application.
	Version *string
	// The list of all tags added to the specified simulation application.
	Tags map[string]*string
	// The Amazon Resource Name (ARN) of the robot simulation application.
	Arn *string
	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite
	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine
	// The name of the simulation application.
	Name *string
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite
	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeSimulationApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSimulationApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSimulationApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSimulationApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeSimulationApplication",
	}
}
