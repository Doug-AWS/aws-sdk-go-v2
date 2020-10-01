// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the asynchronous analysis of an input document for relationships between
// detected items such as key-value pairs, tables, and selection elements.  <p>
// <code>StartDocumentAnalysis</code> can analyze text in documents that are in
// JPEG, PNG, and PDF format. The documents are stored in an Amazon S3 bucket. Use
// <a>DocumentLocation</a> to specify the bucket name and file name of the
// document. </p> <p> <code>StartDocumentAnalysis</code> returns a job identifier
// (<code>JobId</code>) that you use to get the results of the operation. When text
// analysis is finished, Amazon Textract publishes a completion status to the
// Amazon Simple Notification Service (Amazon SNS) topic that you specify in
// <code>NotificationChannel</code>. To get the results of the text analysis
// operation, first check that the status value published to the Amazon SNS topic
// is <code>SUCCEEDED</code>. If so, call <a>GetDocumentAnalysis</a>, and pass the
// job identifier (<code>JobId</code>) from the initial call to
// <code>StartDocumentAnalysis</code>.</p> <p>For more information, see <a
// href="https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html">Document
// Text Analysis</a>.</p>
func (c *Client) StartDocumentAnalysis(ctx context.Context, params *StartDocumentAnalysisInput, optFns ...func(*Options)) (*StartDocumentAnalysisOutput, error) {
	stack := middleware.NewStack("StartDocumentAnalysis", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartDocumentAnalysisMiddlewares(stack)
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
	addOpStartDocumentAnalysisValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDocumentAnalysis(options.Region), middleware.Before)
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
			OperationName: "StartDocumentAnalysis",
			Err:           err,
		}
	}
	out := result.(*StartDocumentAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDocumentAnalysisInput struct {

	// The location of the document to be processed.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables that are detected in the input document. Add FORMS
	// to return detected form data. To perform both types of analysis, add TABLES and
	// FORMS to FeatureTypes. All lines and words detected in the document are included
	// in the response (including text that isn't related to the value of
	// FeatureTypes).
	//
	// This member is required.
	FeatureTypes []types.FeatureType

	// The idempotent token that you use to identify the start request. If you use the
	// same token with multiple StartDocumentAnalysis requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being accidentally
	// started more than once. For more information, see Calling Amazon Textract
	// Asynchronous Operations
	// (https://docs.aws.amazon.com/textract/latest/dg/api-async.html).
	ClientRequestToken *string

	// An identifier that you specify that's included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the completion
	// status of the operation to.
	NotificationChannel *types.NotificationChannel
}

type StartDocumentAnalysisOutput struct {

	// The identifier for the document text detection job. Use JobId to identify the
	// job in a subsequent call to GetDocumentAnalysis. A JobId value is only valid for
	// 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartDocumentAnalysisMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartDocumentAnalysis{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDocumentAnalysis{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDocumentAnalysis(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "StartDocumentAnalysis",
	}
}
