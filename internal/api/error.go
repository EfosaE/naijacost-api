package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Error      string `json:"error"`
	Message    string `json:"message,omitempty"`
}

// Render sets the proper status code before rendering
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// New creates a new ErrorResponse
func New(statusCode int, errorMsg string, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Error:      errorMsg,
		Message:    message,
	}
}

// BadRequest returns a 400 Bad Request error
func BadRequest(err error, message string) *ErrorResponse {
	errMsg := "Bad request"
	if err != nil {
		message = err.Error()
	}
	return New(http.StatusBadRequest, errMsg, message)
}

// NotFound returns a 404 Not Found error
func NotFound(message string) *ErrorResponse {
	return New(http.StatusNotFound, "Resource not found", message)
}

// InternalServerError returns a 500 Internal Server Error
func InternalServerError(err error, message string) *ErrorResponse {
	errMsg := "Internal server error"
	if err != nil && message == "" {
		message = err.Error()
	}
	return New(http.StatusInternalServerError, errMsg, message)
}

// Unauthorized returns a 401 Unauthorized error
func Unauthorized(message string) *ErrorResponse {
	return New(http.StatusUnauthorized, "Unauthorized", message)
}

// Forbidden returns a 403 Forbidden error
func Forbidden(message string) *ErrorResponse {
	return New(http.StatusForbidden, "Forbidden", message)
}

// ValidationError returns a 422 Unprocessable Entity error
func ValidationError(message string) *ErrorResponse {
	return New(http.StatusUnprocessableEntity, "Validation error", message)
}

// // SendError is a convenience function to send an error response
// func SendError(w http.ResponseWriter, r *http.Request, err *ErrorResponse) {
// 	render.Render(w, r, err)
// }


// NotFoundHandler returns a custom 404 handler that responds with JSON
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorResponse := &ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Not Found",
			Message:    fmt.Sprintf("A %s request doesn't exist on URL: '%s' on this server", r.Method, r.URL.Path),
		}
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, errorResponse)
	}
}

// NotAllowedHandler returns a custom 405 handler that responds with JSON
func MethodNotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorResponse := &ErrorResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Error:      "Method Not Allowed",
			Message:    fmt.Sprintf("A %s request is not allowed on URL: '%s' on this server", r.Method, r.URL.Path),
		}
		render.Status(r, http.StatusMethodNotAllowed)
		render.JSON(w, r, errorResponse)
	}
}