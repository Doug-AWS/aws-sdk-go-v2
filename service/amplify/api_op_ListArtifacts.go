// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of artifacts for a specified app, branch, and job.
func (c *Client) ListArtifacts(ctx context.Context, params *ListArtifactsInput, optFns ...func(*Options)) (*ListArtifactsOutput, error) {
	stack := middleware.NewStack("ListArtifacts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListArtifactsMiddlewares(stack)
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
	addOpListArtifactsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListArtifacts(options.Region), middleware.Before)
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
			OperationName: "ListArtifacts",
			Err:           err,
		}
	}
	out := result.(*ListArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes the request structure for the list artifacts request.
type ListArtifactsInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of a branch that is part of an Amplify app.
	//
	// This member is required.
	BranchName *string

	// The unique ID for a job.
	//
	// This member is required.
	JobId *string

	// The maximum number of records to list in a single response.
	MaxResults *int32

	// A pagination token. Set to null to start listing artifacts from start. If a
	// non-null pagination token is returned in a result, pass its value in here to
	// list more artifacts.
	NextToken *string
}

// The result structure for the list artifacts request.
type ListArtifactsOutput struct {

	// A list of artifacts.
	//
	// This member is required.
	Artifacts []*types.Artifact

	// A pagination token. If a non-null pagination token is returned in a result, pass
	// its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListArtifactsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListArtifacts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListArtifacts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListArtifacts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListArtifacts",
	}
}
