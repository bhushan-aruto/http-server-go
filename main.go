package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

func main() {
	mux := http.NewServeMux() //http request router

	mux.HandleFunc("/hello", handleHelo)
	mux.HandleFunc("/post", postHandler)

	fmt.Println("the server is running on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("http server facing isssue while starting...")
	}
}

func handleHelo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "hellooo")

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method isnt allowed", http.StatusMethodNotAllowed)
		return
	}

	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(request)

}
