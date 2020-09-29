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

// Returns the default engine and system parameter information for the specified
// database engine.
func (c *Client) DescribeEngineDefaultParameters(ctx context.Context, params *DescribeEngineDefaultParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error) {
	stack := middleware.NewStack("DescribeEngineDefaultParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeEngineDefaultParametersMiddlewares(stack)
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
	addOpDescribeEngineDefaultParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(options.Region), middleware.Before)
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
			OperationName: "DescribeEngineDefaultParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeEngineDefaultParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeEngineDefaultParametersInput struct {
	// The name of the DB parameter group family.
	DBParameterGroupFamily *string
	// An optional pagination token provided by a previous
	// DescribeEngineDefaultParameters request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// This parameter isn't currently supported.
	Filters []*types.Filter
}

type DescribeEngineDefaultParametersOutput struct {
	// Contains the result of a successful invocation of the
	// DescribeEngineDefaultParameters action.
	EngineDefaults *types.EngineDefaults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeEngineDefaultParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEngineDefaultParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEngineDefaultParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEngineDefaultParameters",
	}
}
