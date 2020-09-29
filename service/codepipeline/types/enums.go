// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionCategory string

// Enum values for ActionCategory
const (
	ActionCategorySource   ActionCategory = "Source"
	ActionCategoryBuild    ActionCategory = "Build"
	ActionCategoryDeploy   ActionCategory = "Deploy"
	ActionCategoryTest     ActionCategory = "Test"
	ActionCategoryInvoke   ActionCategory = "Invoke"
	ActionCategoryApproval ActionCategory = "Approval"
)

type ActionConfigurationPropertyType string

// Enum values for ActionConfigurationPropertyType
const (
	ActionConfigurationPropertyTypeString  ActionConfigurationPropertyType = "String"
	ActionConfigurationPropertyTypeNumber  ActionConfigurationPropertyType = "Number"
	ActionConfigurationPropertyTypeBoolean ActionConfigurationPropertyType = "Boolean"
)

type ActionExecutionStatus string

// Enum values for ActionExecutionStatus
const (
	ActionExecutionStatusInprogress ActionExecutionStatus = "InProgress"
	ActionExecutionStatusAbandoned  ActionExecutionStatus = "Abandoned"
	ActionExecutionStatusSucceeded  ActionExecutionStatus = "Succeeded"
	ActionExecutionStatusFailed     ActionExecutionStatus = "Failed"
)

type ActionOwner string

// Enum values for ActionOwner
const (
	ActionOwnerAws        ActionOwner = "AWS"
	ActionOwnerThirdparty ActionOwner = "ThirdParty"
	ActionOwnerCustom     ActionOwner = "Custom"
)

type ApprovalStatus string

// Enum values for ApprovalStatus
const (
	ApprovalStatusApproved ApprovalStatus = "Approved"
	ApprovalStatusRejected ApprovalStatus = "Rejected"
)

type ArtifactLocationType string

// Enum values for ArtifactLocationType
const (
	ArtifactLocationTypeS3 ArtifactLocationType = "S3"
)

type ArtifactStoreType string

// Enum values for ArtifactStoreType
const (
	ArtifactStoreTypeS3 ArtifactStoreType = "S3"
)

type BlockerType string

// Enum values for BlockerType
const (
	BlockerTypeSchedule BlockerType = "Schedule"
)

type EncryptionKeyType string

// Enum values for EncryptionKeyType
const (
	EncryptionKeyTypeKms EncryptionKeyType = "KMS"
)

type FailureType string

// Enum values for FailureType
const (
	FailureTypeJobfailed           FailureType = "JobFailed"
	FailureTypeConfigurationerror  FailureType = "ConfigurationError"
	FailureTypePermissionerror     FailureType = "PermissionError"
	FailureTypeRevisionoutofsync   FailureType = "RevisionOutOfSync"
	FailureTypeRevisionunavailable FailureType = "RevisionUnavailable"
	FailureTypeSystemunavailable   FailureType = "SystemUnavailable"
)

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated    JobStatus = "Created"
	JobStatusQueued     JobStatus = "Queued"
	JobStatusDispatched JobStatus = "Dispatched"
	JobStatusInprogress JobStatus = "InProgress"
	JobStatusTimedout   JobStatus = "TimedOut"
	JobStatusSucceeded  JobStatus = "Succeeded"
	JobStatusFailed     JobStatus = "Failed"
)

type PipelineExecutionStatus string

// Enum values for PipelineExecutionStatus
const (
	PipelineExecutionStatusInprogress PipelineExecutionStatus = "InProgress"
	PipelineExecutionStatusStopped    PipelineExecutionStatus = "Stopped"
	PipelineExecutionStatusStopping   PipelineExecutionStatus = "Stopping"
	PipelineExecutionStatusSucceeded  PipelineExecutionStatus = "Succeeded"
	PipelineExecutionStatusSuperseded PipelineExecutionStatus = "Superseded"
	PipelineExecutionStatusFailed     PipelineExecutionStatus = "Failed"
)

type StageExecutionStatus string

// Enum values for StageExecutionStatus
const (
	StageExecutionStatusInprogress StageExecutionStatus = "InProgress"
	StageExecutionStatusFailed     StageExecutionStatus = "Failed"
	StageExecutionStatusStopped    StageExecutionStatus = "Stopped"
	StageExecutionStatusStopping   StageExecutionStatus = "Stopping"
	StageExecutionStatusSucceeded  StageExecutionStatus = "Succeeded"
)

type StageRetryMode string

// Enum values for StageRetryMode
const (
	StageRetryModeFailed_actions StageRetryMode = "FAILED_ACTIONS"
)

type StageTransitionType string

// Enum values for StageTransitionType
const (
	StageTransitionTypeInbound  StageTransitionType = "Inbound"
	StageTransitionTypeOutbound StageTransitionType = "Outbound"
)

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeCreatepipeline         TriggerType = "CreatePipeline"
	TriggerTypeStartpipelineexecution TriggerType = "StartPipelineExecution"
	TriggerTypePollforsourcechanges   TriggerType = "PollForSourceChanges"
	TriggerTypeWebhook                TriggerType = "Webhook"
	TriggerTypeCloudwatchevent        TriggerType = "CloudWatchEvent"
	TriggerTypePutactionrevision      TriggerType = "PutActionRevision"
)

type WebhookAuthenticationType string

// Enum values for WebhookAuthenticationType
const (
	WebhookAuthenticationTypeGithub_hmac     WebhookAuthenticationType = "GITHUB_HMAC"
	WebhookAuthenticationTypeIp              WebhookAuthenticationType = "IP"
	WebhookAuthenticationTypeUnauthenticated WebhookAuthenticationType = "UNAUTHENTICATED"
)
