package main

import (
    "fmt"
    "log"
    "net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Steve\n")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Healthy\n")
}

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/health", handleHealth)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
