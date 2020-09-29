// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The request is missing required parameters or has invalid parameters.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}

// You have reached the maximum number of sampling rules.
type RuleLimitExceededException struct {
	Message *string
}

func (e *RuleLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleLimitExceededException) ErrorCode() string             { return "RuleLimitExceededException" }
func (e *RuleLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *RuleLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *RuleLimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The request exceeds the maximum number of requests per second.
type ThrottledException struct {
	Message *string
}

func (e *ThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottledException) ErrorCode() string             { return "ThrottledException" }
func (e *ThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ThrottledException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ThrottledException) HasMessage() bool {
	return e.Message != nil
}
