package common

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SaveUploadedMediaInFolder(w http.ResponseWriter, inputPhotoFilePath string, file multipart.File) (models.ErrorsModel, bool) {
	newPhotoFile, er := os.Create(inputPhotoFilePath)
	if er != nil {

		return errors.SetErrorModel(http.StatusInternalServerError, "Failed to create input photo file path ."), false
	}
	defer newPhotoFile.Close()
	_, er = io.Copy(newPhotoFile, file)
	if er != nil {

		return errors.SetErrorModel(http.StatusInternalServerError, "Failed while copying file "), false
	}
	return models.ErrorsModel{}, true
}
