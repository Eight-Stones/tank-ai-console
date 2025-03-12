package dto

import (
	"encoding/json"
	"net/http"
)

// Error describe custom error.
type Error struct {
	ID      string `json:"id,omitempty"`
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Cause   string `json:"label,omitempty"`
	Trace   string `json:"trace,omitempty"`
	Status  string `json:"status,omitempty"`
}

// Error implements Error() interface.
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// StatusCode return status code for setting response code.
func (e *Error) StatusCode() int {
	return int(e.Code)
}

// NewError generates a custom error.
func NewError(id, label, message string, code int32) error {
	return &Error{
		ID:      id,
		Code:    code,
		Cause:   label,
		Message: message,
		Status:  http.StatusText(int(code)),
	}
}
