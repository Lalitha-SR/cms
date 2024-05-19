package server

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func NewServer() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)
    return mux
}
