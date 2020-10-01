// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation deletes the specified tags for a domain. All tag operations are
// eventually consistent; subsequent operations might not immediately represent all
// issued operations.
func (c *Client) DeleteTagsForDomain(ctx context.Context, params *DeleteTagsForDomainInput, optFns ...func(*Options)) (*DeleteTagsForDomainOutput, error) {
	stack := middleware.NewStack("DeleteTagsForDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteTagsForDomainMiddlewares(stack)
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
	addOpDeleteTagsForDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTagsForDomain(options.Region), middleware.Before)
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
			OperationName: "DeleteTagsForDomain",
			Err:           err,
		}
	}
	out := result.(*DeleteTagsForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DeleteTagsForDomainRequest includes the following elements.
type DeleteTagsForDomainInput struct {

	// The domain for which you want to delete one or more tags.
	//
	// This member is required.
	DomainName *string

	// A list of tag keys to delete.
	//
	// This member is required.
	TagsToDelete []*string
}

type DeleteTagsForDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteTagsForDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTagsForDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTagsForDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTagsForDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "DeleteTagsForDomain",
	}
}
