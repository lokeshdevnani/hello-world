package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.InfoLevel)
}

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    log.Info("Received request for ", name)
    w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
	}
	log.Info("Starting Server...")
	if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
	}
}
