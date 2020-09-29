// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// The classification type that Amazon Macie Classic applies to the associated S3
// resources.
type ClassificationType struct {
	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	OneTime S3OneTimeClassificationType
	// A continuous classification of the objects that are added to a specified S3
	// bucket. Amazon Macie Classic begins performing continuous classification after a
	// bucket is successfully associated with Amazon Macie Classic.
	Continuous S3ContinuousClassificationType
}

// The classification type that Amazon Macie Classic applies to the associated S3
// resources. At least one of the classification types (oneTime or continuous) must
// be specified.
type ClassificationTypeUpdate struct {
	// A continuous classification of the objects that are added to a specified S3
	// bucket. Amazon Macie Classic begins performing continuous classification after a
	// bucket is successfully associated with Amazon Macie Classic.
	Continuous S3ContinuousClassificationType
	// A one-time classification of all of the existing objects in a specified S3
	// bucket.
	OneTime S3OneTimeClassificationType
}

// Includes details about the failed S3 resources.
type FailedS3Resource struct {
	// The status code of a failed item.
	ErrorCode *string
	// The failed S3 resources.
	FailedItem *S3Resource
	// The error message of a failed item.
	ErrorMessage *string
}

// Contains information about the Amazon Macie Classic member account.
type MemberAccount struct {
	// The AWS account ID of the Amazon Macie Classic member account.
	AccountId *string
}

// Contains information about the S3 resource. This data type is used as a request
// parameter in the DisassociateS3Resources action and can be used as a response
// parameter in the AssociateS3Resources and UpdateS3Resources actions.
type S3Resource struct {
	// The prefix of the S3 bucket.
	Prefix *string
	// The name of the S3 bucket.
	BucketName *string
}

// The S3 resources that you want to associate with Amazon Macie Classic for
// monitoring and data classification. This data type is used as a request
// parameter in the AssociateS3Resources action and a response parameter in the
// ListS3Resources action.
type S3ResourceClassification struct {
	// The prefix of the S3 bucket that you want to associate with Amazon Macie
	// Classic.
	Prefix *string
	// The classification type that you want to specify for the resource associated
	// with Amazon Macie Classic.
	ClassificationType *ClassificationType
	// The name of the S3 bucket that you want to associate with Amazon Macie Classic.
	BucketName *string
}

// The S3 resources whose classification types you want to update. This data type
// is used as a request parameter in the UpdateS3Resources action.
type S3ResourceClassificationUpdate struct {
	// The classification type that you want to update for the resource associated with
	// Amazon Macie Classic.
	ClassificationTypeUpdate *ClassificationTypeUpdate
	// The name of the S3 bucket whose classification types you want to update.
	BucketName *string
	// The prefix of the S3 bucket whose classification types you want to update.
	Prefix *string
}
