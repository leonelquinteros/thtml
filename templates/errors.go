package templates

import (
	"runtime"
	"time"
)

// IError is the common error type interface.
type IError interface {
	// errors.Error implementation
	Error() string

	// Stringer implementation
	String() string

	// Dump stack trace
	Stack() string

	// Creation time
	Created() time.Time
}

// TError templates error base class. Implements IError
type TError struct {
	msg     string
	stack   string
	created time.Time
}

// error implementation
func (err TError) Error() string {
	return err.String()
}

// Stringer implementation
func (err TError) String() string {
	return err.msg
}

// Stack getter
func (err TError) Stack() string {
	return err.stack
}

// Created time getter
func (err TError) Created() time.Time {
	return err.created
}

// NewError returns a new TError object
func NewError(message string) TError {
	// Get stack trace
	stack := make([]byte, 1<<16)
	runtime.Stack(stack, false)

	return TError{
		msg:     message,
		stack:   string(stack),
		created: time.Now(),
	}
}

// EmptyTemplateError is returned when templates.Service hasn't been loaded.
type EmptyTemplateError struct {
	// Error composition
	TError
}

// NewEmptyTemplateError returns a new EmptyTemplateError object
func NewEmptyTemplateError() EmptyTemplateError {
	return EmptyTemplateError{
		TError: NewError("Empty template. Call templates.Service.Load(): https://godoc.org/github.com/leonelquinteros/thtml/templates#Service.Load"),
	}
}
