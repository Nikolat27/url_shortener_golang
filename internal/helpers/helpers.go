package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, msg string, statusCode int) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(msg + "\n"))
	return err
}

func DeSerializeJSON(body io.Reader, maxByte int64, obj any) error {
	body = io.LimitReader(body, maxByte)
	err := json.NewDecoder(body).Decode(&obj)
	return err
}
