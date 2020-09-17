// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of orderable DB instance options for the specified engine.
func (c *Client) DescribeOrderableDBInstanceOptions(ctx context.Context, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
	stack := middleware.NewStack("DescribeOrderableDBInstanceOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeOrderableDBInstanceOptionsMiddlewares(stack)
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
	addOpDescribeOrderableDBInstanceOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(options.Region), middleware.Before)

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
			OperationName: "DescribeOrderableDBInstanceOptions",
			Err:           err,
		}
	}
	out := result.(*DescribeOrderableDBInstanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrderableDBInstanceOptionsInput struct {
	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string
	// The VPC filter value. Specify this parameter to show only the available VPC or
	// non-VPC offerings.
	Vpc *bool
	// The name of the engine to retrieve DB instance options for.
	Engine *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string
	// This parameter is not currently supported.
	Filters []*types.Filter
	// The license model filter value. Specify this parameter to show only the
	// available offerings matching the specified license model.
	LicenseModel *string
	// The engine version filter value. Specify this parameter to show only the
	// available offerings matching the specified engine version.
	EngineVersion *string
}

type DescribeOrderableDBInstanceOptionsOutput struct {
	// An OrderableDBInstanceOption () structure containing information about orderable
	// options for the DB instance.
	OrderableDBInstanceOptions []*types.OrderableDBInstanceOption
	// An optional pagination token provided by a previous OrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeOrderableDBInstanceOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeOrderableDBInstanceOptions",
	}
}