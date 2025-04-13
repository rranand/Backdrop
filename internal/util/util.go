package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type JSONResponseWriter struct {
	http.ResponseWriter
}

func (w JSONResponseWriter) SendJSONError(err string, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	jsonStr, _ := json.Marshal(map[string]string{"Error": err})
	stringRep := string(jsonStr)
	http.Error(w, stringRep, httpCode)
}

func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes), nil
}
