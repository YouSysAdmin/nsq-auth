package helpers

import (
	"encoding/json"
	"net/http"
)

// HttpError make HTTP error response
func HttpError(w http.ResponseWriter, message string, status int) {
	resp := struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}{
		Status: status,
		Error:  message,
	}
	r, _ := json.Marshal(resp)
	w.WriteHeader(status)
	w.Write(r)
	return
}
