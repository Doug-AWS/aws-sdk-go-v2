// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the AWS Fargate profiles associated with the specified cluster in your AWS
// account in the specified Region.
func (c *Client) ListFargateProfiles(ctx context.Context, params *ListFargateProfilesInput, optFns ...func(*Options)) (*ListFargateProfilesOutput, error) {
	if params == nil {
		params = &ListFargateProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFargateProfiles", params, optFns, addOperationListFargateProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFargateProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFargateProfilesInput struct {

	// The name of the Amazon EKS cluster that you would like to listFargate profiles
	// in.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of Fargate profile results returned by ListFargateProfiles in
	// paginated output. When you use this parameter, ListFargateProfiles returns only
	// maxResults results in a single page along with a nextToken response element. You
	// can see the remaining results of the initial request by sending another
	// ListFargateProfiles request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListFargateProfiles returns
	// up to 100 results and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListFargateProfiles
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

type ListFargateProfilesOutput struct {

	// A list of all of the Fargate profiles associated with the specified cluster.
	FargateProfileNames []*string

	// The nextToken value to include in a future ListFargateProfiles request. When the
	// results of a ListFargateProfiles request exceed maxResults, you can use this
	// value to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFargateProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFargateProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFargateProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListFargateProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFargateProfiles(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListFargateProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListFargateProfiles",
	}
}
