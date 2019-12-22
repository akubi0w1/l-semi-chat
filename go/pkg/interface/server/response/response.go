package response

import (
	"encoding/json"
	"net/http"
)

// Success status: 200
func Success(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, err.Error())
	}
	w.Write(jsonData)
}

// NoContent
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// BadRequest status: 400
func BadRequest(w http.ResponseWriter, message string) {
	httpError(w, http.StatusBadRequest, message)
}

// MethodNotAllowed
func MethodNotAllowed(w http.ResponseWriter, message string) {
	httpError(w, http.StatusInternalServerError, message)
}

// InternalServerError status: 500
func InternalServerError(w http.ResponseWriter, message string) {
	httpError(w, http.StatusInternalServerError, message)
}

func httpError(w http.ResponseWriter, code int, message string) {
	jsonData, _ := json.Marshal(&errorResponse{
		Code:    code,
		Message: message,
	})
	w.WriteHeader(code)
	w.Write(jsonData)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
