package main

import (
    "log"
    "net/http"
    "gateway/handlers"
    "gateway/middleware"
)

func main() {
    http.HandleFunc("/token", handlers.TokenHandler)
    http.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.APIHandler)))

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
