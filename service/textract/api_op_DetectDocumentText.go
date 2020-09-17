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

// Detects text in the input document. Amazon Textract can detect lines of text and
// the words that make up a line of text. The input document must be an image in
// JPEG or PNG format. DetectDocumentText returns the detected text in an array of
// Block () objects. Each document page has as an associated Block of type PAGE.
// Each PAGE Block object is the parent of LINE Block objects that represent the
// lines of detected text on a page. A LINE Block object is a parent for each word
// that makes up the line. Words are represented by Block objects of type WORD.
// <p> <code>DetectDocumentText</code> is a synchronous operation. To analyze
// documents asynchronously, use <a>StartDocumentTextDetection</a>.</p> <p>For more
// information, see <a
// href="https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html">Document
// Text Detection</a>.</p>
func (c *Client) DetectDocumentText(ctx context.Context, params *DetectDocumentTextInput, optFns ...func(*Options)) (*DetectDocumentTextOutput, error) {
	stack := middleware.NewStack("DetectDocumentText", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectDocumentTextMiddlewares(stack)
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
	addOpDetectDocumentTextValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectDocumentText(options.Region), middleware.Before)

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
			OperationName: "DetectDocumentText",
			Err:           err,
		}
	}
	out := result.(*DetectDocumentTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectDocumentTextInput struct {
	// The input document as base64-encoded bytes or an Amazon S3 object. If you use
	// the AWS CLI to call Amazon Textract operations, you can't pass image bytes. The
	// document must be an image in JPEG or PNG format. If you're using an AWS SDK to
	// call Amazon Textract, you might not need to base64-encode image bytes that are
	// passed using the Bytes field.
	Document *types.Document
}

type DetectDocumentTextOutput struct {
	//
	DetectDocumentTextModelVersion *string
	// Metadata about the document. It contains the number of pages that are detected
	// in the document.
	DocumentMetadata *types.DocumentMetadata
	// An array of Block objects that contain the text that's detected in the document.
	Blocks []*types.Block

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectDocumentTextMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectDocumentText{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectDocumentText{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectDocumentText(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "DetectDocumentText",
	}
}