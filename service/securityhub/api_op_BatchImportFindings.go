// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports security findings generated from an integrated third-party product into
// Security Hub. This action is requested by the integrated product to import its
// findings into Security Hub. The maximum allowed size for a finding is 240 Kb. An
// error is returned for any finding larger than 240 Kb. After a finding is
// created, BatchImportFindings cannot be used to update the following finding
// fields and objects, which Security Hub customers use to manage their
// investigation workflow.
//
//     * Confidence
//
//     * Criticality
//
//     * Note
//
//     *
// RelatedFindings
//
//     * Severity
//
//     * Types
//
//     * UserDefinedFields
//
//     *
// VerificationState
//
//     * Workflow
func (c *Client) BatchImportFindings(ctx context.Context, params *BatchImportFindingsInput, optFns ...func(*Options)) (*BatchImportFindingsOutput, error) {
	stack := middleware.NewStack("BatchImportFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchImportFindingsMiddlewares(stack)
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
	addOpBatchImportFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportFindings(options.Region), middleware.Before)
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
			OperationName: "BatchImportFindings",
			Err:           err,
		}
	}
	out := result.(*BatchImportFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportFindingsInput struct {
	// A list of findings to import. To successfully import a finding, it must follow
	// the AWS Security Finding Format
	// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html).
	// Maximum of 100 findings per request.
	Findings []*types.AwsSecurityFinding
}

type BatchImportFindingsOutput struct {
	// The number of findings that failed to import.
	FailedCount *int32
	// The list of findings that failed to import.
	FailedFindings []*types.ImportFindingsError
	// The number of findings that were successfully imported.
	SuccessCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchImportFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportFindings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchImportFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchImportFindings",
	}
}
