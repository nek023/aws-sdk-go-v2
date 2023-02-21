// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You are not authorized to use this operation with the given parameters.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've exceeded the notification or subscriber limit.
type CreationLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CreationLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreationLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreationLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CreationLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *CreationLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The budget name already exists. Budget names must be unique within an account.
type DuplicateRecordException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicateRecordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateRecordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateRecordException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateRecordException"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateRecordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pagination token expired.
type ExpiredNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ExpiredNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ExpiredNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ExpiredNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error on the server occurred during the processing of your request. Try again
// later.
type InternalErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The pagination token is invalid.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error on the client occurred. Typically, the cause is an invalid input value.
type InvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t locate the resource that you specified.
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was received and recognized by the server, but the server rejected
// that particular method for the requested resource.
type ResourceLockedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceLockedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLockedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLockedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceLockedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceLockedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of API requests has exceeded the maximum allowed API request
// throttling limit for the account.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
