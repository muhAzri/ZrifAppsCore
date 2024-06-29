package response

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func BuildResponse(httpCode int, message string, payload interface{}, errorPayload interface{}, w http.ResponseWriter) {
	var response interface{}
	status := "success"

	if httpCode >= http.StatusBadRequest {
		status = "error"
		response = ResponseError{
			Meta: Meta{
				Message: message,
				Code:    httpCode,
				Status:  status,
			},
			Error: errorPayload,
		}
	} else {
		response = Response{
			Meta: Meta{
				Message: message,
				Code:    httpCode,
				Status:  status,
			},
			Data: payload,
		}
	}

	respondWithJSON(httpCode, response, w)
}
