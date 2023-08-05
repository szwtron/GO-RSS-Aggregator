package main

import "net/http"

func handlerErrors(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Server Error")
}