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

// This operation updates an asset.
func (c *Client) UpdateAsset(ctx context.Context, params *UpdateAssetInput, optFns ...func(*Options)) (*UpdateAssetOutput, error) {
	stack := middleware.NewStack("UpdateAsset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateAssetMiddlewares(stack)
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
	addOpUpdateAssetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAsset(options.Region), middleware.Before)
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
			OperationName: "UpdateAsset",
			Err:           err,
		}
	}
	out := result.(*UpdateAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for UpdateAsset.
type UpdateAssetInput struct {
	// The unique identifier for an asset.
	AssetId *string
	// The name of the asset. When importing from Amazon S3, the S3 object key is used
	// as the asset name. When exporting to Amazon S3, the asset name is used as
	// default target S3 object key.
	Name *string
	// The unique identifier for a revision.
	RevisionId *string
	// The unique identifier for a data set.
	DataSetId *string
}

type UpdateAssetOutput struct {
	// The date and time that the asset was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The unique identifier for the revision associated with this asset.
	RevisionId *string
	// The date and time that the asset was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType types.AssetType
	// The ARN for the asset.
	Arn *string
	// Information about the asset, including its size.
	AssetDetails *types.AssetDetails
	// The name of the asset When importing from Amazon S3, the S3 object key is used
	// as the asset name. When exporting to Amazon S3, the asset name is used as
	// default target S3 object key.
	Name *string
	// The unique identifier for the data set associated with this asset.
	DataSetId *string
	// The unique identifier for the asset.
	Id *string
	// The asset ID of the owned asset corresponding to the entitled asset being
	// viewed. This parameter is returned when an asset owner is viewing the entitled
	// copy of its owned asset.
	SourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateAssetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAsset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAsset{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAsset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "UpdateAsset",
	}
}
