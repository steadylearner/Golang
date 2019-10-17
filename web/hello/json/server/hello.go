package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Page struct {
    Payload string `json:"payload"`
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var page Page
        json.NewDecoder(r.Body).Decode(&page)

        fmt.Fprintf(w, "Hello %s", page.Payload)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        peter := Page{
            Payload: "Hello from server",
        }

        json.NewEncoder(w).Encode(peter)
    })

    http.ListenAndServe(":8080", nil)
}