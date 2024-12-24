package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rishad004/bw01/apiGateway/pkg/domain"
)

func SendJSONResponse(w http.ResponseWriter, message any, statusCode int, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := domain.Response{
		Message:    message,
		StatusCode: statusCode,
	}
	json.NewEncoder(w).Encode(response)
}
