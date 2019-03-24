package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "{\"message\": \"Hello World!!\"}")
}

func main() {
	http.HandleFunc("/", sendHelloWorld)
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
