package util

import (
	"net/http"
	"server/http/response"
)


func HandlerReady(w http.ResponseWriter, r *http.Request) {
	response.RespondeWithJSON(w, 200, struct{}{})
}

func HandleErr(w http.ResponseWriter, r *http.Request) {
	response.RespondeWithError(w, 400, "something went wrong")
}
