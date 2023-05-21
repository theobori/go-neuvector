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
	return fmt.Sprintf("Invalid HTTP status code: %d - ", error.StatusCode) +
		fmt.Sprintf("Reason: %s", error.Reason)
}

func (error *APIError) String() string {
	return error.Error()
}
