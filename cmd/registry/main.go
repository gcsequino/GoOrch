package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
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

func postName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	fmt.Printf("Received name: %s\n", name)
	io.WriteString(w, "Hello, "+name+"!")
}

func longRunningTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Starting long running task...\n")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Printf("Long running task completed!\n")
		io.WriteString(w, "Long running task completed!\n")
	}()
	wg.Wait()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/bye", getBye)
	mux.HandleFunc("/postName", postName)
	mux.HandleFunc("/longRunningTask", longRunningTask)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		os.Exit(1)
	}
}
