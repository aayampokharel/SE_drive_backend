package common

import (
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"net/http"
)

func SaveMediaInfoInMap(w http.ResponseWriter, token string, inputPhotoFilePath string, outputPhotoFilePath string, mediaType string) (models.ErrorsModel, bool) {
	if !global.MediaMap[token].IsSubscribed {
		if modelError, ok := global.AddNewMedia(token, outputPhotoFilePath, mediaType); !ok {

			return modelError, false
		}

	} else {
		global.AddNewMedia(token, inputPhotoFilePath, mediaType)

	}
	return models.ErrorsModel{}, true
}
