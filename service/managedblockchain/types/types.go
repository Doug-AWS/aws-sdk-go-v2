// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// A policy type that defines the voting rules for the network. The rules decide if
// a proposal is approved. Approval may be based on criteria such as the percentage
// of YES votes and the duration of the proposal. The policy applies to all
// proposals and is specified when the network is created.
type ApprovalThresholdPolicy struct {

	// The duration from the time that a proposal is created until it expires. If
	// members cast neither the required number of YES votes to approve the proposal
	// nor the number of NO votes required to reject it before the duration expires,
	// the proposal is EXPIRED and ProposalActions are not carried out.
	ProposalDurationInHours *int32

	// Determines whether the vote percentage must be greater than the
	// ThresholdPercentage or must be greater than or equal to the ThreholdPercentage
	// to be approved.
	ThresholdComparator ThresholdComparator

	// The percentage of votes among all members that must be YES for a proposal to be
	// approved. For example, a ThresholdPercentage value of 50 indicates 50%. The
	// ThresholdComparator determines the precise comparison. If a ThresholdPercentage
	// value of 50 is specified on a network with 10 members, along with a
	// ThresholdComparator value of GREATER_THAN, this indicates that 6 YES votes are
	// required for the proposal to be approved.
	ThresholdPercentage *int32
}

// An invitation to an AWS account to create a member and join the network.
type Invitation struct {

	// The date and time that the invitation was created.
	CreationDate *time.Time

	// The date and time that the invitation expires. This is the CreationDate plus the
	// ProposalDurationInHours that is specified in the ProposalThresholdPolicy. After
	// this date and time, the invitee can no longer create a member and join the
	// network using this InvitationId.
	ExpirationDate *time.Time

	// The unique identifier for the invitation.
	InvitationId *string

	// A summary of network configuration properties.
	NetworkSummary *NetworkSummary

	// The status of the invitation:
	//
	//     * PENDING - The invitee has not created a
	// member to join the network, and the invitation has not yet expired.
	//
	//     *
	// ACCEPTING - The invitee has begun creating a member, and creation has not yet
	// completed.
	//
	//     * ACCEPTED - The invitee created a member and joined the network
	// using the InvitationID.
	//
	//     * REJECTED - The invitee rejected the invitation.
	//
	//
	// * EXPIRED - The invitee neither created a member nor rejected the invitation
	// before the ExpirationDate.
	Status InvitationStatus
}

// An action to invite a specific AWS account to create a member and join the
// network. The InviteAction is carried out when a Proposal is APPROVED.
type InviteAction struct {

	// The AWS account ID to invite.
	//
	// This member is required.
	Principal *string
}

// A configuration for logging events.
type LogConfiguration struct {

	// Indicates whether logging is enabled.
	Enabled *bool
}

// A collection of log configurations.
type LogConfigurations struct {

	// Parameters for publishing logs to Amazon CloudWatch Logs.
	Cloudwatch *LogConfiguration
}

// Member configuration properties.
type Member struct {

	// The date and time that the member was created.
	CreationDate *time.Time

	// An optional description for the member.
	Description *string

	// Attributes relevant to a member for the blockchain framework that the Managed
	// Blockchain network uses.
	FrameworkAttributes *MemberFrameworkAttributes

	// The unique identifier of the member.
	Id *string

	// Configuration properties for logging events associated with a member.
	LogPublishingConfiguration *MemberLogPublishingConfiguration

	// The name of the member.
	Name *string

	// The unique identifier of the network to which the member belongs.
	NetworkId *string

	// The status of a member.
	//
	//     * CREATING - The AWS account is in the process of
	// creating a member.
	//
	//     * AVAILABLE - The member has been created and can
	// participate in the network.
	//
	//     * CREATE_FAILED - The AWS account attempted to
	// create a member and creation failed.
	//
	//     * DELETING - The member and all
	// associated resources are in the process of being deleted. Either the AWS account
	// that owns the member deleted it, or the member is being deleted as the result of
	// an APPROVEDPROPOSAL to remove the member.
	//
	//     * DELETED - The member can no
	// longer participate on the network and all associated resources are deleted.
	// Either the AWS account that owns the member deleted it, or the member is being
	// deleted as the result of an APPROVEDPROPOSAL to remove the member.
	Status MemberStatus
}

// Configuration properties of the member.
type MemberConfiguration struct {

	// Configuration properties of the blockchain framework relevant to the member.
	//
	// This member is required.
	FrameworkConfiguration *MemberFrameworkConfiguration

	// The name of the member.
	//
	// This member is required.
	Name *string

	// An optional description of the member.
	Description *string

	//
	LogPublishingConfiguration *MemberLogPublishingConfiguration
}

// Attributes of Hyperledger Fabric for a member in a Managed Blockchain network
// using the Hyperledger Fabric framework.
type MemberFabricAttributes struct {

	// The user name for the initial administrator user for the member.
	AdminUsername *string

	// The endpoint used to access the member's certificate authority.
	CaEndpoint *string
}

// Configuration properties for Hyperledger Fabric for a member in a Managed
// Blockchain network using the Hyperledger Fabric framework.
type MemberFabricConfiguration struct {

	// The password for the member's initial administrative user. The AdminPassword
	// must be at least eight characters long and no more than 32 characters. It must
	// contain at least one uppercase letter, one lowercase letter, and one digit. It
	// cannot have a single quote(‘), double quote(“), forward slash(/), backward
	// slash(\), @, or a space.
	//
	// This member is required.
	AdminPassword *string

	// The user name for the member's initial administrative user.
	//
	// This member is required.
	AdminUsername *string
}

// Configuration properties for logging events associated with a member of a
// Managed Blockchain network using the Hyperledger Fabric framework.
type MemberFabricLogPublishingConfiguration struct {

	// Configuration properties for logging events associated with a member's
	// Certificate Authority (CA). CA logs help you determine when a member in your
	// account joins the network, or when new peers register with a member CA.
	CaLogs *LogConfigurations
}

// Attributes relevant to a member for the blockchain framework that the Managed
// Blockchain network uses.
type MemberFrameworkAttributes struct {

	// Attributes of Hyperledger Fabric relevant to a member on a Managed Blockchain
	// network that uses Hyperledger Fabric.
	Fabric *MemberFabricAttributes
}

// Configuration properties relevant to a member for the blockchain framework that
// the Managed Blockchain network uses.
type MemberFrameworkConfiguration struct {

	// Attributes of Hyperledger Fabric for a member on a Managed Blockchain network
	// that uses Hyperledger Fabric.
	Fabric *MemberFabricConfiguration
}

// Configuration properties for logging events associated with a member of a
// Managed Blockchain network.
type MemberLogPublishingConfiguration struct {

	// Configuration properties for logging events associated with a member of a
	// Managed Blockchain network using the Hyperledger Fabric framework.
	Fabric *MemberFabricLogPublishingConfiguration
}

// A summary of configuration properties for a member.
type MemberSummary struct {

	// The date and time that the member was created.
	CreationDate *time.Time

	// An optional description of the member.
	Description *string

	// The unique identifier of the member.
	Id *string

	// An indicator of whether the member is owned by your AWS account or a different
	// AWS account.
	IsOwned *bool

	// The name of the member.
	Name *string

	// The status of the member.
	//
	//     * CREATING - The AWS account is in the process of
	// creating a member.
	//
	//     * AVAILABLE - The member has been created and can
	// participate in the network.
	//
	//     * CREATE_FAILED - The AWS account attempted to
	// create a member and creation failed.
	//
	//     * DELETING - The member and all
	// associated resources are in the process of being deleted. Either the AWS account
	// that owns the member deleted it, or the member is being deleted as the result of
	// an APPROVEDPROPOSAL to remove the member.
	//
	//     * DELETED - The member can no
	// longer participate on the network and all associated resources are deleted.
	// Either the AWS account that owns the member deleted it, or the member is being
	// deleted as the result of an APPROVEDPROPOSAL to remove the member.
	Status MemberStatus
}

// Network configuration properties.
type Network struct {

	// The date and time that the network was created.
	CreationDate *time.Time

	// Attributes of the blockchain framework for the network.
	Description *string

	// The blockchain framework that the network uses.
	Framework Framework

	// Attributes of the blockchain framework that the network uses.
	FrameworkAttributes *NetworkFrameworkAttributes

	// The version of the blockchain framework that the network uses.
	FrameworkVersion *string

	// The unique identifier of the network.
	Id *string

	// The name of the network.
	Name *string

	// The current status of the network.
	Status NetworkStatus

	// The voting rules for the network to decide if a proposal is accepted.
	VotingPolicy *VotingPolicy

	// The VPC endpoint service name of the VPC endpoint service of the network.
	// Members use the VPC endpoint service name to create a VPC endpoint to access
	// network resources.
	VpcEndpointServiceName *string
}

// Attributes of Hyperledger Fabric for a network.
type NetworkFabricAttributes struct {

	// The edition of Amazon Managed Blockchain that Hyperledger Fabric uses. For more
	// information, see Amazon Managed Blockchain Pricing
	// (http://aws.amazon.com/managed-blockchain/pricing/).
	Edition Edition

	// The endpoint of the ordering service for the network.
	OrderingServiceEndpoint *string
}

// Hyperledger Fabric configuration properties for the network.
type NetworkFabricConfiguration struct {

	// The edition of Amazon Managed Blockchain that the network uses. For more
	// information, see Amazon Managed Blockchain Pricing
	// (http://aws.amazon.com/managed-blockchain/pricing/).
	//
	// This member is required.
	Edition Edition
}

// Attributes relevant to the network for the blockchain framework that the network
// uses.
type NetworkFrameworkAttributes struct {

	// Attributes of Hyperledger Fabric for a Managed Blockchain network that uses
	// Hyperledger Fabric.
	Fabric *NetworkFabricAttributes
}

// Configuration properties relevant to the network for the blockchain framework
// that the network uses.
type NetworkFrameworkConfiguration struct {

	// Hyperledger Fabric configuration properties for a Managed Blockchain network
	// that uses Hyperledger Fabric.
	Fabric *NetworkFabricConfiguration
}

// A summary of network configuration properties.
type NetworkSummary struct {

	// The date and time that the network was created.
	CreationDate *time.Time

	// An optional description of the network.
	Description *string

	// The blockchain framework that the network uses.
	Framework Framework

	// The version of the blockchain framework that the network uses.
	FrameworkVersion *string

	// The unique identifier of the network.
	Id *string

	// The name of the network.
	Name *string

	// The current status of the network.
	Status NetworkStatus
}

// Configuration properties of a peer node.
type Node struct {

	// The Availability Zone in which the node exists.
	AvailabilityZone *string

	// The date and time that the node was created.
	CreationDate *time.Time

	// Attributes of the blockchain framework being used.
	FrameworkAttributes *NodeFrameworkAttributes

	// The unique identifier of the node.
	Id *string

	// The instance type of the node.
	InstanceType *string

	//
	LogPublishingConfiguration *NodeLogPublishingConfiguration

	// The unique identifier of the member to which the node belongs.
	MemberId *string

	// The unique identifier of the network that the node is in.
	NetworkId *string

	// The status of the node.
	Status NodeStatus
}

// Configuration properties of a peer node.
type NodeConfiguration struct {

	// The Availability Zone in which the node exists.
	//
	// This member is required.
	AvailabilityZone *string

	// The Amazon Managed Blockchain instance type for the node.
	//
	// This member is required.
	InstanceType *string

	//
	LogPublishingConfiguration *NodeLogPublishingConfiguration
}

// Attributes of Hyperledger Fabric for a peer node on a Managed Blockchain network
// that uses Hyperledger Fabric.
type NodeFabricAttributes struct {

	// The endpoint that identifies the peer node for all services except peer
	// channel-based event services.
	PeerEndpoint *string

	// The endpoint that identifies the peer node for peer channel-based event
	// services.
	PeerEventEndpoint *string
}

// Configuration properties for logging events associated with a peer node owned by
// a member in a Managed Blockchain network.
type NodeFabricLogPublishingConfiguration struct {

	// Configuration properties for logging events associated with chaincode execution
	// on a peer node. Chaincode logs contain the results of instantiating, invoking,
	// and querying the chaincode. A peer can run multiple instances of chaincode. When
	// enabled, a log stream is created for all chaincodes, with an individual log
	// stream for each chaincode.
	ChaincodeLogs *LogConfigurations

	// Configuration properties for a peer node log. Peer node logs contain messages
	// generated when your client submits transaction proposals to peer nodes, requests
	// to join channels, enrolls an admin peer, and lists the chaincode instances on a
	// peer node.
	PeerLogs *LogConfigurations
}

// Attributes relevant to a peer node on a Managed Blockchain network for the
// blockchain framework that the network uses.
type NodeFrameworkAttributes struct {

	// Attributes of Hyperledger Fabric for a peer node on a Managed Blockchain network
	// that uses Hyperledger Fabric.
	Fabric *NodeFabricAttributes
}

// Configuration properties for logging events associated with a peer node owned by
// a member in a Managed Blockchain network.
type NodeLogPublishingConfiguration struct {

	// Configuration properties for logging events associated with a node that is owned
	// by a member of a Managed Blockchain network using the Hyperledger Fabric
	// framework.
	Fabric *NodeFabricLogPublishingConfiguration
}

// A summary of configuration properties for a peer node.
type NodeSummary struct {

	// The Availability Zone in which the node exists.
	AvailabilityZone *string

	// The date and time that the node was created.
	CreationDate *time.Time

	// The unique identifier of the node.
	Id *string

	// The EC2 instance type for the node.
	InstanceType *string

	// The status of the node.
	Status NodeStatus
}

// Properties of a proposal on a Managed Blockchain network.
type Proposal struct {

	// The actions to perform on the network if the proposal is APPROVED.
	Actions *ProposalActions

	// The date and time that the proposal was created.
	CreationDate *time.Time

	// The description of the proposal.
	Description *string

	// The date and time that the proposal expires. This is the CreationDate plus the
	// ProposalDurationInHours that is specified in the ProposalThresholdPolicy. After
	// this date and time, if members have not cast enough votes to determine the
	// outcome according to the voting policy, the proposal is EXPIRED and Actions are
	// not carried out.
	ExpirationDate *time.Time

	// The unique identifier of the network for which the proposal is made.
	NetworkId *string

	// The current total of NO votes cast on the proposal by members.
	NoVoteCount *int32

	// The number of votes remaining to be cast on the proposal by members. In other
	// words, the number of members minus the sum of YES votes and NO votes.
	OutstandingVoteCount *int32

	// The unique identifier of the proposal.
	ProposalId *string

	// The unique identifier of the member that created the proposal.
	ProposedByMemberId *string

	// The name of the member that created the proposal.
	ProposedByMemberName *string

	// The status of the proposal. Values are as follows:
	//
	//     * IN_PROGRESS - The
	// proposal is active and open for member voting.
	//
	//     * APPROVED - The proposal
	// was approved with sufficient YES votes among members according to the
	// VotingPolicy specified for the Network. The specified proposal actions are
	// carried out.
	//
	//     * REJECTED - The proposal was rejected with insufficient YES
	// votes among members according to the VotingPolicy specified for the Network. The
	// specified ProposalActions are not carried out.
	//
	//     * EXPIRED - Members did not
	// cast the number of votes required to determine the proposal outcome before the
	// proposal expired. The specified ProposalActions are not carried out.
	//
	//     *
	// ACTION_FAILED - One or more of the specified ProposalActions in a proposal that
	// was approved could not be completed because of an error. The ACTION_FAILED
	// status occurs even if only one ProposalAction fails and other actions are
	// successful.
	Status ProposalStatus

	// The current total of YES votes cast on the proposal by members.
	YesVoteCount *int32
}

// The actions to carry out if a proposal is APPROVED.
type ProposalActions struct {

	// The actions to perform for an APPROVED proposal to invite an AWS account to
	// create a member and join the network.
	Invitations []*InviteAction

	// The actions to perform for an APPROVED proposal to remove a member from the
	// network, which deletes the member and all associated member resources from the
	// network.
	Removals []*RemoveAction
}

// Properties of a proposal.
type ProposalSummary struct {

	// The date and time that the proposal was created.
	CreationDate *time.Time

	// The description of the proposal.
	Description *string

	// The date and time that the proposal expires. This is the CreationDate plus the
	// ProposalDurationInHours that is specified in the ProposalThresholdPolicy. After
	// this date and time, if members have not cast enough votes to determine the
	// outcome according to the voting policy, the proposal is EXPIRED and Actions are
	// not carried out.
	ExpirationDate *time.Time

	// The unique identifier of the proposal.
	ProposalId *string

	// The unique identifier of the member that created the proposal.
	ProposedByMemberId *string

	// The name of the member that created the proposal.
	ProposedByMemberName *string

	// The status of the proposal. Values are as follows:
	//
	//     * IN_PROGRESS - The
	// proposal is active and open for member voting.
	//
	//     * APPROVED - The proposal
	// was approved with sufficient YES votes among members according to the
	// VotingPolicy specified for the Network. The specified proposal actions are
	// carried out.
	//
	//     * REJECTED - The proposal was rejected with insufficient YES
	// votes among members according to the VotingPolicy specified for the Network. The
	// specified ProposalActions are not carried out.
	//
	//     * EXPIRED - Members did not
	// cast the number of votes required to determine the proposal outcome before the
	// proposal expired. The specified ProposalActions are not carried out.
	//
	//     *
	// ACTION_FAILED - One or more of the specified ProposalActions in a proposal that
	// was approved could not be completed because of an error.
	Status ProposalStatus
}

// An action to remove a member from a Managed Blockchain network as the result of
// a removal proposal that is APPROVED. The member and all associated resources are
// deleted from the network.
type RemoveAction struct {

	// The unique identifier of the member to remove.
	//
	// This member is required.
	MemberId *string
}

// Properties of an individual vote that a member cast for a proposal.
type VoteSummary struct {

	// The unique identifier of the member that cast the vote.
	MemberId *string

	// The name of the member that cast the vote.
	MemberName *string

	// The vote value, either YES or NO.
	Vote VoteValue
}

// The voting rules for the network to decide if a proposal is accepted
type VotingPolicy struct {

	// Defines the rules for the network for voting on proposals, such as the
	// percentage of YES votes required for the proposal to be approved and the
	// duration of the proposal. The policy applies to all proposals and is specified
	// when the network is created.
	ApprovalThresholdPolicy *ApprovalThresholdPolicy
}
