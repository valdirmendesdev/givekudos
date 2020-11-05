package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthStatus(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"message":"ok"}`))
}

func users(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"message":"users"}`))
}

func main() {
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(healthStatus))
	server.Handle("/users", http.HandlerFunc(users))

	fmt.Println("Server is running...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal("Server error")
	}
}
