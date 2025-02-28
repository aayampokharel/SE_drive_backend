package media_processing

import (
	"SE_drive_backend/models"
	"mime/multipart"
	"net/http"
)

// # add type here to check for single pic or list of pics(during login .)
func StreamMediaList(w http.ResponseWriter, mediaMapModelRepres models.MediaMap, mediaType string) {
	//var streamingMediaError models.ErrorsModel
	var mediaList []string
	var chunkSize int
	multiPartWriter := multipart.NewWriter(w)
	defer multiPartWriter.Close()
	switch mediaType {
	case "Photo":
		mediaList = mediaMapModelRepres.PhotosList

	case "Video":
		mediaList = mediaMapModelRepres.VideosList
		chunkSize = 2 * 1024 * 1024
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

	for _, mediaListValue := range mediaList {
		UploadStreamInResponse(w, mediaListValue, mediaType, chunkSize)
	}
}
