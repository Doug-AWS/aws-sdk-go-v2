// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a forecast for each item in the TARGET_TIME_SERIES dataset that was used
// to train the predictor. This is known as inference. To retrieve the forecast for
// a single item at low latency, use the operation. To export the complete forecast
// into your Amazon Simple Storage Service (Amazon S3) bucket, use the
// CreateForecastExportJob () operation. The range of the forecast is determined by
// the ForecastHorizon value, which you specify in the CreatePredictor () request.
// When you query a forecast, you can request a specific date range within the
// forecast. To get a list of all your forecasts, use the ListForecasts ()
// operation. The forecasts generated by Amazon Forecast are in the same time zone
// as the dataset that was used to create the predictor. For more information, see
// howitworks-forecast (). The Status of the forecast must be ACTIVE before you can
// query or export the forecast. Use the DescribeForecast () operation to get the
// status.
func (c *Client) CreateForecast(ctx context.Context, params *CreateForecastInput, optFns ...func(*Options)) (*CreateForecastOutput, error) {
	stack := middleware.NewStack("CreateForecast", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateForecastMiddlewares(stack)
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
	addOpCreateForecastValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateForecast(options.Region), middleware.Before)
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
			OperationName: "CreateForecast",
			Err:           err,
		}
	}
	out := result.(*CreateForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateForecastInput struct {
	// A name for the forecast.
	ForecastName *string
	// The Amazon Resource Name (ARN) of the predictor to use to generate the forecast.
	PredictorArn *string
	// The quantiles at which probabilistic forecasts are generated. You can currently
	// specify up to 5 quantiles per forecast. Accepted values include 0.01 to 0.99
	// (increments of .01 only) and mean. The mean forecast is different from the
	// median (0.50) when the distribution is not symmetric (for example, Beta and
	// Negative Binomial). The default value is ["0.1", "0.5", "0.9"].
	ForecastTypes []*string
	// The optional metadata that you apply to the forecast to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per resource - 50.
	//
	//     * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	//     * Maximum key length -
	// 128 Unicode characters in UTF-8.
	//
	//     * Maximum value length - 256 Unicode
	// characters in UTF-8.
	//
	//     * If your tagging schema is used across multiple
	// services and resources, remember that other services may have restrictions on
	// allowed characters. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//
	// * Tag keys and values are case sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for keys as it is reserved
	// for AWS use. You cannot edit or delete tag keys with this prefix. Values can
	// have this prefix. If a tag value has aws as its prefix but the key does not,
	// then Forecast considers it to be a user tag and will count against the limit of
	// 50 tags. Tags with only the key prefix of aws do not count against your tags per
	// resource limit.
	Tags []*types.Tag
}

type CreateForecastOutput struct {
	// The Amazon Resource Name (ARN) of the forecast.
	ForecastArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateForecastMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateForecast{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateForecast{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateForecast(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateForecast",
	}
}
