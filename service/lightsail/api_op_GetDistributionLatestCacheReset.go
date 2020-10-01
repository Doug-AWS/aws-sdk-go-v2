// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the timestamp and status of the last cache reset of a specific Amazon
// Lightsail content delivery network (CDN) distribution.
func (c *Client) GetDistributionLatestCacheReset(ctx context.Context, params *GetDistributionLatestCacheResetInput, optFns ...func(*Options)) (*GetDistributionLatestCacheResetOutput, error) {
	stack := middleware.NewStack("GetDistributionLatestCacheReset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDistributionLatestCacheResetMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributionLatestCacheReset(options.Region), middleware.Before)
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
			OperationName: "GetDistributionLatestCacheReset",
			Err:           err,
		}
	}
	out := result.(*GetDistributionLatestCacheResetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDistributionLatestCacheResetInput struct {

	// The name of the distribution for which to return the timestamp of the last cache
	// reset.  <p>Use the <code>GetDistributions</code> action to get a list of
	// distribution names that you can specify.</p> <p>When omitted, the response
	// includes the latest cache reset timestamp of all your distributions.</p>
	DistributionName *string
}

type GetDistributionLatestCacheResetOutput struct {

	// The timestamp of the last cache reset (e.g., 1479734909.17) in Unix time format.
	CreateTime *time.Time

	// The status of the last cache reset.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDistributionLatestCacheResetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDistributionLatestCacheReset{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDistributionLatestCacheReset{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDistributionLatestCacheReset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetDistributionLatestCacheReset",
	}
}
