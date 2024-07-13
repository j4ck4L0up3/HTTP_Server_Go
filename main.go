package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// HTTP Request Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, readErr := os.ReadFile("index.html")

		if readErr != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// set header
		w.Header().Set("Content-Type", "text/html")

		// write html as response
		_, writeErr := w.Write(html)

		if writeErr != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Request listener
	fmt.Println("Start listening on port 5500...")
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
