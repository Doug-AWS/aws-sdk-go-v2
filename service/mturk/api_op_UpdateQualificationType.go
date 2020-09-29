// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The UpdateQualificationType operation modifies the attributes of an existing
// Qualification type, which is represented by a QualificationType data structure.
// Only the owner of a Qualification type can modify its attributes. Most
// attributes of a Qualification type can be changed after the type has been
// created. However, the Name and Keywords fields cannot be modified. The
// RetryDelayInSeconds parameter can be modified or added to change the delay or to
// enable retries, but RetryDelayInSeconds cannot be used to disable retries. You
// can use this operation to update the test for a Qualification type. The test is
// updated based on the values specified for the Test, TestDurationInSeconds and
// AnswerKey parameters. All three parameters specify the updated test. If you are
// updating the test for a type, you must specify the Test and
// TestDurationInSeconds parameters. The AnswerKey parameter is optional; omitting
// it specifies that the updated test does not have an answer key. If you omit the
// Test parameter, the test for the Qualification type is unchanged. There is no
// way to remove a test from a Qualification type that has one. If the type already
// has a test, you cannot update it to be AutoGranted. If the Qualification type
// does not have a test and one is provided by an update, the type will henceforth
// have a test. If you want to update the test duration or answer key for an
// existing test without changing the questions, you must specify a Test parameter
// with the original questions, along with the updated values. If you provide an
// updated Test but no AnswerKey, the new test will not have an answer key.
// Requests for such Qualifications must be granted manually. You can also update
// the AutoGranted and AutoGrantedValue attributes of the Qualification type.
func (c *Client) UpdateQualificationType(ctx context.Context, params *UpdateQualificationTypeInput, optFns ...func(*Options)) (*UpdateQualificationTypeOutput, error) {
	stack := middleware.NewStack("UpdateQualificationType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateQualificationTypeMiddlewares(stack)
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
	addOpUpdateQualificationTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQualificationType(options.Region), middleware.Before)
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
			OperationName: "UpdateQualificationType",
			Err:           err,
		}
	}
	out := result.(*UpdateQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQualificationTypeInput struct {
	// The ID of the Qualification type to update.
	QualificationTypeId *string
	// The amount of time, in seconds, that Workers must wait after requesting a
	// Qualification of the specified Qualification type before they can retry the
	// Qualification request. It is not possible to disable retries for a Qualification
	// type after it has been created with retries enabled. If you want to disable
	// retries, you must dispose of the existing retry-enabled Qualification type using
	// DisposeQualificationType and then create a new Qualification type with retries
	// disabled using CreateQualificationType.
	RetryDelayInSeconds *int64
	// Specifies whether requests for the Qualification type are granted immediately,
	// without prompting the Worker with a Qualification test. Constraints: If the Test
	// parameter is specified, this parameter cannot be true.
	AutoGranted *bool
	// The answers to the Qualification test specified in the Test parameter, in the
	// form of an AnswerKey data structure.
	AnswerKey *string
	// The number of seconds the Worker has to complete the Qualification test,
	// starting from the time the Worker requests the Qualification.
	TestDurationInSeconds *int64
	// The new status of the Qualification type - Active | Inactive
	QualificationTypeStatus types.QualificationTypeStatus
	// The Qualification value to use for automatically granted Qualifications. This
	// parameter is used only if the AutoGranted parameter is true.
	AutoGrantedValue *int32
	// The questions for the Qualification test a Worker must answer correctly to
	// obtain a Qualification of this type. If this parameter is specified,
	// TestDurationInSeconds must also be specified. Constraints: Must not be longer
	// than 65535 bytes. Must be a QuestionForm data structure. This parameter cannot
	// be specified if AutoGranted is true. Constraints: None. If not specified, the
	// Worker may request the Qualification without answering any questions.
	Test *string
	// The new description of the Qualification type.
	Description *string
}

type UpdateQualificationTypeOutput struct {
	// Contains a QualificationType data structure.
	QualificationType *types.QualificationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateQualificationTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateQualificationType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateQualificationType{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateQualificationType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "UpdateQualificationType",
	}
}
