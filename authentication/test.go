package main

import (
	"net/http"

	mux "github.com/gorilla/mux.git"
)

func main() {
	r := mux.NewRouter()
	// 単純なハンドラ
	r.HandleFunc("/", YourHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillaaaaaaaaadaaaaaaaa!\n"))
}
