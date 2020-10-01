// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A Block represents items that are recognized in a document within a group of
// pixels close to each other. The information returned in a Block object depends
// on the type of operation. In text detection for documents (for example
// DetectDocumentText ()), you get information about the detected words and lines
// of text. In text analysis (for example AnalyzeDocument ()), you can also get
// information about the fields, tables, and selection elements that are detected
// in the document. An array of Block objects is returned by both synchronous and
// asynchronous operations. In synchronous operations, such as DetectDocumentText
// (), the array of Block objects is the entire set of results. In asynchronous
// operations, such as GetDocumentAnalysis (), the array is returned over one or
// more responses. For more information, see How Amazon Textract Works
// (https://docs.aws.amazon.com/textract/latest/dg/how-it-works.html).
type Block struct {

	// The type of text item that's recognized. In operations for text detection, the
	// following types are returned:
	//
	//     * PAGE - Contains a list of the LINE Block
	// objects that are detected on a document page.
	//
	//     * WORD - A word detected on a
	// document page. A word is one or more ISO basic Latin script characters that
	// aren't separated by spaces.
	//
	//     * LINE - A string of tab-delimited, contiguous
	// words that are detected on a document page.
	//
	// In text analysis operations, the
	// following types are returned:
	//
	//     * PAGE - Contains a list of child Block
	// objects that are detected on a document page.
	//
	//     * KEY_VALUE_SET - Stores the
	// KEY and VALUE Block objects for linked text that's detected on a document page.
	// Use the EntityType field to determine if a KEY_VALUE_SET object is a KEY Block
	// object or a VALUE Block object.
	//
	//     * WORD - A word that's detected on a
	// document page. A word is one or more ISO basic Latin script characters that
	// aren't separated by spaces.
	//
	//     * LINE - A string of tab-delimited, contiguous
	// words that are detected on a document page.
	//
	//     * TABLE - A table that's
	// detected on a document page. A table is grid-based information with two or more
	// rows or columns, with a cell span of one row and one column each.
	//
	//     * CELL -
	// A cell within a detected table. The cell is the parent of the block that
	// contains the text in the cell.
	//
	//     * SELECTION_ELEMENT - A selection element
	// such as an option button (radio button) or a check box that's detected on a
	// document page. Use the value of SelectionStatus to determine the status of the
	// selection element.
	BlockType BlockType

	// The column in which a table cell appears. The first column position is 1.
	// ColumnIndex isn't returned by DetectDocumentText and GetDocumentTextDetection.
	ColumnIndex *int32

	// The number of columns that a table cell spans. Currently this value is always 1,
	// even if the number of columns spanned is greater than 1. ColumnSpan isn't
	// returned by DetectDocumentText and GetDocumentTextDetection.
	ColumnSpan *int32

	// The confidence score that Amazon Textract has in the accuracy of the recognized
	// text and the accuracy of the geometry points around the recognized text.
	Confidence *float32

	// The type of entity. The following can be returned:
	//
	//     * KEY - An identifier
	// for a field on the document.
	//
	//     * VALUE - The field text.
	//
	// EntityTypes isn't
	// returned by DetectDocumentText and GetDocumentTextDetection.
	EntityTypes []EntityType

	// The location of the recognized text on the image. It includes an axis-aligned,
	// coarse bounding box that surrounds the text, and a finer-grain polygon for more
	// accurate spatial information.
	Geometry *Geometry

	// The identifier for the recognized text. The identifier is only unique for a
	// single operation.
	Id *string

	// The page on which a block was detected. Page is returned by asynchronous
	// operations. Page values greater than 1 are only returned for multipage documents
	// that are in PDF format. A scanned image (JPEG/PNG), even if it contains multiple
	// document pages, is considered to be a single-page document. The value of Page is
	// always 1. Synchronous operations don't return Page because every input document
	// is considered to be a single-page document.
	Page *int32

	// A list of child blocks of the current block. For example, a LINE object has
	// child blocks for each WORD block that's part of the line of text. There aren't
	// Relationship objects in the list for relationships that don't exist, such as
	// when the current block has no child blocks. The list size can be the
	// following:
	//
	//     * 0 - The block has no child blocks.
	//
	//     * 1 - The block has
	// child blocks.
	Relationships []*Relationship

	// The row in which a table cell is located. The first row position is 1. RowIndex
	// isn't returned by DetectDocumentText and GetDocumentTextDetection.
	RowIndex *int32

	// The number of rows that a table cell spans. Currently this value is always 1,
	// even if the number of rows spanned is greater than 1. RowSpan isn't returned by
	// DetectDocumentText and GetDocumentTextDetection.
	RowSpan *int32

	// The selection status of a selection element, such as an option button or check
	// box.
	SelectionStatus SelectionStatus

	// The word or line of text that's recognized by Amazon Textract.
	Text *string
}

// The bounding box around the detected page, text, key-value pair, table, table
// cell, or selection element on a document page. The left (x-coordinate) and top
// (y-coordinate) are coordinates that represent the top and left sides of the
// bounding box. Note that the upper-left corner of the image is the origin (0,0).
// The top and left values returned are ratios of the overall document page size.
// For example, if the input image is 700 x 200 pixels, and the top-left coordinate
// of the bounding box is 350 x 50 pixels, the API returns a left value of 0.5
// (350/700) and a top value of 0.25 (50/200). The width and height values
// represent the dimensions of the bounding box as a ratio of the overall document
// page dimension. For example, if the document page size is 700 x 200 pixels, and
// the bounding box width is 70 pixels, the width returned is 0.1.
type BoundingBox struct {

	// The height of the bounding box as a ratio of the overall document page height.
	Height *float32

	// The left coordinate of the bounding box as a ratio of overall document page
	// width.
	Left *float32

	// The top coordinate of the bounding box as a ratio of overall document page
	// height.
	Top *float32

	// The width of the bounding box as a ratio of the overall document page width.
	Width *float32
}

// The input document, either as bytes or as an S3 object. You pass image bytes to
// an Amazon Textract API operation by using the Bytes property. For example, you
// would use the Bytes property to pass a document loaded from a local file system.
// Image bytes passed by using the Bytes property must be base64 encoded. Your code
// might not need to encode document file bytes if you're using an AWS SDK to call
// Amazon Textract API operations. You pass images stored in an S3 bucket to an
// Amazon Textract API operation by using the S3Object property. Documents stored
// in an S3 bucket don't need to be base64 encoded. The AWS Region for the S3
// bucket that contains the S3 object must match the AWS Region that you use for
// Amazon Textract operations. If you use the AWS CLI to call Amazon Textract
// operations, passing image bytes using the Bytes property isn't supported. You
// must first upload the document to an Amazon S3 bucket, and then call the
// operation using the S3Object property.  <p>For Amazon Textract to process an S3
// object, the user must have permission to access the S3 object. </p>
type Document struct {

	// A blob of base64-encoded document bytes. The maximum size of a document that's
	// provided in a blob of bytes is 5 MB. The document bytes must be in PNG or JPEG
	// format. If you're using an AWS SDK to call Amazon Textract, you might not need
	// to base64-encode image bytes passed using the Bytes field.
	Bytes []byte

	// Identifies an S3 object as the document source. The maximum size of a document
	// that's stored in an S3 bucket is 5 MB.
	S3Object *S3Object
}

// The Amazon S3 bucket that contains the document to be processed. It's used by
// asynchronous operations such as StartDocumentTextDetection (). The input
// document can be an image file in JPEG or PNG format. It can also be a file in
// PDF format.
type DocumentLocation struct {

	// The Amazon S3 bucket that contains the input document.
	S3Object *S3Object
}

// Information about the input document.
type DocumentMetadata struct {

	// The number of pages that are detected in the document.
	Pages *int32
}

// Information about where the following items are located on a document page:
// detected page, text, key-value pairs, tables, table cells, and selection
// elements.
type Geometry struct {

	// An axis-aligned coarse representation of the location of the recognized item on
	// the document page.
	BoundingBox *BoundingBox

	// Within the bounding box, a fine-grained polygon around the recognized item.
	Polygon []*Point
}

// Shows the results of the human in the loop evaluation. If there is no
// HumanLoopArn, the input did not trigger human review.
type HumanLoopActivationOutput struct {

	// Shows the result of condition evaluations, including those conditions which
	// activated a human review.
	// This value conforms to the media type: application/json
	HumanLoopActivationConditionsEvaluationResults *string

	// Shows if and why human review was needed.
	HumanLoopActivationReasons []*string

	// The Amazon Resource Name (ARN) of the HumanLoop created.
	HumanLoopArn *string
}

// Sets up the human review workflow the document will be sent to if one of the
// conditions is met. You can also set certain attributes of the image before
// review.
type HumanLoopConfig struct {

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// This member is required.
	FlowDefinitionArn *string

	// The name of the human workflow used for this image. This should be kept unique
	// within a region.
	//
	// This member is required.
	HumanLoopName *string

	// Sets attributes of the input data.
	DataAttributes *HumanLoopDataAttributes
}

// Allows you to set attributes of the image. Currently, you can declare an image
// as free of personally identifiable information and adult content.
type HumanLoopDataAttributes struct {

	// Sets whether the input image is free of personally identifiable information or
	// adult content.
	ContentClassifiers []ContentClassifier
}

// The Amazon Simple Notification Service (Amazon SNS) topic to which Amazon
// Textract publishes the completion status of an asynchronous document operation,
// such as StartDocumentTextDetection ().
type NotificationChannel struct {

	// The Amazon Resource Name (ARN) of an IAM role that gives Amazon Textract
	// publishing permissions to the Amazon SNS topic.
	//
	// This member is required.
	RoleArn *string

	// The Amazon SNS topic that Amazon Textract posts the completion status to.
	//
	// This member is required.
	SNSTopicArn *string
}

// The X and Y coordinates of a point on a document page. The X and Y values that
// are returned are ratios of the overall document page size. For example, if the
// input document is 700 x 200 and the operation returns X=0.5 and Y=0.25, then the
// point is at the (350,50) pixel coordinate on the document page.  <p>An array of
// <code>Point</code> objects, <code>Polygon</code>, is returned by
// <a>DetectDocumentText</a>. <code>Polygon</code> represents a fine-grained
// polygon around detected text. For more information, see Geometry in the Amazon
// Textract Developer Guide. </p>
type Point struct {

	// The value of the X coordinate for a point on a Polygon.
	X *float32

	// The value of the Y coordinate for a point on a Polygon.
	Y *float32
}

// Information about how blocks are related to each other. A Block object contains
// 0 or more Relation objects in a list, Relationships. For more information, see
// Block (). The Type element provides the type of the relationship for all blocks
// in the IDs array.
type Relationship struct {

	// An array of IDs for related blocks. You can get the type of the relationship
	// from the Type element.
	Ids []*string

	// The type of relationship that the blocks in the IDs array have with the current
	// block. The relationship can be VALUE or CHILD. A relationship of type VALUE is a
	// list that contains the ID of the VALUE block that's associated with the KEY of a
	// key-value pair. A relationship of type CHILD is a list of IDs that identify WORD
	// blocks.
	Type RelationshipType
}

// The S3 bucket name and file name that identifies the document. The AWS Region
// for the S3 bucket that contains the document must match the Region that you use
// for Amazon Textract operations.  <p>For Amazon Textract to process a file in an
// S3 bucket, the user must have permission to access the S3 bucket and file. </p>
type S3Object struct {

	// The name of the S3 bucket.
	Bucket *string

	// The file name of the input document. Synchronous operations can use image files
	// that are in JPEG or PNG format. Asynchronous operations also support PDF format
	// files.
	Name *string

	// If the bucket has versioning enabled, you can specify the object version.
	Version *string
}

// A warning about an issue that occurred during asynchronous text analysis
// (StartDocumentAnalysis ()) or asynchronous document text detection
// (StartDocumentTextDetection ()).
type Warning struct {

	// The error code for the warning.
	ErrorCode *string

	// A list of the pages that the warning applies to.
	Pages []*int32
}
