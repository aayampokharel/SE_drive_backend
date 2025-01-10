package main

import (
	//"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func uploadFromResponse(w http.ResponseWriter, filePath string, contentType string, chunkSize int) {
	// Create a pipe to stream data
	pr, pw := io.Pipe()
	ch := make(chan bool)

	// Wrap the pipe's writer with a buffered writer to enable flushing
	//bufWriter := bufio.NewWriter(pw)

	writer := multipart.NewWriter(w)

	// Set response headers
	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	w.Header().Set("Transfer-Encoding", "chunked")

	// Ensure resources are closed properly

	defer pw.Close()
	defer writer.Close()

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Error getting file info", http.StatusInternalServerError)
		return
	}

	fileSize := fileInfo.Size()
	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Launch a goroutine to write the file to the pipe
	go func() {
		// Create the form file part
		part, err := writer.CreateFormFile(contentType, fileInfo.Name())
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		// Read and write chunks
		buffer := make([]byte, chunkSize)
		for {
			n, err := file.Read(buffer)
			if err != nil && err != io.EOF {
				pw.CloseWithError(err)
				return
			}
			if n == 0 {
				break
			}

			// Write to the multipart part
			go func() {
				//@to ensure writing completes before flushing .

				if _, err := part.Write(buffer[:n]); err != nil {
					pw.CloseWithError(err)
					return

				}
				ch <- true
			}()
			<-ch //just for ensuring part.write completes

			// Flush the buffer to ensure immediate sending of data
			// if err := w.Flush(); err != nil {
			// 	pw.CloseWithError(err)
			// 	return
			// }
			time.Sleep(4 * time.Second)
		}

		// Close the writer when done
		writer.Close()
		pw.Close()
	}()

	// Copy data from the pipe to the HTTP response
	if _, err := io.Copy(w, pr); err != nil {
		fmt.Printf("Error copying data to response: %v\n", err)
	}
}
