package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, data any, err error) {
	rsp := Response{Data: data}

	if err != nil {
		rsp.ErrorMessage = err.Error()
	} else {
		rsp.ErrorMessage = ""
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		panic(err)
	}
}
