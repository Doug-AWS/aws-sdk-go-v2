// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the assessment targets that are specified by the ARNs of the
// assessment targets.
func (c *Client) DescribeAssessmentTargets(ctx context.Context, params *DescribeAssessmentTargetsInput, optFns ...func(*Options)) (*DescribeAssessmentTargetsOutput, error) {
	stack := middleware.NewStack("DescribeAssessmentTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAssessmentTargetsMiddlewares(stack)
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
	addOpDescribeAssessmentTargetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssessmentTargets(options.Region), middleware.Before)
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
			OperationName: "DescribeAssessmentTargets",
			Err:           err,
		}
	}
	out := result.(*DescribeAssessmentTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssessmentTargetsInput struct {
	// The ARNs that specifies the assessment targets that you want to describe.
	AssessmentTargetArns []*string
}

type DescribeAssessmentTargetsOutput struct {
	// Assessment target details that cannot be described. An error code is provided
	// for each failed item.
	FailedItems map[string]*types.FailedItemDetails
	// Information about the assessment targets.
	AssessmentTargets []*types.AssessmentTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAssessmentTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssessmentTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssessmentTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssessmentTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeAssessmentTargets",
	}
}
