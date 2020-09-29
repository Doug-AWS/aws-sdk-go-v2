// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

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

// Creates a signed, short-term URL that can be passed to a Selenium
// RemoteWebDriver constructor.
func (c *Client) CreateTestGridUrl(ctx context.Context, params *CreateTestGridUrlInput, optFns ...func(*Options)) (*CreateTestGridUrlOutput, error) {
	stack := middleware.NewStack("CreateTestGridUrl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTestGridUrlMiddlewares(stack)
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
	addOpCreateTestGridUrlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTestGridUrl(options.Region), middleware.Before)
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
			OperationName: "CreateTestGridUrl",
			Err:           err,
		}
	}
	out := result.(*CreateTestGridUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTestGridUrlInput struct {
	// ARN (from CreateTestGridProject () or ListTestGridProjects ()) to associate with
	// the short-term URL.
	ProjectArn *string
	// Lifetime, in seconds, of the URL.
	ExpiresInSeconds *int32
}

type CreateTestGridUrlOutput struct {
	// The number of seconds the URL from CreateTestGridUrlResult$url () stays active.
	Expires *time.Time
	// A signed URL, expiring in CreateTestGridUrlRequest$expiresInSeconds () seconds,
	// to be passed to a RemoteWebDriver.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTestGridUrlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTestGridUrl{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTestGridUrl{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTestGridUrl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateTestGridUrl",
	}
}
