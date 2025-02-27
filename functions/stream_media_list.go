package functions

import (
	"SE_drive_backend/models"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func StreamMediaList(w http.ResponseWriter, mediaMapModelRepres models.MediaMap, fieldName string) {
	var streamingMediaError models.ErrorsModel
	var mediaList []string
	multiPartWriter := multipart.NewWriter(w)
	defer multiPartWriter.Close()
	switch fieldName {
	case "Photo":
		mediaList = mediaMapModelRepres.PhotosList

	case "Video":
		mediaList = mediaMapModelRepres.VideosList

	case "Audio":
		mediaList = mediaMapModelRepres.AudiosList
	case "Text":
		mediaList = mediaMapModelRepres.TextsList
	case "Pdf":
		mediaList = mediaMapModelRepres.PdfsList
	}

	w.Header().Set("Content-Type", multiPartWriter.FormDataContentType())

	for i := 0; i < len(mediaList); i++ {

		fullFileName := mediaList[i]
		writerPortion, _ := multiPartWriter.CreateFormFile(fieldName, mediaList[i])
		file, err := os.Open(fullFileName)
		if err != nil {
			streamingMediaError = SetErrorModel(http.StatusBadRequest, fmt.Sprintf("Error while opening file:%s, type:%s", fullFileName, fieldName))
		}
		_, err = io.Copy(writerPortion, file)

		if err != nil {
			streamingMediaError = SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while copying %s, type: %s", fullFileName, fieldName))
			json.NewEncoder(w).Encode(streamingMediaError)
			return
		}
		file.Close()

	}

}
