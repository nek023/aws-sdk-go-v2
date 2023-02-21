// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The input parameters don't match the service's restrictions.
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be processed because of conflict in the current state of
// the resource.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
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
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The client is permanently forbidden from making the request.
type ForbiddenException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ForbiddenException"
	}
	return *e.ErrorCodeOverride
}
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request exceeds the resource limit.
type ResourceLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service encountered an unexpected error.
type ServiceFailureException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *ServiceFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceFailureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The service is currently unavailable.
type ServiceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The client exceeded its request rate limit.
type ThrottledClientException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *ThrottledClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottledClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottledClientException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottledClientException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottledClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The client is not currently authorized to make the request.
type UnauthorizedClientException struct {
	Message *string

	ErrorCodeOverride *string

	Code ErrorCode

	noSmithyDocumentSerde
}

func (e *UnauthorizedClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedClientException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnauthorizedClientException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnauthorizedClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
