package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the go programm"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("error while opening the server")
	}
}
