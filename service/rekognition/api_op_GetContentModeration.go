// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the unsafe content analysis results for a Amazon Rekognition Video analysis
// started by StartContentModeration ().  <p>Unsafe content analysis of a video is
// an asynchronous operation. You start analysis by calling
// <a>StartContentModeration</a> which returns a job identifier
// (<code>JobId</code>). When analysis finishes, Amazon Rekognition Video publishes
// a completion status to the Amazon Simple Notification Service topic registered
// in the initial call to <code>StartContentModeration</code>. To get the results
// of the unsafe content analysis, first check that the status value published to
// the Amazon SNS topic is <code>SUCCEEDED</code>. If so, call
// <code>GetContentModeration</code> and pass the job identifier
// (<code>JobId</code>) from the initial call to
// <code>StartContentModeration</code>. </p> <p>For more information, see Working
// with Stored Videos in the Amazon Rekognition Devlopers Guide.</p> <p>
// <code>GetContentModeration</code> returns detected unsafe content labels, and
// the time they are detected, in an array, <code>ModerationLabels</code>, of
// <a>ContentModerationDetection</a> objects. </p> <p>By default, the moderated
// labels are returned sorted by time, in milliseconds from the start of the video.
// You can also sort them by moderated label by specifying <code>NAME</code> for
// the <code>SortBy</code> input parameter. </p> <p>Since video analysis can return
// a large number of results, use the <code>MaxResults</code> parameter to limit
// the number of labels returned in a single call to
// <code>GetContentModeration</code>. If there are more results than specified in
// <code>MaxResults</code>, the value of <code>NextToken</code> in the operation
// response contains a pagination token for getting the next set of results. To get
// the next page of results, call <code>GetContentModeration</code> and populate
// the <code>NextToken</code> request parameter with the value of
// <code>NextToken</code> returned from the previous call to
// <code>GetContentModeration</code>.</p> <p>For more information, see Detecting
// Unsafe Content in the Amazon Rekognition Developer Guide.</p>
func (c *Client) GetContentModeration(ctx context.Context, params *GetContentModerationInput, optFns ...func(*Options)) (*GetContentModerationOutput, error) {
	stack := middleware.NewStack("GetContentModeration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetContentModerationMiddlewares(stack)
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
	addOpGetContentModerationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetContentModeration(options.Region), middleware.Before)
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
			OperationName: "GetContentModeration",
			Err:           err,
		}
	}
	out := result.(*GetContentModerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContentModerationInput struct {

	// The identifier for the unsafe content job. Use JobId to identify the job in a
	// subsequent call to GetContentModeration.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of unsafe content labels.
	NextToken *string

	// Sort to use for elements in the ModerationLabelDetections array. Use TIMESTAMP
	// to sort array elements by the time labels are detected. Use NAME to
	// alphabetically group elements for a label together. Within each label group, the
	// array element are sorted by detection confidence. The default sort is by
	// TIMESTAMP.
	SortBy types.ContentModerationSortBy
}

type GetContentModerationOutput struct {

	// The current status of the unsafe content analysis job.
	JobStatus types.VideoJobStatus

	// The detected unsafe content labels and the time(s) they were detected.
	ModerationLabels []*types.ContentModerationDetection

	// Version number of the moderation detection model that was used to detect unsafe
	// content.
	ModerationModelVersion *string

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of unsafe content
	// labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from GetContentModeration.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetContentModerationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetContentModeration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContentModeration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetContentModeration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetContentModeration",
	}
}
