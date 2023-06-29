package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to plain text
		w.Header().Set("Content-Type", "text/plain")
		// Write the response body
		fmt.Fprint(w, "You created a waypoint golang template!")
	}

	// Register the handler function to the default ServeMux
	http.HandleFunc("/", handler)

	// Start the server and listen on port 3000
	log.Println("Server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

