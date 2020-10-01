// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new document classifier that you can use to categorize documents. To
// create a classifier, you provide a set of training documents that labeled with
// the categories that you want to use. After the classifier is trained you can use
// it to categorize a set of labeled documents into the categories. For more
// information, see how-document-classification ().
func (c *Client) CreateDocumentClassifier(ctx context.Context, params *CreateDocumentClassifierInput, optFns ...func(*Options)) (*CreateDocumentClassifierOutput, error) {
	stack := middleware.NewStack("CreateDocumentClassifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDocumentClassifierMiddlewares(stack)
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
	addIdempotencyToken_opCreateDocumentClassifierMiddleware(stack, options)
	addOpCreateDocumentClassifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDocumentClassifier(options.Region), middleware.Before)
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
			OperationName: "CreateDocumentClassifier",
			Err:           err,
		}
	}
	out := result.(*CreateDocumentClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDocumentClassifierInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Management (IAM) role
	// that grants Amazon Comprehend read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The name of the document classifier.
	//
	// This member is required.
	DocumentClassifierName *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.DocumentClassifierInputDataConfig

	// The language of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"), Spanish
	// ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must
	// be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// Indicates the mode in which the classifier will be trained. The classifier can
	// be trained in multi-class mode, which identifies one and only one class for each
	// document, or multi-label mode, which identifies one or more labels for each
	// document. In multi-label mode, multiple labels for an individual document are
	// separated by a delimiter. The default delimiter between labels is a pipe (|).
	Mode types.DocumentClassifierMode

	// Enables the addition of output results configuration parameters for custom
	// classifier jobs.
	OutputDataConfig *types.DocumentClassifierOutputDataConfig

	// Tags to be associated with the document classifier being created. A tag is a
	// key-value pair that adds as a metadata to a resource used by Amazon Comprehend.
	// For example, a tag with "Sales" as the key might be added to a resource to
	// indicate its use by the sales department.
	Tags []*types.Tag

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt data on the storage volume attached to the ML compute instance(s) that
	// process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	//     * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//     * Amazon
	// Resource Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your custom classifier. For more
	// information, see Amazon VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *types.VpcConfig
}

type CreateDocumentClassifierOutput struct {

	// The Amazon Resource Name (ARN) that identifies the document classifier.
	DocumentClassifierArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDocumentClassifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDocumentClassifier{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDocumentClassifier{}, middleware.After)
}

type idempotencyToken_initializeOpCreateDocumentClassifier struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDocumentClassifier) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDocumentClassifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDocumentClassifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDocumentClassifierInput ")
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
func addIdempotencyToken_opCreateDocumentClassifierMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDocumentClassifier{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDocumentClassifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "CreateDocumentClassifier",
	}
}
