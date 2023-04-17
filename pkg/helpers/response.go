package helpers

import "net/http"

type ResponseDefault struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// APIResponse is for generating template responses
func APIResponse(message string, statusCode int, data interface{}) ResponseDefault {
	status := "Success"
	switch statusCode {
	case http.StatusOK:
		status = "Success"
	case http.StatusCreated:
		status = "Created"
	case http.StatusBadRequest:
		status = "Bad Request"
	case http.StatusUnauthorized:
		status = "Unauthorized"
	case http.StatusNotFound:
		status = "Not Found"
	case http.StatusInternalServerError:
		status = "Internal Server Error"
	default:
		status = "Error"
	}

	return ResponseDefault{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
