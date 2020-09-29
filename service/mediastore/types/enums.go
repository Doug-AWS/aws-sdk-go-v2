// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContainerLevelMetrics string

// Enum values for ContainerLevelMetrics
const (
	ContainerLevelMetricsEnabled  ContainerLevelMetrics = "ENABLED"
	ContainerLevelMetricsDisabled ContainerLevelMetrics = "DISABLED"
)

type ContainerStatus string

// Enum values for ContainerStatus
const (
	ContainerStatusActive   ContainerStatus = "ACTIVE"
	ContainerStatusCreating ContainerStatus = "CREATING"
	ContainerStatusDeleting ContainerStatus = "DELETING"
)

type MethodName string

// Enum values for MethodName
const (
	MethodNamePut    MethodName = "PUT"
	MethodNameGet    MethodName = "GET"
	MethodNameDelete MethodName = "DELETE"
	MethodNameHead   MethodName = "HEAD"
)
