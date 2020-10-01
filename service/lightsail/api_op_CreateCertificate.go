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

// Creates an SSL/TLS certificate for a Amazon Lightsail content delivery network
// (CDN) distribution.  <p>After the certificate is created, use the
// <code>AttachCertificateToDistribution</code> action to attach the certificate to
// your distribution.</p> <important> <p>Only certificates created in the
// <code>us-east-1</code> AWS Region can be attached to Lightsail distributions.
// Lightsail distributions are global resources that can reference an origin in any
// AWS Region, and distribute its content globally. However, all distributions are
// located in the <code>us-east-1</code> Region.</p> </important>
func (c *Client) CreateCertificate(ctx context.Context, params *CreateCertificateInput, optFns ...func(*Options)) (*CreateCertificateOutput, error) {
	stack := middleware.NewStack("CreateCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCertificateMiddlewares(stack)
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
	addOpCreateCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificate(options.Region), middleware.Before)
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
			OperationName: "CreateCertificate",
			Err:           err,
		}
	}
	out := result.(*CreateCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateInput struct {

	// The name for the certificate.
	//
	// This member is required.
	CertificateName *string

	// The domain name (e.g., example.com) for the certificate.
	//
	// This member is required.
	DomainName *string

	// An array of strings that specify the alternate domains (e.g., example2.com) and
	// subdomains (e.g., blog.example.com) for the certificate.  <p>You can specify a
	// maximum of nine alternate domains (in addition to the primary domain name).</p>
	// <p>Wildcard domain entries (e.g., <code>*.example.com</code>) are not
	// supported.</p>
	SubjectAlternativeNames []*string

	// The tag keys and optional values to add to the certificate during create.
	// <p>Use the <code>TagResource</code> action to tag a resource after it's
	// created.</p>
	Tags []*types.Tag
}

type CreateCertificateOutput struct {

	// An object that describes the certificate created.
	Certificate *types.CertificateSummary

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateCertificate",
	}
}
