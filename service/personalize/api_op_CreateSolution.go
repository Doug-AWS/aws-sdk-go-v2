// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates the configuration for training a model. A trained model is known as a
// solution. After the configuration is created, you train the model (create a
// solution) by calling the CreateSolutionVersion () operation. Every time you call
// CreateSolutionVersion, a new version of the solution is created. After creating
// a solution version, you check its accuracy by calling GetSolutionMetrics ().
// When you are satisfied with the version, you deploy it using CreateCampaign ().
// The campaign provides recommendations to a client through the GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API. To train a model, Amazon Personalize requires training data and a recipe.
// The training data comes from the dataset group that you provide in the request.
// A recipe specifies the training algorithm and a feature transformation. You can
// specify one of the predefined recipes provided by Amazon Personalize.
// Alternatively, you can specify performAutoML and Amazon Personalize will analyze
// your data and select the optimum USER_PERSONALIZATION recipe for you. Status A
// solution can be in one of the following states:
//
//     * CREATE PENDING > CREATE
// IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//     * DELETE PENDING > DELETE
// IN_PROGRESS
//
// To get the status of the solution, call DescribeSolution (). Wait
// until the status shows as ACTIVE before calling CreateSolutionVersion.  <p
// class="title"> <b>Related APIs</b> </p> <ul> <li> <p> <a>ListSolutions</a> </p>
// </li> <li> <p> <a>CreateSolutionVersion</a> </p> </li> <li> <p>
// <a>DescribeSolution</a> </p> </li> <li> <p> <a>DeleteSolution</a> </p> </li>
// </ul> <ul> <li> <p> <a>ListSolutionVersions</a> </p> </li> <li> <p>
// <a>DescribeSolutionVersion</a> </p> </li> </ul>
func (c *Client) CreateSolution(ctx context.Context, params *CreateSolutionInput, optFns ...func(*Options)) (*CreateSolutionOutput, error) {
	stack := middleware.NewStack("CreateSolution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateSolutionMiddlewares(stack)
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
	addOpCreateSolutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolution(options.Region), middleware.Before)
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
			OperationName: "CreateSolution",
			Err:           err,
		}
	}
	out := result.(*CreateSolutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolutionInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that provides the training
	// data.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name for the solution.
	//
	// This member is required.
	Name *string

	// When your have multiple event types (using an EVENT_TYPE schema field), this
	// parameter specifies which event type (for example, 'click' or 'like') is used
	// for training the model.
	EventType *string

	// Whether to perform automated machine learning (AutoML). The default is false.
	// For this case, you must specify recipeArn. When set to true, Amazon Personalize
	// analyzes your training data and selects the optimal USER_PERSONALIZATION recipe
	// and hyperparameters. In this case, you must omit recipeArn. Amazon Personalize
	// determines the optimal recipe by running tests with different values for the
	// hyperparameters. AutoML lengthens the training process as compared to selecting
	// a specific recipe.
	PerformAutoML *bool

	// Whether to perform hyperparameter optimization (HPO) on the specified or
	// selected recipe. The default is false. When performing AutoML, this parameter is
	// always true and you should not set it to false.
	PerformHPO *bool

	// The ARN of the recipe to use for model training. Only specified when
	// performAutoML is false.
	RecipeArn *string

	// The configuration to use with the solution. When performAutoML is set to true,
	// Amazon Personalize only evaluates the autoMLConfig section of the solution
	// configuration.
	SolutionConfig *types.SolutionConfig
}

type CreateSolutionOutput struct {

	// The ARN of the solution.
	SolutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateSolutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSolution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSolution{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSolution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateSolution",
	}
}
