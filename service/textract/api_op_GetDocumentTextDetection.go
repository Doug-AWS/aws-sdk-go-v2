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

// Gets the results for an Amazon Textract asynchronous operation that detects text
// in a document. Amazon Textract can detect lines of text and the words that make
// up a line of text. You start asynchronous text detection by calling
// StartDocumentTextDetection (), which returns a job identifier (JobId). When the
// text detection operation finishes, Amazon Textract publishes a completion status
// to the Amazon Simple Notification Service (Amazon SNS) topic that's registered
// in the initial call to StartDocumentTextDetection. To get the results of the
// text-detection operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If so, call GetDocumentTextDetection, and pass
// the job identifier (JobId) from the initial call to StartDocumentTextDetection.
// GetDocumentTextDetection returns an array of Block () objects. Each document
// page has as an associated Block of type PAGE. Each PAGE Block object is the
// parent of LINE Block objects that represent the lines of detected text on a
// page. A LINE Block object is a parent for each word that makes up the line.
// Words are represented by Block objects of type WORD.  <p>Use the MaxResults
// parameter to limit the number of blocks that are returned. If there are more
// results than specified in <code>MaxResults</code>, the value of
// <code>NextToken</code> in the operation response contains a pagination token for
// getting the next set of results. To get the next page of results, call
// <code>GetDocumentTextDetection</code>, and populate the <code>NextToken</code>
// request parameter with the token value that's returned from the previous call to
// <code>GetDocumentTextDetection</code>.</p> <p>For more information, see <a
// href="https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html">Document
// Text Detection</a>.</p>
func (c *Client) GetDocumentTextDetection(ctx context.Context, params *GetDocumentTextDetectionInput, optFns ...func(*Options)) (*GetDocumentTextDetectionOutput, error) {
	stack := middleware.NewStack("GetDocumentTextDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDocumentTextDetectionMiddlewares(stack)
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
	addOpGetDocumentTextDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentTextDetection(options.Region), middleware.Before)
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
			OperationName: "GetDocumentTextDetection",
			Err:           err,
		}
	}
	out := result.(*GetDocumentTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentTextDetectionInput struct {

	// A unique identifier for the text detection job. The JobId is returned from
	// StartDocumentTextDetection. A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 1,000. If you specify a value greater than 1,000, a maximum
	// of 1,000 results is returned. The default value is 1,000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more blocks to
	// retrieve), Amazon Textract returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of blocks.
	NextToken *string
}

type GetDocumentTextDetectionOutput struct {

	// The results of the text-detection operation.
	Blocks []*types.Block

	//
	DetectDocumentTextModelVersion *string

	// Information about a document that Amazon Textract processed. DocumentMetadata is
	// returned in every page of paginated responses from an Amazon Textract video
	// operation.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the text detection job.
	JobStatus types.JobStatus

	// If the response is truncated, Amazon Textract returns this token. You can use
	// this token in the subsequent request to retrieve the next set of text-detection
	// results.
	NextToken *string

	// The current status of an asynchronous text-detection operation for the document.
	StatusMessage *string

	// A list of warnings that occurred during the text-detection operation for the
	// document.
	Warnings []*types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDocumentTextDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDocumentTextDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDocumentTextDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDocumentTextDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetDocumentTextDetection",
	}
}
