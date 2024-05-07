package main

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, "Hello nerds ;)\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("get /hello request\n")
    io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)

    err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error startsing the server: %s\n", err)
        os.Exit(1)
    }
}
