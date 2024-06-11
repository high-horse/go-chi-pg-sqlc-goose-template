package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct{
	Error string 	`json:"error"`
}

func RespondeWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Reading with 5xx error: ", message)
	}

	RespondeWithJSON(w, code, errResponse{
		Error: message,
	})
	 
}

func RespondeWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshell Json response : %v", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}