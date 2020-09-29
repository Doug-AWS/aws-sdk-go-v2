// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets operations for a specific resource (e.g., an instance or a static IP).
func (c *Client) GetOperationsForResource(ctx context.Context, params *GetOperationsForResourceInput, optFns ...func(*Options)) (*GetOperationsForResourceOutput, error) {
	stack := middleware.NewStack("GetOperationsForResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetOperationsForResourceMiddlewares(stack)
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
	addOpGetOperationsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOperationsForResource(options.Region), middleware.Before)
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
			OperationName: "GetOperationsForResource",
			Err:           err,
		}
	}
	out := result.(*GetOperationsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOperationsForResourceInput struct {
	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetOperationsForResource request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string
	// The name of the resource for which you are requesting information.
	ResourceName *string
}

type GetOperationsForResourceOutput struct {
	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetOperationsForResource request and specify
	// the next page token using the pageToken parameter.
	NextPageToken *string
	// (Deprecated) Returns the number of pages of results that remain. In releases
	// prior to June 12, 2017, this parameter returned null by the API. It is now
	// deprecated, and the API returns the next page token parameter instead.
	NextPageCount *string
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetOperationsForResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetOperationsForResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOperationsForResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOperationsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetOperationsForResource",
	}
}
