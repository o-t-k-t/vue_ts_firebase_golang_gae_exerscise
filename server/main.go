package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://vue-go-exercise.firebaseapp.com/")

	w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private!\n"))
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", private)

	http.Handle("/", r)
}
