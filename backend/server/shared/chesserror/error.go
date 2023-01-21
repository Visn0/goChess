package chesserror

import (
	"encoding/json"
	"net/http"
)

type chessError struct {
	action  string
	key     string
	message string
	cause   string
}

func NewError(key, message string) *chessError {
	return &chessError{
		key:     key,
		message: message,
		cause:   "",
	}
}

func (c *chessError) MarshalJSON() ([]byte, error) {
	type err struct {
		HttpCode int    `json:"httpCode"`
		Key      string `json:"key"`
		Message  string `json:"message"`
		Cause    string `json:"cause,omitempty"`
	}

	m := make(map[string]interface{})

	if c.action != "" {
		m["action"] = c.action
	}

	m["error"] = &err{
		HttpCode: c.getHttpCode(c.key),
		Key:      c.key,
		Message:  c.message,
		Cause:    c.cause,
	}

	return json.Marshal(m)
}

func (c *chessError) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *chessError) Error() string {
	return c.String()
}

func (c *chessError) WithCause(err error) *chessError {
	c.cause = err.Error()
	return c
}

func (c *chessError) getHttpCode(key string) int {
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

func WithAction(err error, action string) *chessError {
	if e, ok := err.(*chessError); ok {
		e.action = action
		return e
	}
	return nil
}
