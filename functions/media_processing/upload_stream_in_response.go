package media_processing

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"time"
)

// Progress update structure

func UploadStreamInResponse(w http.ResponseWriter, filePath string, _ string, chunkSize int, writer *multipart.Writer) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File does not exist", http.StatusNotFound)
		return
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting file info: %v", err), http.StatusInternalServerError)
		return
	}

	// Create a multipart writer

	// Create a form file part
	formFileWriter, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileInfo.Name())},
		"Content-Type":        []string{"image/jpeg"}, // Use the provided content type
	})
	if err != nil {
		http.Error(w, "Error while creating form file", http.StatusBadRequest)
		return
	}

	// Stream the file data in chunks
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break // End of file
		}

		// Write the chunk to the form file part
		if _, err := formFileWriter.Write(buffer[:n]); err != nil {
			fmt.Printf("Error writing to response: %v\n", err)
			return
		}

		// Flush the response to ensure immediate delivery
		w.(http.Flusher).Flush()

		time.Sleep(time.Second * 4)
	}

}
