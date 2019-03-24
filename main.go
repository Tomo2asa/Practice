package main

import (
	"fmt"
	"log"
	"net/http"
)

func sendHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "{\"message\": \"Hello World!!\"}")
}

func main() {
	http.HandleFunc("/", sendHelloWorld)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
