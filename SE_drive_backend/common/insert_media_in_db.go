package common

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"fmt"
	"net/http"
)

func InsertMediaInDb(w http.ResponseWriter, token string, inputMediaFilePath string, outputMediaFilePath string, mediaType string) (models.ErrorsModel, bool) {
	var query string
	db, err := functions.DbConnect(w)

	if err != nil {

		return errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while connecting to db while uploading photo.%s", err)), false

	}
	switch mediaType {
	case "Photo":
		query = `INSERT INTO PhotoTable(token,originalPhotoFileName,outputPhotoFileName) VALUES(?,?,?)`
	}

	if _, err = db.Exec(query, token, inputMediaFilePath, outputMediaFilePath); err != nil {
		if !global.MediaMap[token].IsSubscribed {
			global.MediaMap[token].TrialsLeft += 1
		}

		//! make separate folders for each user and store media there .
		//! test for issubscribed as well .

		return errors.SetErrorModel(http.StatusBadGateway, fmt.Sprintf("Error while executing insertion in db for photo.%s", err)), false
	}
	return models.ErrorsModel{}, true
}
