// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the Data Catalog resource policy for access control.
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	stack := middleware.NewStack("PutResourcePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutResourcePolicyMiddlewares(stack)
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
	addOpPutResourcePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before)
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
			OperationName: "PutResourcePolicy",
			Err:           err,
		}
	}
	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// Contains the policy document to set, in JSON format.
	//
	// This member is required.
	PolicyInJson *string

	// Allows you to specify if you want to use both resource-level and
	// account/catalog-level resource policies. A resource-level policy is a policy
	// attached to an individual resource such as a database or a table.  <p>The
	// default value of <code>NO</code> indicates that resource-level policies cannot
	// co-exist with an account-level policy. A value of <code>YES</code> means the use
	// of both resource-level and account/catalog-level resource policies is
	// allowed.</p>
	EnableHybrid types.EnableHybridValues

	// A value of MUST_EXIST is used to update a policy. A value of NOT_EXIST is used
	// to create a new policy. If a value of NONE or a null value is used, the call
	// will not depend on the existence of a policy.
	PolicyExistsCondition types.ExistCondition

	// The hash value returned when the previous policy was set using
	// PutResourcePolicy. Its purpose is to prevent concurrent modifications of a
	// policy. Do not use this parameter if no previous policy has been set.
	PolicyHashCondition *string

	// The ARN of the AWS Glue resource for the resource policy to be set. For more
	// information about AWS Glue resource ARNs, see the AWS Glue ARN string pattern
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html#aws-glue-api-regex-aws-glue-arn-id)
	ResourceArn *string
}

type PutResourcePolicyOutput struct {

	// A hash of the policy that has just been set. This must be included in a
	// subsequent call that overwrites or updates this policy.
	PolicyHash *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutResourcePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourcePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourcePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "PutResourcePolicy",
	}
}
