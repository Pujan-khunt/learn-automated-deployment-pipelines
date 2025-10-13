package main

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("you just reached the main route."))
}

func main() {
	http.HandleFunc("/", RootHandler)
	fmt.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
