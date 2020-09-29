// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details about a project in AWS Mobile Hub.
func (c *Client) DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) {
	stack := middleware.NewStack("DescribeProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeProjectMiddlewares(stack)
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
	addOpDescribeProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProject(options.Region), middleware.Before)
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
			OperationName: "DescribeProject",
			Err:           err,
		}
	}
	out := result.(*DescribeProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used to request details about a project.
type DescribeProjectInput struct {
	// Unique project identifier.
	ProjectId *string
	// If set to true, causes AWS Mobile Hub to synchronize information from other
	// services, e.g., update state of AWS CloudFormation stacks in the AWS Mobile Hub
	// project.
	SyncFromResources *bool
}

// Result structure used for requests of project details.
type DescribeProjectOutput struct {
	// Detailed information about an AWS Mobile Hub project.
	Details *types.ProjectDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "DescribeProject",
	}
}
