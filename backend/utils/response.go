package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Printf("could not marshal response")
		return
	}
	w.Write(jsonResponse)
}

func JsonResponse(w http.ResponseWriter, res ...any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
