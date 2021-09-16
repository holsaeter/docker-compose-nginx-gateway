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
	w.Write([]byte("This is model 1. Hostname: " + servant))
}

func funcOneHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("This is what the function 1 handler is returning. Hostname: " + servant))
}

func funcTwoHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("This is another function handler. Hostname:" + servant))
}

func main() {
	http.HandleFunc("/model1", modelHandler)
	http.HandleFunc("/model1/func1", funcOneHandler)
	http.HandleFunc("/model1/func2", funcTwoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
