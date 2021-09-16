package main

import (
	"log"
	"net/http"
	"os"
)

func modelHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("This is a another seperate microservice. Hostname: " + servant))
}

func main() {
	http.HandleFunc("/model2", modelHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
