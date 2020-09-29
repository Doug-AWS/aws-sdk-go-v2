// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// Exception thrown when the customer does not have a valid subscription for the
// product.
type CustomerNotEntitledException struct {
	Message *string
}

func (e *CustomerNotEntitledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomerNotEntitledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomerNotEntitledException) ErrorCode() string             { return "CustomerNotEntitledException" }
func (e *CustomerNotEntitledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *CustomerNotEntitledException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *CustomerNotEntitledException) HasMessage() bool {
	return e.Message != nil
}

// The API is disabled in the Region.
type DisabledApiException struct {
	Message *string
}

func (e *DisabledApiException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DisabledApiException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DisabledApiException) ErrorCode() string             { return "DisabledApiException" }
func (e *DisabledApiException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DisabledApiException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DisabledApiException) HasMessage() bool {
	return e.Message != nil
}

// A metering record has already been emitted by the same EC2 instance, ECS task,
// or EKS pod for the given {usageDimension, timestamp} with a different
// usageQuantity.
type DuplicateRequestException struct {
	Message *string
}

func (e *DuplicateRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateRequestException) ErrorCode() string             { return "DuplicateRequestException" }
func (e *DuplicateRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DuplicateRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DuplicateRequestException) HasMessage() bool {
	return e.Message != nil
}

// The submitted registration token has expired. This can happen if the buyer's
// browser takes too long to redirect to your page, the buyer has resubmitted the
// registration token, or your application has held on to the registration token
// for too long. Your SaaS registration website should redeem this token as soon as
// it is submitted by the buyer's browser.
type ExpiredTokenException struct {
	Message *string
}

func (e *ExpiredTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredTokenException) ErrorCode() string             { return "ExpiredTokenException" }
func (e *ExpiredTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ExpiredTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ExpiredTokenException) HasMessage() bool {
	return e.Message != nil
}

// An internal error has occurred. Retry your request. If the problem persists,
// post a message with details on the AWS forums.
type InternalServiceErrorException struct {
	Message *string
}

func (e *InternalServiceErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceErrorException) ErrorCode() string             { return "InternalServiceErrorException" }
func (e *InternalServiceErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServiceErrorException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServiceErrorException) HasMessage() bool {
	return e.Message != nil
}

// You have metered usage for a CustomerIdentifier that does not exist.
type InvalidCustomerIdentifierException struct {
	Message *string
}

func (e *InvalidCustomerIdentifierException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCustomerIdentifierException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCustomerIdentifierException) ErrorCode() string {
	return "InvalidCustomerIdentifierException"
}
func (e *InvalidCustomerIdentifierException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidCustomerIdentifierException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidCustomerIdentifierException) HasMessage() bool {
	return e.Message != nil
}

// The endpoint being called is in a AWS Region different from your EC2 instance,
// ECS task, or EKS pod. The Region of the Metering Service endpoint and the AWS
// Region of the resource must match.
type InvalidEndpointRegionException struct {
	Message *string
}

func (e *InvalidEndpointRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEndpointRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEndpointRegionException) ErrorCode() string             { return "InvalidEndpointRegionException" }
func (e *InvalidEndpointRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidEndpointRegionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidEndpointRegionException) HasMessage() bool {
	return e.Message != nil
}

// The product code passed does not match the product code used for publishing the
// product.
type InvalidProductCodeException struct {
	Message *string
}

func (e *InvalidProductCodeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidProductCodeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidProductCodeException) ErrorCode() string             { return "InvalidProductCodeException" }
func (e *InvalidProductCodeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidProductCodeException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidProductCodeException) HasMessage() bool {
	return e.Message != nil
}

// Public Key version is invalid.
type InvalidPublicKeyVersionException struct {
	Message *string
}

func (e *InvalidPublicKeyVersionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPublicKeyVersionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPublicKeyVersionException) ErrorCode() string {
	return "InvalidPublicKeyVersionException"
}
func (e *InvalidPublicKeyVersionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidPublicKeyVersionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidPublicKeyVersionException) HasMessage() bool {
	return e.Message != nil
}

// RegisterUsage must be called in the same AWS Region the ECS task was launched
// in. This prevents a container from hardcoding a Region (e.g.
// withRegion(“us-east-1”) when calling RegisterUsage.
type InvalidRegionException struct {
	Message *string
}

func (e *InvalidRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRegionException) ErrorCode() string             { return "InvalidRegionException" }
func (e *InvalidRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRegionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRegionException) HasMessage() bool {
	return e.Message != nil
}

// Registration token is invalid.
type InvalidTokenException struct {
	Message *string
}

func (e *InvalidTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTokenException) ErrorCode() string             { return "InvalidTokenException" }
func (e *InvalidTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidTokenException) HasMessage() bool {
	return e.Message != nil
}

// The usage dimension does not match one of the UsageDimensions associated with
// products.
type InvalidUsageDimensionException struct {
	Message *string
}

func (e *InvalidUsageDimensionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUsageDimensionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUsageDimensionException) ErrorCode() string             { return "InvalidUsageDimensionException" }
func (e *InvalidUsageDimensionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidUsageDimensionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidUsageDimensionException) HasMessage() bool {
	return e.Message != nil
}

// AWS Marketplace does not support metering usage from the underlying platform.
// Currently, Amazon ECS, Amazon EKS, and AWS Fargate are supported.
type PlatformNotSupportedException struct {
	Message *string
}

func (e *PlatformNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PlatformNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PlatformNotSupportedException) ErrorCode() string             { return "PlatformNotSupportedException" }
func (e *PlatformNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PlatformNotSupportedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PlatformNotSupportedException) HasMessage() bool {
	return e.Message != nil
}

// The calls to the API are throttled.
type ThrottlingException struct {
	Message *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ThrottlingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ThrottlingException) HasMessage() bool {
	return e.Message != nil
}

// The timestamp value passed in the meterUsage() is out of allowed range.
type TimestampOutOfBoundsException struct {
	Message *string
}

func (e *TimestampOutOfBoundsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TimestampOutOfBoundsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TimestampOutOfBoundsException) ErrorCode() string             { return "TimestampOutOfBoundsException" }
func (e *TimestampOutOfBoundsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TimestampOutOfBoundsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TimestampOutOfBoundsException) HasMessage() bool {
	return e.Message != nil
}
