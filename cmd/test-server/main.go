package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverPort = "8001"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working\n")
}
func main() {
	http.HandleFunc("/", index)
	fmt.Println("Running test-server on port: ", serverPort)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", serverPort), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error: ", err)
	}
}
