// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation creates a data set.
func (c *Client) CreateDataSet(ctx context.Context, params *CreateDataSetInput, optFns ...func(*Options)) (*CreateDataSetOutput, error) {
	stack := middleware.NewStack("CreateDataSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDataSetMiddlewares(stack)
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
	addOpCreateDataSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSet(options.Region), middleware.Before)
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
			OperationName: "CreateDataSet",
			Err:           err,
		}
	}
	out := result.(*CreateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for CreateDataSet.
type CreateDataSetInput struct {
	// A description for the data set. This value can be up to 16,348 characters long.
	Description *string
	// The name of the data set.
	Name *string
	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType types.AssetType
	// A data set tag is an optional label that you can assign to a data set when you
	// create it. Each tag consists of a key and an optional value, both of which you
	// define. When you use tagging, you can also use tag-based access control in IAM
	// policies to control access to these data sets and revisions.
	Tags map[string]*string
}

type CreateDataSetOutput struct {
	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType types.AssetType
	// The date and time that the data set was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The tags for the data set.
	Tags map[string]*string
	// The data set ID of the owned data set corresponding to the entitled data set
	// being viewed. This parameter is returned when a data set owner is viewing the
	// entitled copy of its owned data set.
	SourceId *string
	// If the origin of this data set is ENTITLED, includes the details for the product
	// on AWS Marketplace.
	OriginDetails *types.OriginDetails
	// The name of the data set.
	Name *string
	// The date and time that the data set was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
	// The description for the data set.
	Description *string
	// The ARN for the data set.
	Arn *string
	// The unique identifier for the data set.
	Id *string
	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin types.Origin

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDataSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "CreateDataSet",
	}
}
