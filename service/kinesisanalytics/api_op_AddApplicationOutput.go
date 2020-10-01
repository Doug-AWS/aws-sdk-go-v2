// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Adds an external destination to your Amazon
// Kinesis Analytics application. If you want Amazon Kinesis Analytics to deliver
// data from an in-application stream within your application to an external
// destination (such as an Amazon Kinesis stream, an Amazon Kinesis Firehose
// delivery stream, or an AWS Lambda function), you add the relevant configuration
// to your application using this operation. You can configure one or more outputs
// for your application. Each output configuration maps an in-application stream
// and an external destination. You can use one of the output configurations to
// deliver data from your in-application error stream to an external destination so
// that you can analyze the errors. For more information, see Understanding
// Application Output (Destination)
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the
// DescribeApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
// operation to find the current application version. For the limits on the number
// of application inputs and outputs you can configure, see Limits
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html). This
// operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
func (c *Client) AddApplicationOutput(ctx context.Context, params *AddApplicationOutputInput, optFns ...func(*Options)) (*AddApplicationOutputOutput, error) {
	stack := middleware.NewStack("AddApplicationOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddApplicationOutputMiddlewares(stack)
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
	addOpAddApplicationOutputValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationOutput(options.Region), middleware.Before)
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
			OperationName: "AddApplicationOutput",
			Err:           err,
		}
	}
	out := result.(*AddApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type AddApplicationOutputInput struct {

	// Name of the application to which you want to add the output configuration.
	//
	// This member is required.
	ApplicationName *string

	// Version of the application to which you want to add the output configuration.
	// You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified is
	// not the current version, the ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream,
	// or an AWS Lambda function), and record the formation to use when writing to the
	// destination.
	//
	// This member is required.
	Output *types.Output
}

//
type AddApplicationOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddApplicationOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationOutput{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddApplicationOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationOutput",
	}
}
