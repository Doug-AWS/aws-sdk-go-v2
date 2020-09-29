// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// An AWS service limit was exceeded for the calling AWS account.
type AccountLimitExceededException struct {
	Message *string
}

func (e *AccountLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountLimitExceededException) ErrorCode() string             { return "AccountLimitExceededException" }
func (e *AccountLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccountLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccountLimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The input value that was provided is not valid.
type InvalidInputException struct {
	Message *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidInputException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidInputException) HasMessage() bool {
	return e.Message != nil
}

// There was a problem with the underlying OAuth provider.
type OAuthProviderException struct {
	Message *string
}

func (e *OAuthProviderException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OAuthProviderException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OAuthProviderException) ErrorCode() string             { return "OAuthProviderException" }
func (e *OAuthProviderException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *OAuthProviderException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *OAuthProviderException) HasMessage() bool {
	return e.Message != nil
}

// The specified AWS resource cannot be created, because an AWS resource with the
// same settings already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// The specified AWS resource cannot be found.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
