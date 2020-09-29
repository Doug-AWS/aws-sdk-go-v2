// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains information about a backup of an AWS CloudHSM cluster. All backup
// objects contain the BackupId, BackupState, ClusterId, and CreateTimestamp
// parameters. Backups that were copied into a destination region additionally
// contain the CopyTimestamp, SourceBackup, SourceCluster, and SourceRegion
// paramters. A backup that is pending deletion will include the DeleteTimestamp
// parameter.
type Backup struct {
	// The AWS Region that contains the source backup from which the new backup was
	// copied.
	SourceRegion *string
	// The identifier (ID) of the backup.
	BackupId *string
	// The date and time when the backup was copied from a source backup.
	CopyTimestamp *time.Time
	// The date and time when the backup will be permanently deleted.
	DeleteTimestamp *time.Time
	// The identifier (ID) of the cluster that was backed up.
	ClusterId *string
	// The identifier (ID) of the source backup from which the new backup was copied.
	SourceBackup *string
	// The date and time when the backup was created.
	CreateTimestamp *time.Time
	// The list of tags for the backup.
	TagList []*Tag
	// The identifier (ID) of the cluster containing the source backup from which the
	// new backup was copied.
	SourceCluster *string
	// The state of the backup.
	BackupState BackupState
}

// Contains one or more certificates or a certificate signing request (CSR).
type Certificates struct {
	// The cluster's certificate signing request (CSR). The CSR exists only when the
	// cluster's state is UNINITIALIZED.
	ClusterCsr *string
	// The HSM certificate issued (signed) by the HSM hardware.
	HsmCertificate *string
	// The HSM hardware certificate issued (signed) by the hardware manufacturer.
	ManufacturerHardwareCertificate *string
	// The cluster certificate issued (signed) by the issuing certificate authority
	// (CA) of the cluster's owner.
	ClusterCertificate *string
	// The HSM hardware certificate issued (signed) by AWS CloudHSM.
	AwsHardwareCertificate *string
}

// Contains information about an AWS CloudHSM cluster.
type Cluster struct {
	// Contains one or more certificates or a certificate signing request (CSR).
	Certificates *Certificates
	// The cluster's backup policy.
	BackupPolicy BackupPolicy
	// The cluster's identifier (ID).
	ClusterId *string
	// The cluster's state.
	State ClusterState
	// Contains information about the HSMs in the cluster.
	Hsms []*Hsm
	// The identifier (ID) of the virtual private cloud (VPC) that contains the
	// cluster.
	VpcId *string
	// The identifier (ID) of the backup used to create the cluster. This value exists
	// only when the cluster was created from a backup.
	SourceBackupId *string
	// The identifier (ID) of the cluster's security group.
	SecurityGroup *string
	// The list of tags for the cluster.
	TagList []*Tag
	// The default password for the cluster's Pre-Crypto Officer (PRECO) user.
	PreCoPassword *string
	// A map from availability zone to the cluster’s subnet in that availability zone.
	SubnetMapping map[string]*string
	// A description of the cluster's state.
	StateMessage *string
	// The type of HSM that the cluster contains.
	HsmType *string
	// The date and time when the cluster was created.
	CreateTimestamp *time.Time
}

// Contains information about the backup that will be copied and created by the
// CopyBackupToRegion () operation.
type DestinationBackup struct {
	// The identifier (ID) of the cluster containing the source backup from which the
	// new backup was copied.
	SourceCluster *string
	// The date and time when both the source backup was created.
	CreateTimestamp *time.Time
	// The identifier (ID) of the source backup from which the new backup was copied.
	SourceBackup *string
	// The AWS region that contains the source backup from which the new backup was
	// copied.
	SourceRegion *string
}

// Contains information about a hardware security module (HSM) in an AWS CloudHSM
// cluster.
type Hsm struct {
	// The subnet that contains the HSM's elastic network interface (ENI).
	SubnetId *string
	// The HSM's identifier (ID).
	HsmId *string
	// A description of the HSM's state.
	StateMessage *string
	// The identifier (ID) of the HSM's elastic network interface (ENI).
	EniId *string
	// The HSM's state.
	State HsmState
	// The IP address of the HSM's elastic network interface (ENI).
	EniIp *string
	// The Availability Zone that contains the HSM.
	AvailabilityZone *string
	// The identifier (ID) of the cluster that contains the HSM.
	ClusterId *string
}

// Contains a tag. A tag is a key-value pair.
type Tag struct {
	// The key of the tag.
	Key *string
	// The value of the tag.
	Value *string
}
