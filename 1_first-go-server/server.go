package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> Go to /api/v0/[endpoint_goes_here]")
}

func main() {
	http.HandleFunc("/", rootHandle)
	fmt.Println("- Server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}