package main

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/common"
	"SE_drive_backend/global"
	"encoding/json"

	"net/http"
	"os"

	"path/filepath"
	"strings"
)

const (
	maxUploadSize = 20 << 20 //20mb
	photoType     = "Photo"
	chunkSize     = 1024 * 250
)

//@ its good to always defer the uploadFromResponse.

func uploadPhoto(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "file size exceeds the alloted limit of 20mb ."))
		return
	}

	//# token_id sent in multipart itself ,
	token := r.FormValue("token_id")
	if token == "" {

		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "Empty Token found .")) //direct .
		return

	}

	file, header, er := r.FormFile(photoType)
	if er != nil {

		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "Failed to retrieve file"))
		return
	}
	defer file.Close()
	userDir := filepath.Join("./uploadedPhotos", token)
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusInternalServerError, "Failed to create user directory"))
		return
	}
	originalFileName := replaceSpaceInFileName(header.Filename)
	inputPhotoFilePath := filepath.Join(userDir, replaceSpaceInFileName(header.Filename))

	//!make a create and save function to be reused in other .
	if errorValue, ok := common.SaveUploadedMediaInFolder(w, inputPhotoFilePath, file); !ok {
		json.NewEncoder(w).Encode(errorValue)
		return
	}
	outputPhotoFilePath := filepath.Join(userDir, "output_"+strings.TrimSuffix(originalFileName, filepath.Ext(originalFileName))+".jpeg")
	//ffmpeg -i input.png output.jpg
	//! extensions checking left ....! undone !
	if errorValue, ok := common.CompressMediaUsingFfmpeg(inputPhotoFilePath, outputPhotoFilePath, photoType); !ok {
		json.NewEncoder(w).Encode(errorValue)
		return
	}

	if mediaMapErrValue, ok := common.SaveMediaInfoInMap(w, token, inputPhotoFilePath, outputPhotoFilePath, photoType); !ok {
		json.NewEncoder(w).Encode(mediaMapErrValue)
		return
	}
	//--database execution --//
	if insertMediaDbError, ok := common.InsertMediaInDb(w, token, inputPhotoFilePath, outputPhotoFilePath, photoType); !ok {
		json.NewEncoder(w).Encode(insertMediaDbError)
		return
	}

	if !global.MediaMap[token].IsSubscribed {
		uploadFromResponse(w, outputPhotoFilePath, photoType, chunkSize)
	} else {

		uploadFromResponse(w, inputPhotoFilePath, photoType, chunkSize)
	}

}
