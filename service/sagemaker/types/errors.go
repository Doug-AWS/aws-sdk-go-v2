// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// There was a conflict when you attempted to modify an experiment, trial, or trial
// component.
type ConflictException struct {
	Message *string
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConflictException) HasMessage() bool {
	return e.Message != nil
}

// Resource being accessed is in use.
type ResourceInUse struct {
	Message *string
}

func (e *ResourceInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUse) ErrorCode() string             { return "ResourceInUse" }
func (e *ResourceInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceInUse) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceInUse) HasMessage() bool {
	return e.Message != nil
}

// You have exceeded an Amazon SageMaker resource limit. For example, you might
// have too many training jobs created.
type ResourceLimitExceeded struct {
	Message *string
}

func (e *ResourceLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceeded) ErrorCode() string             { return "ResourceLimitExceeded" }
func (e *ResourceLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceLimitExceeded) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceLimitExceeded) HasMessage() bool {
	return e.Message != nil
}

// Resource being access is not found.
type ResourceNotFound struct {
	Message *string
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string             { return "ResourceNotFound" }
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFound) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFound) HasMessage() bool {
	return e.Message != nil
}
