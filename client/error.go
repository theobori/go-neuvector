package client

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Reason     string
}

func NewAPIError(statusCode int, reason string) APIError {
	return APIError{
		StatusCode: statusCode,
		Reason:     reason,
	}
}

func (error *APIError) Error() string {
	return fmt.Sprintf(
		"%d - %s",
		error.StatusCode,
		error.Reason,
	)
}

func (error *APIError) String() string {
	return error.Error()
}
