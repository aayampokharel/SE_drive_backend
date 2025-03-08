package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	//insert all thngs from added maps into db .

	//remove key from 2 maps .
	var signOutRequestModel models.SignOutRequestModel
	// var signOutResponseModel models.SignOutResponseModel

	jsonError := json.NewDecoder(r.Body).Decode(&signOutRequestModel)
	if jsonError != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "error while decoding the token from request . "))

	}

}
