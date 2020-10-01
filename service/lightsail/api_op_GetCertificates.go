// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more Amazon Lightsail SSL/TLS certificates.
// <note> <p>To get a summary of a certificate, ommit
// <code>includeCertificateDetails</code> from your request. The response will
// include only the certificate Amazon Resource Name (ARN), certificate name,
// domain name, and tags.</p> </note>
func (c *Client) GetCertificates(ctx context.Context, params *GetCertificatesInput, optFns ...func(*Options)) (*GetCertificatesOutput, error) {
	stack := middleware.NewStack("GetCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCertificatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificates(options.Region), middleware.Before)
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
			OperationName: "GetCertificates",
			Err:           err,
		}
	}
	out := result.(*GetCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificatesInput struct {

	// The name for the certificate for which to return information.  <p>When omitted,
	// the response includes all of your certificates in the AWS region where the
	// request is made.</p>
	CertificateName *string

	// The status of the certificates for which to return information.  <p>For example,
	// specify <code>ISSUED</code> to return only certificates with an
	// <code>ISSUED</code> status.</p> <p>When omitted, the response includes all of
	// your certificates in the AWS region where the request is made, regardless of
	// their current status.</p>
	CertificateStatuses []types.CertificateStatus

	// Indicates whether to include detailed information about the certificates in the
	// response.  <p>When omitted, the response includes only the certificate names,
	// Amazon Resource Names (ARNs), domain names, and tags.</p>
	IncludeCertificateDetails *bool
}

type GetCertificatesOutput struct {

	// An object that describes certificates.
	Certificates []*types.CertificateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetCertificates",
	}
}
