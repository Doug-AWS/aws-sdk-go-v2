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

// Updates a specified stack. Required Permissions: To use this action, an IAM user
// must have a Manage permissions level for the stack, or an attached policy that
// explicitly grants permissions. For more information on user permissions, see
// Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateStack(ctx context.Context, params *UpdateStackInput, optFns ...func(*Options)) (*UpdateStackOutput, error) {
	stack := middleware.NewStack("UpdateStack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateStackMiddlewares(stack)
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
	addOpUpdateStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStack(options.Region), middleware.Before)
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
			OperationName: "UpdateStack",
			Err:           err,
		}
	}
	out := result.(*UpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStackInput struct {

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The default AWS OpsWorks Stacks agent version. You have the following options:
	//
	//
	// * Auto-update - Set this parameter to LATEST. AWS OpsWorks Stacks automatically
	// installs new agent versions on the stack's instances as soon as they are
	// available.
	//
	//     * Fixed version - Set this parameter to your preferred agent
	// version. To update the agent version, you must edit the stack configuration and
	// specify a new version. AWS OpsWorks Stacks then automatically installs that
	// version on the stack's instances.
	//
	// The default setting is LATEST. To specify an
	// agent version, you must use the complete version number, not the abbreviated
	// number shown on the console. For a list of available agent version numbers, call
	// DescribeAgentVersions (). AgentVersion cannot be set to Chef 12.2. You can also
	// specify an agent version when you create or update an instance, which overrides
	// the stack's default setting.
	AgentVersion *string

	// One or more user-defined key-value pairs to be added to the stack attributes.
	Attributes map[string]*string

	// A ChefConfiguration object that specifies whether to enable Berkshelf and the
	// Berkshelf version on Chef 11.10 stacks. For more information, see Create a New
	// Stack
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html).
	ChefConfiguration *types.ChefConfiguration

	// The configuration manager. When you update a stack, we recommend that you use
	// the configuration manager to specify the Chef version: 12, 11.10, or 11.4 for
	// Linux stacks, or 12.2 for Windows stacks. The default value for Linux stacks is
	// currently 12.
	ConfigurationManager *types.StackConfigurationManager

	// Contains the information required to retrieve an app or cookbook from a
	// repository. For more information, see Adding Apps
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html)
	// or Cookbooks and Recipes
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook.html).
	CustomCookbooksSource *types.Source

	// A string that contains user-defined, custom JSON. It can be used to override the
	// corresponding default stack configuration JSON values or to pass data to
	// recipes. The string should be in the following format: "{\"key1\": \"value1\",
	// \"key2\": \"value2\",...}" For more information about custom JSON, see Use
	// Custom JSON to Modify the Stack Configuration Attributes
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-json.html).
	CustomJson *string

	// The stack's default Availability Zone, which must be in the stack's region. For
	// more information, see Regions and Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html). If you also specify
	// a value for DefaultSubnetId, the subnet must be in the same zone. For more
	// information, see CreateStack ().
	DefaultAvailabilityZone *string

	// The ARN of an IAM profile that is the default profile for all of the stack's EC2
	// instances. For more information about IAM ARNs, see Using Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	DefaultInstanceProfileArn *string

	// The stack's operating system, which must be set to one of the following:
	//
	//     *
	// A supported Linux operating system: An Amazon Linux version, such as Amazon
	// Linux 2018.03, Amazon Linux 2017.09, Amazon Linux 2017.03, Amazon Linux 2016.09,
	// Amazon Linux 2016.03, Amazon Linux 2015.09, or Amazon Linux 2015.03.
	//
	//     * A
	// supported Ubuntu operating system, such as Ubuntu 16.04 LTS, Ubuntu 14.04 LTS,
	// or Ubuntu 12.04 LTS.
	//
	//     * CentOS Linux 7
	//
	//     * Red Hat Enterprise Linux 7
	//
	//
	// * A supported Windows operating system, such as Microsoft Windows Server 2012 R2
	// Base, Microsoft Windows Server 2012 R2 with SQL Server Express, Microsoft
	// Windows Server 2012 R2 with SQL Server Standard, or Microsoft Windows Server
	// 2012 R2 with SQL Server Web.
	//
	//     * A custom AMI: Custom. You specify the custom
	// AMI you want to use when you create instances. For more information about how to
	// use custom AMIs with OpsWorks, see Using Custom AMIs
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	//
	// The
	// default option is the stack's current operating system. For more information
	// about supported operating systems, see AWS OpsWorks Stacks Operating Systems
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	DefaultOs *string

	// The default root device type. This value is used by default for all instances in
	// the stack, but you can override it when you create an instance. For more
	// information, see Storage for the Root Device
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device).
	DefaultRootDeviceType types.RootDeviceType

	// A default Amazon EC2 key-pair name. The default value is none. If you specify a
	// key-pair name, AWS OpsWorks Stacks installs the public key on the instance and
	// you can use the private key with an SSH client to log in to the instance. For
	// more information, see  Using SSH to Communicate with an Instance
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-ssh.html)
	// and  Managing SSH Access
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/security-ssh-access.html).
	// You can override this setting by specifying a different key pair, or no key
	// pair, when you  create an instance
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html).
	DefaultSshKeyName *string

	// The stack's default VPC subnet ID. This parameter is required if you specify a
	// value for the VpcId parameter. All instances are launched into this subnet
	// unless you specify otherwise when you create the instance. If you also specify a
	// value for DefaultAvailabilityZone, the subnet must be in that zone. For
	// information on default values and when this parameter is required, see the VpcId
	// parameter description.
	DefaultSubnetId *string

	// The stack's new host name theme, with spaces replaced by underscores. The theme
	// is used to generate host names for the stack's instances. By default,
	// HostnameTheme is set to Layer_Dependent, which creates host names by appending
	// integers to the layer's short name. The other themes are:
	//
	//     * Baked_Goods
	//
	//
	// * Clouds
	//
	//     * Europe_Cities
	//
	//     * Fruits
	//
	//     * Greek_Deities_and_Titans
	//
	//
	// * Legendary_creatures_from_Japan
	//
	//     * Planets_and_Moons
	//
	//     * Roman_Deities
	//
	//
	// * Scottish_Islands
	//
	//     * US_Cities
	//
	//     * Wild_Cats
	//
	// To obtain a generated host
	// name, call GetHostNameSuggestion, which returns a host name based on the current
	// theme.
	HostnameTheme *string

	// The stack's new name.
	Name *string

	// Do not use this parameter. You cannot update a stack's service role.
	ServiceRoleArn *string

	// Whether the stack uses custom cookbooks.
	UseCustomCookbooks *bool

	// Whether to associate the AWS OpsWorks Stacks built-in security groups with the
	// stack's layers. AWS OpsWorks Stacks provides a standard set of built-in security
	// groups, one for each layer, which are associated with layers by default.
	// UseOpsworksSecurityGroups allows you to provide your own custom security groups
	// instead of using the built-in groups. UseOpsworksSecurityGroups has the
	// following settings:
	//
	//     * True - AWS OpsWorks Stacks automatically associates
	// the appropriate built-in security group with each layer (default setting). You
	// can associate additional security groups with a layer after you create it, but
	// you cannot delete the built-in security group.
	//
	//     * False - AWS OpsWorks
	// Stacks does not associate built-in security groups with layers. You must create
	// appropriate EC2 security groups and associate a security group with each layer
	// that you create. However, you can still manually associate a built-in security
	// group with a layer on. Custom security groups are required only for those layers
	// that need custom settings.
	//
	// For more information, see Create a New Stack
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-creating.html).
	UseOpsworksSecurityGroups *bool
}

type UpdateStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateStackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStack{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateStack",
	}
}
