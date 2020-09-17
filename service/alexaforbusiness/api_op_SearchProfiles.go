// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches room profiles and lists the ones that meet a set of filter criteria.
func (c *Client) SearchProfiles(ctx context.Context, params *SearchProfilesInput, optFns ...func(*Options)) (*SearchProfilesOutput, error) {
	stack := middleware.NewStack("SearchProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchProfilesMiddlewares(stack)
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
	addOpSearchProfilesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProfiles(options.Region), middleware.Before)

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
			OperationName: "SearchProfiles",
			Err:           err,
		}
	}
	out := result.(*SearchProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProfilesInput struct {
	// The filters to use to list a specified set of room profiles. Supported filter
	// keys are ProfileName and Address. Required.
	Filters []*types.Filter
	// The sort order to use in listing the specified set of room profiles. Supported
	// sort keys are ProfileName and Address.
	SortCriteria []*types.Sort
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string
}

type SearchProfilesOutput struct {
	// The token returned to indicate that there is more data available.
	NextToken *string
	// The total number of room profiles returned.
	TotalCount *int32
	// The profiles that meet the specified set of filter criteria, in sort order.
	Profiles []*types.ProfileData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchProfiles",
	}
}