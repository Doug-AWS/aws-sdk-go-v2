// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the user import.
func (c *Client) StartUserImportJob(ctx context.Context, params *StartUserImportJobInput, optFns ...func(*Options)) (*StartUserImportJobOutput, error) {
	stack := middleware.NewStack("StartUserImportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartUserImportJobMiddlewares(stack)
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
	addOpStartUserImportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartUserImportJob(options.Region), middleware.Before)
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
			OperationName: "StartUserImportJob",
			Err:           err,
		}
	}
	out := result.(*StartUserImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to start the user import job.
type StartUserImportJobInput struct {
	// The job ID for the user import job.
	JobId *string
	// The user pool ID for the user pool that the users are being imported into.
	UserPoolId *string
}

// Represents the response from the server to the request to start the user import
// job.
type StartUserImportJobOutput struct {
	// The job object that represents the user import job.
	UserImportJob *types.UserImportJobType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartUserImportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartUserImportJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartUserImportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartUserImportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "StartUserImportJob",
	}
}
