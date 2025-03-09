package main

import (
	errors "SE_drive_backend/Errors"
	"SE_drive_backend/functions"
	"SE_drive_backend/global"
	"encoding/json"

	"SE_drive_backend/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//@ its good to always defer the uploadFromResponse.

func uploadPhoto(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(20 << 20)
	if err != nil {

		log.Fatal("Size not enough . ")
	}
	//# token_id sent in multipart itself ,
	token := r.FormValue("token_id")
	if token == "" {

		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, "Empty Token found .")) //direct .
		return

	}
	isSubscribed := global.MediaMap[token].IsSubscribed

	photoRequestModel := models.PhotoRequestModel{Token: token}

	file, header, er := r.FormFile("Photo")
	if er != nil {

		log.Fatal(er)
	}
	defer file.Close()

	originalFileName := replaceSpaceInFileName(header.Filename)
	inputPhotoFileStr := "./uploadedPhotos/" + originalFileName
	newPhotoFile, er := os.Create(inputPhotoFileStr)
	if er != nil {

		log.Fatal(er)
	}
	defer newPhotoFile.Close()
	_, er = io.Copy(newPhotoFile, file)
	if er != nil {

		log.Fatal(er)
	}
	baseName := strings.TrimSuffix(originalFileName, filepath.Ext(originalFileName))
	outputPhotoFileStr := "./uploadedPhotos/" + "output_" + baseName + ".jpeg"
	//ffmpeg -i input.png output.jpg
	//@ extensions checking left ....! undone !
	cmdStr := fmt.Sprintf("ffmpeg -i %s -qscale:v 31 -f image2 -vcodec mjpeg %s", newPhotoFile.Name(), outputPhotoFileStr)
	cmd := exec.Command("cmd", "/C", cmdStr)

	er = cmd.Run()
	if er != nil {

		log.Fatal(er)
	}
	//can add below logic inside uploadFromResponse itself as well .

	defer fmt.Print("done")

	//--database execution --//
	db, err := functions.DbConnect(w)

	if err != nil {

		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while connecting to db while uploading photo.%s", err)))

		// create table if not exists VideoTable(
		// 	count int AUTO_INCREMENT primary key,
		// 	token varchar(50) not null,
		// 	videoFileName varchar(150) not null

		// 	);

	}
	if !isSubscribed {
		if modelError, ok := global.AddNewMedia(token, outputPhotoFileStr, "Photo"); !ok {
			json.NewEncoder(w).Encode(modelError)
			return
		}

	} else {
		global.AddNewMedia(token, newPhotoFile.Name(), "Photo")

	}

	query := `INSERT INTO PhotoTable(token,originalPhotoFileName,outputPhotoFileName) VALUES(?,?,?)`

	_, err = db.Exec(query, photoRequestModel.Token, newPhotoFile.Name(), outputPhotoFileStr)
	if err != nil {
		print("error 2")
		json.NewEncoder(w).Encode(errors.SetErrorModel(http.StatusBadGateway, fmt.Sprintf("Error while executing insertion in db for photo.%s", err)))

		//! make separate folders for each user and store media there .
		//! test for issubscribed as well .

		return
	}
	if !isSubscribed {

		uploadFromResponse(w, outputPhotoFileStr, "Photo", 1024*250)
	} else {

		uploadFromResponse(w, newPhotoFile.Name(), "Photo", 1024*250)
	}

	// json.NewEncoder(w).Encode(models.LogInResponseModel{MessageStatus: "Photo  uploaded  successfully!", OriginalPhotoFileName: newPhotoFile.Name(), OutputPhotoFileName: outputPhotoFileStr})

}
