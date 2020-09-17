// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about provisioned RDS instances. This API supports
// pagination. This operation can also return information for Amazon Neptune DB
// instances and Amazon DocumentDB instances.
func (c *Client) DescribeDBInstances(ctx context.Context, params *DescribeDBInstancesInput, optFns ...func(*Options)) (*DescribeDBInstancesOutput, error) {
	stack := middleware.NewStack("DescribeDBInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBInstancesMiddlewares(stack)
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
	addOpDescribeDBInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstances(options.Region), middleware.Before)

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
			OperationName: "DescribeDBInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBInstancesInput struct {
	// The user-supplied instance identifier. If this parameter is specified,
	// information from only the specific DB instance is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * If supplied, must match the identifier of an
	// existing DBInstance.
	DBInstanceIdentifier *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// A filter that specifies one or more DB instances to describe. Supported
	// filters:
	//
	//     * db-cluster-id - Accepts DB cluster identifiers and DB cluster
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB instances associated with the DB clusters identified by these
	// ARNs.
	//
	//     * db-instance-id - Accepts DB instance identifiers and DB instance
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB instances identified by these ARNs.
	//
	//     * dbi-resource-id -
	// Accepts DB instance resource identifiers. The results list will only include
	// information about the DB instances identified by these DB instance resource
	// identifiers.
	//
	//     * domain - Accepts Active Directory directory IDs. The results
	// list will only include information about the DB instances associated with these
	// domains.
	//
	//     * engine - Accepts engine names. The results list will only
	// include information about the DB instances for these engines.
	Filters []*types.Filter
	// An optional pagination token provided by a previous DescribeDBInstances request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string
}

// Contains the result of a successful invocation of the DescribeDBInstances
// action.
type DescribeDBInstancesOutput struct {
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string
	// A list of DBInstance instances.
	DBInstances []*types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBInstances",
	}
}