package handlers

import (
	"SE_drive_backend/models"
	"encoding/json"
	"mime/multipart"
	"net/http"
)

func GetHistory(w http.ResponseWriter, r *http.Request) {
	//know if response can be multipart
	//multipart writer .
	//write .
	CORSFix(w, r)
	var failure models.ErrorsModel
	if mediaType := r.Header.Get("Accept"); mediaType == "multipart/form-data" {
		failure = models.ErrorsModel{Err: "Header not able to accept multipart/form-data mediaTypes.", StatusCode: http.StatusBadRequest}
		json.NewEncoder(w).Encode(failure)
		return

	}
	mw := multipart.NewWriter(w)

}
