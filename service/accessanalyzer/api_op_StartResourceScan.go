// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Immediately starts a scan of the policies applied to the specified resource.
func (c *Client) StartResourceScan(ctx context.Context, params *StartResourceScanInput, optFns ...func(*Options)) (*StartResourceScanOutput, error) {
	stack := middleware.NewStack("StartResourceScan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartResourceScanMiddlewares(stack)
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
	addOpStartResourceScanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartResourceScan(options.Region), middleware.Before)
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
			OperationName: "StartResourceScan",
			Err:           err,
		}
	}
	out := result.(*StartResourceScanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Starts a scan of the policies applied to the specified resource.
type StartResourceScanInput struct {
	// The ARN of the analyzer to use to scan the policies applied to the specified
	// resource.
	AnalyzerArn *string
	// The ARN of the resource to scan.
	ResourceArn *string
}

type StartResourceScanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartResourceScanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartResourceScan{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartResourceScan{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartResourceScan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "StartResourceScan",
	}
}
