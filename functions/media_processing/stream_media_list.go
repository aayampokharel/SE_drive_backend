package media_processing

import (
	"SE_drive_backend/models"
	"mime/multipart"
	"net/http"
)

// # add type here to check for single pic or list of pics(during login .)
func StreamMediaList(w http.ResponseWriter, mediaMapModelRepres models.MediaMap, mediaType string) {

	var mediaList []string
	var chunkSize int

	switch mediaType {
	case "Photo":
		mediaList = mediaMapModelRepres.PhotosList

		chunkSize = 500 * 1024
	case "Video":
		mediaList = mediaMapModelRepres.VideosList
		chunkSize = 2 * 1024 * 1024 //1024 *1024 =1024 kb
	case "Audio":
		mediaList = mediaMapModelRepres.AudiosList
		chunkSize = 256 * 1024
	case "Text":
		mediaList = mediaMapModelRepres.TextsList
		chunkSize = 16 * 1024
	case "Pdf":
		mediaList = mediaMapModelRepres.PdfsList
		chunkSize = 256 * 1024
	default:
		mediaList = mediaMapModelRepres.PhotosList
		chunkSize = 128 * 1024
		// 256KB
	}

	//w.Header().Set("Content-Type", multiPartWriter.FormDataContentType())
	writer := multipart.NewWriter(w)
	defer writer.Close()
	boundary := writer.Boundary()

	// Set response headers
	w.Header().Set("Content-Type", "multipart/form-data; boundary="+boundary)
	w.Header().Set("Transfer-Encoding", "chunked")

	// Flush the headers
	w.(http.Flusher).Flush()

	for _, mediaListValue := range mediaList {
		UploadStreamInResponse(w, mediaListValue, mediaType, chunkSize, writer)
	}
}
