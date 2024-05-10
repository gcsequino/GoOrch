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

func getBye(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get /bye request\n")
	io.WriteString(w, "Bye, HTTP!\n")
	os.Exit(1)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/bye", getBye)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
