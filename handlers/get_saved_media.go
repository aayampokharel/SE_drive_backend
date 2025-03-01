package handlers

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions/media_processing"
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

// this method is called whenever the person just logs in , without any associating with UI , i wnat this to be called , this is for initial saved history loading whcih should be done when user logs in immediately .
// query ma esko lagi we require a specific format adding type query.
func GetSavedMedia(w http.ResponseWriter, r *http.Request) {
	//@ Token is required .
	//@ so only after proper Login as savedhistory is only of the user .

	var getSavedPhotosModel models.GetSavedPhotosModel
	var getSavedPhotosError models.ErrorsModel
	var mediaType string = r.URL.Query().Get("type")

	if mediaType == "" { //! describe enums for mediaType .
		mediaType = "Photo"
	}

	json.NewDecoder(r.Body).Decode(&getSavedPhotosModel)

	mediaMapModelRepres, ok := global.MediaMap[getSavedPhotosModel.Token]
	if !ok {
		getSavedPhotosError = errors.SetErrorModel(http.StatusBadRequest, "Token invalid while getting saved Photos.")
		json.NewEncoder(w).Encode(getSavedPhotosError)
		return

	}
	defer media_processing.StreamMediaList(w, mediaMapModelRepres, mediaType)

}
