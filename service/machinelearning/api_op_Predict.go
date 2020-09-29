// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a prediction for the observation using the specified ML Model. Note:
// Not all response parameters will be populated. Whether a response parameter is
// populated depends on the type of model requested.
func (c *Client) Predict(ctx context.Context, params *PredictInput, optFns ...func(*Options)) (*PredictOutput, error) {
	stack := middleware.NewStack("Predict", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPredictMiddlewares(stack)
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
	addOpPredictValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPredict(options.Region), middleware.Before)
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
			OperationName: "Predict",
			Err:           err,
		}
	}
	out := result.(*PredictOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PredictInput struct {
	// A map of variable name-value pairs that represent an observation.
	Record map[string]*string
	// A unique identifier of the MLModel.
	MLModelId *string
	// The predicted endpoint for the input.
	PredictEndpoint *string
}

type PredictOutput struct {
	// The output from a Predict operation:  <ul> <li> <p> <code>Details</code> -
	// Contains the following attributes: <code>DetailsAttributes.PREDICTIVE_MODEL_TYPE
	// - REGRESSION | BINARY | MULTICLASS</code> <code>DetailsAttributes.ALGORITHM -
	// SGD</code> </p> </li> <li> <p> <code>PredictedLabel</code> - Present for either
	// a <code>BINARY</code> or <code>MULTICLASS</code> <code>MLModel</code> request.
	// </p> </li> <li> <p> <code>PredictedScores</code> - Contains the raw
	// classification score corresponding to each label. </p> </li> <li> <p>
	// <code>PredictedValue</code> - Present for a <code>REGRESSION</code>
	// <code>MLModel</code> request. </p> </li> </ul>
	Prediction *types.Prediction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPredictMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPredict{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPredict{}, middleware.After)
}

func newServiceMetadataMiddleware_opPredict(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "Predict",
	}
}
