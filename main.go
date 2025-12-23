package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Request struct {
	Title string `json:"title"`
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	log.Printf("Searching and streaming for title: %s", req.Title)

	// Construct the yt-dlp command
	// -f bestaudio: Best available audio quality
	// -o -: Output to stdout
	// --quiet: Suppress progress output
	// --no-warnings: Suppress warnings
	// ytsearch1:<query>: Search and pick the first result
	// Remove quiet flags and connect stderr for debugging
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", "-", "ytsearch1:"+req.Title)

	// Connect stderr to the server logs so we can see why yt-dlp fails
	cmd.Stderr = os.Stderr

	// Set headers
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Transfer-Encoding", "chunked")

	// Connect stdout to the response writer
	cmd.Stdout = w

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting yt-dlp: %v", err)
		http.Error(w, "Failed to start audio stream", http.StatusInternalServerError)
		return
	}

	// Wait for the command to finish.
	// usage of 'w' happens in the background as the process writes to the pipe.
	if err := cmd.Wait(); err != nil {
		log.Printf("yt-dlp finished with error (might be just an interruption): %v", err)
		// We can't really write an HTTP error here if headers were already flushed.
	}
}

// Renamed to avoid usage conflict with dj.go
func main() {
	http.HandleFunc("/stream", streamHandler)
	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
