// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified ingress rules from a security group. To remove a rule, the
// values that you specify (for example, ports) must match the existing rule's
// values exactly. [EC2-Classic only] If the values you specify do not match the
// existing rule's values, no error is returned. Use DescribeSecurityGroups () to
// verify that the rule has been removed. Each rule consists of the protocol and
// the CIDR range or source security group. For the TCP and UDP protocols, you must
// also specify the destination port or range of ports. For the ICMP protocol, you
// must also specify the ICMP type and code. If the security group rule has a
// description, you do not have to specify the description to revoke the rule. Rule
// changes are propagated to instances within the security group as quickly as
// possible. However, a small delay might occur.
func (c *Client) RevokeSecurityGroupIngress(ctx context.Context, params *RevokeSecurityGroupIngressInput, optFns ...func(*Options)) (*RevokeSecurityGroupIngressOutput, error) {
	stack := middleware.NewStack("RevokeSecurityGroupIngress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRevokeSecurityGroupIngressMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSecurityGroupIngress(options.Region), middleware.Before)
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
			OperationName: "RevokeSecurityGroupIngress",
			Err:           err,
		}
	}
	out := result.(*RevokeSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeSecurityGroupIngressInput struct {

	// The CIDR IP address range. You can't specify this parameter when specifying a
	// source security group.
	CidrIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The start of port range for the TCP and UDP protocols, or an ICMP type number.
	// For the ICMP type number, use -1 to specify all ICMP types.
	FromPort *int32

	// The ID of the security group. You must specify either the security group ID or
	// the security group name in the request. For security groups in a nondefault VPC,
	// you must specify the security group ID.
	GroupId *string

	// [EC2-Classic, default VPC] The name of the security group. You must specify
	// either the security group ID or the security group name in the request.
	GroupName *string

	// The sets of IP permissions. You can't specify a source security group and a CIDR
	// IP address range in the same set of permissions.
	IpPermissions []*types.IpPermission

	// The IP protocol name (tcp, udp, icmp) or number (see Protocol Numbers
	// (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). Use
	// -1 to specify all.
	IpProtocol *string

	// [EC2-Classic, default VPC] The name of the source security group. You can't
	// specify this parameter in combination with the following parameters: the CIDR IP
	// address range, the start of the port range, the IP protocol, and the end of the
	// port range. For EC2-VPC, the source security group must be in the same VPC. To
	// revoke a specific rule for an IP protocol and port range, use a set of IP
	// permissions instead.
	SourceSecurityGroupName *string

	// [EC2-Classic] The AWS account ID of the source security group, if the source
	// security group is in a different account. You can't specify this parameter in
	// combination with the following parameters: the CIDR IP address range, the IP
	// protocol, the start of the port range, and the end of the port range. To revoke
	// a specific rule for an IP protocol and port range, use a set of IP permissions
	// instead.
	SourceSecurityGroupOwnerId *string

	// The end of port range for the TCP and UDP protocols, or an ICMP code number. For
	// the ICMP code number, use -1 to specify all ICMP codes for the ICMP type.
	ToPort *int32
}

type RevokeSecurityGroupIngressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRevokeSecurityGroupIngressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRevokeSecurityGroupIngress{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRevokeSecurityGroupIngress{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokeSecurityGroupIngress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RevokeSecurityGroupIngress",
	}
}
