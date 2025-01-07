package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFromResponse(w http.ResponseWriter, filePath string, contentType string, chunkSize int) {

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	w.Header().Set("Transfer-Encoding", "chunked")

	// Get file info for file size
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Error getting file info", http.StatusInternalServerError)
		return
	}
	fileSize := fileInfo.Size()
	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))

	// Create a multipart form part
	// part, err := writer.CreateFormFile(contentType, fileInfo.Name())
	if err != nil {
		http.Error(w, "Error creating form file", http.StatusInternalServerError)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		// Write the chunk to the multipart part
		_, writeErr := w.Write(buffer[:n])
		if writeErr != nil {
			http.Error(w, "Error writing file data", http.StatusInternalServerError)
			return
		}

		// Flush the response to send the chunk immediately
		flusher.Flush()

		// For debugging
		fmt.Printf("Chunk of size %d sent\n", n)
	}

	// Close the multipart writer
	if err := writer.Close(); err != nil {
		http.Error(w, "Error closing multipart writer", http.StatusInternalServerError)
		return
	}
}
