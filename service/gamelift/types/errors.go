// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The requested operation would cause a conflict with the current state of a
// service resource associated with the request. Resolve the conflict before
// retrying this request.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

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

// The specified fleet has no available instances to fulfill a CreateGameSession
// request. Clients can retry such requests immediately or after a waiting period.
type FleetCapacityExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FleetCapacityExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FleetCapacityExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FleetCapacityExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FleetCapacityExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *FleetCapacityExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The game instance is currently full and cannot allow the requested player(s) to
// join. Clients can retry such requests immediately or after a waiting period.
type GameSessionFullException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *GameSessionFullException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GameSessionFullException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GameSessionFullException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "GameSessionFullException"
	}
	return *e.ErrorCodeOverride
}
func (e *GameSessionFullException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A game session with this custom ID string already exists in this fleet. Resolve
// this conflict before retrying this request.
type IdempotentParameterMismatchException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IdempotentParameterMismatchException"
	}
	return *e.ErrorCodeOverride
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The service encountered an unrecoverable internal failure while processing the
// request. Clients can retry such requests immediately or after a waiting period.
type InternalServiceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *InternalServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The requested operation would cause a conflict with the current state of a
// resource associated with the request and/or the fleet. Resolve the conflict
// before retrying.
type InvalidFleetStatusException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidFleetStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFleetStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFleetStatusException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidFleetStatusException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidFleetStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation would cause a conflict with the current state of a
// resource associated with the request and/or the game instance. Resolve the
// conflict before retrying.
type InvalidGameSessionStatusException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidGameSessionStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGameSessionStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGameSessionStatusException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidGameSessionStatusException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidGameSessionStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameter values in the request are invalid. Correct the invalid
// parameter values before retrying.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation would cause the resource to exceed the allowed service
// limit. Resolve the issue before retrying.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
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
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// THe requested resources was not found. The resource was either not created yet
// or deleted.
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

// The specified game server group has no available game servers to fulfill a
// ClaimGameServer request. Clients can retry such requests immediately or after a
// waiting period.
type OutOfCapacityException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *OutOfCapacityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OutOfCapacityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OutOfCapacityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OutOfCapacityException"
	}
	return *e.ErrorCodeOverride
}
func (e *OutOfCapacityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested tagging operation did not succeed. This may be due to invalid tag
// format or the maximum tag limit may have been exceeded. Resolve the issue before
// retrying.
type TaggingFailedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TaggingFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TaggingFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TaggingFailedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TaggingFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *TaggingFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is unable to resolve the routing for a particular alias because it
// has a terminal RoutingStrategy associated with it. The message returned in this
// exception is the message defined in the routing strategy itself. Such requests
// should only be retried if the routing strategy for the specified alias is
// modified.
type TerminalRoutingStrategyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TerminalRoutingStrategyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TerminalRoutingStrategyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TerminalRoutingStrategyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TerminalRoutingStrategyException"
	}
	return *e.ErrorCodeOverride
}
func (e *TerminalRoutingStrategyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The client failed authentication. Clients should not retry such requests.
type UnauthorizedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnauthorizedException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation is not supported in the Region specified.
type UnsupportedRegionException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedRegionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedRegionException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
