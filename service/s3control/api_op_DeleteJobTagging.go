// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Removes the entire tag set from the specified S3 Batch Operations job. To use
// this operation, you must have permission to perform the s3:DeleteJobTagging
// action. For more information, see Controlling access and labeling jobs using
// tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
// in the Amazon Simple Storage Service Developer Guide. Related actions
// include:
//
// * CreateJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
// *
// GetJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html)
//
// *
// PutJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html)
func (c *Client) DeleteJobTagging(ctx context.Context, params *DeleteJobTaggingInput, optFns ...func(*Options)) (*DeleteJobTaggingOutput, error) {
	if params == nil {
		params = &DeleteJobTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJobTagging", params, optFns, addOperationDeleteJobTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJobTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobTaggingInput struct {

	// The AWS account ID associated with the S3 Batch Operations job.
	//
	// This member is required.
	AccountId *string

	// The ID for the S3 Batch Operations job whose tags you want to delete.
	//
	// This member is required.
	JobId *string
}

type DeleteJobTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteJobTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteJobTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteJobTagging{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opDeleteJobTaggingMiddleware(stack)
	addOpDeleteJobTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobTagging(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opDeleteJobTaggingMiddleware struct {
}

func (*endpointPrefix_opDeleteJobTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteJobTaggingMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteJobTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteJobTaggingMiddleware(stack *middleware.Stack) (err error) {
	err = stack.Serialize.Insert(&endpointPrefix_opDeleteJobTaggingMiddleware{}, `OperationSerializer`, middleware.Before)
	if err != nil {
		return err
	}
	err = stack.Build.Add(&smithyhttp.HostPrefixMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteJobTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteJobTagging",
	}
}

func (in DeleteJobTaggingInput) backfillAccountID(v string) (interface{}, error) {
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return &in, fmt.Errorf("error backfilling account id")
		}
		return &in, nil
	}
	in.AccountId = &v
	return &in, nil
}
