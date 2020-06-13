package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello web3! THis is n3 or n2")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check. n3 or n2")
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe error: ", err.Error())
		return
	}
}
