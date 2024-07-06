package server

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	config := NewConfig()
	c := cors.New(config.CORS)

	handler := c.Handler(mux)

	fmt.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
