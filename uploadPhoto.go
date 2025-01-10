package main

import (
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
	CORSFix(w, r)
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {

		log.Fatal(err)
	}

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
	defer uploadFromResponse(w, outputPhotoFileStr, "image", 1024*250)
	defer fmt.Print("done")

}
