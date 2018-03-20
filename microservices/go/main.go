package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Version string
	Message string
	Host    string
}

func hello(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Hello, instance: %s - %s", os.Getenv("SERVICE_VERSION"), os.Getenv("HOSTNAME"))
	w.Write([]byte(message))
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	msg := &Message{
		Version: os.Getenv("SERVICE_VERSION"),
		Message: "Hello from the binary",
		Host:    os.Getenv("HOSTNAME"),
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("%s\nListening in 8080...", os.Getenv("SERVICE_VERSION"))

	http.HandleFunc("/version", version)
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
