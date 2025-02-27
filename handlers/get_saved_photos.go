package handlers

import (
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"SE_drive_backend/models"
	"encoding/json"
	"net/http"
)

func GetSavedPhotos(w http.ResponseWriter, r *http.Request) {
	//@ Token is required .
	//@ so only after proper Login as savedhistory is only of the user .
	CORSFix(w, r)
	var getSavedPhotosModel models.GetSavedPhotosModel
	var getSavedPhotosError models.ErrorsModel
	var mediaType string = r.URL.Query().Get("type")

	if mediaType == "" { //! describe enums for mediaType .
		mediaType = "Photo"
	}

	json.NewDecoder(r.Body).Decode(&getSavedPhotosModel)

	mediaMapModelRepres, ok := global.MediaMap[getSavedPhotosModel.Token]
	if !ok {
		getSavedPhotosError = functions.SetErrorModel(http.StatusBadRequest, "Token invalid while getting saved Photos.")
		json.NewEncoder(w).Encode(getSavedPhotosError)
		return

	}
	functions.StreamMediaList(w, mediaMapModelRepres, mediaType)

}
