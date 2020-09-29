// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the recovery points for a specified gateway. This operation is only
// supported in the cached volume gateway type.  <p>Each cache volume has one
// recovery point. A volume recovery point is a point in time at which all data of
// the volume is consistent and from which you can create a snapshot or clone a new
// cached volume from a source volume. To create a snapshot from a volume recovery
// point use the <a>CreateSnapshotFromVolumeRecoveryPoint</a> operation.</p>
func (c *Client) ListVolumeRecoveryPoints(ctx context.Context, params *ListVolumeRecoveryPointsInput, optFns ...func(*Options)) (*ListVolumeRecoveryPointsOutput, error) {
	stack := middleware.NewStack("ListVolumeRecoveryPoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListVolumeRecoveryPointsMiddlewares(stack)
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
	addOpListVolumeRecoveryPointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumeRecoveryPoints(options.Region), middleware.Before)
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
			OperationName: "ListVolumeRecoveryPoints",
			Err:           err,
		}
	}
	out := result.(*ListVolumeRecoveryPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVolumeRecoveryPointsInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type ListVolumeRecoveryPointsOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// An array of VolumeRecoveryPointInfo () objects.
	VolumeRecoveryPointInfos []*types.VolumeRecoveryPointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListVolumeRecoveryPointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumeRecoveryPoints{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumeRecoveryPoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVolumeRecoveryPoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumeRecoveryPoints",
	}
}
