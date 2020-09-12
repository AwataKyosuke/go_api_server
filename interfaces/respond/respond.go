package respond

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type errorResponse struct {
	Message string
	Code    int
}

func Success(w rest.ResponseWriter, ret interface{}) {
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&ret)
}

func InternalServerError(w rest.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.WriteJson(&errorResponse{
		Message: message,
		Code:    http.StatusInternalServerError,
	})
}

func BadRequest(w rest.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.WriteJson(&errorResponse{
		Message: message,
		Code:    http.StatusBadRequest,
	})
}
