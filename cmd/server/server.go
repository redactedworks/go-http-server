package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Create a new HTTP server.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a new Response struct.
		response := &Response{
			Message: "Hello, world!",
		}

		// Marshal the Response struct to JSON.
		data, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Write the JSON to the HTTP response writer.
		w.Write(data)
	})

	// Listen on port 8080.
	http.ListenAndServe(":8080", nil)
}
