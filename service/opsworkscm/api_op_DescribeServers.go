// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all configuration management servers that are identified with your
// account. Only the stored results from Amazon DynamoDB are returned. AWS OpsWorks
// CM does not query other services. This operation is synchronous. A
// ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid.
func (c *Client) DescribeServers(ctx context.Context, params *DescribeServersInput, optFns ...func(*Options)) (*DescribeServersOutput, error) {
	if params == nil {
		params = &DescribeServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServers", params, optFns, addOperationDescribeServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServersInput struct {

	// This is not currently implemented for DescribeServers requests.
	MaxResults *int32

	// This is not currently implemented for DescribeServers requests.
	NextToken *string

	// Describes the server with the specified ServerName.
	ServerName *string
}

type DescribeServersOutput struct {

	// This is not currently implemented for DescribeServers requests.
	NextToken *string

	// Contains the response to a DescribeServers request. For Chef Automate servers:
	// If DescribeServersResponse$Servers$EngineAttributes includes
	// CHEF_MAJOR_UPGRADE_AVAILABLE, you can upgrade the Chef Automate server to Chef
	// Automate 2. To be eligible for upgrade, a server running Chef Automate 1 must
	// have had at least one successful maintenance run after November 1, 2019. For
	// Puppet Server: DescribeServersResponse$Servers$EngineAttributes contains
	// PUPPET_API_CA_CERT. This is the PEM-encoded CA certificate that is used by the
	// Puppet API over TCP port number 8140. The CA certificate is also used to sign
	// node certificates.
	Servers []*types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServers{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeServers",
	}
}
