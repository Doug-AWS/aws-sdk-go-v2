// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs the SQL query statements contained in the Query. Requires you to have
// access to the workgroup in which the query ran. Running queries against an
// external catalog requires GetDataCatalog () permission to the catalog. For code
// samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
func (c *Client) StartQueryExecution(ctx context.Context, params *StartQueryExecutionInput, optFns ...func(*Options)) (*StartQueryExecutionOutput, error) {
	stack := middleware.NewStack("StartQueryExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartQueryExecutionMiddlewares(stack)
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
	addIdempotencyToken_opStartQueryExecutionMiddleware(stack, options)
	addOpStartQueryExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartQueryExecution(options.Region), middleware.Before)
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
			OperationName: "StartQueryExecution",
			Err:           err,
		}
	}
	out := result.(*StartQueryExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQueryExecutionInput struct {
	// The database within which the query executes.
	QueryExecutionContext *types.QueryExecutionContext
	// Specifies information about where and how to save the results of the query
	// execution. If the query runs in a workgroup, then workgroup's settings may
	// override query settings. This affects the query results location. The workgroup
	// settings override is specified in EnforceWorkGroupConfiguration (true/false) in
	// the WorkGroupConfiguration. See
	// WorkGroupConfiguration$EnforceWorkGroupConfiguration ().
	ResultConfiguration *types.ResultConfiguration
	// The SQL query statements to be executed.
	QueryString *string
	// A unique case-sensitive string used to ensure the request to create the query is
	// idempotent (executes only once). If another StartQueryExecution request is
	// received, the same response is returned and another query is not created. If a
	// parameter has changed, for example, the QueryString, an error is returned. This
	// token is listed as not required because AWS SDKs (for example the AWS SDK for
	// Java) auto-generate the token for users. If you are not using the AWS SDK or the
	// AWS CLI, you must provide this token or the action will fail.
	ClientRequestToken *string
	// The name of the workgroup in which the query is being started.
	WorkGroup *string
}

type StartQueryExecutionOutput struct {
	// The unique ID of the query that ran as a result of this request.
	QueryExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartQueryExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartQueryExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartQueryExecution{}, middleware.After)
}

type idempotencyToken_initializeOpStartQueryExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartQueryExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartQueryExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartQueryExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartQueryExecutionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartQueryExecutionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpStartQueryExecution{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartQueryExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "StartQueryExecution",
	}
}
