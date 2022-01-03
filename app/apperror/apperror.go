package apperror

import (
	"errors"
	"fmt"
)

// TODO: stack trace
type AppError struct {
	err error // expected to set original error

	params map[string]interface{} // Expected userInput params (Not to be translated)
	fields map[string]interface{} // Expected app managed values (To be translated)

	msgLog string // a logging message (not for a user message)
}

func New(err error) *AppError {
	e := &AppError{
		fields: map[string]interface{}{},
		params: map[string]interface{}{},
	}

	if errors.As(err, &e) {
		return e
	}
	e.err = err
	return e
}

// WithParams - params
func (e *AppError) WithParams(params map[string]interface{}) *AppError {
	for k, v := range params {
		e.params[k] = v
	}
	return e
}

// WithFields - fields
func (e *AppError) WithFields(fields map[string]interface{}) *AppError {
	for k, v := range fields {
		e.fields[k] = v
	}
	return e
}

// WithMsgLog - msgLog
func (e *AppError) WithMsgLog(msg string) *AppError {
	if e.msgLog == "" {
		e.msgLog = msg
		return e
	}
	e.msgLog = fmt.Sprintf("%s\n%s", e.msgLog, msg)
	return e
}

// Error - error interface
func (e *AppError) Error() string {
	return e.msgLog
}

// Is - errors.Is
func (e *AppError) Is(err error) bool {
	return errors.Is(e.err, err)
}

// Unwrap - errors.Unwrap
func (e *AppError) Unwrap() error {
	return e.err
}
