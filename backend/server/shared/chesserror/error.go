package chesserror

import (
	"encoding/json"
	"net/http"
)

type ChessError struct {
	action  string
	key     string
	message string
	cause   string
}

func NewError(key, message string) *ChessError {
	return &ChessError{
		key:     key,
		message: message,
		cause:   "",
	}
}

func (c *ChessError) MarshalJSON() ([]byte, error) {
	type err struct {
		HTTPCode int    `json:"httpCode"`
		Key      string `json:"key"`
		Message  string `json:"message"`
		Cause    string `json:"cause,omitempty"`
	}

	m := make(map[string]interface{})

	if c.action != "" {
		m["action"] = c.action
	}

	m["error"] = &err{
		HTTPCode: c.getHTTPCode(c.key),
		Key:      c.key,
		Message:  c.message,
		Cause:    c.cause,
	}

	return json.Marshal(m)
}

func (c *ChessError) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *ChessError) Error() string {
	return c.String()
}

func (c *ChessError) WithCause(err error) *ChessError {
	c.cause = err.Error()
	return c
}

func (c *ChessError) getHTTPCode(key string) int {
	switch key {
	case ResourceAlreadyExists:
		return http.StatusConflict
	case ResourceNotFound:
		return http.StatusNotFound
	case WrongInputParameter:
		return http.StatusUnprocessableEntity
	case GenericError:
		return http.StatusTeapot
	default:
		return http.StatusInternalServerError
	}
}

func WithAction(err error, action string) *ChessError {
	if e, ok := err.(*ChessError); ok {
		e.action = action
		return e
	}
	return nil
}
