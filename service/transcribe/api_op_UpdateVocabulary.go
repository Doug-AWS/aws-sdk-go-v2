// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates an existing vocabulary with new values. The UpdateVocabulary operation
// overwrites all of the existing information with the values that you provide in
// the request.
func (c *Client) UpdateVocabulary(ctx context.Context, params *UpdateVocabularyInput, optFns ...func(*Options)) (*UpdateVocabularyOutput, error) {
	stack := middleware.NewStack("UpdateVocabulary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateVocabularyMiddlewares(stack)
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
	addOpUpdateVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVocabulary(options.Region), middleware.Before)
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
			OperationName: "UpdateVocabulary",
			Err:           err,
		}
	}
	out := result.(*UpdateVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVocabularyInput struct {

	// The language code of the vocabulary entries.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The name of the vocabulary to update. The name is case-sensitive. If you try to
	// update a vocabulary with the same name as a previous vocabulary you will receive
	// a ConflictException error.
	//
	// This member is required.
	VocabularyName *string

	// An array of strings containing the vocabulary entries.
	Phrases []*string

	// The S3 location of the text file that contains the definition of the custom
	// vocabulary. The URI must be in the same region as the API endpoint that you are
	// calling. The general form is  <p>For example:</p> <p>For more information about
	// S3 object names, see <a
	// href="http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys">Object
	// Keys</a> in the <i>Amazon S3 Developer Guide</i>.</p> <p>For more information
	// about custom vocabularies, see <a
	// href="http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary">Custom
	// Vocabularies</a>.</p>
	VocabularyFileUri *string
}

type UpdateVocabularyOutput struct {

	// The language code of the vocabulary entries.
	LanguageCode types.LanguageCode

	// The date and time that the vocabulary was updated.
	LastModifiedTime *time.Time

	// The name of the vocabulary that was updated.
	VocabularyName *string

	// The processing state of the vocabulary. When the VocabularyState field contains
	// READY the vocabulary is ready to be used in a StartTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateVocabularyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVocabulary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVocabulary{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateVocabulary",
	}
}
