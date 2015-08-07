package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello world from GO!")
}

func main() {
    http.HandlerFunc("/", helloHandler)

    fmt.Println("Started, serving at 8080");
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
