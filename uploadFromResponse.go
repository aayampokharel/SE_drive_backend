package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func uploadFromResponse(w http.ResponseWriter, filePath string, fieldName string, chunkSize int) {
	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Error getting file info", http.StatusInternalServerError)
		return
	}

	fileSize := fileInfo.Size()
	fmt.Println("File size:", fileSize)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a multipart writer
	writer := multipart.NewWriter(w)
	defer writer.Close()
	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))
	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	w.WriteHeader(http.StatusOK) //to flush headers .
	fmt.Print("\n", w.Header(), "\n")
	part, err := writer.CreateFormFile(fieldName, fileInfo.Name())
	if err != nil {
		http.Error(w, "Error creating form file part", http.StatusInternalServerError)
		return
	}

	// Read and write chunks
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}

		// Write the chunk to the multipart part
		if _, err := part.Write(buffer[:n]); err != nil {
			http.Error(w, "Error writing chunk", http.StatusInternalServerError)
			return
		}

		// Flush the response to send the data

		time.Sleep(4 * time.Second) // Adjust this value to control the upload speed
	}
}
