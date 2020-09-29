// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Stores account attributes.
type AccountAttribute struct {
	// The attribute name. The following are supported attribute names.
	//
	//     *
	// ServerLimit: The number of current servers/maximum number of servers allowed. By
	// default, you can have a maximum of 10 servers.
	//
	//     * ManualBackupLimit: The
	// number of current manual backups/maximum number of backups allowed. By default,
	// you can have a maximum of 50 manual backups saved.
	Name *string
	// The maximum allowed value.
	Maximum *int32
	// The current usage, such as the current number of servers that are associated
	// with the account.
	Used *int32
}

// Describes a single backup.
type Backup struct {
	// The generated ID of the backup. Example: myServerName-yyyyMMddHHmmssSSS
	BackupId *string
	// An informational message about backup status.
	StatusDescription *string
	// The Amazon S3 URL of the backup's log file.
	S3LogUrl *string
	// The status of a backup while in progress.
	Status BackupStatus
	// The version of AWS OpsWorks CM-specific tools that is obtained from the server
	// when the backup is created.
	ToolsVersion *string
	// The engine type that is obtained from the server when the backup is created.
	Engine *string
	// The backup type. Valid values are automated or manual.
	BackupType BackupType
	// The engine model that is obtained from the server when the backup is created.
	EngineModel *string
	// The instance type that is obtained from the server when the backup is created.
	InstanceType *string
	// A user-provided description for a manual backup. This field is empty for
	// automated backups.
	Description *string
	// The key pair that is obtained from the server when the backup is created.
	KeyPair *string
	// The IAM user ARN of the requester for manual backups. This field is empty for
	// automated backups.
	UserArn *string
	// This field is deprecated and is no longer used.
	S3DataSize *int32
	// The service role ARN that is obtained from the server when the backup is
	// created.
	ServiceRoleArn *string
	// The name of the server from which the backup was made.
	ServerName *string
	// The time stamp when the backup was created in the database. Example:
	// 2016-07-29T13:38:47.520Z
	CreatedAt *time.Time
	// The ARN of the backup.
	BackupArn *string
	// The engine version that is obtained from the server when the backup is created.
	EngineVersion *string
	// The security group IDs that are obtained from the server when the backup is
	// created.
	SecurityGroupIds []*string
	// The subnet IDs that are obtained from the server when the backup is created.
	SubnetIds []*string
	// The preferred maintenance period that is obtained from the server when the
	// backup is created.
	PreferredMaintenanceWindow *string
	// The preferred backup period that is obtained from the server when the backup is
	// created.
	PreferredBackupWindow *string
	// This field is deprecated and is no longer used.
	S3DataUrl *string
	// The EC2 instance profile ARN that is obtained from the server when the backup is
	// created. Because this value is stored, you are not required to provide the
	// InstanceProfileArn again if you restore a backup.
	InstanceProfileArn *string
}

// A name and value pair that is specific to the engine of the server.
type EngineAttribute struct {
	// The name of the engine attribute.
	Name *string
	// The value of the engine attribute.
	Value *string
}

// Describes a configuration management server.
type Server struct {
	// A DNS name that can be used to access the engine. Example:
	// myserver-asdfghjkl.us-east-1.opsworks.io. You cannot access the server by using
	// the Endpoint value if the server has a CustomDomain specified.
	Endpoint *string
	// The engine model of the server. Valid values in this release include Monolithic
	// for Puppet and Single for Chef.
	EngineModel *string
	// The preferred maintenance period specified for the server.
	PreferredMaintenanceWindow *string
	// The response of a createServer() request returns the master credential to access
	// the server in EngineAttributes. These credentials are not stored by AWS OpsWorks
	// CM; they are returned only as part of the result of createServer(). Attributes
	// returned in a createServer response for Chef
	//
	//     * CHEF_AUTOMATE_PIVOTAL_KEY: A
	// base64-encoded RSA private key that is generated by AWS OpsWorks for Chef
	// Automate. This private key is required to access the Chef API.
	//
	//     *
	// CHEF_STARTER_KIT: A base64-encoded ZIP file. The ZIP file contains a Chef
	// starter kit, which includes a README, a configuration file, and the required RSA
	// private key. Save this file, unzip it, and then change to the directory where
	// you've unzipped the file contents. From this directory, you can run Knife
	// commands.
	//
	// Attributes returned in a createServer response for Puppet
	//
	//     *
	// PUPPET_STARTER_KIT: A base64-encoded ZIP file. The ZIP file contains a Puppet
	// starter kit, including a README and a required private key. Save this file,
	// unzip it, and then change to the directory where you've unzipped the file
	// contents.
	//
	//     * PUPPET_ADMIN_PASSWORD: An administrator password that you can
	// use to sign in to the Puppet Enterprise console after the server is online.
	EngineAttributes []*EngineAttribute
	// The name of the server.
	ServerName *string
	// Disables automated backups. The number of stored backups is dependent on the
	// value of PreferredBackupCount.
	DisableAutomatedBackup *bool
	// The service role ARN used to create the server.
	ServiceRoleArn *string
	// Time stamp of server creation. Example 2016-07-29T13:38:47.520Z
	CreatedAt *time.Time
	// The key pair associated with the server.
	KeyPair *string
	// The instance profile ARN of the server.
	InstanceProfileArn *string
	// Associate a public IP address with a server that you are launching.
	AssociatePublicIpAddress *bool
	// The subnet IDs specified in a CreateServer request.
	SubnetIds []*string
	// The preferred backup period specified for the server.
	PreferredBackupWindow *string
	// Depending on the server status, this field has either a human-readable message
	// (such as a create or backup error), or an escaped block of JSON (used for health
	// check results).
	StatusReason *string
	// The engine type of the server. Valid values in this release include ChefAutomate
	// and Puppet.
	Engine *string
	// The number of automated backups to keep.
	BackupRetentionCount *int32
	// The security group IDs for the server, as specified in the CloudFormation stack.
	// These might not be the same security groups that are shown in the EC2 console.
	SecurityGroupIds []*string
	// The engine version of the server. For a Chef server, the valid value for
	// EngineVersion is currently 2. For a Puppet server, the valid value is 2017.
	EngineVersion *string
	// An optional public endpoint of a server, such as https://aws.my-company.com. You
	// cannot access the server by using the Endpoint value if the server has a
	// CustomDomain specified.
	CustomDomain *string
	// The ARN of the server.
	ServerArn *string
	// The ARN of the CloudFormation stack that was used to create the server.
	CloudFormationStackArn *string
	// The server's status. This field displays the states of actions in progress, such
	// as creating, running, or backing up the server, as well as the server's health
	// state.
	Status ServerStatus
	// The status of the most recent server maintenance run. Shows SUCCESS or FAILED.
	MaintenanceStatus MaintenanceStatus
	// The instance type for the server, as specified in the CloudFormation stack. This
	// might not be the same instance type that is shown in the EC2 console.
	InstanceType *string
}

// An event that is related to the server, such as the start of maintenance or
// backup.
type ServerEvent struct {
	// The time when the event occurred.
	CreatedAt *time.Time
	// A human-readable informational or status message.
	Message *string
	// The Amazon S3 URL of the event's log file.
	LogUrl *string
	// The name of the server on or for which the event occurred.
	ServerName *string
}

// A map that contains tag keys and tag values to attach to an AWS OpsWorks for
// Chef Automate or AWS OpsWorks for Puppet Enterprise server. Leading and trailing
// white spaces are trimmed from both the key and value. A maximum of 50
// user-applied tags is allowed for tag-supported AWS OpsWorks-CM resources.
type Tag struct {
	// A tag key, such as Stage or Name. A tag key cannot be empty. The key can be a
	// maximum of 127 characters, and can contain only Unicode letters, numbers, or
	// separators, or the following special characters: + - = . _ : /
	Key *string
	// An optional tag value, such as Production or test-owcm-server. The value can be
	// a maximum of 255 characters, and contain only Unicode letters, numbers, or
	// separators, or the following special characters: + - = . _ : /
	Value *string
}
