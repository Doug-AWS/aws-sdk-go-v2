// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the merge options available for merging two specified
// branches. For details about why a merge option is not available, use
// GetMergeConflicts or DescribeMergeConflicts.
func (c *Client) GetMergeOptions(ctx context.Context, params *GetMergeOptionsInput, optFns ...func(*Options)) (*GetMergeOptionsOutput, error) {
	stack := middleware.NewStack("GetMergeOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetMergeOptionsMiddlewares(stack)
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
	addOpGetMergeOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMergeOptions(options.Region), middleware.Before)
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
			OperationName: "GetMergeOptions",
			Err:           err,
		}
	}
	out := result.(*GetMergeOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMergeOptionsInput struct {
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	DestinationCommitSpecifier *string
	// The name of the repository that contains the commits about which you want to get
	// merge options.
	RepositoryName *string
	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	SourceCommitSpecifier *string
	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum
}

type GetMergeOptionsOutput struct {
	// The merge option or strategy used to merge the code.
	MergeOptions []types.MergeOptionTypeEnum
	// The commit ID of the merge base.
	BaseCommitId *string
	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	SourceCommitId *string
	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	DestinationCommitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetMergeOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetMergeOptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMergeOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMergeOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetMergeOptions",
	}
}
