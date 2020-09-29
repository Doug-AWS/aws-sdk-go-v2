// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the contents of a data set as pre-signed URIs.
func (c *Client) GetDatasetContent(ctx context.Context, params *GetDatasetContentInput, optFns ...func(*Options)) (*GetDatasetContentOutput, error) {
	stack := middleware.NewStack("GetDatasetContent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDatasetContentMiddlewares(stack)
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
	addOpGetDatasetContentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDatasetContent(options.Region), middleware.Before)
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
			OperationName: "GetDatasetContent",
			Err:           err,
		}
	}
	out := result.(*GetDatasetContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDatasetContentInput struct {
	// The name of the data set whose contents are retrieved.
	DatasetName *string
	// The version of the data set whose contents are retrieved. You can also use the
	// strings "$LATEST" or "$LATEST_SUCCEEDED" to retrieve the contents of the latest
	// or latest successfully completed data set. If not specified, "$LATEST_SUCCEEDED"
	// is the default.
	VersionId *string
}

type GetDatasetContentOutput struct {
	// A list of "DatasetEntry" objects.
	Entries []*types.DatasetEntry
	// The time when the request was made.
	Timestamp *time.Time
	// The status of the data set content.
	Status *types.DatasetContentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDatasetContentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDatasetContent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDatasetContent{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDatasetContent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "GetDatasetContent",
	}
}
