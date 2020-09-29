// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The target of the operation is currently being modified by a different request.
// Try again later.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConcurrentModificationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConcurrentModificationException) HasMessage() bool {
	return e.Message != nil
}

// The request was denied because performing this operation violates a constraint.
// Some of the reasons in the following list might not apply to this specific
// operation.
//
//     * You must meet the prerequisites for using tag policies. For
// information, see Prerequisites and Permissions for Using Tag Policies
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies-prereqs.html)
// in the AWS Organizations User Guide.
//
//     * You must enable the tag policies
// service principal (tagpolicies.tag.amazonaws.com) to integrate with AWS
// Organizations For information, see EnableAWSServiceAccess
// (http://docs.aws.amazon.com/organizations/latest/APIReference/API_EnableAWSServiceAccess.html).
//
//
// * You must have a tag policy attached to the organization root, an OU, or an
// account.
type ConstraintViolationException struct {
	Message *string
}

func (e *ConstraintViolationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConstraintViolationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConstraintViolationException) ErrorCode() string             { return "ConstraintViolationException" }
func (e *ConstraintViolationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConstraintViolationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConstraintViolationException) HasMessage() bool {
	return e.Message != nil
}

// The request processing failed because of an unknown error, exception, or
// failure. You can retry the request.
type InternalServiceException struct {
	Message *string
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceException) HasMessage() bool {
	return e.Message != nil
}

// This error indicates one of the following:
//
//     * A parameter is missing.
//
//     *
// A malformed string was supplied for the request parameter.
//
//     * An
// out-of-range value was supplied for the request parameter.
//
//     * The target ID
// is invalid, unsupported, or doesn't exist.
//
//     * You can't access the Amazon S3
// bucket for report storage. For more information, see Additional Requirements for
// Organization-wide Tag Compliance Reports
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies-prereqs.html#bucket-policies-org-report)
// in the AWS Organizations User Guide.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}

// A PaginationToken is valid for a maximum of 15 minutes. Your request was denied
// because the specified PaginationToken has expired.
type PaginationTokenExpiredException struct {
	Message *string
}

func (e *PaginationTokenExpiredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PaginationTokenExpiredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PaginationTokenExpiredException) ErrorCode() string {
	return "PaginationTokenExpiredException"
}
func (e *PaginationTokenExpiredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PaginationTokenExpiredException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PaginationTokenExpiredException) HasMessage() bool {
	return e.Message != nil
}

// The request was denied to limit the frequency of submitted requests.
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
