// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves violations for a resource based on the specified AWS Firewall Manager
// policy and AWS account.
func (c *Client) GetViolationDetails(ctx context.Context, params *GetViolationDetailsInput, optFns ...func(*Options)) (*GetViolationDetailsOutput, error) {
	stack := middleware.NewStack("GetViolationDetails", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetViolationDetailsMiddlewares(stack)
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
	addOpGetViolationDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetViolationDetails(options.Region), middleware.Before)

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
			OperationName: "GetViolationDetails",
			Err:           err,
		}
	}
	out := result.(*GetViolationDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetViolationDetailsInput struct {
	// The ID of the resource that has violations.
	ResourceId *string
	// The ID of the AWS Firewall Manager policy that you want the details for. This
	// currently only supports security group content audit policies.
	PolicyId *string
	// The AWS account ID that you want the details for.
	MemberAccount *string
	// The resource type. This is in the format shown in the AWS Resource Types
	// Reference
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	// Supported resource types are: AWS::EC2::Instance, AWS::EC2::NetworkInterface, or
	// AWS::EC2::SecurityGroup.
	ResourceType *string
}

type GetViolationDetailsOutput struct {
	// Violation detail for a resource.
	ViolationDetail *types.ViolationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetViolationDetailsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetViolationDetails{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetViolationDetails{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetViolationDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetViolationDetails",
	}
}