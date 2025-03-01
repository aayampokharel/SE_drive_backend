package errors

import "SE_drive_backend/models"

func SetErrorModel(statusCode int, err string) models.ErrorsModel {

	return models.ErrorsModel{Err: err, StatusCode: statusCode}

}
