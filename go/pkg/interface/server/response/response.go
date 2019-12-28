package response

import (
	"encoding/json"
	"l-semi-chat/pkg/domain"
	"net/http"
)

// Success status: 200
func Success(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		HttpError(w, domain.InternalServerError(err))
	}
	w.Write(jsonData)
}

// NoContent
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// HttpError
func HttpError(w http.ResponseWriter, err error) {
	e, ok := err.(domain.Error)
	if !ok {
		e = domain.InternalServerError(err)
	}
	jsonData, _ := json.Marshal(&errorResponse{
		Code:    e.GetStatusCode(),
		Message: e.Error(),
	})
	w.WriteHeader(e.GetStatusCode())
	w.Write(jsonData)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
