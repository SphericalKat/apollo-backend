package utils

import (
	"encoding/json"
	"net/http"
)

// Message Constructs a new map to use as a response
func Message(stat int, msg string) map[string]interface{} {
	return map[string]interface{}{"code": stat, "message": msg}
}

// Respond Convenience function for marshaling and sending a JSON response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	d, _ := json.Marshal(data)
	_, _ = w.Write(d)
}
