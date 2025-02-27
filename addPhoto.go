package main

import (
	"SE_drive_backend/functions"
	"SE_drive_backend/handlers"
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
	handlers.CORSFix(w, r)

	err := r.ParseMultipartForm(20 << 20)
	if err != nil {

		log.Fatal("Size not enough . ")
	}
	token := r.FormValue("token_id")
	if token == "" {

		json.NewEncoder(w).Encode(functions.SetErrorModel(http.StatusBadRequest, "Empty Token found .")) //direct .
		return

	}
	photoRequestModel := models.PhotoRequestModel{Token: token}

	file, header, er := r.FormFile("Photo")
	if er != nil {

		log.Fatal(er)
	}
	defer file.Close()

	fileName := "photo_" + replaceSpaceInFileName(header.Filename)
	inputPhotoFileStr := "./uploadedPhotos/" + fileName
	newPhotoFile, er := os.Create(inputPhotoFileStr)
	if er != nil {

		log.Fatal(er)
	}
	defer newPhotoFile.Close()
	_, er = io.Copy(newPhotoFile, file)
	if er != nil {
		fmt.Print("4")
		log.Fatal(er)
	}
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outputPhotoFileStr := "./uploadedPhotos/" + "output_" + baseName + ".jpeg"
	//ffmpeg -i input.png output.jpg
	//@ extensions checking left ....! undone !
	cmdStr := fmt.Sprintf("ffmpeg -i %s -qscale:v 31 -f image2 -vcodec mjpeg %s", newPhotoFile.Name(), outputPhotoFileStr)
	cmd := exec.Command("cmd", "/C", cmdStr)

	er = cmd.Run()
	if er != nil {

		log.Fatal(er)
	}
	//! what does this do? to know .
	//defer uploadFromResponse(w, outputPhotoFileStr, "image", 1024*250)
	defer fmt.Print("done")

	//--database execution --//
	db, err := functions.DbConnect(w)

	if err != nil {

		json.NewEncoder(w).Encode(functions.SetErrorModel(http.StatusBadRequest, fmt.Sprintf("error while connecting to db while uploading photo.%s", err)))

		// create table if not exists VideoTable(
		// 	count int AUTO_INCREMENT primary key,
		// 	token varchar(50) not null,
		// 	videoFileName varchar(150) not null

		// 	);

	}
	query := `INSERT INTO PhotoTable(token,originalPhotoFileName,outputPhotoFileName) VALUES(?,?,?)`

	_, err = db.Exec(query, photoRequestModel.Token, newPhotoFile.Name(), outputPhotoFileStr)
	if err != nil {
		print("error 2")
		json.NewEncoder(w).Encode(functions.SetErrorModel(http.StatusBadGateway, fmt.Sprintf("Error while executing insertion in db for photo.%s", err)))
		//! i can also throw error by making PHOTOFILENAME UNIQUE AS WHY 2 OF SAME NAME  AND THAT IS NOT POSSIBLE AS WELL . SO TEI HO .FRON FRONTEND TELL AAKASH TO CHECK IF THE FILENAME IS SAME AS OTHER THEN ONLY SEND ELSE ERROR WILL BE THROWN .

		//! ALSO DEDUCE -1 from trial photos .
		return
	}

	json.NewEncoder(w).Encode(models.LogInResponseModel{MessageStatus: "Photo  uploaded  successfully!", OriginalPhotoFileName: newPhotoFile.Name(), OutputPhotoFileName: outputPhotoFileStr})

}
