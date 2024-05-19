package main

import (
    "log"
    "net/http"
    "github.com/Lalitha-SR/cms/internal/server"
)

func main() {
    srv := server.NewServer()
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", srv)
    if err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}
