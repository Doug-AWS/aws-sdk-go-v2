// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

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

// Retrieves information about the criteria and other settings for a custom data
// identifier.
func (c *Client) GetCustomDataIdentifier(ctx context.Context, params *GetCustomDataIdentifierInput, optFns ...func(*Options)) (*GetCustomDataIdentifierOutput, error) {
	stack := middleware.NewStack("GetCustomDataIdentifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCustomDataIdentifierMiddlewares(stack)
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
	addOpGetCustomDataIdentifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCustomDataIdentifier(options.Region), middleware.Before)
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
			OperationName: "GetCustomDataIdentifier",
			Err:           err,
		}
	}
	out := result.(*GetCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCustomDataIdentifierInput struct {

	// The unique identifier for the Amazon Macie resource or account that the request
	// applies to.
	//
	// This member is required.
	Id *string
}

type GetCustomDataIdentifierOutput struct {

	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn *string

	// The date and time, in UTC and extended ISO 8601 format, when the custom data
	// identifier was created.
	CreatedAt *time.Time

	// Specifies whether the custom data identifier was deleted. If you delete a custom
	// data identifier, Amazon Macie doesn't delete it permanently. Instead, it soft
	// deletes the identifier.
	Deleted *bool

	// The custom description of the custom data identifier.
	Description *string

	// The unique identifier for the custom data identifier.
	Id *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression is the same as any
	// string in this array, Amazon Macie ignores it.
	IgnoreWords []*string

	// An array that lists specific character sequences (keywords), one of which must
	// be within proximity (maximumMatchDistance) of the regular expression to match.
	Keywords []*string

	// The maximum number of characters that can exist between text that matches the
	// regex pattern and the character sequences specified by the keywords array. Macie
	// includes or excludes a result based on the proximity of a keyword to text that
	// matches the regex pattern.
	MaximumMatchDistance *int32

	// The custom name of the custom data identifier.
	Name *string

	// The regular expression (regex) that defines the pattern to match.
	Regex *string

	// A map of key-value pairs that identifies the tags (keys and values) that are
	// associated with the custom data identifier.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCustomDataIdentifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCustomDataIdentifier{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCustomDataIdentifier{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCustomDataIdentifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetCustomDataIdentifier",
	}
}
