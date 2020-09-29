// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The specified action cannot be found.
type ActionNotFoundException struct {
	Message *string
}

func (e *ActionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActionNotFoundException) ErrorCode() string             { return "ActionNotFoundException" }
func (e *ActionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ActionNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ActionNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The specified action type cannot be found.
type ActionTypeNotFoundException struct {
	Message *string
}

func (e *ActionTypeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActionTypeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActionTypeNotFoundException) ErrorCode() string             { return "ActionTypeNotFoundException" }
func (e *ActionTypeNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ActionTypeNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ActionTypeNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The approval action has already been approved or rejected.
type ApprovalAlreadyCompletedException struct {
	Message *string
}

func (e *ApprovalAlreadyCompletedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ApprovalAlreadyCompletedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ApprovalAlreadyCompletedException) ErrorCode() string {
	return "ApprovalAlreadyCompletedException"
}
func (e *ApprovalAlreadyCompletedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ApprovalAlreadyCompletedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ApprovalAlreadyCompletedException) HasMessage() bool {
	return e.Message != nil
}

// Unable to modify the tag due to a simultaneous update request.
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

// The pipeline execution is already in a Stopping state. If you already chose to
// stop and wait, you cannot make that request again. You can choose to stop and
// abandon now, but be aware that this option can lead to failed tasks or out of
// sequence tasks. If you already chose to stop and abandon, you cannot make that
// request again.
type DuplicatedStopRequestException struct {
	Message *string
}

func (e *DuplicatedStopRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicatedStopRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicatedStopRequestException) ErrorCode() string             { return "DuplicatedStopRequestException" }
func (e *DuplicatedStopRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DuplicatedStopRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DuplicatedStopRequestException) HasMessage() bool {
	return e.Message != nil
}

// The action declaration was specified in an invalid format.
type InvalidActionDeclarationException struct {
	Message *string
}

func (e *InvalidActionDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidActionDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidActionDeclarationException) ErrorCode() string {
	return "InvalidActionDeclarationException"
}
func (e *InvalidActionDeclarationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidActionDeclarationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidActionDeclarationException) HasMessage() bool {
	return e.Message != nil
}

// The approval request already received a response or has expired.
type InvalidApprovalTokenException struct {
	Message *string
}

func (e *InvalidApprovalTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidApprovalTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidApprovalTokenException) ErrorCode() string             { return "InvalidApprovalTokenException" }
func (e *InvalidApprovalTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidApprovalTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidApprovalTokenException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource ARN is invalid.
type InvalidArnException struct {
	Message *string
}

func (e *InvalidArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArnException) ErrorCode() string             { return "InvalidArnException" }
func (e *InvalidArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidArnException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidArnException) HasMessage() bool {
	return e.Message != nil
}

// Reserved for future use.
type InvalidBlockerDeclarationException struct {
	Message *string
}

func (e *InvalidBlockerDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidBlockerDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidBlockerDeclarationException) ErrorCode() string {
	return "InvalidBlockerDeclarationException"
}
func (e *InvalidBlockerDeclarationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidBlockerDeclarationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidBlockerDeclarationException) HasMessage() bool {
	return e.Message != nil
}

// The client token was specified in an invalid format
type InvalidClientTokenException struct {
	Message *string
}

func (e *InvalidClientTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClientTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClientTokenException) ErrorCode() string             { return "InvalidClientTokenException" }
func (e *InvalidClientTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidClientTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidClientTokenException) HasMessage() bool {
	return e.Message != nil
}

// The job was specified in an invalid format or cannot be found.
type InvalidJobException struct {
	Message *string
}

func (e *InvalidJobException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobException) ErrorCode() string             { return "InvalidJobException" }
func (e *InvalidJobException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidJobException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidJobException) HasMessage() bool {
	return e.Message != nil
}

// The job state was specified in an invalid format.
type InvalidJobStateException struct {
	Message *string
}

func (e *InvalidJobStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobStateException) ErrorCode() string             { return "InvalidJobStateException" }
func (e *InvalidJobStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidJobStateException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidJobStateException) HasMessage() bool {
	return e.Message != nil
}

// The next token was specified in an invalid format. Make sure that the next token
// you provide is the token returned by a previous call.
type InvalidNextTokenException struct {
	Message *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidNextTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidNextTokenException) HasMessage() bool {
	return e.Message != nil
}

// The nonce was specified in an invalid format.
type InvalidNonceException struct {
	Message *string
}

func (e *InvalidNonceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNonceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNonceException) ErrorCode() string             { return "InvalidNonceException" }
func (e *InvalidNonceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidNonceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidNonceException) HasMessage() bool {
	return e.Message != nil
}

// The stage declaration was specified in an invalid format.
type InvalidStageDeclarationException struct {
	Message *string
}

func (e *InvalidStageDeclarationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStageDeclarationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStageDeclarationException) ErrorCode() string {
	return "InvalidStageDeclarationException"
}
func (e *InvalidStageDeclarationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidStageDeclarationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidStageDeclarationException) HasMessage() bool {
	return e.Message != nil
}

// The structure was specified in an invalid format.
type InvalidStructureException struct {
	Message *string
}

func (e *InvalidStructureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStructureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStructureException) ErrorCode() string             { return "InvalidStructureException" }
func (e *InvalidStructureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidStructureException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidStructureException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource tags are invalid.
type InvalidTagsException struct {
	Message *string
}

func (e *InvalidTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagsException) ErrorCode() string             { return "InvalidTagsException" }
func (e *InvalidTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidTagsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidTagsException) HasMessage() bool {
	return e.Message != nil
}

// The specified authentication type is in an invalid format.
type InvalidWebhookAuthenticationParametersException struct {
	Message *string
}

func (e *InvalidWebhookAuthenticationParametersException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorCode() string {
	return "InvalidWebhookAuthenticationParametersException"
}
func (e *InvalidWebhookAuthenticationParametersException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidWebhookAuthenticationParametersException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidWebhookAuthenticationParametersException) HasMessage() bool {
	return e.Message != nil
}

// The specified event filter rule is in an invalid format.
type InvalidWebhookFilterPatternException struct {
	Message *string
}

func (e *InvalidWebhookFilterPatternException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidWebhookFilterPatternException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidWebhookFilterPatternException) ErrorCode() string {
	return "InvalidWebhookFilterPatternException"
}
func (e *InvalidWebhookFilterPatternException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *InvalidWebhookFilterPatternException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidWebhookFilterPatternException) HasMessage() bool {
	return e.Message != nil
}

// The job was specified in an invalid format or cannot be found.
type JobNotFoundException struct {
	Message *string
}

func (e *JobNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *JobNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *JobNotFoundException) ErrorCode() string             { return "JobNotFoundException" }
func (e *JobNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *JobNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *JobNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The number of pipelines associated with the AWS account has exceeded the limit
// allowed for the account.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The stage has failed in a later run of the pipeline and the pipelineExecutionId
// associated with the request is out of date.
type NotLatestPipelineExecutionException struct {
	Message *string
}

func (e *NotLatestPipelineExecutionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotLatestPipelineExecutionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotLatestPipelineExecutionException) ErrorCode() string {
	return "NotLatestPipelineExecutionException"
}
func (e *NotLatestPipelineExecutionException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *NotLatestPipelineExecutionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotLatestPipelineExecutionException) HasMessage() bool {
	return e.Message != nil
}

// Exceeded the total size limit for all variables in the pipeline.
type OutputVariablesSizeExceededException struct {
	Message *string
}

func (e *OutputVariablesSizeExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OutputVariablesSizeExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OutputVariablesSizeExceededException) ErrorCode() string {
	return "OutputVariablesSizeExceededException"
}
func (e *OutputVariablesSizeExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *OutputVariablesSizeExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *OutputVariablesSizeExceededException) HasMessage() bool {
	return e.Message != nil
}

// The pipeline execution was specified in an invalid format or cannot be found, or
// an execution ID does not belong to the specified pipeline.
type PipelineExecutionNotFoundException struct {
	Message *string
}

func (e *PipelineExecutionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineExecutionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineExecutionNotFoundException) ErrorCode() string {
	return "PipelineExecutionNotFoundException"
}
func (e *PipelineExecutionNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *PipelineExecutionNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PipelineExecutionNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// Unable to stop the pipeline execution. The execution might already be in a
// Stopped state, or it might no longer be in progress.
type PipelineExecutionNotStoppableException struct {
	Message *string
}

func (e *PipelineExecutionNotStoppableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineExecutionNotStoppableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineExecutionNotStoppableException) ErrorCode() string {
	return "PipelineExecutionNotStoppableException"
}
func (e *PipelineExecutionNotStoppableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *PipelineExecutionNotStoppableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PipelineExecutionNotStoppableException) HasMessage() bool {
	return e.Message != nil
}

// The specified pipeline name is already in use.
type PipelineNameInUseException struct {
	Message *string
}

func (e *PipelineNameInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineNameInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineNameInUseException) ErrorCode() string             { return "PipelineNameInUseException" }
func (e *PipelineNameInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PipelineNameInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PipelineNameInUseException) HasMessage() bool {
	return e.Message != nil
}

// The pipeline was specified in an invalid format or cannot be found.
type PipelineNotFoundException struct {
	Message *string
}

func (e *PipelineNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineNotFoundException) ErrorCode() string             { return "PipelineNotFoundException" }
func (e *PipelineNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PipelineNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PipelineNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The pipeline version was specified in an invalid format or cannot be found.
type PipelineVersionNotFoundException struct {
	Message *string
}

func (e *PipelineVersionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PipelineVersionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PipelineVersionNotFoundException) ErrorCode() string {
	return "PipelineVersionNotFoundException"
}
func (e *PipelineVersionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PipelineVersionNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PipelineVersionNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The resource was specified in an invalid format.
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

// The stage was specified in an invalid format or cannot be found.
type StageNotFoundException struct {
	Message *string
}

func (e *StageNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StageNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StageNotFoundException) ErrorCode() string             { return "StageNotFoundException" }
func (e *StageNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *StageNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *StageNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// Unable to retry. The pipeline structure or stage state might have changed while
// actions awaited retry, or the stage contains no failed actions.
type StageNotRetryableException struct {
	Message *string
}

func (e *StageNotRetryableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StageNotRetryableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StageNotRetryableException) ErrorCode() string             { return "StageNotRetryableException" }
func (e *StageNotRetryableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *StageNotRetryableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *StageNotRetryableException) HasMessage() bool {
	return e.Message != nil
}

// The tags limit for a resource has been exceeded.
type TooManyTagsException struct {
	Message *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TooManyTagsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TooManyTagsException) HasMessage() bool {
	return e.Message != nil
}

// The validation was specified in an invalid format.
type ValidationException struct {
	Message *string
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ValidationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ValidationException) HasMessage() bool {
	return e.Message != nil
}

// The specified webhook was entered in an invalid format or cannot be found.
type WebhookNotFoundException struct {
	Message *string
}

func (e *WebhookNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WebhookNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WebhookNotFoundException) ErrorCode() string             { return "WebhookNotFoundException" }
func (e *WebhookNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
