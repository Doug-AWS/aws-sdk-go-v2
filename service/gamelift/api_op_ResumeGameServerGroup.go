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

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Reinstates activity on a game
// server group after it has been suspended. A game server group may be suspended
// by calling SuspendGameServerGroup (), or it may have been involuntarily
// suspended due to a configuration problem. You can manually resume activity on
// the group once the configuration problem has been resolved. Refer to the game
// server group status and status reason for more information on why group activity
// is suspended. To resume activity, specify a game server group ARN and the type
// of activity to be resumed. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * CreateGameServerGroup ()
//
//     * ListGameServerGroups
// ()
//
//     * DescribeGameServerGroup ()
//
//     * UpdateGameServerGroup ()
//
//     *
// DeleteGameServerGroup ()
//
//     * ResumeGameServerGroup ()
//
//     *
// SuspendGameServerGroup ()
func (c *Client) ResumeGameServerGroup(ctx context.Context, params *ResumeGameServerGroupInput, optFns ...func(*Options)) (*ResumeGameServerGroupOutput, error) {
	stack := middleware.NewStack("ResumeGameServerGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResumeGameServerGroupMiddlewares(stack)
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
	addOpResumeGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeGameServerGroup(options.Region), middleware.Before)
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
			OperationName: "ResumeGameServerGroup",
			Err:           err,
		}
	}
	out := result.(*ResumeGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeGameServerGroupInput struct {
	// The action to resume for this game server group.
	ResumeActions []types.GameServerGroupAction
	// The unique identifier of the game server group to resume activity on. Use either
	// the GameServerGroup () name or ARN value.
	GameServerGroupName *string
}

type ResumeGameServerGroupOutput struct {
	// An object that describes the game server group resource, with the
	// SuspendedActions property updated to reflect the resumed activity.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResumeGameServerGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResumeGameServerGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeGameServerGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opResumeGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ResumeGameServerGroup",
	}
}
