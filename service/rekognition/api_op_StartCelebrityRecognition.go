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

// Starts asynchronous recognition of celebrities in a stored video. Amazon
// Rekognition Video can detect celebrities in a video must be stored in an Amazon
// S3 bucket. Use Video () to specify the bucket name and the filename of the
// video. StartCelebrityRecognition returns a job identifier (JobId) which you use
// to get the results of the analysis. When celebrity recognition analysis is
// finished, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic that you specify in NotificationChannel. To
// get the results of the celebrity recognition analysis, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetCelebrityRecognition () and pass the job identifier (JobId) from the initial
// call to StartCelebrityRecognition.  <p>For more information, see Recognizing
// Celebrities in the Amazon Rekognition Developer Guide.</p>
func (c *Client) StartCelebrityRecognition(ctx context.Context, params *StartCelebrityRecognitionInput, optFns ...func(*Options)) (*StartCelebrityRecognitionOutput, error) {
	stack := middleware.NewStack("StartCelebrityRecognition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartCelebrityRecognitionMiddlewares(stack)
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
	addOpStartCelebrityRecognitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartCelebrityRecognition(options.Region), middleware.Before)
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
			OperationName: "StartCelebrityRecognition",
			Err:           err,
		}
	}
	out := result.(*StartCelebrityRecognitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCelebrityRecognitionInput struct {

	// The video in which you want to recognize celebrities. The video must be stored
	// in an Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartCelebrityRecognition requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish the
	// completion status of the celebrity recognition analysis to.
	NotificationChannel *types.NotificationChannel
}

type StartCelebrityRecognitionOutput struct {

	// The identifier for the celebrity recognition analysis job. Use JobId to identify
	// the job in a subsequent call to GetCelebrityRecognition.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartCelebrityRecognitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartCelebrityRecognition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartCelebrityRecognition{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartCelebrityRecognition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartCelebrityRecognition",
	}
}
