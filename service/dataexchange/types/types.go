// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The destination for the asset.
type AssetDestinationEntry struct {
	// The S3 bucket that is the destination for the asset.
	Bucket *string
	// The unique identifier for the asset.
	AssetId *string
	// The name of the object in Amazon S3 for the asset.
	Key *string
}

type AssetDetails struct {
	// The S3 object that is the asset.
	S3SnapshotAsset *S3SnapshotAsset
}

// An asset in AWS Data Exchange is a piece of data that can be stored as an S3
// object. The asset can be a structured data file, an image file, or some other
// data file. When you create an import job for your files, you create an asset in
// AWS Data Exchange for each of those files.
type AssetEntry struct {
	// The asset ID of the owned asset corresponding to the entitled asset being
	// viewed. This parameter is returned when an asset owner is viewing the entitled
	// copy of its owned asset.
	SourceId *string
	// The name of the asset. When importing from Amazon S3, the S3 object key is used
	// as the asset name. When exporting to Amazon S3, the asset name is used as
	// default target S3 object key.
	Name *string
	// Information about the asset, including its size.
	AssetDetails *AssetDetails
	// The date and time that the asset was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType AssetType
	// The unique identifier for the revision associated with this asset.
	RevisionId *string
	// The date and time that the asset was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The unique identifier for the asset.
	Id *string
	// The ARN for the asset.
	Arn *string
	// The unique identifier for the data set associated with this asset.
	DataSetId *string
}

// The source of the assets.
type AssetSourceEntry struct {
	// The S3 bucket that's part of the source of the asset.
	Bucket *string
	// The name of the object in Amazon S3 for the asset.
	Key *string
}

// A data set is an AWS resource with one or more revisions.
type DataSetEntry struct {
	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin Origin
	// The date and time that the data set was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The ARN for the data set.
	Arn *string
	// The description for the data set.
	Description *string
	// The unique identifier for the data set.
	Id *string
	// The name of the data set.
	Name *string
	// The data set ID of the owned data set corresponding to the entitled data set
	// being viewed. This parameter is returned when a data set owner is viewing the
	// entitled copy of its owned data set.
	SourceId *string
	// The date and time that the data set was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
	// If the origin of this data set is ENTITLED, includes the details for the product
	// on AWS Marketplace.
	OriginDetails *OriginDetails
	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType AssetType
}

type Details struct {
	// The list of sources for the assets.
	ImportAssetsFromS3JobErrorDetails       []*AssetSourceEntry
	ImportAssetFromSignedUrlJobErrorDetails *ImportAssetFromSignedUrlJobErrorDetails
}

// Details of the operation to be performed by the job.
type ExportAssetsToS3RequestDetails struct {
	// The destination for the asset.
	AssetDestinations []*AssetDestinationEntry
	// The unique identifier for the data set associated with this export job.
	DataSetId *string
	// Encryption configuration for the export job.
	Encryption *ExportServerSideEncryption
	// The unique identifier for the revision associated with this export request.
	RevisionId *string
}

// Details about the export to Amazon S3 response.
type ExportAssetsToS3ResponseDetails struct {
	// The destination in Amazon S3 where the asset is exported.
	AssetDestinations []*AssetDestinationEntry
	// The unique identifier for the revision associated with this export response.
	RevisionId *string
	// The unique identifier for the data set associated with this export job.
	DataSetId *string
	// Encryption configuration of the export job.
	Encryption *ExportServerSideEncryption
}

// Details of the operation to be performed by the job.
type ExportAssetToSignedUrlRequestDetails struct {
	// The unique identifier for the revision associated with this export request.
	RevisionId *string
	// The unique identifier for the asset that is exported to a signed URL.
	AssetId *string
	// The unique identifier for the data set associated with this export job.
	DataSetId *string
}

// The details of the export to signed URL response.
type ExportAssetToSignedUrlResponseDetails struct {
	// The unique identifier for the asset associated with this export job.
	AssetId *string
	// The unique identifier for the data set associated with this export job.
	DataSetId *string
	// The date and time that the signed URL expires, in ISO 8601 format.
	SignedUrlExpiresAt *time.Time
	// The unique identifier for the revision associated with this export response.
	RevisionId *string
	// The signed URL for the export request.
	SignedUrl *string
}

// Encryption configuration of the export job. Includes the encryption type as well
// as the AWS KMS key. The KMS key is only necessary if you chose the KMS
// encryption type.
type ExportServerSideEncryption struct {
	// The Amazon Resource Name (ARN) of the the AWS KMS key you want to use to encrypt
	// the Amazon S3 objects. This parameter is required if you choose aws:kms as an
	// encryption type.
	KmsKeyArn *string
	// The type of server side encryption used for encrypting the objects in Amazon S3.
	Type ServerSideEncryptionTypes
}

type ImportAssetFromSignedUrlJobErrorDetails struct {
	// The name of the asset. When importing from Amazon S3, the S3 object key is used
	// as the asset name. When exporting to Amazon S3, the asset name is used as
	// default target S3 object key.
	AssetName *string
}

// Details of the operation to be performed by the job.
type ImportAssetFromSignedUrlRequestDetails struct {
	// The unique identifier for the data set associated with this import job.
	DataSetId *string
	// The Base64-encoded Md5 hash for the asset, used to ensure the integrity of the
	// file at that location.
	Md5Hash *string
	// The unique identifier for the revision associated with this import request.
	RevisionId *string
	// The name of the asset. When importing from Amazon S3, the S3 object key is used
	// as the asset name.
	AssetName *string
}

// The details in the response for an import request, including the signed URL and
// other information.
type ImportAssetFromSignedUrlResponseDetails struct {
	// The signed URL.
	SignedUrl *string
	// The unique identifier for the data set associated with this import job.
	DataSetId *string
	// The time and date at which the signed URL expires, in ISO 8601 format.
	SignedUrlExpiresAt *time.Time
	// The Base64-encoded Md5 hash for the asset, used to ensure the integrity of the
	// file at that location.
	Md5Hash *string
	// The name for the asset associated with this import response.
	AssetName *string
	// The unique identifier for the revision associated with this import response.
	RevisionId *string
}

// Details of the operation to be performed by the job.
type ImportAssetsFromS3RequestDetails struct {
	// Is a list of S3 bucket and object key pairs.
	AssetSources []*AssetSourceEntry
	// The unique identifier for the data set associated with this import job.
	DataSetId *string
	// The unique identifier for the revision associated with this import request.
	RevisionId *string
}

// Details from an import from Amazon S3 response.
type ImportAssetsFromS3ResponseDetails struct {
	// The unique identifier for the data set associated with this import job.
	DataSetId *string
	// Is a list of Amazon S3 bucket and object key pairs.
	AssetSources []*AssetSourceEntry
	// The unique identifier for the revision associated with this import response.
	RevisionId *string
}

// AWS Data Exchange Jobs are asynchronous import or export operations used to
// create or copy assets. A data set owner can both import and export as they see
// fit. Someone with an entitlement to a data set can only export. Jobs are deleted
// 90 days after they are created.
type JobEntry struct {
	// The ARN for the job.
	Arn *string
	// The unique identifier for the job.
	Id *string
	// The state of the job.
	State State
	// Errors for jobs.
	Errors []*JobError
	// The date and time that the job was created, in ISO 8601 format.
	CreatedAt *time.Time
	// Details of the operation to be performed by the job, such as export destination
	// details or import source details.
	Details *ResponseDetails
	// The job type.
	Type Type
	// The date and time that the job was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
}

// An error that occurred with the job request.
type JobError struct {
	// The type of resource related to the error.
	ResourceType JobErrorResourceTypes
	Details      *Details
	// The value of the exceeded limit.
	LimitValue *float64
	// The code for the job error.
	Code Code
	// The name of the limit that was reached.
	LimitName JobErrorLimitName
	// The unique identifier for the resource related to the error.
	ResourceId *string
	// The message related to the job error.
	Message *string
}

type OriginDetails struct {
	ProductId *string
}

// The details for the request.
type RequestDetails struct {
	// Details about the export to signed URL request.
	ExportAssetToSignedUrl *ExportAssetToSignedUrlRequestDetails
	// Details about the import from Amazon S3 request.
	ImportAssetsFromS3 *ImportAssetsFromS3RequestDetails
	// Details about the import from signed URL request.
	ImportAssetFromSignedUrl *ImportAssetFromSignedUrlRequestDetails
	// Details about the export to Amazon S3 request.
	ExportAssetsToS3 *ExportAssetsToS3RequestDetails
}

// Details for the response.
type ResponseDetails struct {
	// Details for the export to Amazon S3 response.
	ExportAssetsToS3 *ExportAssetsToS3ResponseDetails
	// Details for the import from signed URL response.
	ImportAssetFromSignedUrl *ImportAssetFromSignedUrlResponseDetails
	// Details for the import from Amazon S3 response.
	ImportAssetsFromS3 *ImportAssetsFromS3ResponseDetails
	// Details for the export to signed URL response.
	ExportAssetToSignedUrl *ExportAssetToSignedUrlResponseDetails
}

// A revision is a container for one or more assets.
type RevisionEntry struct {
	// The unique identifier for the data set associated with this revision.
	DataSetId *string
	// The revision ID of the owned revision corresponding to the entitled revision
	// being viewed. This parameter is returned when a revision owner is viewing the
	// entitled copy of its owned revision.
	SourceId *string
	// The unique identifier for the revision.
	Id *string
	// The date and time that the revision was created, in ISO 8601 format.
	CreatedAt *time.Time
	// An optional comment about the revision.
	Comment *string
	// The ARN for the revision.
	Arn *string
	// To publish a revision to a data set in a product, the revision must first be
	// finalized. Finalizing a revision tells AWS Data Exchange that your changes to
	// the assets in the revision are complete. After it's in this read-only state, you
	// can publish the revision to your products. Finalized revisions can be published
	// through the AWS Data Exchange console or the AWS Marketplace Catalog API, using
	// the StartChangeSet AWS Marketplace Catalog API action. When using the API,
	// revisions are uniquely identified by their ARN.
	Finalized *bool
	// The date and time that the revision was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
}

// The S3 object that is the asset.
type S3SnapshotAsset struct {
	// The size of the S3 object that is the object.
	Size *float64
}
