// package main

// import (
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func uploadFromResponse(w http.ResponseWriter, filePath string, fieldName string, chunkSize int) {
// 	// Get file info
// 	fileInfo, err := os.Stat(filePath)
// 	if err != nil {
// 		http.Error(w, "Error getting file info", http.StatusInternalServerError)
// 		return
// 	}

// 	fileSize := fileInfo.Size()
// 	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))

// 	// Open the file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		http.Error(w, "Error opening file", http.StatusInternalServerError)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a multipart writer
// 	writer := multipart.NewWriter(w)
// 	defer writer.Close()

// 	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
// 	//flusher, ok := w.(http.Flusher)
// 	// if !ok {
// 	// 	http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// Create the form file part
// 	part, err := writer.CreateFormFile(fieldName, fileInfo.Name())
// 	if err != nil {
// 		http.Error(w, "Error creating form file part", http.StatusInternalServerError)
// 		return
// 	}

// 	// Read and write chunks
// 	buffer := make([]byte, chunkSize)
// 	for {
// 		n, err := file.Read(buffer)
// 		if err != nil && err != io.EOF {
// 			http.Error(w, "Error reading file", http.StatusInternalServerError)
// 			return
// 		}
// 		if n == 0 {
// 			break
// 		}

// 		// Write the chunk to the multipart part
// 		if _, err := part.Write(buffer[:n]); err != nil {
// 			http.Error(w, "Error writing chunk", http.StatusInternalServerError)
// 			return
// 		}

// 		// Flush the response to send the data

// 		time.Sleep(4 * time.Second) // Adjust this value to control the upload speed
// 	}
// }

///////

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
	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))

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

	// Set the headers for the multipart response
	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

	// Verify that the ResponseWriter supports flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Create the form file part
	part, err := writer.CreateFormFile(fieldName, fileInfo.Name())
	if err != nil {
		http.Error(w, "Error creating form file part", http.StatusInternalServerError)
		return
	}

	// Read and write chunks
	buffer := make([]byte, chunkSize)
	for {
		// Read exactly chunkSize bytes
		n, err := io.ReadFull(file, buffer)
		if err != nil {
			// If it's EOF and we read some bytes, write the partial chunk
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				if n > 0 {
					if _, writeErr := part.Write(buffer[:n]); writeErr != nil {
						http.Error(w, "Error writing chunk", http.StatusInternalServerError)
						return
					}
					flusher.Flush()
				}
			}
			break
		}

		// Write the chunk
		if _, err := part.Write(buffer[:n]); err != nil {
			http.Error(w, "Error writing chunk", http.StatusInternalServerError)
			return
		}
		flusher.Flush()
		time.Sleep(4 * time.Second) //just for testing .
	}
}
