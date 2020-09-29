// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a specified instance. Required Permissions: To use this action, an IAM
// user must have a Manage permissions level for the stack, or an attached policy
// that explicitly grants permissions. For more information on user permissions,
// see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateInstance(ctx context.Context, params *UpdateInstanceInput, optFns ...func(*Options)) (*UpdateInstanceOutput, error) {
	stack := middleware.NewStack("UpdateInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateInstanceMiddlewares(stack)
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
	addOpUpdateInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstance(options.Region), middleware.Before)
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
			OperationName: "UpdateInstance",
			Err:           err,
		}
	}
	out := result.(*UpdateInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceInput struct {
	// The instance ID.
	InstanceId *string
	// For load-based or time-based instances, the type. Windows stacks can use only
	// time-based instances.
	AutoScalingType types.AutoScalingType
	// The instance architecture. Instance types do not necessarily support both
	// architectures. For a list of the architectures that are supported by the
	// different instance types, see Instance Families and Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	Architecture types.Architecture
	// The instance type, such as t2.micro. For a list of supported instance types,
	// open the stack in the console, choose Instances, and choose + Instance. The Size
	// list contains the currently supported types. For more information, see Instance
	// Families and Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html). The
	// parameter values that you use to specify the various types are in the API Name
	// column of the Available Instance Types table.
	InstanceType *string
	// The default AWS OpsWorks Stacks agent version. You have the following options:
	//
	//
	// * INHERIT - Use the stack's default agent version setting.
	//
	//     * version_number
	// - Use the specified agent version. This value overrides the stack's default
	// setting. To update the agent version, you must edit the instance configuration
	// and specify a new version. AWS OpsWorks Stacks then automatically installs that
	// version on the instance.
	//
	// The default setting is INHERIT. To specify an agent
	// version, you must use the complete version number, not the abbreviated number
	// shown on the console. For a list of available agent version numbers, call
	// DescribeAgentVersions (). AgentVersion cannot be set to Chef 12.2.
	AgentVersion *string
	// Whether to install operating system and package updates when the instance boots.
	// The default value is true. To control when updates are installed, set this value
	// to false. You must then update your instances manually by using CreateDeployment
	// () to run the update_dependencies stack command or by manually running yum
	// (Amazon Linux) or apt-get (Ubuntu) on the instances. We strongly recommend using
	// the default value of true, to ensure that your instances have the latest
	// security updates.
	InstallUpdatesOnBoot *bool
	// The instance host name.
	Hostname *string
	// This property cannot be updated.
	EbsOptimized *bool
	// The instance's layer IDs.
	LayerIds []*string
	// The ID of the AMI that was used to create the instance. The value of this
	// parameter must be the same AMI ID that the instance is already using. You cannot
	// apply a new AMI to an instance by running UpdateInstance. UpdateInstance does
	// not work on instances that are using custom AMIs.
	AmiId *string
	// The instance's operating system, which must be set to one of the following. You
	// cannot update an instance that is using a custom AMI.
	//
	//     * A supported Linux
	// operating system: An Amazon Linux version, such as Amazon Linux 2018.03, Amazon
	// Linux 2017.09, Amazon Linux 2017.03, Amazon Linux 2016.09, Amazon Linux 2016.03,
	// Amazon Linux 2015.09, or Amazon Linux 2015.03.
	//
	//     * A supported Ubuntu
	// operating system, such as Ubuntu 16.04 LTS, Ubuntu 14.04 LTS, or Ubuntu 12.04
	// LTS.
	//
	//     * CentOS Linux 7
	//
	//     * Red Hat Enterprise Linux 7
	//
	//     * A supported
	// Windows operating system, such as Microsoft Windows Server 2012 R2 Base,
	// Microsoft Windows Server 2012 R2 with SQL Server Express, Microsoft Windows
	// Server 2012 R2 with SQL Server Standard, or Microsoft Windows Server 2012 R2
	// with SQL Server Web.
	//
	// For more information about supported operating systems,
	// see AWS OpsWorks Stacks Operating Systems
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	// The default option is the current Amazon Linux version. If you set this
	// parameter to Custom, you must use the AmiId parameter to specify the custom AMI
	// that you want to use. For more information about supported operating systems,
	// see Operating Systems
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	// For more information about how to use custom AMIs with OpsWorks, see Using
	// Custom AMIs
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	// You can specify a different Linux operating system for the updated stack, but
	// you cannot change from Linux to Windows or Windows to Linux.
	Os *string
	// The instance's Amazon EC2 key name.
	SshKeyName *string
}

type UpdateInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateInstance",
	}
}
