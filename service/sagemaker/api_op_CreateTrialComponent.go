// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a trial component, which is a stage of a machine learning trial. A trial
// is composed of one or more trial components. A trial component can be used in
// multiple trials. Trial components include pre-processing jobs, training jobs,
// and batch transform jobs. When you use Amazon SageMaker Studio or the Amazon
// SageMaker Python SDK, all experiments, trials, and trial components are
// automatically tracked, logged, and indexed. When you use the AWS SDK for Python
// (Boto), you must use the logging APIs provided by the SDK. You can add tags to a
// trial component and then use the Search () API to search for the tags.
// CreateTrialComponent can only be invoked from within an Amazon SageMaker managed
// environment. This includes Amazon SageMaker training jobs, processing jobs,
// transform jobs, and Amazon SageMaker notebooks. A call to CreateTrialComponent
// from outside one of these environments results in an error.
func (c *Client) CreateTrialComponent(ctx context.Context, params *CreateTrialComponentInput, optFns ...func(*Options)) (*CreateTrialComponentOutput, error) {
	stack := middleware.NewStack("CreateTrialComponent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTrialComponentMiddlewares(stack)
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
	addOpCreateTrialComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrialComponent(options.Region), middleware.Before)
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
			OperationName: "CreateTrialComponent",
			Err:           err,
		}
	}
	out := result.(*CreateTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrialComponentInput struct {

	// The name of the component. The name must be unique in your AWS account and is
	// not case-sensitive.
	//
	// This member is required.
	TrialComponentName *string

	// The name of the component as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialComponentName is displayed.
	DisplayName *string

	// When the component ended.
	EndTime *time.Time

	// The input artifacts for the component. Examples of input artifacts are datasets,
	// algorithms, hyperparameters, source code, and instance types.
	InputArtifacts map[string]*types.TrialComponentArtifact

	// The output artifacts for the component. Examples of output artifacts are
	// metrics, snapshots, logs, and images.
	OutputArtifacts map[string]*types.TrialComponentArtifact

	// The hyperparameters for the component.
	Parameters map[string]*types.TrialComponentParameterValue

	// When the component started.
	StartTime *time.Time

	// The status of the component. States include:
	//
	//     * InProgress
	//
	//     *
	// Completed
	//
	//     * Failed
	Status *types.TrialComponentStatus

	// A list of tags to associate with the component. You can use Search () API to
	// search on the tags.
	Tags []*types.Tag
}

type CreateTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTrialComponentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrialComponent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrialComponent{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrialComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTrialComponent",
	}
}
