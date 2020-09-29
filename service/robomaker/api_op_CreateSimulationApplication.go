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

// Creates a simulation application.
func (c *Client) CreateSimulationApplication(ctx context.Context, params *CreateSimulationApplicationInput, optFns ...func(*Options)) (*CreateSimulationApplicationOutput, error) {
	stack := middleware.NewStack("CreateSimulationApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateSimulationApplicationMiddlewares(stack)
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
	addOpCreateSimulationApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSimulationApplication(options.Region), middleware.Before)
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
			OperationName: "CreateSimulationApplication",
			Err:           err,
		}
	}
	out := result.(*CreateSimulationApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSimulationApplicationInput struct {
	// A map that contains tag keys and tag values that are attached to the simulation
	// application.
	Tags map[string]*string
	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite
	// The sources of the simulation application.
	Sources []*types.SourceConfig
	// The robot software suite (ROS distribution) used by the simulation application.
	RobotSoftwareSuite *types.RobotSoftwareSuite
	// The name of the simulation application.
	Name *string
}

type CreateSimulationApplicationOutput struct {
	// The list of all tags added to the simulation application.
	Tags map[string]*string
	// The name of the simulation application.
	Name *string
	// The version of the simulation application.
	Version *string
	// The revision id of the simulation application.
	RevisionId *string
	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite
	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite
	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time
	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine
	// The sources of the simulation application.
	Sources []*types.Source

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateSimulationApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateSimulationApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSimulationApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSimulationApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateSimulationApplication",
	}
}
