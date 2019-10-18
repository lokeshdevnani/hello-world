package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
	}
	srv.ListenAndServe();
}
