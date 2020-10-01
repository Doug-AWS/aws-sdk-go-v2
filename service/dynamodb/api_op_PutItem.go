// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new item, or replaces an old item with a new item. If an item that has
// the same primary key as the new item already exists in the specified table, the
// new item completely replaces the existing item. You can perform a conditional
// put operation (add a new item if one with the specified primary key doesn't
// exist), or replace an existing item if it has certain attribute values. You can
// return the item's attribute values in the same operation, using the ReturnValues
// parameter. This topic provides general information about the PutItem API. For
// information on how to call the PutItem API using the AWS SDK in specific
// languages, see the following:
//
//     * PutItem in the AWS Command Line Interface
// (http://docs.aws.amazon.com/goto/aws-cli/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for .NET
// (http://docs.aws.amazon.com/goto/DotNetSDKV3/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for C++
// (http://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for Go
// (http://docs.aws.amazon.com/goto/SdkForGoV1/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for Java
// (http://docs.aws.amazon.com/goto/SdkForJava/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for JavaScript
// (http://docs.aws.amazon.com/goto/AWSJavaScriptSDK/dynamodb-2012-08-10/PutItem)
//
//
// * PutItem in the AWS SDK for PHP V3
// (http://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for Python
// (http://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/PutItem)
//
//     *
// PutItem in the AWS SDK for Ruby V2
// (http://docs.aws.amazon.com/goto/SdkForRubyV2/dynamodb-2012-08-10/PutItem)
//
//
// <p>When you add an item, the primary key attributes are the only required
// attributes. Attribute values cannot be null.</p> <p>Empty String and Binary
// attribute values are allowed. Attribute values of type String and Binary must
// have a length greater than zero if the attribute is used as a key attribute for
// a table or index. Set type attributes cannot be empty. </p> <p>Invalid Requests
// with empty values will be rejected with a <code>ValidationException</code>
// exception.</p> <note> <p>To prevent a new item from replacing an existing item,
// use a conditional expression that contains the <code>attribute_not_exists</code>
// function with the name of the attribute being used as the partition key for the
// table. Since every record must contain that attribute, the
// <code>attribute_not_exists</code> function will only succeed if no matching item
// exists.</p> </note> <p>For more information about <code>PutItem</code>, see <a
// href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html">Working
// with Items</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
func (c *Client) PutItem(ctx context.Context, params *PutItemInput, optFns ...func(*Options)) (*PutItemOutput, error) {
	stack := middleware.NewStack("PutItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpPutItemMiddlewares(stack)
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
	addOpPutItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutItem(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "PutItem",
			Err:           err,
		}
	}
	out := result.(*PutItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutItem operation.
type PutItemInput struct {

	// A map of attribute name/value pairs, one for each attribute. Only the primary
	// key attributes are required; you can optionally provide other attribute
	// name-value pairs for the item. You must provide all of the attributes for the
	// primary key. For example, with a simple primary key, you only need to provide a
	// value for the partition key. For a composite primary key, you must provide both
	// values for both the partition key and the sort key. If you specify any
	// attributes that are part of an index key, then the data types for those
	// attributes must match those of the schema in the table's attribute definition.
	// Empty String and Binary attribute values are allowed. Attribute values of type
	// String and Binary must have a length greater than zero if the attribute is used
	// as a key attribute for a table or index.  <p>For more information about primary
	// keys, see <a
	// href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html#HowItWorks.CoreComponents.PrimaryKey">Primary
	// Key</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>Each element in
	// the <code>Item</code> map is an <code>AttributeValue</code> object.</p>
	//
	// This member is required.
	Item map[string]*types.AttributeValue

	// The name of the table to contain the item.
	//
	// This member is required.
	TableName *string

	// A condition that must be satisfied in order for a conditional PutItem operation
	// to succeed. An expression can contain any of the following:
	//
	//     * Functions:
	// attribute_exists | attribute_not_exists | attribute_type | contains |
	// begins_with | size These function names are case-sensitive.
	//
	//     * Comparison
	// operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//
	//     * Logical operators: AND
	// | OR | NOT
	//
	// For more information on condition expressions, see Condition
	// Expressions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string

	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see ConditionalOperator
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator types.ConditionalOperator

	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see Expected
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]*types.ExpectedAttributeValue

	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames:
	//
	//     * To
	// access an attribute whose name conflicts with a DynamoDB reserved word.
	//
	//     *
	// To create a placeholder for repeating occurrences of an attribute name in an
	// expression.
	//
	//     * To prevent special characters in an attribute name from being
	// misinterpreted in an expression.
	//
	// Use the # character in an expression to
	// dereference an attribute name. For example, consider the following attribute
	// name:
	//
	//     * Percentile
	//
	// The name of this attribute conflicts with a reserved
	// word, so it cannot be used directly in an expression. (For the complete list of
	// reserved words, see Reserved Words
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//     * {"#P":"Percentile"}
	//
	// You
	// could then use this substitution in an expression, as in this example:
	//
	//     * #P
	// = :val
	//
	// Tokens that begin with the : character are expression attribute values,
	// which are placeholders for the actual value at runtime. For more information on
	// expression attribute names, see Specifying Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]*string

	// One or more values that can be substituted in an expression. Use the : (colon)
	// character in an expression to dereference an attribute value. For example,
	// suppose that you wanted to check whether the value of the ProductStatus
	// attribute was one of the following: Available | Backordered | Discontinued You
	// would first need to specify ExpressionAttributeValues as follows: {
	// ":avail":{"S":"Available"}, ":back":{"S":"Backordered"},
	// ":disc":{"S":"Discontinued"} } You could then use these values in an expression,
	// such as this: ProductStatus IN (:avail, :back, :disc) For more information on
	// expression attribute values, see Condition Expressions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]*types.AttributeValue

	// Determines the level of detail about provisioned throughput consumption that is
	// returned in the response:
	//
	//     * INDEXES - The response includes the aggregate
	// ConsumedCapacity for the operation, together with ConsumedCapacity for each
	// table and secondary index that was accessed. Note that some operations, such as
	// GetItem and BatchGetItem, do not access any indexes at all. In these cases,
	// specifying INDEXES will only return ConsumedCapacity information for table(s).
	//
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity for the
	// operation.
	//
	//     * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	// Determines whether item collection metrics are returned. If set to SIZE, the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the default),
	// no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics

	// Use ReturnValues if you want to get the item attributes as they appeared before
	// they were updated with the PutItem request. For PutItem, the valid values are:
	//
	//
	// * NONE - If ReturnValues is not specified, or if its value is NONE, then nothing
	// is returned. (This setting is the default for ReturnValues.)
	//
	//     * ALL_OLD - If
	// PutItem overwrote an attribute name-value pair, then the content of the old item
	// is returned.
	//
	// The ReturnValues parameter is used by several DynamoDB operations;
	// however, PutItem does not recognize any values other than NONE or ALL_OLD.
	ReturnValues types.ReturnValue
}

// Represents the output of a PutItem operation.
type PutItemOutput struct {

	// The attribute values as they appeared before the PutItem operation, but only if
	// ReturnValues is specified as ALL_OLD in the request. Each element consists of an
	// attribute name and an attribute value.
	Attributes map[string]*types.AttributeValue

	// The capacity units consumed by the PutItem operation. The data returned includes
	// the total provisioned throughput consumed, along with statistics for the table
	// and any indexes involved in the operation. ConsumedCapacity is only returned if
	// the ReturnConsumedCapacity parameter was specified. For more information, see
	// Read/Write Capacity Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity

	// Information about item collections, if any, that were affected by the PutItem
	// operation. ItemCollectionMetrics is only returned if the
	// ReturnItemCollectionMetrics parameter was specified. If the table does not have
	// any local secondary indexes, this information is not returned in the response.
	// Each ItemCollectionMetrics element consists of:
	//
	//     * ItemCollectionKey - The
	// partition key value of the item collection. This is the same as the partition
	// key value of the item itself.
	//
	//     * SizeEstimateRangeGB - An estimate of item
	// collection size, in gigabytes. This value is a two-element array containing a
	// lower bound and an upper bound for the estimate. The estimate includes the size
	// of all the items in the table, plus the size of all attributes projected into
	// all of the local secondary indexes on that table. Use this estimate to measure
	// whether a local secondary index is approaching its size limit. The estimate is
	// subject to change over time; therefore, do not rely on the precision or accuracy
	// of the estimate.
	ItemCollectionMetrics *types.ItemCollectionMetrics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpPutItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpPutItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "PutItem",
	}
}
