// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates a VPC link.
func (c *Client) UpdateVpcLink(ctx context.Context, params *UpdateVpcLinkInput, optFns ...func(*Options)) (*UpdateVpcLinkOutput, error) {
	stack := middleware.NewStack("UpdateVpcLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateVpcLinkMiddlewares(stack)
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
	addOpUpdateVpcLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVpcLink(options.Region), middleware.Before)
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
			OperationName: "UpdateVpcLink",
			Err:           err,
		}
	}
	out := result.(*UpdateVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a VPC link.
type UpdateVpcLinkInput struct {
	// The ID of the VPC link.
	VpcLinkId *string
	// The name of the VPC link.
	Name *string
}

type UpdateVpcLinkOutput struct {
	// The name of the VPC link.
	Name *string
	// The version of the VPC link.
	VpcLinkVersion types.VpcLinkVersion
	// A message summarizing the cause of the status of the VPC link.
	VpcLinkStatusMessage *string
	// The timestamp when the VPC link was created.
	CreatedDate *time.Time
	// A list of security group IDs for the VPC link.
	SecurityGroupIds []*string
	// The status of the VPC link.
	VpcLinkStatus types.VpcLinkStatus
	// Tags for the VPC link.
	Tags map[string]*string
	// The ID of the VPC link.
	VpcLinkId *string
	// A list of subnet IDs to include in the VPC link.
	SubnetIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateVpcLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVpcLink{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVpcLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateVpcLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateVpcLink",
	}
}
