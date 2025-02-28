// package media_processing

// import (
// 	//"bufio"
// 	errors "SE_drive_backend/Errors"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// )

// func UploadStreamInResponse(w http.ResponseWriter, filePath string, contentType string, chunkSize int) {
// 	// Create a pipe to stream data
// 	pr, pw := io.Pipe()
// 	ch := make(chan bool)

// 	// Wrap the pipe's writer with a buffered writer to enable flushing
// 	//bufWriter := bufio.NewWriter(pw)

// 	writer := multipart.NewWriter(w)

// 	// Set response headers
// 	w.Header().Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
// 	w.Header().Set("Transfer-Encoding", "chunked")

// 	// Ensure resources are closed properly

// 	defer pw.Close()
// 	defer writer.Close()

// 	// Get file info
// 	fileInfo, err := os.Stat(filePath)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while getting fileinfo during streaming of file data %s", err)))
// 		return
// 	}

// 	fileSize := fileInfo.Size()
// 	w.Header().Set("File-Size", fmt.Sprintf("%d", fileSize))

// 	// Open the file for reading
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while opening file during streaming of file data %s", err)))
// 		return
// 	}
// 	defer file.Close()

// 	// Launch a goroutine to write the file to the pipe
// 	go func() {
// 		// Create the form file part
// 		part, err := writer.CreateFormFile(contentType, fileInfo.Name())
// 		if err != nil {
// 			pw.CloseWithError(err)
// 			return
// 		}

// 		// Read and write chunks
// 		buffer := make([]byte, chunkSize)
// 		for {
// 			n, err := file.Read(buffer)
// 			if err != nil && err != io.EOF {
// 				pw.CloseWithError(err)
// 				return
// 			}
// 			if n == 0 {
// 				break
// 			}

// 			// Write to the multipart part
// 			go func() {
// 				//@to ensure writing completes before flushing .

// 				if _, err := part.Write(buffer[:n]); err != nil {
// 					pw.CloseWithError(err)
// 					return

// 				}
// 				ch <- true
// 			}()
// 			<-ch //just for ensuring part.write completes

// 			// time.Sleep(4 * time.Second)
// 		}

// 		// Close the writer when done
// 		writer.Close()
// 		pw.Close()
// 	}()

// 	// Copy data from the pipe to the HTTP response
// 	if _, err := io.Copy(w, pr); err != nil {
// 		fmt.Printf("Error copying data to response: %v\n", err)
// 	}
// }

package media_processing

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadStreamInResponse(w http.ResponseWriter, filePath string, _ string, chunkSize int) {
	chunkSize = 16 * 1024
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File does not exist", http.StatusNotFound)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("error opening file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting file info: %v", err), http.StatusInternalServerError)
		return
	}

	fileSize := fileInfo.Size()
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))

	// Streaming with delay
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, fmt.Sprintf("error reading file: %v", err), http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}

		// Write chunk to response
		if _, err := w.Write(buffer[:n]); err != nil {
			fmt.Printf("Error writing to response: %v\n", err)
			return
		}

		// Flush the response to ensure immediate delivery
		w.(http.Flusher).Flush()

		// Introduce delay to slow down streaming
		time.Sleep(time.Second * 4)
	}
}
