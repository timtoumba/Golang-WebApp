package main

import (
    "fmt"
    "net/http"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s !", r.URL.Path[1:]) // send data to client side
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World") // send data to client side
}

func main() {
    http.HandleFunc("/", sayhelloName) // set router

    http.HandleFunc("/hello", hello)

    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
