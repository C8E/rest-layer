package rest

import (
	"github.com/rs/rest-layer/resource"
	"golang.org/x/net/context"
)

var (
	// ErrNotFound represents a 404 HTTP error.
	ErrNotFound = &Error{404, "Not Found", nil}
	// ErrPreconditionFailed happends when a conditional request condition is not met.
	ErrPreconditionFailed = &Error{412, "Precondition Failed", nil}
	// ErrConflict happens when another thread or node modified the data concurrently
	// with our own thread in such a way we can't securely apply the requested changes.
	ErrConflict = &Error{409, "Conflict", nil}
	// ErrInvalidMethod happends when the used HTTP method is not supported for this
	// resource.
	ErrInvalidMethod = &Error{405, "Invalid Method", nil}
	// ErrClientClosedRequest is returned when the client closed the connection before
	// the server was able to finish processing the request.
	ErrClientClosedRequest = &Error{499, "Client Closed Request", nil}
	// ErrNotImplemented happends when a requested feature is not implemented.
	ErrNotImplemented = &Error{501, "Not Implemented", nil}
	// ErrGatewayTimeout is returned when the specified timeout for the request has been
	// reached before the server was able to process it.
	ErrGatewayTimeout = &Error{504, "Deadline Exceeded", nil}
	// ErrUnknown is thrown when the origine of the error can't be identified.
	ErrUnknown = &Error{520, "Unknown Error", nil}
)

// Error defines a REST error with optional per fields error details
type Error struct {
	// Code defines the error code to be used for the error and for the HTTP status
	Code int
	// Message is the error message
	Message string
	// Issues holds per fields errors if any
	Issues map[string][]interface{}
}

// NewError returns a rest.Error from an standard error.
//
// If the the inputed error is recognized, the appropriate rest.Error is mapped.
func NewError(err error) *Error {
	switch err {
	case context.Canceled:
		return ErrClientClosedRequest
	case context.DeadlineExceeded:
		return ErrGatewayTimeout
	case resource.ErrNotFound:
		return ErrNotFound
	case resource.ErrConflict:
		return ErrConflict
	case resource.ErrNotImplemented:
		return ErrNotImplemented
	case resource.ErrNoStorage:
		return &Error{501, err.Error(), nil}
	case nil:
		return nil
	default:
		return &Error{520, err.Error(), nil}
	}
}

// Error returns the error as string
func (e *Error) Error() string {
	return e.Message
}
