// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation replaces the current set of name servers for the domain with the
// specified set of name servers. If you use Amazon Route 53 as your DNS service,
// specify the four name servers in the delegation set for the hosted zone for the
// domain. If successful, this operation returns an operation ID that you can use
// to track the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Client) UpdateDomainNameservers(ctx context.Context, params *UpdateDomainNameserversInput, optFns ...func(*Options)) (*UpdateDomainNameserversOutput, error) {
	stack := middleware.NewStack("UpdateDomainNameservers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDomainNameserversMiddlewares(stack)
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
	addOpUpdateDomainNameserversValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainNameservers(options.Region), middleware.Before)
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
			OperationName: "UpdateDomainNameservers",
			Err:           err,
		}
	}
	out := result.(*UpdateDomainNameserversOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Replaces the current set of name servers for the domain with the specified set
// of name servers. If you use Amazon Route 53 as your DNS service, specify the
// four name servers in the delegation set for the hosted zone for the domain. If
// successful, this operation returns an operation ID that you can use to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
type UpdateDomainNameserversInput struct {

	// The name of the domain that you want to change name servers for.
	//
	// This member is required.
	DomainName *string

	// A list of new name servers for the domain.
	//
	// This member is required.
	Nameservers []*types.Nameserver

	// The authorization key for .fi domains
	FIAuthKey *string
}

// The UpdateDomainNameservers response includes the following element.
type UpdateDomainNameserversOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDomainNameserversMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainNameservers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainNameservers{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDomainNameservers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "UpdateDomainNameservers",
	}
}
